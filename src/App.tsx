import { createBrowserRouter, RouterProvider } from "react-router-dom";

import Home from "./pages/home";
import Login from "./pages/login";
import Register from "./pages/register";
import RegisterSuccess from "./pages/register-success";
import AccountVerificationSuccess from "./pages/account-verification-success";
import ForgotPassword from "./pages/forgot-password";
import ResetPassword from "./pages/reset-password";
import ResetPasswordSuccess from "./pages/reset-password-success";

import "./App.css";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/register",
    element: <Register />,
  },
  {
    path: "/register-success",
    element: <RegisterSuccess />,
  },
  {
    path: "/account-verification-success",
    element: <AccountVerificationSuccess />,
  },
  {
    path: "/forgot-password",
    element: <ForgotPassword />,
  },
  {
    path: "/reset-password",
    element: <ResetPassword />,
  },
  {
    path: "/reset-password-success",
    element: <ResetPasswordSuccess />,
  },
]);

function App() {
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
