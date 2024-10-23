const { fontFamily } = require("tailwindcss/defaultTheme");

const colors = {
  background: "hsl(var(--background))",
  foreground: "hsl(var(--foreground))",
  primary: {
    light: "hsl(var(--primary-light))",
    dark: "hsl(var(--primary-dark))",
    DEFAULT: "hsl(var(--primary))",
    foreground: "hsl(var(--primary-foreground))",
    "foreground-light": "hsl(var(--primary-foreground-light))",
    "foreground-dark": "hsl(var(--primary-foreground-dark))",
  },
  secondary: {
    light: "hsl(var(--secondary-light))",
    dark: "hsl(var(--secondary-dark))",
    DEFAULT: "hsl(var(--secondary))",
    foreground: "hsl(var(--secondary-foreground))",
    "foreground-light": "hsl(var(--secondary-foreground-light))",
    "foreground-dark": "hsl(var(--secondary-foreground-dark))",
  },
  destructive: {
    DEFAULT: "hsl(var(--destructive))",
    foreground: "hsl(var(--destructive-foreground))",
  },
  muted: {
    DEFAULT: "hsl(var(--muted))",
    foreground: "hsl(var(--muted-foreground))",
  },
  accent: {
    light: "hsl(var(--accent-light))",
    dark: "hsl(var(--accent-dark))",
    DEFAULT: "hsl(var(--accent))",
    foreground: "hsl(var(--accent-foreground))",
    "foreground-light": "hsl(var(--accent-foreground-light))",
    "foreground-dark": "hsl(var(--accent-foreground-dark))",
  },
  popover: {
    DEFAULT: "hsl(var(--popover))",
    foreground: "hsl(var(--popover-foreground))",
  },
  card: {
    DEFAULT: "hsl(var(--card))",
    foreground: "hsl(var(--card-foreground))",
  },
};

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "./src/components/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    container: {
      center: "true",
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      borderColor: {
        ...colors,
      },
      backgroundImage: {
        "home-background": "linear-gradient(116.82deg, #010A13, #021B34)",
      },
      colors: {
        ...colors,
      },
      fontFamily: {
        poppins: ["Poppins", ...fontFamily.sans],
        sans: ["var(--font-sans)", ...fontFamily.sans],
      },
    },
  },
  plugins: [require("tailwindcss-animate")],
};
