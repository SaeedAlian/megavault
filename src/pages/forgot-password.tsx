import { FormEventHandler, useState } from "react";
import { FaUser } from "react-icons/fa";

import logo from "../assets/images/logo.webp";

import { Button, Input } from "../components/ui";

type FormData = {
  username: string;
};

function ForgotPassword() {
  const [formData, setFormData] = useState<FormData>({} as FormData);

  const onChangeData = (name: string, value: string | number | boolean) => {
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const onSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    console.log(formData);
  };

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
            If you forget your password and you cannot login into your account,
            please provide your username or email. We will send you an email
            containing the instructions to reset your password.
          </p>

          <Input
            required
            autoFocus
            label="Enter your username (or email)"
            type="text"
            variant="contained-primary"
            placeholder="JohnDoe (or JohnDoe@email.com)"
            Icon={FaUser}
            containerClassName="mb-16 w-full"
            name="username"
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

export default ForgotPassword;
