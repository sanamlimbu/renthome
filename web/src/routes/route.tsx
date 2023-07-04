import { useContext } from "react";
import {
  createBrowserRouter,
  Navigate,
  RouterProvider,
} from "react-router-dom";
import { UserContext } from "../context/user";
import { AccountPage } from "../pages/account";
import ConfirmForgotPasswordPage from "../pages/auth/confirmForgotPassword";
import EmailUpdatePage from "../pages/auth/emailUpdate";
import FacebookAuthRedirectPage from "../pages/auth/facebookAuth";
import ForgotPasswordPage from "../pages/auth/forgotPassword";
import GoogleAuthRedirectPage from "../pages/auth/googleAuth";
import LoginPage from "../pages/auth/login";
import PasswordUpdatePage from "../pages/auth/passwordUpdate";
import SignoutAllPage from "../pages/auth/signoutAll";
import SignupPage from "../pages/auth/signup";
import BuyPage from "../pages/buy";
import FindAgentPage from "../pages/findAgent";
import HomePage from "../pages/home";
import ProfilePage from "../pages/profile";
import RentPage from "../pages/rent";
import PersonalDetails from "../pages/renter-profile/personalDetails";
import RenterProfile from "../pages/renterProfile";
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
          path: "/my-real-estate/account",
          element: (
            <RequireAuth>
              <AccountPage />
            </RequireAuth>
          ),
        },
        {
          path: "/me",
          element: (
            <RequireAuth>
              <ProfilePage />
            </RequireAuth>
          ),
        },
        {
          path: "rent",
          element: <RentPage />,
        },
        {
          path: "rent/renter-profile",
          element: (
            <RequireAuth>
              <RenterProfile />
            </RequireAuth>
          ),
        },
        {
          path: "rent/renter-profile/personal-details",
          element: <PersonalDetails />,
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
      element: (
        <RequireAuth>
          <PasswordUpdatePage />
        </RequireAuth>
      ),
    },
    {
      path: "/email-update",
      element: (
        <RequireAuth>
          <EmailUpdatePage />
        </RequireAuth>
      ),
    },
    {
      path: "/signout-all",
      element: (
        <RequireAuth>
          <SignoutAllPage />
        </RequireAuth>
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

const RequireAuth = (props: { children: JSX.Element }) => {
  const { user } = useContext(UserContext);

  if (!user) {
    return <Navigate to="/login" />;
  } else {
    return props.children;
  }
};
