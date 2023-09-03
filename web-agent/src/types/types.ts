import { AustraliaState, PropertyCategory, PropertyType } from "./enum";

export interface User {
  id: string;
  name: string;
  email: string;
  title?: string;
  description?: string;
  role: string;
  mobile?: string;
  is_verified: boolean;
  avatar_id?: string;
  created_at: Date;
  update_at: Date;
  deleted_at?: Date;
}

export interface Property {
  id: string;
  type: PropertyType;
  category: PropertyCategory;
  street: string;
  suburb: string;
  postcode: string;
  state: string;
  bed_count: number;
  bath_count: number;
  car_count: number;
  has_aircon: boolean;
  is_furnished: boolean;
  is_pets_considered: boolean;
  available_at: Date;
  created_at: Date;
  updated_at: Date;
  deleted_at?: Date;
}

export interface CreatePropertyRequest {
  type: PropertyType;
  category: PropertyCategory;
  street: string;
  suburb: string;
  postcode: number;
  state: AustraliaState;
  bed_count: number;
  bath_count: number;
  car_count: number;
  has_aircon: boolean;
  is_furnished: boolean;
  is_pets_considered: boolean;
  available_at?: Date;
  open_at?: Date;
  price: number;
}
