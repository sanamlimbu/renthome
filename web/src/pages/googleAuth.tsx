import { Typography } from "@mui/material";
import { useLocation } from "react-router-dom";
import { API_ADDRESS, GOOGLE_REDIRECT_URI } from "../config";
import { GoogleAuthRequest } from "../types/types";

export default function GoogleAuthRedirectPage() {
  const location = useLocation();
  const searchParams = new URLSearchParams(location.search);
  const code = searchParams.get("code");
  const data: GoogleAuthRequest = {
    code: code || "",
    redirect_uri: GOOGLE_REDIRECT_URI,
  };

  (async function () {
    try {
      const res = await fetch(`${API_ADDRESS}/auth/google`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      const user = await res.json();
      console.log(user);
    } catch (error) {}
  })();

  // const { loading, payload, error, query } = useQuery({
  //   method: "POST",
  //   endpoint: `${API_ADDRESS}/auth/google`,
  // });

  // console.log(payload);
  return (
    <>
      <Typography>Hello Google</Typography>
    </>
  );
}
