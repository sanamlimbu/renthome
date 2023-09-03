import { Action } from "react-fetching-library";
import { Property } from "../types/types";

const getPropertyList: Action<Property[]> = {
  method: "GET",
  endpoint: "/users",
};

const getProperty: Action<Property[]> = {
  method: "GET",
  endpoint: "/users",
};

export const query = {
  getPropertyList,
  getProperty,
};
