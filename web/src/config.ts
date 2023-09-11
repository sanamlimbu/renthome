export const API_ADDRESS =
  import.meta.env.VITE_RENTHOME_API_ADDRESS || "localhost:8000";
export const MY_ADDRESS =
  import.meta.env.VITE_RENTHOME_PUBLIC_ADDRESS || "localhost:3000";

// OAuth Google
export const GOOGLE_OAUTH_CLIENT_ID = import.meta.env
  .VITE_RENTHOME_GOOGLE_OAUTH_CLIENT_ID;
export const GOOGLE_OAUTH_REDIRECT_URI = `${MY_ADDRESS}/oauth/google`;
export const GOOGLE_OAUTH_URL = "https://accounts.google.com/o/oauth2/v2/auth";
export const GOOGLE_OAUTH_USER_URL =
  "https://www.googleapis.com/oauth2/v3/userinfo";

// OAuth Facebook
export const FACEBOOK_OAUTH_CLIENT_ID =
  import.meta.env.VITE_RENTHOME_FACEBOOK_OAUTH_CLIENT_ID || "";
export const FACEBOOK_OAUTH_REDIRECT_URI =
  import.meta.env.VITE_RENTHOME_FACEBOOK_OAUTH_REDIRECT_URI ||
  `${MY_ADDRESS}/oauth/facebook`;
export const FACEBOOK_OAUTH_URL = "https://www.facebook.com/v16.0/dialog/oauth";
export const FACEBOOK_OAUTH_USER_URL = "https://graph.facebook.com/me";
