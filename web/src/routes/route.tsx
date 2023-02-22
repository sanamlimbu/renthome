import { useContext } from "react";
import {
  createBrowserRouter,
  Navigate,
  RouterProvider,
} from "react-router-dom";
import { UserContext } from "../context/user";
import BuyPage from "../pages/buy";
import FacebookAuthRedirectPage from "../pages/facebookAuth";
import FindAgentPage from "../pages/findAgent";
import ForgotPasswordPage from "../pages/forgotPassword";
import GoogleAuthRedirectPage from "../pages/googleAuth";
import HomePage from "../pages/home";
import LoginPage from "../pages/login";
import { MePage } from "../pages/me";
import RentPage from "../pages/rent";
import RootPage from "../pages/root";
import SignupPage from "../pages/signup";
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
      element:
        user === undefined ? <LoginPage /> : <Navigate to="/" replace={true} />,
    },
    {
      path: "/signup",
      element:
        user === undefined ? (
          <SignupPage />
        ) : (
          <Navigate to="/" replace={true} />
        ),
    },
    {
      path: "forgot-password",
      element:
        user === undefined ? (
          <ForgotPasswordPage />
        ) : (
          <Navigate to="/" replace={true} />
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
