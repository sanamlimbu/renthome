import { Action } from "react-fetching-library";
import { GoogleAuthResponse, GoogleUser } from "../types/types";

const fileUpload = (values: {
  file: File;
  public?: boolean;
  filePath?: string;
}): Action<{ id: string; msg: string }> => {
  const formData = new FormData();
  formData.set("file", values.file, values.filePath);
  return {
    method: "POST",
    endpoint: `/files/upload${values.public ? "?public=true" : ""}`,
    credentials: "include",
    body: formData,
    responseType: "json",
  };
};

const googleOAuth = (values: GoogleUser): Action<GoogleAuthResponse> => {
  return {
    method: "POST",
    endpoint: "/auth/google",
    credentials: "include",
    body: values,
    responseType: "json",
    headers: {
      "Content-Type": "application/json",
    },
  };
};

export const mutation = {
  fileUpload,
  googleOAuth,
};
