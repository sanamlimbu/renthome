import react from "@vitejs/plugin-react";
import { defineConfig, loadEnv } from "vite";

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd(), "");

  return {
    plugins: [react()],
    define: {
      "process.env.VITE_RENTHOME_API_ADDRESS": JSON.stringify(
        env.VITE_RENTHOME_API_ADDRESS
      ),
      "process.env.VITE_RENTHOME_PUBLIC_ADDRESS": JSON.stringify(
        env.VITE_RENTHOME_PUBLIC_ADDRESS
      ),
      "process.env.VITE_RENTHOME_GOOGLE_OAUTH_CLIENT_ID": JSON.stringify(
        env.VITE_RENTHOME_GOOGLE_OAUTH_CLIENT_ID
      ),
      "process.env.VITE_RENTHOME_FACEBOOK_OAUTH_CLIENT_ID": JSON.stringify(
        env.VITE_RENTHOME_FACEBOOK_OAUTH_CLIENT_ID
      ),
    },
  };
});
