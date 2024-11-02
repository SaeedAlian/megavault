package blog

import (
	"database/sql"
	"fmt"

	"megavault/api/types/blog"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateBlog(blog types_blog.CreateBlogPayload) (*types_blog.Blog, error) {
	rowId := ""
	err := s.db.QueryRow(
		"INSERT INTO blogs (title,description,slug,mdFilename,pictureName) VALUES ($1,$2,$3,$4,$5) RETURNING id;",
		blog.Title,
		blog.Description,
		blog.Slug,
		blog.MDFilename,
		blog.PictureName,
	).Scan(&rowId)
	if err != nil {
		return nil, err
	}

	b, err := s.GetBlogById(rowId)
	if err != nil || b == nil {
		return nil, err
	}

	return b, nil
}

func (s *Store) GetBlogs(query types_blog.SearchBlogQuery) ([]types_blog.Blog, error) {
	rows, err := s.db.Query(
		"SELECT * FROM blogs WHERE title ILIKE $1 OR description ILIKE $2;",
		fmt.Sprintf("%%%s%%", query.Keyword),
		fmt.Sprintf("%%%s%%", query.Keyword),
	)
	if err != nil {
		return nil, err
	}

	blogs := []types_blog.Blog{}

	for rows.Next() {
		blog, err := scanRow(rows)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, *blog)
	}

	return blogs, nil
}

func (s *Store) GetBlogById(id string) (*types_blog.Blog, error) {
	rows, err := s.db.Query("SELECT * FROM blogs WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}

	blog := new(types_blog.Blog)

	for rows.Next() {
		blog, err = scanRow(rows)
		if err != nil {
			return nil, err
		}
	}

	if blog.Id == "" {
		return nil, fmt.Errorf("Blog not found")
	}

	return blog, nil
}

func (s *Store) GetBlogBySlug(slug string) (*types_blog.Blog, error) {
	rows, err := s.db.Query("SELECT * FROM blogs WHERE slug = $1;", slug)
	if err != nil {
		return nil, err
	}

	blog := new(types_blog.Blog)

	for rows.Next() {
		blog, err = scanRow(rows)
		if err != nil {
			return nil, err
		}
	}

	if blog.Id == "" {
		return nil, fmt.Errorf("Blog not found")
	}

	return blog, nil
}

func (s *Store) UpdateBlog(id string, blog types_blog.UpdateBlogPayload) error {
	_, err := s.db.Exec(
		"UPDATE blogs SET title = $1, description = $2, slug = $3, pictureName = $4, mdFilename = $5, updatedAt = $6 WHERE id = $7",
		blog.Title,
		blog.Description,
		blog.Slug,
		blog.PictureName,
		blog.MDFilename,
		blog.UpdatedAt,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteBlogById(id string) error {
	_, err := s.db.Exec("DELETE FROM blogs WHERE id = $1;", id)
	if err != nil {
		return err
	}

	return nil
}

func scanRow(rows *sql.Rows) (*types_blog.Blog, error) {
	blog := new(types_blog.Blog)

	err := rows.Scan(
		&blog.Id,
		&blog.Title,
		&blog.Description,
		&blog.Slug,
		&blog.PictureName,
		&blog.MDFilename,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return blog, nil
}
