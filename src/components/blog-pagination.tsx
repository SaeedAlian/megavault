import { forwardRef, useEffect, useMemo, useState } from "react";

import { Button } from "./ui";

import { cn } from "../lib/utils";

export interface BlogCardProps {
  defaultPage: number;
  totalPages: number;
  onChangePage?: (newPage: number) => any;
}

const BlogPagination = forwardRef<HTMLDivElement, BlogCardProps>(
  ({ totalPages, defaultPage, onChangePage }, ref) => {
    const [currentPage, setCurrentPage] = useState<number>(defaultPage);
    const [displayStartIndex, setDisplayStartIndex] = useState<number>(0);
    const [displayEndIndex, setDisplayEndIndex] = useState<number>(3);

    const allPages = useMemo<number[]>(() => {
      const array: number[] = [];

      for (let i = 1; i <= totalPages; i++) {
        array.push(i);
      }

      return array;
    }, [totalPages]);

    const onChange = (page: number) => {
      if (onChangePage != null) onChangePage(page);
      setCurrentPage(page);
    };

    const handleSetDisplayIndex = () => {
      const start = currentPage === 1 ? 0 : currentPage - 2;
      const end = currentPage === 1 ? 3 : currentPage + 1;

      setDisplayStartIndex(start < 0 ? 0 : start);
      setDisplayEndIndex(end >= totalPages ? totalPages : end);
    };

    useEffect(() => {
      if (currentPage > totalPages && totalPages !== 0) {
        setCurrentPage(totalPages);
      }
    }, [totalPages]);

    useEffect(handleSetDisplayIndex, [currentPage, totalPages]);

    return (
      <div
        ref={ref}
        className="flex flex-wrap justify-evenly items-center gap-5"
      >
        {allPages.slice(displayStartIndex, displayEndIndex).map((i) => (
          <>
            <Button
              key={`blog-pagination-${i}`}
              onClick={() => onChange(i)}
              variant={currentPage === i ? "contained-primary" : "muted"}
              className={cn(currentPage === i ? "hover:bg-primary" : "")}
            >
              {i}
            </Button>
          </>
        ))}
      </div>
    );
  },
);
BlogPagination.displayName = "BlogPagination";

export default BlogPagination;
