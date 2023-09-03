import { useContext } from "react";
import {
  Navigate,
  RouterProvider,
  createBrowserRouter,
} from "react-router-dom";
import { AuthContext } from "../context/authContext";
import PasswordResetPage from "../pages/auth/passwordReset";
import SignInPage from "../pages/auth/signin";
import SignUpPage from "../pages/auth/signup";
import HomePage from "../pages/home";
import CreatePropertyPage from "../pages/property/create";

export default function RoutesProvider() {
  const { currentUser } = useContext(AuthContext);

  const RequireAuth = (props: { children: JSX.Element }) => {
    if (!currentUser) {
      return <Navigate to="/sign-in" />;
    } else {
      return props.children;
    }
  };

  const router = createBrowserRouter([
    {
      path: "/",
      element: (
        <RequireAuth>
          <HomePage />
        </RequireAuth>
      ),
    },
    {
      path: "sign-in",
      element: <SignInPage />,
    },
    {
      path: "sign-up",
      element: <SignUpPage />,
    },
    {
      path: "password-reset",
      element: <PasswordResetPage />,
    },
    {
      path: "/properties/create",
      element: <CreatePropertyPage />,
    },
  ]);

  return <RouterProvider router={router} />;
}
