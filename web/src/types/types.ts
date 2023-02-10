import { Perm, PropertyCategory, PropertyType } from "./enums";

export interface Property {
  id: string;
  type: PropertyType;
  category: PropertyCategory;
  street: string;
  suburb: string;
  postcode: string;
  state: string;
  bedCount: number;
  bathCount: number;
  carCount: number;
  hasAircon: boolean;
  isFurnished: boolean;
  isPetsConsidered: boolean;
  availableAt: Date;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
}

export interface Agent {
  id: string;
  name: string;
  description: string;
  street: string;
  suburb: string;
  postcode: string;
  state: string;
  phone: string;
  email: string;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
}

export interface Manager {
  id: string;
  name: string;
  mobile: string;
  email: string;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
}

export interface User {
  name: string;
  mobile: string;
  email: string;
  isVerified: boolean;
  avatarID: string;
  roleID: string;
  role: Role;

  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
}

export interface Role {
  id: string;
  name: string;
  permissions: Perm[];
  tier: number;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
}

export interface GoogleAuthRequest {
  code: string;
  redirect_uri: string;
}
