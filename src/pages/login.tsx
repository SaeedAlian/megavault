import { FormEventHandler, useState } from "react";
import { FaUser } from "react-icons/fa";
import { RiLockPasswordFill } from "react-icons/ri";

import logo from "../assets/images/logo.webp";

import { Button, Input } from "../components/ui";

type FormData = {
  usernameOrEmail: string;
  password: string;
};

function Login() {
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
          <h1 className="text-center text-lg font-bold mb-12">Login</h1>

          <Input
            required
            autoFocus
            label="Username (or email)"
            type="text"
            variant="contained-primary"
            placeholder="JohnDoe"
            Icon={FaUser}
            containerClassName="mb-6 w-full"
            name="usernameOrEmail"
            onChange={(e) => {
              onChangeData(e.target.name, e.target.value);
            }}
          />

          <Input
            required
            label="Password"
            type="password"
            variant="contained-primary"
            placeholder="*********"
            Icon={RiLockPasswordFill}
            containerClassName="mb-1 w-full"
            name="password"
            onChange={(e) => {
              onChangeData(e.target.name, e.target.value);
            }}
          />
          <Button
            variant="link-accent"
            size="sm"
            className="self-start mb-8"
            asLink
            type="button"
            linkHref="/forgot-password"
          >
            I've forgot my password
          </Button>

          <Button
            variant="contained-accent"
            className="w-full max-w-36 self-center mb-12"
            type="submit"
          >
            Submit
          </Button>

          <span className="flex flex-col items-center text-center">
            Don't have an account?{" "}
            <Button
              variant="link-primary"
              asLink
              type="button"
              linkHref="/register"
            >
              Register NOW!
            </Button>
          </span>
        </form>
      </div>
    </div>
  );
}

export default Login;
