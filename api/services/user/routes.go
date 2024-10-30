package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"

	"megavault/api/services/auth"
	types_user "megavault/api/types"
	"megavault/api/utils"
)

type Handler struct {
	store types_user.UserStore
}

func NewHandler(store types_user.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", auth.WithJWTAuth(h.getUsers, h.store)).Methods("GET")
	router.HandleFunc("/me", auth.WithJWTAuth(h.getMe, h.store)).Methods("GET")
	router.HandleFunc("/{id}", auth.WithJWTAuth(h.getUser, h.store)).Methods("GET")

	router.HandleFunc("/register", h.register).Methods("POST")
	router.HandleFunc("/login", h.login).Methods("POST")
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var credentials types_user.LoginUserPayload
	if err := utils.ParseJSONFromRequest(r, &credentials); err != nil {
		utils.WriteErrorInResponse(w, http.StatusBadRequest, "Invalid login payload")
		return
	}

	if err := utils.Validator.Struct(credentials); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Invalid payload: %v", errors),
		)
		return
	}

	user, err := h.store.GetUserByUsernameOrEmail(
		credentials.UsernameOrEmail,
		credentials.UsernameOrEmail,
	)
	if err != nil || user == nil {
		utils.WriteErrorInResponse(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	if isPasswordCorrect := auth.ComparePassword(credentials.Password, user.Password); !isPasswordCorrect {
		utils.WriteErrorInResponse(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	jwt, err := auth.GenerateJWT(jwt.MapClaims{"userId": user.Id}, 1*24*60)
	if err != nil {
		utils.WriteErrorInResponse(w, http.StatusInternalServerError, "An error occurred")
		return
	}

	utils.WriteJSONInResponse(w, http.StatusOK, map[string]string{"token": jwt}, nil)
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var user types_user.RegisterUserPayload
	if err := utils.ParseJSONFromRequest(r, &user); err != nil {
		utils.WriteErrorInResponse(w, http.StatusBadRequest, "Invalid user payload")
		return
	}

	if err := utils.Validator.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Invalid payload: %v", errors),
		)
		return
	}

	if u, _ := h.store.GetUserByUsernameOrEmail(user.Username, user.Email); u != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			"Another user with this email or username already exists",
		)
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusInternalServerError,
			"An error occurred",
		)
		return
	}

	created_user, err := h.store.CreateUser(types_user.RegisterUserPayload{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusInternalServerError,
			"An error occurred",
		)
		return
	}

	utils.WriteJSONInResponse(w, http.StatusCreated, created_user, nil)
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	usernameQuery := r.URL.Query().Get("username")

	query := types_user.SearchUserQuery{
		Username: usernameQuery,
	}

	users, err := h.store.GetUsers(query)
	if err != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusInternalServerError,
			"An error occurred",
		)
		return
	}

	if len(users) == 0 {
		utils.WriteErrorInResponse(
			w,
			http.StatusNotFound,
			"No user found",
		)
		return
	}

	payload := map[string][]types_user.User{
		"result": users,
	}

	utils.WriteJSONInResponse(w, http.StatusOK, payload, nil)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, ok := vars["id"]
	if !ok {
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			"User id not found",
		)
		return
	}

	u, err := h.store.GetUserById(userId)
	if err != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusNotFound,
			"User not found",
		)
		return
	}

	utils.WriteJSONInResponse(w, http.StatusOK, u, nil)
}

func (h *Handler) getMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId")
	if userId == nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusBadRequest,
			"User id not found within the authorization token",
		)
		return
	}

	u, err := h.store.GetUserById(string(userId.(string)))
	if err != nil {
		utils.WriteErrorInResponse(
			w,
			http.StatusNotFound,
			"Invalid user id",
		)
		return
	}

	utils.WriteJSONInResponse(w, http.StatusOK, u, nil)
}
