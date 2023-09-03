import { User } from "../types/types";

/**
 * Retrieves the user information from localStorage.
 * @returns The user object if found, otherwise null.
 */
export function getUserFromLocalStorage(): User | null {
  try {
    const userKey = Object.keys(window.localStorage).find((key) =>
      key.startsWith("renthome:authUser")
    );
    if (userKey) {
      const userString: string | null = localStorage.getItem(userKey);
      if (userString) {
        const currentUser: User = JSON.parse(userString);
        return currentUser;
      }
    }
  } catch (error) {
    console.error("Error retrieving user from localStorage:", error);
  }
  return null;
}
