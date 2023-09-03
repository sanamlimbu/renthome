// Inspired by: https://github.com/HamedBahram/dropzone/blob/main/components/Dropzone.jsx

import { UploadFileOutlined } from "@mui/icons-material";
import { Box, Button, Typography } from "@mui/material";
import { useCallback, useEffect, useState } from "react";
import { FileRejection, useDropzone } from "react-dropzone";
import FileCard from "./fileCard";

interface FileWithPreview extends File {
  preview: string;
}
interface ImagesUploadProps {
  onFilesUpload: (files: File[]) => void;
}
export default function FilesUpload({ onFilesUpload }: ImagesUploadProps) {
  const [files, setFiles] = useState<FileWithPreview[]>([]);
  const [rejected, setRejected] = useState<FileRejection[]>([]);

  useEffect(() => {
    onFilesUpload(files);
  }, [files]);

  const onDrop = useCallback(
    (acceptedFiles: File[], rejectedFiles: FileRejection[]) => {
      if (acceptedFiles?.length) {
        setFiles((previousFiles) => [
          ...previousFiles,
          ...acceptedFiles.map((file) =>
            Object.assign(file, { preview: URL.createObjectURL(file) })
          ),
        ]);
      }

      if (rejectedFiles?.length) {
        setRejected((previousFiles) => [...previousFiles, ...rejectedFiles]);
      }
    },
    []
  );

  useEffect(() => {
    // Revoke the data uris to avoid memory leaks
    return () => files.forEach((file) => URL.revokeObjectURL(file.preview));
  }, [files]);

  const checkDuplicateImage = (file: File) => {
    if (files.some((existingFile) => existingFile.name === file.name)) {
      return {
        message: "Duplicate image",
        code: "duplicate",
      };
    }
    return null;
  };

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: {
      "image/*": [],
    },
    maxFiles: 5,
    maxSize: 1024 * 1000,
    validator: checkDuplicateImage,
  });

  const removeFile = (name: string) => {
    setFiles((files) => files.filter((file) => file.name !== name));
  };

  const removeAll = () => {
    setFiles([]);
    setRejected([]);
  };

  const removeRejected = (name: string) => {
    setRejected((files) => files.filter(({ file }) => file.name !== name));
  };

  return (
    <div>
      <div
        {...getRootProps({
          style: {
            padding: "64px",
            border: "solid 1px lightgray",
          },
        })}
      >
        <input {...getInputProps()} className="padding:" />
        <div
          style={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            flexDirection: "column",
          }}
        >
          <UploadFileOutlined />
          {isDragActive ? (
            <p>Drop the files here ...</p>
          ) : (
            <p>Drag & drop some files here, or click to select files</p>
          )}
        </div>
      </div>
      {(files.length !== 0 || rejected.length !== 0) && (
        <Box
          sx={{
            marginTop: "1em",
          }}
        >
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              gap: "1em",
              marginBottom: "1em",
            }}
          >
            <Typography fontWeight="bold">Preview</Typography>
            <Button
              variant="contained"
              color="warning"
              sx={{ textTransform: "none" }}
              onClick={removeAll}
            >
              Remove all files
            </Button>
          </Box>
          {files.length !== 0 && (
            <div>
              <Typography fontWeight={"bold"}>Accepted files</Typography>
              {files.map((file) => (
                <FileCard
                  key={file.name}
                  src={file.preview}
                  name={file.name}
                  remove={removeFile}
                />
              ))}
            </div>
          )}

          {rejected.length !== 0 && (
            <div>
              <Typography fontWeight={"bold"}>Rejected files</Typography>
              <div>
                {rejected.map(({ file, errors }) => (
                  <div key={file.name} style={{ display: "flex", gap: "1em" }}>
                    <div>
                      <Typography>{file.name}</Typography>
                      <div>
                        {errors.map((error) => (
                          <div key={error.code}>
                            <Typography color="red" marginBottom={"1em"}>
                              {error.message}
                            </Typography>
                          </div>
                        ))}
                      </div>
                    </div>
                    <div>
                      <Button
                        variant="outlined"
                        onClick={() => removeRejected(file.name)}
                      >
                        Remove
                      </Button>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}
        </Box>
      )}
    </div>
  );
}
