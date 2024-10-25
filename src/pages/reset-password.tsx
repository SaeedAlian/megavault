import { FormEventHandler, useEffect, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { MdOutlinePassword } from "react-icons/md";
import { RiLockPasswordFill } from "react-icons/ri";

import { Button, Input } from "../components/ui";

import logo from "../assets/images/logo.webp";

type FormData = {
  password: string;
  confirmPassword: string;
};

function ResetPassword() {
  const [formData, setFormData] = useState<FormData>({} as FormData);

  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const [token, setToken] = useState<string>("");

  const onChangeData = (name: string, value: string | number | boolean) => {
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const onSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    console.log(formData);
  };

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
    <div className="min-h-screen w-screen bg-home-background flex flex-col items-center justify-center py-16 max-md:py-0">
      <div className="flex flex-col bg-popover items-center px-6 pt-12 pb-6 w-full max-w-[600px] relative rounded-lg max-md:max-w-full max-md:min-h-screen">
        <form
          onSubmit={onSubmit}
          className="max-w-[460px] w-full flex flex-col items-center"
        >
          <img
            className="w-20 absolute -top-10 left-[50%] translate-x-[-50%] max-md:static max-md:translate-x-0 max-md:mb-3"
            src={logo}
            alt="MegaVault"
          />
          <h1 className="text-center text-lg font-bold mb-2">Reset Password</h1>
          <p className="mb-12 text-center text-sm">
            Enter your new password down below.
          </p>

          <Input
            required
            autoFocus
            label="Password"
            type="password"
            variant="contained-primary"
            placeholder="***********"
            containerClassName="mb-6 w-full"
            name="password"
            Icon={RiLockPasswordFill}
            value={formData.password}
            onChange={(e) => {
              onChangeData(e.target.name, e.target.value);
            }}
          />

          <Input
            required
            label="Confirm Password"
            type="password"
            variant="contained-primary"
            placeholder="***********"
            containerClassName="mb-12 w-full"
            name="confirmPassword"
            Icon={MdOutlinePassword}
            value={formData.confirmPassword}
            onChange={(e) => {
              onChangeData(e.target.name, e.target.value);
            }}
          />

          <Button
            variant="contained-accent"
            className="w-full max-w-36 self-center mb-12"
            type="submit"
          >
            Submit
          </Button>
        </form>
      </div>
    </div>
  );
}

export default ResetPassword;
