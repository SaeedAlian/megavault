import { forwardRef, InputHTMLAttributes } from "react";
import { IconType } from "react-icons";
import { cva, type VariantProps } from "class-variance-authority";

import { cn } from "../../lib/utils";

export const inputLabelVariants = cva(
  "flex ml-2 mb-2 font-medium items-center gap-2",
  {
    variants: {
      labelColor: {
        dark: "text-background",
        light: "text-foreground",
      },
      fontSize: {
        default: "text-base [&_svg]:size-5",
        sm: "text-sm [&_svg]:size-4",
        lg: "text-lg [&_svg]:size-6",
      },
    },
    defaultVariants: {
      labelColor: "light",
      fontSize: "default",
    },
  },
);

export const inputVariants = cva(
  "inline-flex disabled:opacity-50 font-medium disabled:pointer-events-none transition-colors outline-none",
  {
    variants: {
      radius: {
        normal: "rounded",
        full: "rounded-full",
      },
      variant: {
        // Contained
        "contained-primary":
          "bg-primary text-primary-foreground hover:bg-primary/90 focus:bg-primary/80 placeholder:text-primary-foreground/50",
        "contained-secondary":
          "bg-secondary text-secondary-foreground hover:bg-secondary/90 focus:bg-secondary/80 placeholder:text-secondary-foreground/60",
        "contained-accent":
          "bg-accent text-accent-foreground hover:bg-accent/90 focus:bg-accent/80 placeholder:text-accent-foreground/50",
        "contained-muted":
          "bg-muted text-muted-foreground hover:bg-muted/90 focus:bg-muted/80 placeholder:text-muted-foreground/50",
        "contained-card":
          "bg-card text-card-foreground hover:bg-card/90 focus:bg-card/80 placeholder:text-card-foreground/50",

        // Outlined
        "outlined-primary":
          "bg-transparent border border-primary text-foreground hover:bg-primary/10 focus:bg-primary placeholder:text-foreground/50 focus:text-primary-foreground",
        "outlined-secondary":
          "bg-transparent border border-secondary text-foreground hover:bg-secondary/10 focus:bg-secondary placeholder:text-foreground/50 focus:text-secondary-foreground",
        "outlined-accent":
          "bg-transparent border border-accent text-foreground hover:bg-accent/10 focus:bg-accent placeholder:text-foreground/50 focus:text-accent-foreground",
        "outlined-muted":
          "bg-transparent border border-white/70 text-foreground hover:bg-muted/10 focus:bg-muted placeholder:text-foreground/50 focus:text-muted-foreground",
        "outlined-card":
          "bg-transparent border border-primary/50 text-foreground hover:bg-card/10 focus:bg-card placeholder:text-foreground/50 focus:text-card-foreground",
      },
      fontSize: {
        default: "px-4 py-2 text-base",
        sm: "px-3 py-1 text-sm",
        lg: "px-5 py-3 text-lg",
      },
    },
    defaultVariants: {
      variant: "contained-primary",
      fontSize: "default",
      radius: "normal",
    },
  },
);

export interface InputProps
  extends InputHTMLAttributes<HTMLInputElement>,
  VariantProps<typeof inputVariants>,
  VariantProps<typeof inputLabelVariants>,
  VariantProps<typeof inputLabelVariants> {
  label?: string;
  Icon?: IconType;
  containerClassName?: string;
  labelClassName?: string;
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  (
    {
      className,
      labelClassName,
      containerClassName,
      type,
      fontSize,
      variant,
      radius,
      label,
      labelColor,
      Icon,
      ...props
    },
    ref,
  ) => {
    return (
      <div className={cn("flex flex-col", containerClassName)}>
        <label
          className={cn(
            inputLabelVariants({
              fontSize,
              labelColor,
              className: labelClassName,
            }),
          )}
        >
          {Icon != null && <Icon />}
          <span>{label}</span>
        </label>
        <input
          type={type}
          className={cn(
            inputVariants({ className, fontSize, variant, radius }),
          )}
          ref={ref}
          {...props}
        />
      </div>
    );
  },
);
Input.displayName = "Input";

export default Input;
