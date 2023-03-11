import { PropertyCategory, PropertyType } from "./enums";

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

export interface SocialType {
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

export interface ForgotPasswordResponse {
  reset_token: string;
}
