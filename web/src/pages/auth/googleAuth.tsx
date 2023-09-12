import queryString from "query-string";
import React, { useContext, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { API_ADDRESS } from "../../config";
import { UserContext } from "../../context/user";
import {
  getGoogleUser,
  getOAuthState,
  saveTokenInLocalStorage,
} from "../../helpers/auth";
import { User } from "../../types/types";

interface GoogleAuthResponse {
  user: User;
  token: string;
}

function GoogleAuthRedirectPage() {
  const parsed = queryString.parse(window.location.hash);
  const token = parsed.access_token;
  const state = parsed.state;
  const navigate = useNavigate();
  const { setUser } = useContext(UserContext);

  if (state !== getOAuthState("GOOGLE_OAUTH_STATE")) {
    navigate("/");
  }

  useEffect(() => {
    console.log("i fire once");
    (async function () {
      try {
        // retrieve Google user
        const googleUser = await getGoogleUser(token as string);
        if (googleUser) {
          const res = await fetch(`${API_ADDRESS}/api/auth/google`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(googleUser),
          });

          if (res.ok) {
            const data: GoogleAuthResponse = await res.json();
            saveTokenInLocalStorage(data.token);
            setUser(data.user);
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
