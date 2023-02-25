import { useContext } from "react";
import {
  createBrowserRouter,
  Navigate,
  RouterProvider,
} from "react-router-dom";
import { UserContext } from "../context/user";
import ConfirmForgotPasswordPage from "../pages/auth/confirmForgotPassword";
import EmailUpdatePage from "../pages/auth/emailUpdate";
import FacebookAuthRedirectPage from "../pages/auth/facebookAuth";
import ForgotPasswordPage from "../pages/auth/forgotPassword";
import GoogleAuthRedirectPage from "../pages/auth/googleAuth";
import LoginPage from "../pages/auth/login";
import PasswordUpdatePage from "../pages/auth/passwordUpdate";
import SignupPage from "../pages/auth/signup";
import BuyPage from "../pages/buy";
import FindAgentPage from "../pages/findAgent";
import HomePage from "../pages/home";
import { MePage } from "../pages/me";
import RentPage from "../pages/rent";
import RootPage from "../pages/root";
import SoldPage from "../pages/sold";

export default function RoutesProvider() {
  const { user } = useContext(UserContext);
  const router = createBrowserRouter([
    {
      path: "/",
      element: <RootPage />,
      children: [
        {
          path: "",
          element: <HomePage />,
        },
        {
          path: "/me",
          element:
            user === undefined ? (
              <Navigate to="/login" replace={true} />
            ) : (
              <MePage />
            ),
        },
        {
          path: "rent",
          element: <RentPage />,
        },
        {
          path: "/sold",
          element: <SoldPage />,
        },
        {
          path: "/buy",
          element: <BuyPage />,
        },
        {
          path: "find-agent",
          element: <FindAgentPage />,
        },
      ],
    },
    {
      path: "/login",
      element: user ? <Navigate to="/" replace={true} /> : <LoginPage />,
    },
    {
      path: "/signup",
      element: user ? <Navigate to="/" replace={true} /> : <SignupPage />,
    },
    {
      path: "/forgot-password",
      element: <ForgotPasswordPage />,
    },
    {
      path: "/confirm-forgot-password",
      element: <ConfirmForgotPasswordPage />,
    },
    {
      path: "/password-update",
      element: user ? (
        <PasswordUpdatePage />
      ) : (
        <Navigate to="/login" replace={true} />
      ),
    },
    {
      path: "/email-update",
      element: user ? (
        <EmailUpdatePage />
      ) : (
        <Navigate to="/login" replace={true} />
      ),
    },
    {
      path: "/oauth/google",
      element: <GoogleAuthRedirectPage />,
    },
    {
      path: "/oauth/facebook",
      element: <FacebookAuthRedirectPage />,
    },
  ]);
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}
