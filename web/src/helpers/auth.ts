import jwt_decode, { JwtPayload } from "jwt-decode";
import { v4 as uuidv4 } from "uuid";

import {
  FACEBOOK_OAUTH_CLIENT_ID,
  FACEBOOK_OAUTH_REDIRECT_URI,
  FACEBOOK_OAUTH_URL,
  GOOGLE_OAUTH_CLIENT_ID,
  GOOGLE_OAUTH_REDIRECT_URI,
  GOOGLE_OAUTH_URL,
  GOOGLE_OAUTH_USER_URL,
} from "../config";
import { GoogleUser, SocialType, User } from "../types/types";

export function saveOAuthState(key: string, state: string) {
  sessionStorage.setItem(key, state);
}

// generates state for OAuth
export const generateOAuthState = () => {
  const uuid = uuidv4();
  return uuid;
};

// retrieves Google user info
export async function getGoogleUser(token: string): Promise<GoogleUser | null> {
  try {
    let res = await fetch(`${GOOGLE_OAUTH_USER_URL}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      return null;
    }

    const data: GoogleUser = await res.json();
    return data;
  } catch (error) {
    console.log(error);
    return null;
  }
}

// getSocialURL returns social OAuth url and saves CSRF state in session
export function getSocialURL(type: SocialType): string {
  let url = "";
  switch (type.name) {
    case "Google": {
      const state = generateOAuthState();
      saveOAuthState("GOOGLE_OAUTH_STATE", state);
      const options = {
        client_id: GOOGLE_OAUTH_CLIENT_ID,
        redirect_uri: GOOGLE_OAUTH_REDIRECT_URI,
        response_type: "token",
        scope: [
          "https://www.googleapis.com/auth/userinfo.profile",
          "https://www.googleapis.com/auth/userinfo.email",
        ].join(" "),
        include_granted_scopes: "true",
        state: state,
        prompt: "consent",
      };

      const qs = new URLSearchParams(options);
      url = `${GOOGLE_OAUTH_URL}?${qs.toString()}`;
      break;
    }
    case "Facebook": {
      const state = generateOAuthState();
      saveOAuthState("FACEBOOK_OAUTH_STATE", state);
      const options = {
        redirect_uri: FACEBOOK_OAUTH_REDIRECT_URI as string,
        client_id: FACEBOOK_OAUTH_CLIENT_ID as string,
        response_type: "token",
        state: state,
      };

      const qs = new URLSearchParams(options);
      url = `${FACEBOOK_OAUTH_URL}?${qs.toString()}`;
      break;
    }
    case "Apple": {
      break;
    }
  }

  return url;
}

// retrieves user from local storage
export function getUserFromLocalStorage(): User | undefined {
  const storedUser = localStorage.getItem("renthome_user");
  const parsedUser = storedUser ? JSON.parse(storedUser) : undefined;
  return parsedUser;
}

// saves user in local storage
export function saveUserInLocalStorage(user: User) {
  localStorage.setItem("renthome_user", JSON.stringify(user));
}

// removes user from local storage
export function removeUserFromLocalStorage() {
  localStorage.removeItem("renthome_user");
}

// retrieves token from local storage
export function getTokenFromLocalStorage(): string | null {
  const storedToken = localStorage.getItem("renthome_token");
  const parsedToken = storedToken ? JSON.parse(storedToken) : null;
  return parsedToken;
}

// saves token in local storage
export function saveTokenInLocalStorage(token: string) {
  localStorage.setItem("renthome_token", JSON.stringify(token));
}

// retrieves reset password token from local storage
export function getResetPasswordTokenFromLocalStorage(): string | null {
  const storedToken = localStorage.getItem("renthome_reset_password_token");
  const parsedToken = storedToken ? JSON.parse(storedToken) : null;
  return parsedToken;
}
// saves reset password token in local storage
export function saveResetPasswordTokenInLocalStorage(token: string) {
  localStorage.setItem("renthome_reset_password_token", JSON.stringify(token));
}

// removes token from local storage
export function removeTokenFromLocalStorage() {
  localStorage.removeItem("renthome_token");
}

// removes reset token from local storage
export function removeResetTokenFromLocalStorage() {
  localStorage.removeItem("renthome_reset_password_token");
}

// checks if JWT access token is expired
export function isTokenExpired() {
  try {
    const token = getTokenFromLocalStorage();
    if (!token) {
      return true;
    }
    const date = new Date(0);
    const decoded: JwtPayload = jwt_decode(token);
    date.setUTCSeconds(decoded.exp as number);

    return new Date().valueOf() > date.valueOf();
  } catch (error) {
    return true;
  }
}
