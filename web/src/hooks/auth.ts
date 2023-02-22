import { useState } from "react";

export function useOAuth() {
  const [googleState, setGoogleState] = useState();

  function getGoogleState() {}

  function getFacebookState() {}

  return {
    getGoogleState,
    getFacebookState,
  };
}
