import React, { useReducer } from "react";
import { User } from "../types/types";
import { getUserFromLocalStorage } from "../utils/auth";
import AuthReducer, { IAuthAction } from "./authReducer";

interface IAuthContext {
  currentUser: User | null;
  dispatch: React.Dispatch<IAuthAction>;
}

const INITIAL_STATE: IAuthContext = {
  currentUser: getUserFromLocalStorage(),
  dispatch: () => {},
};

const AuthContext = React.createContext<IAuthContext>(INITIAL_STATE);

const AuthContextProvider = (props: { children: React.ReactNode }) => {
  const [state, dispatch] = useReducer(AuthReducer, INITIAL_STATE);

  return (
    <AuthContext.Provider value={{ currentUser: state.currentUser, dispatch }}>
      {props.children}
    </AuthContext.Provider>
  );
};

export { AuthContext, AuthContextProvider };
