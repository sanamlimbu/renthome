import { Action } from "react-fetching-library";
import { Property } from "../types/types";

const getPropertyList: Action<Property[]> = {
  method: "GET",
  endpoint: "/users",
};

const getSuggestedLocations: Action<Location[]> = {
  method: "GET",
  endpoint: "/suggest-locations",
};

export const query = {
  getPropertyList,
  getSuggestedLocations,
};
