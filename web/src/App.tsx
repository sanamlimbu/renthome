import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { UserContextProvider } from "./context/user";
import BuyPage from "./pages/buy";
import FacebookAuthRedirectPage from "./pages/facebookAuth";
import FindAgentPage from "./pages/findAgent";
import ForgotPasswordPage from "./pages/forgotPassword";
import GoogleAuthRedirectPage from "./pages/googleAuth";
import HomePage from "./pages/home";
import LoginPage from "./pages/login";
import { MePage } from "./pages/me";
import RentPage from "./pages/rent";
import RootPage from "./pages/root";
import SignupPage from "./pages/signup";
import SoldPage from "./pages/sold";

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
        element: <MePage />,
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
    element: <LoginPage />,
  },
  {
    path: "/signup",
    element: <SignupPage />,
  },
  {
    path: "forgot-password",
    element: <ForgotPasswordPage />,
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

function App() {
  return (
    <UserContextProvider>
      <RouterProvider router={router} />
    </UserContextProvider>
  );
}

export default App;
