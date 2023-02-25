import { Typography } from "@mui/material";
import { useLocation } from "react-router-dom";

export default function FacebookAuthRedirectPage() {
  const location = useLocation();
  const searchParams = new URLSearchParams(location.search);
  const token = searchParams.get("token");
  console.log(token);

  //   if (code === null) {
  //     return <></>;
  //   }

  //   (async function () {
  //     try {
  //       // retrieve Google access token
  //       const token = await getGoogleAccessToken(code);
  //       console.log(token);

  //       if (token === null) {
  //         return <></>;
  //       }

  //       // retrieve Google user
  //       const googleUser = await getGoogleUser(token);
  //       console.log(googleUser);
  //     } catch (error) {
  //       return <></>;
  //     }
  //   })();

  return (
    <>
      <Typography>Hello Facebook</Typography>
    </>
  );
}
