import { createBrowserRouter, RouterProvider } from "react-router-dom";
import ForgotPasswordPage from "./pages/forgotPassword";
import HomePage from "./pages/home";
import LoginPage from "./pages/login";
import RentPage from "./pages/rentPage";
import SignupPage from "./pages/signup";

const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
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
    path: "rent",
    element: <RentPage />,
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;