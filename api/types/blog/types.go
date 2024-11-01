package types_blog

import (
	"time"
)

type BlogStore interface {
	CreateBlog(blog CreateBlogPayload) (*Blog, error)
	GetBlogs(query SearchBlogQuery) ([]Blog, error)
	GetBlogById(id string) (*Blog, error)
	GetBlogBySlug(slug string) (*Blog, error)
	UpdateBlog(id string, blog UpdateBlogPayload) error
	DeleteBlogById(id string) error
}

type Blog struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	PictureName string    `json:"pictureName"`
	MDFilename  string    `json:"mdFilename"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateBlogPayload struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"       validate:"required"`
	Description string `json:"description" validate:"required"`
	PictureName string `json:"pictureName" validate:"required"`
	MDFilename  string `json:"mdFilename"  validate:"required"`
}

type UpdateBlogPayload struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PictureName string    `json:"pictureName"`
	MDFilename  string    `json:"mdFilename"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type SearchBlogQuery struct {
	Keyword string `json:"keyword"`
}
