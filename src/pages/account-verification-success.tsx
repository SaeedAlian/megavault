import { useEffect, useState } from "react";
import { useSearchParams } from "react-router-dom";

import { Button } from "../components/ui";

import icon from "../assets/svgs/register-success-icon.svg";

function AccountVerificationSuccess() {
  const [searchParams] = useSearchParams();
  const [token, setToken] = useState<string>("");

  useEffect(() => {
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
          Your account has verified successfully!
        </h1>
        <p className="text-center font-normal text-lg mb-16">
          Now you can use all of the features of MegaVault without any worry in
          the world.
        </p>
        <Button asLink variant="contained-accent" size="lg" linkHref="/login">
          Continue
        </Button>
      </div>
    </div>
  );
}

export default AccountVerificationSuccess;
