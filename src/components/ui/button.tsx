import { forwardRef } from "react";
import { Link } from "react-router-dom";
import { cva, type VariantProps } from "class-variance-authority";

import { cn } from "../../lib/utils";

export const buttonVariants = cva(
  "relative inline-flex items-center justify-center gap-2 whitespace-nowrap text-sm font-medium transition-colors outline-none disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0",
  {
    variants: {
      radius: {
        normal: "rounded-lg",
        full: "rounded-full",
      },
      variant: {
        // Contained
        "contained-primary":
          "bg-primary text-primary-foreground hover:bg-primary/80",
        "contained-secondary":
          "bg-secondary text-secondary-foreground hover:bg-secondary/80",
        "contained-accent":
          "bg-accent text-accent-foreground hover:bg-accent/80",

        // Outlined
        "outlined-primary":
          "border border-primary text-foreground hover:bg-primary/80 hover:text-primary-foreground",
        "outlined-secondary":
          "border border-secondary text-foreground hover:bg-secondary/80 hover:text-secondary-foreground",
        "outlined-accent":
          "border border-accent text-foreground hover:bg-accent/80 hover:text-accent-foreground",

        // Ghost
        "ghost-primary":
          "text-primary hover:bg-primary/90 hover:text-primary-foreground",
        "ghost-secondary":
          "text-secondary hover:bg-secondary/90 hover:text-secondary-foreground",
        "ghost-accent":
          "text-accent hover:bg-accent/90 hover:text-accent-foreground",

        // Link
        "link-primary": "text-primary hover:underline",
        "link-secondary": "text-secondary hover:underline",
        "link-accent": "text-accent hover:underline",

        // Others
        muted: "bg-muted text-muted-foreground hover:bg-muted/80",
        disabled: "bg-muted/60 text-muted-foreground/50",
      },
      size: {
        default: "h-8 px-4 py-2 text-sm",
        sm: "h-6 px-2 py-1 text-xs",
        lg: "h-9 px-5 py-2 text-base",
        icon: "h-8 w-8 [&_svg]:size-5",
        "icon-sm": "h-6 w-6 [&_svg]:size-4",
        "icon-lg": "h-10 w-10 [&_svg]:size-6",
      },
    },
    defaultVariants: {
      variant: "contained-primary",
      size: "default",
      radius: "normal",
    },
  },
);

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
  VariantProps<typeof buttonVariants> {
  asLink?: boolean;
  linkTarget?: React.HTMLAttributeAnchorTarget;
  linkHref?: string;
}

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  (
    {
      className,
      asLink,
      linkTarget,
      linkHref,
      children,
      variant,
      size,
      radius,
      ...props
    },
    ref,
  ) => {
    const Comp = "button";
    return (
      <Comp
        className={cn(buttonVariants({ radius, variant, size, className }))}
        ref={ref}
        {...props}
      >
        {asLink ? (
          <>
            <Link
              className="top-0 left-0 right-0 bottom-0 absolute w-full h-full"
              to={linkHref ?? ""}
              target={linkTarget ? linkTarget : "_self"}
            ></Link>
            {children}
          </>
        ) : (
          children
        )}
      </Comp>
    );
  },
);
Button.displayName = "Button";

export default Button;
