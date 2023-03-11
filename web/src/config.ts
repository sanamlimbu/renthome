// OAuth Google
export const GOOGLE_OAUTH_CLIENT_ID =
  process.env.REACT_APP_RENTHOME_GOOGLE_OAUTH_CLIENT_ID || "";
export const GOOGLE_OAUTH_URL = "https://accounts.google.com/o/oauth2/v2/auth";
export const GOOGLE_OAUTH_USER_URL =
  "https://www.googleapis.com/oauth2/v3/userinfo";
export const GOOGLE_OAUTH_REDIRECT_URI = "http://localhost:3000/oauth/google";

// OAuth Facebook
export const FACEBOOK_OAUTH_CLIENT_ID =
  process.env.REACT_APP_RENTHOME_FACEBOOK_OAUTH_CLIENT_ID || "";
export const FACEBOOK_OAUTH_URL = "https://www.facebook.com/v16.0/dialog/oauth";
export const FACEBOOK_OAUTH_REDIRECT_URI =
  "http://localhost:3000/oauth/facebook";
export const FACEBOOK_OAUTH_USER_URL = "https://graph.facebook.com/me";

// OAuth Apple
export const APPLE_OAUTH_CLIENT_ID =
  process.env.REACT_APP_RENTHOME_APPLE_OAUTH_CLIENT_ID || "";
export const APPLE_OAUTH_CLIENT_SECRET =
  process.env.REACT_APP_RENTHOME_APPLE_OAUTH_CLIENT_SECRET || "";
export const APPLE_OAUTH_URL = "https://www.facebook.com/v16.0/dialog/oauth";
export const APPLE_OAUTH_TOKEN_URL = "";
export const APPLE_OAUTH_USER_URL = "";
export const APPLE_OAUTH_REDIRECT_URI = "http://localhost/oauth/apple";

export const API_ADDRESS = process.env.REACT_APP_RENTHOME_API_ADDR;
export const OAUTH_STATE_KEY =
  process.env.REACT_APP_RENTHOME_OAUTH_STATE_KEY ||
  "11041b9d-e0d0-4b4d-8bb6-287904bd469e";
