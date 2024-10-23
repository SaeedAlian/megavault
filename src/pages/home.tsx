import { useEffect, useState } from "react";
import { MdOutlineClose } from "react-icons/md";
import { GiHamburgerMenu } from "react-icons/gi";
import { FaInstagram, FaTwitter, FaGithub } from "react-icons/fa";

import logo from "../assets/images/logo.webp";
import detailsImg from "../assets/images/home-details-image.webp";
import homeIllustration from "../assets/images/home-illustration.png";
import speedIcon from "../assets/svgs/speed-icon.svg";
import secureDataIcon from "../assets/svgs/secure-data-icon.svg";
import securityIcon from "../assets/svgs/security-icon.svg";

import { BlogCard, BlogPagination } from "../components";
import { Button } from "../components/ui";

import { cn } from "../lib/utils";

import tempBlogCard from "../assets/images/temp-blog-image.jpg";

type NavbarLink = { href: string; title: string };

type Blog = {
  title: string;
  description: string;
  date: Date;
  link: string;
  image: string;
};

const navbarLinks: NavbarLink[] = [
  {
    title: "Login",
    href: "/login",
  },
  {
    title: "Register",
    href: "/register",
  },
  {
    title: "Blog",
    href: "#blog",
  },
];

const tempBlogs = [
  {
    title: "Lorem ipsum dolor sit amet1",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore Lorem ipsum dolor tur.",
    date: new Date(),
    link: "/blog/1",
    image: tempBlogCard,
  },
  {
    title: "Lorem ipsum dolor sit amet2",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore Lorem ipsum dolor tur.",
    date: new Date(),
    link: "/blog/1",
    image: tempBlogCard,
  },
  {
    title: "Lorem ipsum dolor sit amet3",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore Lorem ipsum dolor tur.",
    date: new Date(),
    link: "/blog/1",
    image: tempBlogCard,
  },
  {
    title: "Lorem ipsum dolor sit amet4",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore Lorem ipsum dolor tur.",
    date: new Date(),
    link: "/blog/1",
    image: tempBlogCard,
  },
  {
    title: "Lorem ipsum dolor sit amet5",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore Lorem ipsum dolor tur.",
    date: new Date(),
    link: "/blog/1",
    image: tempBlogCard,
  },
];

function Home() {
  const [isSidebarMenuOpen, setIsSidebarMenuOpen] = useState<boolean>(false);
  const [blogs, setBlogs] = useState<Blog[]>([]);
  const [displayBlogs, setDisplayBlogs] = useState<Blog[]>([]);
  const [totalBlogPage, setTotalBlogPage] = useState<number>(0);
  const [currentBlogPage, setCurrentBlogPage] = useState<number>(1);
  const [blogDisplayNumber, setBlogDisplayNumber] = useState<number>(3);

  const [prevScrollY, setPrevScrollY] = useState<number>(0);
  const [currentScrollY, setCurrentScrollY] = useState<number>(0);
  const [isScrollingDown, setIsScrollingDown] = useState<boolean>(false);

  const onOpenSidebar = () => {
    setIsSidebarMenuOpen(true);
  };

  const onCloseSidebar = () => {
    setIsSidebarMenuOpen(false);
  };

  const onChangeBlogPage = (page: number) => {
    if (page <= 0 || page > totalBlogPage) return;
    setCurrentBlogPage(page);
  };

  const handleSetDisplayBlogs = (blogs: Blog[], page: number) => {
    const startIndex = (page - 1) * blogDisplayNumber;
    const endIndex = page * blogDisplayNumber;
    const newDisplayBlogs = blogs.slice(startIndex, endIndex);

    setDisplayBlogs(newDisplayBlogs);
  };

  const handleSetBlogDisplayNumber = () => {
    const width = window.innerWidth;

    if (width <= 500) {
      setBlogDisplayNumber(1);
    } else if (width > 500 && width < 1024) {
      setBlogDisplayNumber(2);
    } else {
      setBlogDisplayNumber(3);
    }
  };

  const handleSetScrollDownState = () => {
    const scrollY = window.scrollY;
    setCurrentScrollY(scrollY);
  };

  useEffect(() => {
    if (window == undefined) return;
    if (window == null) return;

    handleSetBlogDisplayNumber();

    setPrevScrollY(() => window.scrollY);
    setCurrentScrollY(() => window.scrollY);

    window.addEventListener("resize", handleSetBlogDisplayNumber);
    window.addEventListener("scroll", handleSetScrollDownState);

    return () => {
      window.removeEventListener("resize", handleSetBlogDisplayNumber);
      window.removeEventListener("scroll", handleSetScrollDownState);
    };
  }, []);

  useEffect(() => {
    setBlogs(tempBlogs);
  }, []);

  useEffect(() => {
    if (prevScrollY < currentScrollY) {
      setIsScrollingDown(true);
    } else {
      setIsScrollingDown(false);
    }

    setPrevScrollY(currentScrollY);
  }, [currentScrollY]);

  useEffect(() => {
    if (currentBlogPage <= 0 || currentBlogPage > totalBlogPage) return;
    handleSetDisplayBlogs(blogs, currentBlogPage);
  }, [blogs, currentBlogPage, totalBlogPage, blogDisplayNumber]);

  useEffect(() => {
    if (blogs.length === 0) return;

    setTotalBlogPage(parseInt((blogs.length / blogDisplayNumber).toFixed()));
    setCurrentBlogPage(1);
  }, [blogs, blogDisplayNumber]);

  return (
    <div className="min-h-screen w-screen bg-home-background flex flex-col items-center">
      <div className="w-full gap-y-24 flex flex-col">
        {/* Navbar */}
        <header
          className={cn(
            "w-full fixed duration-300 top-0 left-0 right-0 px-12 transition-all py-6 flex justify-center max-md:px-8 max-sm:px-6 z-[5000]",
            isScrollingDown && !isSidebarMenuOpen
              ? "translate-y-[-100%]"
              : "translate-y-[0%]",
            currentScrollY > 0 ? "backdrop-blur-lg bg-black/30" : "",
          )}
        >
          <nav className="w-full flex items-center gap-4 max-w-[1440px]">
            <img className="w-12" src={logo} alt="MegaVault" />
            <h1 className="font-bold text-lg">MegaVault</h1>
            <ul className="ml-auto flex items-center gap-12 max-md:hidden">
              {navbarLinks.map((l) => (
                <li key={l.title}>
                  <Button asLink variant="link" linkHref={l.href}>
                    {l.title}
                  </Button>
                </li>
              ))}
            </ul>

            <Button
              onClick={onOpenSidebar}
              variant="link"
              size="icon"
              className="hidden ml-auto max-md:inline-flex"
            >
              <GiHamburgerMenu />
            </Button>

            <div
              className={cn(
                "hidden flex-col items-center transition-transform py-5 px-4 duration-200 ease-linear z-[5000] right-0 top-0 h-screen bg-popover fixed w-screen max-md:flex",
                isSidebarMenuOpen ? "translate-x-[0%]" : "translate-x-[100%]",
              )}
            >
              <img className="w-12" src={logo} alt="MegaVault" />
              <h1 className="font-bold text-center text-lg">MegaVault</h1>

              <Button
                variant="link"
                size="icon"
                className="absolute top-7 right-6"
                onClick={onCloseSidebar}
              >
                <MdOutlineClose />
              </Button>

              <ul className="flex flex-col items-center mt-6 gap-4">
                {navbarLinks.map((l) => (
                  <li key={l.title}>
                    <Button asLink variant="link" linkHref={l.href}>
                      {l.title}
                    </Button>
                  </li>
                ))}
              </ul>
              <span className="font-normal text-sm text-foreground/50 mt-auto text-center">
                Copyright &#169; MegaVault 2024
              </span>
            </div>
          </nav>
        </header>

        {/* Main Section */}
        <section className="flex w-full justify-center">
          <div className="w-full max-w-[1440px] flex flex-row-reverse gap-x-20 gap-y-16 items-center justify-between max-md:flex-col px-12 py-6 pt-32 max-md:px-8 max-sm:px-6">
            <div>
              <img
                src={homeIllustration}
                alt="secure cloud"
                className="w-full max-w-[445px] max-lg:max-w-[400px] flex-1 max-md:max-w-[350px]"
              />
            </div>
            <div className="flex-1 max-w-[700px]">
              <h2 className="font-extrabold text-3xl tracking-wide">
                Your Data, Your Way. Securely Stored. Easily Accessed.
              </h2>
              <p className="font-normal text-sm text-foreground/90 mt-5">
                MegaVault provides you with a secure, fast, and reliable cloud
                storage solution that allows you to access your data anytime,
                anywhere. Whether you need to back up your files, share
                documents with colleagues, or simply have peace of mind knowing
                your data is safe, MegaVault is the perfect solution for you.
                Our advanced security features, lightning-fast speeds, and
                reliable infrastructure ensure that your data is always
                protected and readily available. With MegaVault, you can focus
                on what matters most - your work, your life, and your peace of
                mind.
              </p>
              <div className="mt-14 flex items-center gap-4 max-sm:flex-col">
                <Button
                  radius="full"
                  asLink
                  linkHref="/"
                  className="max-sm:w-full"
                >
                  Try it out for free!
                </Button>
                <span>Or</span>
                <Button
                  variant="outlined-accent"
                  radius="full"
                  asLink
                  linkHref="/login"
                  className="max-sm:w-full"
                >
                  Log In
                </Button>
              </div>
            </div>
          </div>
        </section>
        {/* Details Section */}
        <section className="flex justify-center bg-popover w-full">
          <div className="w-full max-w-[1440px] flex gap-x-20 gap-y-16 items-center justify-between my-8 max-md:flex-col px-12 py-6 max-md:px-8 max-sm:px-6">
            <div className="z-[1]">
              <img
                className="w-full max-w-[545px] max-xl:max-w-[445px] max-lg:max-w-[400px] flex-1 max-md:max-w-[350px]"
                src={detailsImg}
                alt="secure, fast and simple"
              />
            </div>

            <div className="flex flex-col flex-1 max-w-[700px] z-[1] max-lg:max-w-full">
              <h2 className="text-right font-bold text-2xl tracking-wide">
                Experience the Power of Encryption. Enjoy the Freedom of Free
                Storage.
              </h2>
              <p className="text-right font-normal text-sm mt-5">
                Whether it's important documents, precious photos, or
                irreplaceable memories, MegaVault keeps your data safe from
                prying eyes and unauthorized access. We employ industry-leading
                encryption protocols to ensure the highest level of security,
                giving you peace of mind knowing your data is protected. And the
                best part? MegaVault offers a generous amount of free storage,
                so you can start protecting your data without any upfront costs.
                Experience the power of secure, reliable cloud storage without
                breaking the bank.
              </p>
              <div className="flex items-center gap-7 text-center justify-center w-fit self-end mt-9">
                <div className="flex flex-col items-center">
                  <img src={securityIcon} className="w-9 mb-1" />
                  <span className="text-sm">Secure & Reliable</span>
                </div>
                <div className="flex flex-col items-center">
                  <img src={speedIcon} className="w-9 mb-1" />
                  <span className="text-sm">Fast & Smooth</span>
                </div>
                <div className="flex flex-col items-center">
                  <img src={secureDataIcon} className="w-9 mb-1" />
                  <span className="text-sm">Safe Data Encryption</span>
                </div>
              </div>
            </div>
          </div>
        </section>
        {/* Blog Section */}
        <section id="blog" className="flex w-full justify-center">
          <div className="w-full max-w-[1440px] flex items-center flex-col px-12 py-6 max-md:px-8 max-sm:px-6">
            <h3 className="text-5xl font-extrabold">Blog</h3>

            <div className="max-w-[1000px] grid grid-cols-3 items-center justify-center gap-x-10 gap-y-7 mt-16 max-lg:grid-cols-2 max-[500px]:grid-cols-1">
              {displayBlogs.map((b) => (
                <BlogCard
                  key={b.title}
                  title={b.title}
                  description={b.description}
                  date={b.date}
                  link={b.link}
                  image={b.image}
                />
              ))}
            </div>

            <div className="mt-12">
              <BlogPagination
                defaultPage={currentBlogPage}
                totalPages={totalBlogPage}
                onChangePage={onChangeBlogPage}
              />
            </div>
          </div>
        </section>
        {/* Footer */}
        <footer className="justify-center flex w-full bg-popover">
          <div className="w-full gap-x-20 gap-y-16 max-w-[1440px] flex justify-between items-center flex-col px-12 py-6 max-md:px-8 max-sm:px-6">
            <div className="z-[1] flex flex-col items-center">
              <img src={logo} alt="MegaVault" className="w-16" />
              <h4 className="font-bold text-center text-xl mt-4">MegaVault</h4>
              <p className="text-sm text-foreground/90 text-center font-normal mt-4">
                Your Digital Fortress, Made Easy.
              </p>
              <ul className="flex items-center gap-5 mt-4">
                <li>
                  <Button
                    asLink
                    linkHref="#"
                    linkTarget="_blank"
                    variant="link"
                    size="icon"
                  >
                    <FaInstagram />
                  </Button>
                </li>
                <li>
                  <Button
                    asLink
                    linkHref="#"
                    linkTarget="_blank"
                    variant="link"
                    size="icon"
                  >
                    <FaGithub />
                  </Button>
                </li>
                <li>
                  <Button
                    asLink
                    linkHref="#"
                    linkTarget="_blank"
                    variant="link"
                    size="icon"
                  >
                    <FaTwitter />
                  </Button>
                </li>
              </ul>
              <ul className="flex items-center gap-10 max-sm:gap-6 flex-wrap justify-center mt-10">
                {navbarLinks.map((l) => (
                  <li key={l.title}>
                    <Button asLink variant="link-primary" linkHref={l.href}>
                      {l.title}
                    </Button>
                  </li>
                ))}
              </ul>

              <span className="font-normal text-sm text-foreground/50 mt-16 text-center">
                Copyright &#169; MegaVault 2024
              </span>
            </div>
          </div>
        </footer>
      </div>
    </div>
  );
}

export default Home;
