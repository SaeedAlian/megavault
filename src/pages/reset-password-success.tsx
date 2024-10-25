import { useEffect, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

import { Button } from "../components/ui";

import icon from "../assets/svgs/reset-password-success.svg";

function ResetPasswordSuccess() {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const [token, setToken] = useState<string>("");

  useEffect(() => {
    const token = searchParams.get("token");

    if (token == null || token.length === 0) {
      navigate("/");
    }

    setToken(searchParams.get("token") ?? "");
  }, []);

  useEffect(() => {
    console.log(token);
  }, [token]);

  return (
    <div className="min-h-screen w-screen bg-home-background flex flex-col items-center">
      <div className="flex flex-col items-center w-full justify-center px-4 pt-12 pb-24 min-h-screen max-w-[900px]">
        <img src={icon} alt="successful register" className="w-32 mb-6" />
        <h1 className="font-bold text-2xl text-center mb-4">
          Your password has been changed successfully.
        </h1>
        <p className="text-center font-normal text-lg mb-16">
          Now you can login into your account with the new password.
        </p>
        <Button asLink variant="contained-accent" size="lg" linkHref="/login">
          Back To Login Page
        </Button>
      </div>
    </div>
  );
}

export default ResetPasswordSuccess;
