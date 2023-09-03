package api

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"renthome/boiler"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/h2non/filetype"
	"github.com/minio/minio-go/v7"
	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const MAX_UPLOAD_SIZE = 32<<20 + 1024 // 32MB

var (
	ErrUnknownFileType = fmt.Errorf("file type is unknown")
	ErrInvalidFileType = fmt.Errorf("file type is invalid")
)

type FileUploadResponse struct {
	ID string `json:"id"`
}

func (api *APIController) FileUpload(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	defer r.Body.Close()

	// Get blob
	blob, _, err := parseUploadRequest(w, r, nil)
	if errors.Is(err, ErrUnknownFileType) {
		return http.StatusBadRequest, terror.Error(err, "file type not allowed")
	}
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "something went wrong, please try again")
	}

	if blob == nil {
		return http.StatusInternalServerError, terror.Error(err, "file is required")
	}

	// Get arguments
	public := r.URL.Query().Get("public")
	if public == "true" {
		blob.Public = true
	}

	// Upload to object storage
	if api.ObjectStorage == nil {
		return http.StatusInternalServerError, terror.Error(err, "error uploading to object storage")
	}

	reader := bytes.NewReader(blob.File)
	_, err = api.ObjectStorage.PutObject(
		context.Background(),
		api.ObjectStorage.BucketName,
		blob.FileName,
		reader,
		-1,
		minio.PutObjectOptions{
			UserMetadata: map[string]string{
				"filename": blob.FileName,
				"filesize": strconv.FormatInt(blob.FileSizeBytes, 10),
			},
			ContentType: blob.MimeType,
		},
	)

	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "error uploading to object storage")
	}

	err = blob.Insert(api.Conn, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	fileUploadResponse := &FileUploadResponse{
		ID: blob.ID,
	}

	if err = json.NewEncoder(w).Encode(fileUploadResponse); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	return http.StatusOK, nil

}

// parseUploadRequest will read a multipart form request that includes both a file, and a request body
// returns a blob struct, ready to be inserted, as well as decoding json into supplied interface when present
func parseUploadRequest(w http.ResponseWriter, r *http.Request, req interface{}) (*boiler.Blob, map[string]string, error) {
	// Limit size to 32MB
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	mr, err := r.MultipartReader()
	if err != nil {
		return nil, nil, terror.Error(err, "parse error")
	}

	var blob *boiler.Blob
	params := make(map[string]string)

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, terror.Error(err, "parse error")
		}

		data, err := ioutil.ReadAll(part)
		if err != nil {
			return nil, nil, terror.Error(terror.ErrParse, "parse error")
		}

		// handle file
		if part.FormName() == "file" {
			// get mime type
			kind, err := filetype.Match(data)
			if err != nil {
				return nil, nil, terror.Error(terror.ErrParse, "parse error")
			}

			if kind == filetype.Unknown {
				return nil, nil, terror.Error(ErrUnknownFileType, "unsupported file type")
			}

			mimeType := kind.MIME.Value
			extension := kind.Extension

			hasher := md5.New()
			_, err = hasher.Write(data)
			if err != nil {
				return nil, nil, terror.Error(err, "hash error")
			}

			hashResult := hasher.Sum(nil)
			hash := hex.EncodeToString(hashResult)

			// the filepath is sent with ! because part returns filepath.Base(filename) removes the filePath
			fileName := strings.Replace(part.FileName(), "!", "/", -1)
			id, err := uuid.NewV4()
			if err != nil {
				return nil, nil, terror.Error(err, "uuid error")
			}

			blob = &boiler.Blob{
				ID:            id.String(),
				FileName:      fileName + id.String(),
				MimeType:      mimeType,
				Extension:     extension,
				FileSizeBytes: int64(len(data)),
				File:          data,
				Hash:          hash,
			}
		} else {
			params[part.FormName()] = string(data)
		}

		// handle JSON body
		if req != nil && part.FormName() == "json" {
			err = json.NewDecoder(part).Decode(req)
			if err != nil {
				return nil, nil, terror.Error(err, "parse error")
			}
		}
	}

	return blob, params, nil
}
