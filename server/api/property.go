package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"renthome/boiler"

	"github.com/h2non/filetype"
	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Filter struct {
	All               bool   `json:"all"`
	House             bool   `json:"house"`
	Apartment         bool   `json:"apartment"`
	Unit              bool   `json:"unit"`
	TownHouse         bool   `json:"town_house"`
	Villa             bool   `json:"villa"`
	PriceMin          string `json:"price_min"`
	PriceMax          string `json:"price_max"`
	BedMin            string `json:"bed_min"`
	BedMax            string `json:"bed_max"`
	BathroomCount     string `json:"bathroom_count"`
	CarCount          string `json:"car_count"`
	AvailableDate     string `json:"available_date"`
	IsFurnished       bool   `json:"is_furnished"`
	IsPetsConsidered  bool   `json:"is_pets_conisdereed"`
	HasAirConditioner bool   `json:"has_air_conditioner"`
	HasDishwasher     bool   `json:"has_dishwasher"`
}
type GetPropertiesRequest struct {
	Search string `json:"search"`
	Filter Filter `json:"filter"`
}

func (api *APIController) GetProperties(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetPropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	return 200, nil

}

func (api *APIController) GetProperty(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetPropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	return 200, nil

}

type CreatePropertyRequest struct {
	Type             string    `json:"type" validate:"required"`
	Category         string    `json:"category" validate:"required"`
	Street           string    `json:"street" validate:"required"`
	Suburb           string    `json:"suburb" validate:"required"`
	Postcode         string    `json:"postcode" validate:"required"`
	State            string    `json:"state" validate:"required"`
	BedCount         int       `json:"bed_count" validate:"required"`
	BathCount        int       `json:"bath_count" validate:"required"`
	CarCount         int       `json:"car_count" validate:"required"`
	HasAirCon        bool      `json:"has_aircon" validate:"required"`
	IsFurnished      bool      `json:"is_furnished" validate:"required"`
	IsPetsConsidered bool      `json:"is_pets_considered" validate:"required"`
	AvailableAt      null.Time `json:"available_at" validate:""`
	OpenAt           null.Time `json:"open_at" validate:""`
	Price            int       `json:"price" validate:"required"`
}

type CreatePropertyResponse struct {
	Property *boiler.Property `json:"property"`
}

func (api *APIController) CreateProperty(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &CreatePropertyRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	// check of zero value

	// if req.AvailableAt.Valid {
	// 	availableAt, err = time.Parse(time.RFC3339, req.AvailableAt)
	// 	if err != nil {
	// 		return http.StatusBadRequest, terror.Error(err, "Invalid date.")
	// 	}
	// }

	// openAt, err := time.Parse(time.RFC3339, req.AvailableAt.String())
	// if err != nil {
	// 	return http.StatusBadRequest, terror.Error(err, "Invalid date.")
	// }

	// propertyID, err := uuid.NewV4()
	// if err != nil {
	// 	return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	// }

	fmt.Println("available", req.AvailableAt)
	fmt.Println("Open", req.OpenAt)

	property := &boiler.Property{
		Type:             req.Type,
		Category:         req.Category,
		Street:           req.Street,
		Suburb:           req.Suburb,
		Postcode:         req.Postcode,
		State:            req.State,
		BedCount:         req.BedCount,
		BathCount:        req.BathCount,
		CarCount:         req.CarCount,
		HasAircon:        req.HasAirCon,
		IsFurnished:      req.IsFurnished,
		IsPetsConsidered: req.IsPetsConsidered,
		AvailableAt:      req.AvailableAt,
		OpenAt:           req.OpenAt,
		Price:            req.Price,
		ManagerID:        "90b71c18-c836-421b-9e17-0bb119019baa",
		AgencyID:         "5d621a17-6ea0-430a-98f0-ea419097c751",
	}

	// begin transaction
	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	err = property.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "Unable to create property.")
	}

	// for _, image := range images {
	// 	image.PropertyID = propertyID.String()
	// 	image.UploaderID = propertyID.String()

	// 	image.Insert(tx, boil.Infer())
	// 	if err != nil {
	// 		return http.StatusInternalServerError, terror.Error(err, "Unable to create property.")
	// 	}
	// }

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrCommitTransaction)
	}

	createPropertyResponse := &CreatePropertyResponse{
		Property: property,
	}

	if err = json.NewEncoder(w).Encode(createPropertyResponse); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	return http.StatusCreated, nil
}

func (api *APIController) UpdateProperty(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	req := &GetPropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	return 200, nil

}

func (api *APIController) DeleteProperty(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	req := &GetPropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	return 200, nil

}

// parseImages will read a multipart form request and returns slice of Image
func parseImages(r *http.Request) ([]*boiler.Image, error) {

	var images []*boiler.Image

	multipartReader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	for {
		part, err := multipartReader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		data, err := ioutil.ReadAll(part)
		if err != nil {
			return nil, err
		}

		// handle file
		if part.FormName() == "file" {
			// get mime type
			kind, err := filetype.Match(data)
			if err != nil {
				return nil, err
			}

			mimeType := kind.MIME.Value
			extension := kind.Extension

			if kind == filetype.Unknown {
				return nil, err
			}

			image := &boiler.Image{
				FileSizeBytes: int64(len(data)),
				Extension:     extension,
				MimeType:      mimeType,
			}

			images = append(images, image)

		}
	}
	return images, nil
}
