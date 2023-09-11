import queryString from "query-string";
import React, { useContext, useEffect } from "react";
import { useMutation } from "react-fetching-library";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../../context/user";
import { fetching } from "../../fetching";
import {
  getGoogleUser,
  getOAuthState,
  saveTokenInLocalStorage,
} from "../../helpers/auth";

function GoogleAuthRedirectPage() {
  const parsed = queryString.parse(window.location.hash);
  const token = parsed.access_token;
  const state = parsed.state;
  const navigate = useNavigate();
  const { setUser } = useContext(UserContext);
  const { mutate } = useMutation(fetching.mutation.googleOAuth);

  if (state !== getOAuthState("GOOGLE_OAUTH_STATE")) {
    navigate("/");
  }

  useEffect(() => {
    (async function () {
      try {
        // retrieve Google user
        const googleUser = await getGoogleUser(token as string);
        if (googleUser) {
          const { error, payload } = await mutate(googleUser);
          if (!error && payload) {
            saveTokenInLocalStorage(payload.token);
            setUser(payload.user);
            navigate("/");
          }
        }
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  return <></>;
}

export default React.memo(GoogleAuthRedirectPage);
