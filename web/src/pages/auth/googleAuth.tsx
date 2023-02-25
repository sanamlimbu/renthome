import { Typography } from "@mui/material";
import queryString from "query-string";
import { useEffect } from "react";
import { API_ADDRESS } from "../../config";
import { getGoogleUser } from "../../helpers/auth";

export default function GoogleAuthRedirectPage() {
  const parsed = queryString.parse(window.location.hash);
  const token = parsed.access_token;
  useEffect(() => {
    (async function () {
      try {
        // retrieve Google user
        const googleUser = await getGoogleUser(token as string);
        const res = await fetch(`${API_ADDRESS}/auth/google`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(googleUser),
        });

        const user = await res.json();
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  return (
    <>
      <Typography>Hello Google</Typography>
    </>
  );
}
