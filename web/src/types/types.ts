import {
  AvailableDateCondition,
  PropertyCategory,
  PropertyType,
} from "./enums";

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
  created_at: Date;
  updated_at: Date;
  deleted_at?: Date;
}

export interface Manager {
  id: string;
  name: string;
  mobile: string;
  email: string;
  created_at: Date;
  updated_at: Date;
  deleted_at?: Date;
}

export interface Location {
  id: string;
  suburb_name: string;
  post_code: string;
  state: string;
}

export interface User {
  id: string;
  name: string;
  email: string;
  facebook_id?: string;
  google_id?: string;
  apple_id?: string;
  title?: string;
  description?: string;
  role: string;
  mobile?: string;
  is_verified: boolean;
  avatar_id?: string;
  created_at: Date;
  update_at: Date;
  deleted_at?: Date;
  notifications: Notification[];
}

export interface GoogleUser {
  sub: string;
  email: string;
  email_verified: boolean;
  name: string;
  given_name: string;
  family_name: string;
  picture: string;
  locale: string;
}

export interface FacebookUser {
  id: string;
  email: string;
  name: string;
}

export interface OAuth2Provider {
  name: string;
  bgColor: string;
  color: string;
}

export interface EmailAuthResponse {
  user: User;
  token: string;
}

export interface ErrorResponse {
  code: string;
  message: string;
}

export interface Notification {
  id: string;
  name: string;
  slug: string;
  description: string;
  category: string;
  method: NotificationMethod;
  created_at: Date;
  updated_at: Date;
  deleted_at?: Date;
}

enum NotificationMethod {
  Email = "Email",
  Push = "Push",
}

export interface Privacy {
  id: string;
  name: string;
  slug: string;
  description: string;
  created_at: Date;
  updated_at: Date;
  deleted_at?: Date;
}

export interface PrivacyState {
  privacy: Privacy;
  state: string;
}

export interface NotificationState {
  notification: Notification;
  state: string;
}

export interface GetNotificationsResponse {
  notifications: Notification[];
}

export interface ForgotPasswordResponse {
  reset_token: string;
}

interface AvailableDate {
  date: Date;
  condition: AvailableDateCondition;
}

export interface SearchFilter {
  property_types: string[];
  property_types_any: boolean;
  price_min: number;
  price_max: number;
  price_min_any: boolean;
  price_max_any: boolean;
  bed_min: number;
  bed_max: number;
  bed_min_any: boolean;
  bed_max_any: boolean;
  bathroom_count: number;
  bathroom_count_any: boolean;
  car_count: number;
  car_count_any: boolean;
  available_date: AvailableDate;
  available_date_any: boolean;
  available_now: boolean;
  is_furnished: boolean;
  has_airconditioner: boolean;
  is_pets_considered: boolean;
  has_dishwasher: boolean;
}
