package blog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"megavault/api/types/blog"
	"megavault/api/types/user"
	"megavault/api/utils"
)

func TestBlogService(t *testing.T) {
	userStore := MockUserStore{}
	blogStore := MockBlogStore{
		DefaultBlogs: []types_blog.Blog{
			{
				Id:          "1",
				Title:       "Blog1",
				Description: "Blog1",
				Slug:        "blog1",
				PictureName: "blog1-pic.jpg",
				MDFilename:  "blog1.md",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Id:          "2",
				Title:       "Blog2",
				Description: "Blog2",
				Slug:        "blog2",
				PictureName: "blog2-pic.jpg",
				MDFilename:  "blog2.md",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Id:          "3",
				Title:       "Blog3",
				Description: "Blog3",
				Slug:        "blog3",
				PictureName: "blog3-pic.jpg",
				MDFilename:  "blog3.md",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
	}

	mdFileUploadDir := "testuploads/blogs/mds"
	imageUploadDir := "testuploads/blogs/images"
	handler := NewHandler(&blogStore, &userStore, mdFileUploadDir, imageUploadDir)

	t.Run("should get all blogs successfully", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/blog", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog", handler.getBlogs).Methods("GET")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should get no blogs at all because of wrong keyword", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/blog?keyword=wrongkeyword", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog", handler.getBlogs).Methods("GET")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected code %d, received %d", http.StatusNotFound, rr.Code)
		}
	})

	t.Run("should get only one blog with getBlogs", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/blog?keyword=Blog1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog", handler.getBlogs).Methods("GET")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should get blog by slug", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/blog/blog1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog/{slug}", handler.getBlog).Methods("GET")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should fail to get a blog because of wrong slug", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/blog/blog19", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog/{slug}", handler.getBlog).Methods("GET")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected code %d, received %d", http.StatusNotFound, rr.Code)
		}
	})

	t.Run("should create a blog successfully", func(t *testing.T) {
		payload := types_blog.CreateBlogPayload{
			Title:       "My Test Blog",
			Slug:        "",
			Description: "This is a test blog",
			PictureName: "test.jpg",
			MDFilename:  "test.md",
		}

		payload.Slug = utils.CreateSlug(payload.Title)

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/blog", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog", handler.createBlog).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected code %d, received %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("should fail to create a blog because of wrong picture name", func(t *testing.T) {
		payload := types_blog.CreateBlogPayload{
			Title:       "My Test Blog 2",
			Slug:        "",
			Description: "This is a test blog",
			PictureName: "test123.jpg",
			MDFilename:  "test.md",
		}

		payload.Slug = utils.CreateSlug(payload.Title)

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/blog", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog", handler.createBlog).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to create a blog because of wrong md file name", func(t *testing.T) {
		payload := types_blog.CreateBlogPayload{
			Title:       "My Test Blog 3",
			Slug:        "",
			Description: "This is a test blog",
			PictureName: "test.jpg",
			MDFilename:  "test123.md",
		}

		payload.Slug = utils.CreateSlug(payload.Title)

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/blog", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog", handler.createBlog).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to create a blog because of duplicated title", func(t *testing.T) {
		payload := types_blog.CreateBlogPayload{
			Title:       "Blog1",
			Slug:        "",
			Description: "This is a test blog",
			PictureName: "test.jpg",
			MDFilename:  "test.md",
		}

		payload.Slug = utils.CreateSlug(payload.Title)

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/blog", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog", handler.createBlog).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to create a blog because of wrong payload", func(t *testing.T) {
		payload := types_blog.CreateBlogPayload{
			Title:       "",
			Slug:        "",
			Description: "This is a test blog",
			PictureName: "test.jpg",
			MDFilename:  "test.md",
		}

		payload.Slug = utils.CreateSlug(payload.Title)

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/blog", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog", handler.createBlog).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should delete blog successfully", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/blog/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog/{id}", handler.deleteBlog).Methods("DELETE")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should fail to delete blog because of wrong id", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/blog/11234", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/blog/{id}", handler.deleteBlog).Methods("DELETE")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected code %d, received %d", http.StatusNotFound, rr.Code)
		}
	})
}

type MockBlogStore struct {
	DefaultBlogs []types_blog.Blog
}

type MockGetBlogsResult struct {
	Result map[string][]types_blog.Blog
}

type MockUserStore struct{}

func (m *MockUserStore) GetUserById(id string) (*types_user.User, error) {
	return nil, nil
}

func (m *MockUserStore) GetUserByUsername(username string) (*types_user.User, error) {
	return nil, nil
}

func (m *MockUserStore) GetUserByEmail(email string) (*types_user.User, error) {
	return nil, nil
}

func (m *MockUserStore) GetUserByUsernameOrEmail(
	username string,
	email string,
) (*types_user.User, error) {
	return nil, nil
}

func (m *MockUserStore) CreateUser(u types_user.RegisterUserPayload) (*types_user.User, error) {
	return nil, nil
}

func (m *MockUserStore) GetUsers(
	query types_user.SearchUserQuery,
) ([]types_user.User, error) {
	return nil, nil
}

func (m *MockUserStore) DeleteUserById(
	id string,
) error {
	return nil
}

func (m *MockUserStore) DeleteUserByUsername(
	username string,
) error {
	return nil
}

func (m *MockBlogStore) GetBlogById(id string) (*types_blog.Blog, error) {
	for i := range m.DefaultBlogs {
		b := m.DefaultBlogs[i]

		if b.Id == id {
			return &b, nil
		}
	}

	return nil, fmt.Errorf("Cannot find blog")
}

func (m *MockBlogStore) GetBlogBySlug(slug string) (*types_blog.Blog, error) {
	for i := range m.DefaultBlogs {
		b := m.DefaultBlogs[i]

		if b.Slug == slug {
			return &b, nil
		}
	}

	return nil, fmt.Errorf("Cannot find blog")
}

func (m *MockBlogStore) CreateBlog(b types_blog.CreateBlogPayload) (*types_blog.Blog, error) {
	created := types_blog.Blog{
		Id:          strconv.Itoa(rand.Int()),
		Title:       b.Title,
		Description: b.Description,
		Slug:        b.Slug,
		PictureName: b.PictureName,
		MDFilename:  b.MDFilename,
		CreatedAt:   time.Now(),
	}

	m.DefaultBlogs = append(m.DefaultBlogs, created)

	return &created, nil
}

func (m *MockBlogStore) GetBlogs(
	query types_blog.SearchBlogQuery,
) ([]types_blog.Blog, error) {
	var res []types_blog.Blog

	for i := range m.DefaultBlogs {
		b := m.DefaultBlogs[i]

		if len(query.Keyword) > 0 {
			if strings.Contains(strings.ToLower(b.Description), query.Keyword) ||
				strings.Contains(strings.ToLower(b.Title), query.Keyword) {
				res = append(res, b)
			}
		} else {
			res = append(res, b)
		}
	}

	return res, nil
}

func (m *MockBlogStore) UpdateBlog(id string, payload types_blog.UpdateBlogPayload) error {
	var res []types_blog.Blog

	isUpdated := false

	for i := range m.DefaultBlogs {
		b := m.DefaultBlogs[i]

		if id == b.Id {
			updatedBlog := types_blog.Blog{
				Id:          b.Id,
				Title:       payload.Title,
				Description: payload.Description,
				Slug:        payload.Slug,
				PictureName: payload.PictureName,
				MDFilename:  payload.MDFilename,
				CreatedAt:   b.CreatedAt,
				UpdatedAt:   payload.UpdatedAt,
			}

			res = append(res, updatedBlog)
			isUpdated = true
		}
	}

	if !isUpdated {
		return fmt.Errorf("Blog not found to update")
	}

	m.DefaultBlogs = res

	return nil
}

func (m *MockBlogStore) DeleteBlogById(
	id string,
) error {
	var res []types_blog.Blog

	for i := range m.DefaultBlogs {
		u := m.DefaultBlogs[i]

		if u.Id != id {
			res = append(res, u)
		}
	}

	if len(res) == len(m.DefaultBlogs) {
		return fmt.Errorf("Blog not found to delete")
	}

	m.DefaultBlogs = res

	return nil
}
