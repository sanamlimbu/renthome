import { Action } from "react-fetching-library";

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

export const mutation = {
  fileUpload,
};
