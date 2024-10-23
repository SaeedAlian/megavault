import { forwardRef, useMemo } from "react";
import { FaCalendar } from "react-icons/fa";

import { Button } from "./ui";

export interface BlogCardProps {
  image: string;
  title: string;
  description: string;
  date: Date;
  link: string;
}

const BlogCard = forwardRef<HTMLDivElement, BlogCardProps>(
  ({ image, title, description, date, link }, ref) => {
    const dateYear = useMemo(() => date.getFullYear(), [date]);
    const dateMonth = useMemo(() => date.getMonth() + 1, [date]);
    const dateDay = useMemo(() => date.getDate(), [date]);

    return (
      <div
        ref={ref}
        className="max-w-full bg-gradient-to-r from-secondary-dark via-primary-dark to-secondary-light w-full flex-[1_0_300px] rounded-lg"
      >
        <img
          alt={title}
          src={image}
          className="w-full aspect-video max-h-[270px] rounded-t-lg"
        />
        <div className="px-3 py-3 w-full text-card-foreground">
          <h3 className="font-bold text-lg max-sm:text-sm">{title}</h3>
          <p className="font-normal text-sm text-foreground/80 mt-1 max-sm:text-xs">
            {description.substring(0, 100)}...
          </p>
          <div className="mt-7 flex justify-between flex-wrap">
            <Button
              variant="ghost-accent"
              className="px-0 py-0 hover:bg-transparent hover:text-accent-dark"
              asLink
              linkHref={link}
            >
              Read More...
            </Button>
            <span className="flex items-center gap-1 max-sm:text-xs">
              <FaCalendar />
              {`${dateYear}-${dateMonth < 10 ? "0" + dateMonth : dateMonth}-${dateDay < 10 ? "0" + dateDay : dateDay}`}
            </span>
          </div>
        </div>
      </div>
    );
  },
);
BlogCard.displayName = "BlogCard";

export default BlogCard;
