import {
  ComponentPropsWithoutRef,
  ElementRef,
  forwardRef,
  ReactNode,
} from "react";
import { cva, VariantProps } from "class-variance-authority";
import { FaCheck } from "react-icons/fa";
import * as CheckboxPrimitive from "@radix-ui/react-checkbox";

import { cn } from "../../lib/utils";

export const checkboxVariants = cva(
  "peer shrink-0 border flex items-center justify-center ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
  {
    variants: {
      radius: {
        normal: "rounded-md",
        full: "rounded-full",
      },
      variant: {
        primary:
          "border-primary data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground",
        secondary:
          "border-secondary data-[state=checked]:bg-secondary data-[state=checked]:text-secondary-foreground",
        accent:
          "border-accent data-[state=checked]:bg-accent data-[state=checked]:text-accent-foreground",
        muted:
          "border-muted data-[state=checked]:bg-muted data-[state=checked]:text-muted-foreground",
      },
      size: {
        default: "w-1 h-1 p-[0.65rem] [&_svg]:size-4",
        sm: "w-1 h-1 p-[0.55rem] [&_svg]:size-3",
        lg: "w-1 h-1 p-[0.85rem] [&_svg]:size-5",
      },
    },
    defaultVariants: {
      variant: "primary",
      size: "default",
      radius: "normal",
    },
  },
);

export const checkboxLabelVariants = cva("font-medium", {
  variants: {
    labelColor: {
      dark: "text-background",
      light: "text-foreground",
    },
    size: {
      default: "text-base",
      sm: "text-sm",
      lg: "text-lg",
    },
  },
  defaultVariants: {
    labelColor: "light",
    size: "default",
  },
});

export interface CheckboxProps
  extends ComponentPropsWithoutRef<typeof CheckboxPrimitive.Root>,
  VariantProps<typeof checkboxVariants>,
  VariantProps<typeof checkboxLabelVariants> {
  label?: string | ReactNode;
  containerClassName?: string;
  labelClassName?: string;
}

const Checkbox = forwardRef<
  ElementRef<typeof CheckboxPrimitive.Root>,
  CheckboxProps
>(
  (
    {
      className,
      containerClassName,
      labelClassName,
      variant,
      labelColor,
      size,
      radius,
      label,
      id,
      ...props
    },
    ref,
  ) => (
    <div className={cn("flex flex-row-reverse gap-2", containerClassName)}>
      {label != null && (
        <label
          htmlFor={id}
          className={checkboxLabelVariants({
            labelColor,
            className: labelClassName,
            size,
          })}
        >
          {label}
        </label>
      )}
      <CheckboxPrimitive.Root
        ref={ref}
        className={checkboxVariants({
          className,
          variant,
          size,
          radius,
        })}
        id={id}
        {...props}
      >
        <CheckboxPrimitive.Indicator
          className={cn("flex items-center justify-center text-current")}
        >
          <FaCheck />
        </CheckboxPrimitive.Indicator>
      </CheckboxPrimitive.Root>
    </div>
  ),
);
Checkbox.displayName = CheckboxPrimitive.Root.displayName;

export default Checkbox;
