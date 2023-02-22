import React, { useEffect, useState } from "react";
import {
  getUserFromLocalStorage,
  isTokenExpired,
  removeTokenFromLocalStorage,
  removeUserFromLocalStorage,
} from "../helpers/auth";
import { User } from "../types/types";

interface IUserContext {
  user: User | undefined;
  setUser: React.Dispatch<React.SetStateAction<User | undefined>>;
}

const UserContext = React.createContext<IUserContext>({} as IUserContext);

const UserContextProvider = (props: { children: React.ReactNode }) => {
  const [user, setUser] = useState<User | undefined>(getUserFromLocalStorage());
  useEffect(() => {
    if (!user) {
      return;
    }

    if (isTokenExpired()) {
      setUser(undefined);
      removeUserFromLocalStorage();
      removeTokenFromLocalStorage();
    }
  }, [user]);

  return (
    <UserContext.Provider value={{ user, setUser }}>
      {props.children}
    </UserContext.Provider>
  );
};

export { UserContext, UserContextProvider };
