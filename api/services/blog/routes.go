package blog

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"megavault/api/config"
	"megavault/api/services/auth"
	"megavault/api/types/blog"
	"megavault/api/types/user"
	"megavault/api/utils"
)

type Handler struct {
	store           types_blog.BlogStore
	userStore       types_user.UserStore
	mdFileUploadDir string
	imageUploadDir  string
}

func NewHandler(store types_blog.BlogStore, userStore types_user.UserStore) *Handler {
	return &Handler{
		store:           store,
		userStore:       userStore,
		mdFileUploadDir: fmt.Sprintf("%s/blogs/mds", config.Env.UploadsRootDir),
		imageUploadDir:  fmt.Sprintf("%s/blogs/images", config.Env.UploadsRootDir),
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	blogMdUploadHandler := utils.FileUploadHandler(
		"mdFile",
		5,
		[]string{"text/markdown"},
		h.mdFileUploadDir,
	)

	blogImageUploadHandler := utils.FileUploadHandler(
		"image",
		3,
		[]string{"image/jpeg", "image/png", "image/jpg", "image/webp"},
		h.imageUploadDir,
	)

	router.HandleFunc("/", auth.WithJWTAuth(h.getBlogs, h.userStore)).Methods("GET")
	router.HandleFunc("/{slug}", auth.WithJWTAuth(h.getBlog, h.userStore)).Methods("GET")
	router.HandleFunc("/", auth.WithJWTAuth(h.createBlog, h.userStore)).Methods("POST")
	router.HandleFunc("/md", auth.WithJWTAuth(blogMdUploadHandler, h.userStore)).Methods("POST")
	router.HandleFunc("/image", auth.WithJWTAuth(blogImageUploadHandler, h.userStore)).
		Methods("POST")
	router.HandleFunc("/{id}", auth.WithJWTAuth(h.updateBlog, h.userStore)).Methods("PATCH")
	router.HandleFunc("/{id}", auth.WithJWTAuth(h.deleteBlog, h.userStore)).Methods("DELETE")
}

func (h *Handler) createBlog(w http.ResponseWriter, r *http.Request) {
	var payload types_blog.CreateBlogPayload
	if err := utils.ParseJSONFromRequest(r, &payload); err != nil {
		utils.WriteErrorInResponse(w, http.StatusBadRequest, "Invalid blog payload")
		return
	}

	if err := utils.Validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Invalid payload: %v", errors),
		)
		return
	}

	slug := utils.CreateSlug(payload.Title)

	if b, _ := h.store.GetBlogBySlug(slug); b != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			"Another blog with that title already exists",
		)
		return
	}

	isMdExists, err := utils.PathExists(fmt.Sprintf("%s/%s", h.mdFileUploadDir, payload.MDFilename))
	if err != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusInternalServerError,
			"An error occurred",
		)
		return
	}

	if !isMdExists {
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			"MD file doesn't exist, please make sure to upload it first",
		)
		return
	}

	isPictureExists, err := utils.PathExists(
		fmt.Sprintf("%s/%s", h.imageUploadDir, payload.PictureName),
	)
	if err != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusInternalServerError,
			"An error occurred",
		)
		return
	}

	if !isPictureExists {
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			"Picture file doesn't exist, please make sure to upload it first",
		)
		return
	}

	b, err := h.store.CreateBlog(types_blog.CreateBlogPayload{
		Title:       payload.Title,
		Description: payload.Description,
		PictureName: payload.PictureName,
		MDFilename:  payload.MDFilename,
		Slug:        slug,
	})
	if err != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusInternalServerError,
			"An error occurred",
		)
		return
	}

	utils.WriteJSONInResponse(w, http.StatusCreated, b, nil)
}

func (h *Handler) getBlogs(w http.ResponseWriter, r *http.Request) {
	keywordQuery := r.URL.Query().Get("keyword")

	query := types_blog.SearchBlogQuery{
		Keyword: strings.ToLower(keywordQuery),
	}

	blogs, err := h.store.GetBlogs(query)
	if err != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusInternalServerError,
			"An error occurred",
		)
		return
	}

	if len(blogs) == 0 {
		utils.WriteErrorInResponse(
			w,
			http.StatusNotFound,
			"No blog found",
		)
		return
	}

	payload := map[string][]types_blog.Blog{
		"result": blogs,
	}

	utils.WriteJSONInResponse(w, http.StatusOK, payload, nil)
}

func (h *Handler) getBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			"Blog slug not found",
		)
		return
	}

	b, err := h.store.GetBlogBySlug(slug)
	if err != nil || b == nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusNotFound,
			"Blog not found",
		)
		return
	}

	utils.WriteJSONInResponse(w, http.StatusOK, b, nil)
}

func (h *Handler) updateBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId, ok := vars["id"]
	if !ok {
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			"Blog id not found",
		)
		return
	}

	var payload types_blog.UpdateBlogPayload
	if err := utils.ParseJSONFromRequest(r, &payload); err != nil {
		utils.WriteErrorInResponse(w, http.StatusBadRequest, "Invalid blog payload")
		return
	}

	if err := utils.Validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Invalid payload: %v", errors),
		)
		return
	}

	b, err := h.store.GetBlogById(blogId)
	if err != nil || b == nil {
		utils.WriteErrorInResponse(w, http.StatusNotFound, "Blog not found")
		return
	}

	updatedDate := time.Now()
	updatePayload := types_blog.UpdateBlogPayload{
		Slug:        b.Slug,
		Title:       b.Title,
		Description: b.Description,
		PictureName: b.PictureName,
		MDFilename:  b.MDFilename,
		UpdatedAt:   updatedDate,
	}

	if payload.Title != "" {
		newSlug := utils.CreateSlug(payload.Title)

		if b, _ := h.store.GetBlogBySlug(newSlug); b != nil {
			utils.WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				"Another blog with that title already exists",
			)
			return
		}

		updatePayload.Title = payload.Title
		updatePayload.Slug = newSlug
	}

	if payload.Description != "" {
		updatePayload.Description = payload.Description
	}

	if payload.PictureName != "" {
		isPictureExists, err := utils.PathExists(
			fmt.Sprintf("%s/%s", h.imageUploadDir, payload.PictureName),
		)
		if err != nil {
			utils.WriteErrorInResponse(
				w,
				http.StatusInternalServerError,
				"An error occurred",
			)
			return
		}

		if !isPictureExists {
			utils.WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				"Picture file doesn't exist, please make sure to upload it first",
			)
			return
		}

		updatePayload.PictureName = payload.PictureName
	}

	if payload.MDFilename != "" {
		isMdExists, err := utils.PathExists(
			fmt.Sprintf("%s/%s", h.mdFileUploadDir, payload.MDFilename),
		)
		if err != nil {
			utils.WriteErrorInResponse(
				w,
				http.StatusInternalServerError,
				"An error occurred",
			)
			return
		}

		if !isMdExists {
			utils.WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				"MD file doesn't exist, please make sure to upload it first",
			)
			return
		}

		updatePayload.MDFilename = payload.MDFilename
	}

	if err := h.store.UpdateBlog(b.Id, updatePayload); err != nil {
		utils.WriteErrorInResponse(w, http.StatusInternalServerError, "An error occurred")
		return
	}

	utils.WriteJSONInResponse(
		w,
		http.StatusOK,
		map[string]string{
			"message": fmt.Sprintf("Blog with id %s has been updated successfully", b.Id),
		},
		nil,
	)
}

func (h *Handler) deleteBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId, ok := vars["id"]
	if !ok {
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			"Blog id not found",
		)
		return
	}

	b, err := h.store.GetBlogById(blogId)
	if err != nil || b == nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusNotFound,
			"Blog not found",
		)
		return
	}

	if err := h.store.DeleteBlogById(b.Id); err != nil {
		utils.WriteErrorInResponse(w, http.StatusInternalServerError, "An error occurred")
		return
	}

	utils.WriteJSONInResponse(
		w,
		http.StatusOK,
		map[string]string{
			"message": fmt.Sprintf("Blog with id %s has been deleted successfully", b.Id),
		},
		nil,
	)
}
