import { NotificationState, PrivacyState } from "../types/types";

// type guard
export const isNotificationState = (
  item: NotificationState | PrivacyState
): item is NotificationState => {
  return "notification" in item;
};

// type guard
export const isPrivacyState = (
  item: NotificationState | PrivacyState
): item is PrivacyState => {
  return "privacy" in item;
};
