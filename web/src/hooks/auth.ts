import React from "react";

export function useOAuth() {
  const [googleState, setGoogleState] = React.useState();

  function getGoogleState() {}

  function getFacebookState() {}

  return {
    getGoogleState,
    getFacebookState,
  };
}
