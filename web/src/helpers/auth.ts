import jwt_decode, { JwtPayload } from "jwt-decode";
import { v4 as uuidv4 } from "uuid";

import {
  FACEBOOK_OAUTH_CLIENT_ID,
  FACEBOOK_OAUTH_REDIRECT_URI,
  FACEBOOK_OAUTH_URL,
  FACEBOOK_OAUTH_USER_URL,
  GOOGLE_OAUTH_CLIENT_ID,
  GOOGLE_OAUTH_REDIRECT_URI,
  GOOGLE_OAUTH_URL,
  GOOGLE_OAUTH_USER_URL,
} from "../config";
import { FacebookUser, GoogleUser, OAuth2Provider, User } from "../types/types";

export function saveOAuthState(key: string, state: string) {
  sessionStorage.setItem(key, state);
}

export function getOAuthState(key: string) {
  return sessionStorage.getItem(key);
}

// Generates unique state for OAuth
export const generateOAuthState = () => {
  const uuid = uuidv4();
  return uuid;
};

// Retrieves Google user info using token sent to redirect URL by Google OAuth 2.0 server
// by fetching to Google OAuth 2.0 userinfo endpoint
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

// Retrieves Facebook user info using token sent to redirect URL by Facebook OAuth 2.0 server
// by fetching to Facebook OAuth 2.0 me endpoint
export async function getFacebookUser(
  token: string
): Promise<FacebookUser | null> {
  try {
    const params = new URLSearchParams({
      fields: "id,name,email",
    });

    let res = await fetch(`${FACEBOOK_OAUTH_USER_URL}?${params}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      return null;
    }

    const data: FacebookUser = await res.json();
    return data;
  } catch (error) {
    console.log(error);
    return null;
  }
}

// Returns OAuth 2.0 endpoint based on the provider and while saving CSRF state in session
export function getOAuth2Endpoint(type: OAuth2Provider): string {
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
      // Aborted, Apple Developer Account costs $135 yearly
      break;
    }
  }

  return url;
}

// retrieves user from local storage
// export function getUserFromLocalStorage(): User | undefined {
//   const storedUser = localStorage.getItem("renthome_user");
//   const parsedUser = storedUser ? JSON.parse(storedUser) : undefined;
//   return parsedUser;
// }

// Retrieves renthome_token from local storage
export function getTokenFromLocalStorage(): string | null {
  const storedToken = localStorage.getItem("renthome_token");
  const parsedToken = storedToken ? JSON.parse(storedToken) : null;
  return parsedToken;
}

// Saves renthome_token in local storage
export function saveTokenInLocalStorage(token: string) {
  localStorage.setItem("renthome_token", JSON.stringify(token));
}

// Retrieves renthome_reset_password_token token from local storage
export function getResetPasswordTokenFromLocalStorage(): string | null {
  const storedToken = localStorage.getItem("renthome_reset_password_token");
  const parsedToken = storedToken ? JSON.parse(storedToken) : null;
  return parsedToken;
}
// Saves renthome_reset_password_token in local storage
export function saveResetPasswordTokenInLocalStorage(token: string) {
  localStorage.setItem("renthome_reset_password_token", JSON.stringify(token));
}

// Removes renthome_token from local storage
export function removeTokenFromLocalStorage() {
  localStorage.removeItem("renthome_token");
}

// Removes renthome_reset_password_token from local storage
export function removeResetTokenFromLocalStorage() {
  localStorage.removeItem("renthome_reset_password_token");
}

interface JwtPayloadWithCustomClaims extends JwtPayload {
  email: string;
  name: string;
  is_verified: string;
}
// Retrives basic user from renthome_token
export function getUserFromToken() {
  try {
    const token = getTokenFromLocalStorage();
    if (!token) {
      return undefined;
    }
    const date = new Date(0);
    const decoded: JwtPayloadWithCustomClaims = jwt_decode(token);
    date.setUTCSeconds(decoded.exp as number);

    // expried
    if (new Date().valueOf() > date.valueOf()) {
      return undefined;
    }

    const user = {
      id: decoded.sub,
      name: decoded.name,
      email: decoded.email,
      is_verified: decoded.is_verified,
    };

    return user as unknown as User;
  } catch (error) {
    return undefined;
  }
}

// Checks if JWT access token (renthome_token) is expired
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
