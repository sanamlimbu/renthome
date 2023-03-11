import queryString from "query-string";
import { useContext, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { API_ADDRESS } from "../../config";
import { UserContext } from "../../context/user";
import {
  getFacebookUser,
  getOAuthState,
  saveTokenInLocalStorage,
} from "../../helpers/auth";
import { User } from "../../types/types";

interface FacebookAuthResponse {
  user: User;
  token: string;
}

export default function FacebookAuthRedirectPage() {
  const parsed = queryString.parse(window.location.hash);
  const token = parsed.access_token;
  const state = parsed.state;
  const navigate = useNavigate();
  const { setUser } = useContext(UserContext);

  if (state !== getOAuthState("FACEBOOK_OAUTH_STATE")) {
    navigate("/");
  }

  useEffect(() => {
    (async function () {
      try {
        // retrieve Facebook user
        const facebookUser = await getFacebookUser(token as string);
        // OAuth app is not verfied so Facebook doesnot send user email
        if (facebookUser) {
          facebookUser.email = "subbasanam08@gmail.com";
        }

        const res = await fetch(`${API_ADDRESS}/auth/facebook`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(facebookUser),
        });
        console.log(res);
        if (res.ok) {
          const data: FacebookAuthResponse = await res.json();
          saveTokenInLocalStorage(data.token);
          setUser(data.user);
          navigate("/");
          console.log(data);
        }
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  return <></>;
}
