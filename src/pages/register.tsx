import { FormEventHandler, useState } from "react";
import { Link } from "react-router-dom";
import { FaUser } from "react-icons/fa";
import { MdEmail, MdOutlinePassword } from "react-icons/md";
import { RiLockPasswordFill } from "react-icons/ri";
import { IoIosArrowBack } from "react-icons/io";

import { Button, Checkbox, Input } from "../components/ui";

import logo from "../assets/images/logo.webp";

type FormData = {
  firstname: string;
  lastname: string;
  username: string;
  email: string;
  password: string;
  confirmPassword: string;
  serviceCheck: boolean;
};

const MAX_PAGE = 3;

function Register() {
  const [page, setPage] = useState<number>(3);
  const [formData, setFormData] = useState<FormData>({} as FormData);

  const onChangePage = (page: number) => {
    if (page < 1 && page > MAX_PAGE) return;
    setPage(page);
  };

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
        {page === 1 ? (
          <FirstPage
            maxPage={MAX_PAGE}
            formData={formData}
            onSubmit={() => onChangePage(2)}
            onChangeData={onChangeData}
          />
        ) : page === 2 ? (
          <SecondPage
            maxPage={MAX_PAGE}
            formData={formData}
            onSubmit={() => onChangePage(3)}
            onGoBack={() => onChangePage(1)}
            onChangeData={onChangeData}
          />
        ) : page === 3 ? (
          <ThirdPage
            maxPage={MAX_PAGE}
            formData={formData}
            onSubmit={onSubmit}
            onChangeData={onChangeData}
            onGoBack={() => onChangePage(2)}
          />
        ) : null}
      </div>
    </div>
  );
}

function FirstPage({
  maxPage,
  formData,
  onChangeData,
  onSubmit,
}: {
  formData: FormData;
  maxPage: number;
  onChangeData: (name: string, value: string | boolean | number) => void;
  onSubmit: FormEventHandler<HTMLFormElement>;
}) {
  return (
    <form
      onSubmit={onSubmit}
      className="max-w-[460px] w-full flex flex-col items-center"
    >
      <img
        className="w-20 absolute -top-10 left-[50%] translate-x-[-50%] max-md:static max-md:translate-x-0 max-md:mb-3"
        src={logo}
        alt="MegaVault"
      />
      <h1 className="text-center text-lg font-bold mb-2">Register</h1>
      <p className="text-center text-xs font-normal mb-12">Page 1/{maxPage}</p>

      <Input
        required
        label="First Name"
        type="text"
        variant="contained-primary"
        placeholder="John"
        containerClassName="mb-6 w-full"
        name="firstname"
        value={formData.firstname}
        onChange={(e) => {
          onChangeData(e.target.name, e.target.value);
        }}
      />

      <Input
        required
        label="Last Name"
        type="text"
        variant="contained-primary"
        placeholder="Doe"
        containerClassName="mb-8 w-full"
        name="lastname"
        value={formData.lastname}
        onChange={(e) => {
          onChangeData(e.target.name, e.target.value);
        }}
      />

      <Button
        variant="contained-accent"
        className="w-full max-w-36 self-center mb-12"
        type="submit"
      >
        Next
      </Button>

      <span className="flex flex-col items-center text-center">
        Already have an account?{" "}
        <Button variant="link-primary" asLink type="button" linkHref="/login">
          Login NOW!
        </Button>
      </span>
    </form>
  );
}

function SecondPage({
  maxPage,
  formData,
  onChangeData,
  onGoBack,
  onSubmit,
}: {
  formData: FormData;
  maxPage: number;
  onChangeData: (name: string, value: string | boolean | number) => void;
  onGoBack: () => void;
  onSubmit: FormEventHandler<HTMLFormElement>;
}) {
  return (
    <form
      onSubmit={onSubmit}
      className="max-w-[460px] w-full flex flex-col items-center"
    >
      <Button
        type="button"
        size="icon"
        variant="link-accent"
        onClick={onGoBack}
        className="self-start"
      >
        <IoIosArrowBack />
      </Button>
      <img
        className="w-20 absolute -top-10 left-[50%] translate-x-[-50%] max-md:static max-md:translate-x-0 max-md:mb-3"
        src={logo}
        alt="MegaVault"
      />
      <h1 className="text-center text-lg font-bold mb-2">Register</h1>
      <p className="text-center text-xs font-normal mb-12">Page 2/{maxPage}</p>

      <Input
        required
        label="Email"
        type="email"
        variant="contained-primary"
        placeholder="JohnDoe@email.com"
        containerClassName="mb-6 w-full"
        name="email"
        Icon={MdEmail}
        value={formData.email}
        onChange={(e) => {
          onChangeData(e.target.name, e.target.value);
        }}
      />

      <Input
        required
        label="Username"
        type="text"
        variant="contained-primary"
        placeholder="JohnDoe"
        containerClassName="mb-8 w-full"
        name="username"
        Icon={FaUser}
        value={formData.username}
        onChange={(e) => {
          onChangeData(e.target.name, e.target.value);
        }}
      />

      <Button
        variant="contained-accent"
        className="w-full max-w-36 self-center mb-12"
        type="submit"
      >
        Next
      </Button>
    </form>
  );
}

function ThirdPage({
  maxPage,
  formData,
  onChangeData,
  onGoBack,
  onSubmit,
}: {
  formData: FormData;
  maxPage: number;
  onChangeData: (name: string, value: string | boolean | number) => void;
  onGoBack: () => void;
  onSubmit: FormEventHandler<HTMLFormElement>;
}) {
  return (
    <form
      onSubmit={onSubmit}
      className="relative max-w-[460px] w-full flex flex-col items-center"
    >
      <Button
        type="button"
        size="icon"
        variant="link-accent"
        onClick={onGoBack}
        className="self-start"
      >
        <IoIosArrowBack />
      </Button>
      <img
        className="w-20 absolute -top-10 left-[50%] translate-x-[-50%] max-md:static max-md:translate-x-0 max-md:mb-3"
        src={logo}
        alt="MegaVault"
      />
      <h1 className="text-center text-lg font-bold mb-2">Register</h1>
      <p className="text-center text-xs font-normal mb-12">Page 3/{maxPage}</p>

      <Input
        required
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
        containerClassName="mb-3 w-full"
        name="confirmPassword"
        Icon={MdOutlinePassword}
        value={formData.confirmPassword}
        onChange={(e) => {
          onChangeData(e.target.name, e.target.value);
        }}
      />
      <Checkbox
        radius="full"
        variant="accent"
        size="sm"
        containerClassName="mb-12"
        id="serviceCheck"
        name="serviceCheck"
        checked={formData.serviceCheck}
        onCheckedChange={(c) => {
          onChangeData("serviceCheck", c);
        }}
        label={
          <>
            I agree to the{" "}
            <Link
              to="/privacy-policy"
              target="_blank"
              className="text-accent p-0"
            >
              privacy & policy
            </Link>{" "}
            and{" "}
            <Link
              to="/terms-of-service"
              target="_blank"
              className="text-accent p-0"
            >
              terms of service
            </Link>{" "}
            of <span className="text-primary">MegaVault</span>
          </>
        }
      />

      <Button
        variant="contained-accent"
        className="w-full max-w-36 self-center mb-12"
        type="submit"
      >
        Submit
      </Button>
    </form>
  );
}

export default Register;
