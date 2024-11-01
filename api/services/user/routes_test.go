package user

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

	"megavault/api/services/auth"
	"megavault/api/types/user"
)

func TestUserService(t *testing.T) {
	hashedPassword, err := auth.HashPassword("password")
	if err != nil {
		t.Errorf("Error on hashing password: %v", err)
	}

	userStore := MockUserStore{
		DefaultUsers: []types_user.User{
			{
				Id:        "1",
				Username:  "johndoe",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "johndoe@gmail.com",
				Password:  hashedPassword,
				CreatedAt: time.Now(),
			},
			{
				Id:        "2",
				Username:  "maryjane12",
				FirstName: "Mary",
				LastName:  "Jane",
				Email:     "maryjane@gmail.com",
				Password:  hashedPassword,
				CreatedAt: time.Now(),
			},
			{
				Id:        "3",
				Username:  "alanturing00",
				FirstName: "Alan",
				LastName:  "Turing",
				Email:     "alanturing0090@gmail.com",
				Password:  hashedPassword,
				CreatedAt: time.Now(),
			},
		},
	}

	handler := NewHandler(&userStore)

	t.Run("should get all users", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/user", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/user", handler.getUsers).Methods("GET")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run(
		"should fail on getting all users with 404 status because of wrong query",
		func(t *testing.T) {
			req, err := http.NewRequest("GET", "/user?username=wrongusername", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			router.HandleFunc("/user", handler.getUsers).Methods("GET")

			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusNotFound {
				t.Errorf("Expected code %d, received %d", http.StatusNotFound, rr.Code)
			}
		},
	)

	t.Run(
		"should get only one user from get users method",
		func(t *testing.T) {
			req, err := http.NewRequest("GET", "/user?username=Alanturing00", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			router.HandleFunc("/user", handler.getUsers).Methods("GET")

			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("Expected code %d, received %d", http.StatusOK, rr.Code)
			}
		},
	)

	t.Run(
		"should get a user",
		func(t *testing.T) {
			req, err := http.NewRequest("GET", "/user/1", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			router.HandleFunc("/user/{id}", handler.getUser).Methods("GET")

			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("Expected code %d, received %d", http.StatusOK, rr.Code)
			}
		},
	)

	t.Run(
		"should not get a user",
		func(t *testing.T) {
			req, err := http.NewRequest("GET", "/user/1312321313", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			router.HandleFunc("/user/{id}", handler.getUser).Methods("GET")

			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusNotFound {
				t.Errorf("Expected code %d, received %d", http.StatusNotFound, rr.Code)
			}
		},
	)

	t.Run("should login correctly with username", func(t *testing.T) {
		payload := types_user.LoginUserPayload{
			UsernameOrEmail: strings.ToLower("alanturing00"),
			Password:        "password",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.login).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should login correctly with email", func(t *testing.T) {
		payload := types_user.LoginUserPayload{
			UsernameOrEmail: strings.ToLower("johndoe@gmail.com"),
			Password:        "password",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.login).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should fail to login because of invalid username", func(t *testing.T) {
		payload := types_user.LoginUserPayload{
			UsernameOrEmail: "",
			Password:        "password",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.login).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to login because of invalid password", func(t *testing.T) {
		payload := types_user.LoginUserPayload{
			UsernameOrEmail: strings.ToLower("alanturing00"),
			Password:        "",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.login).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to login because password is short", func(t *testing.T) {
		payload := types_user.LoginUserPayload{
			UsernameOrEmail: strings.ToLower("alanturing00"),
			Password:        "1",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.login).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to login because password is long", func(t *testing.T) {
		payload := types_user.LoginUserPayload{
			UsernameOrEmail: "",
			Password:        strings.Repeat("s", 100000000),
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.login).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to login because username doesn't exist", func(t *testing.T) {
		payload := types_user.LoginUserPayload{
			UsernameOrEmail: strings.ToLower("awdnwoidwaoidjwa"),
			Password:        "password",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.login).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to login because password is incorrect", func(t *testing.T) {
		payload := types_user.LoginUserPayload{
			UsernameOrEmail: strings.ToLower("alanturing00"),
			Password:        "password112",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.login).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should register successfully", func(t *testing.T) {
		payload := types_user.RegisterUserPayload{
			FirstName: "FirstUser",
			LastName:  "FirstUser",
			Username:  strings.ToLower("FirstUser123"),
			Email:     strings.ToLower("FirstUserEmail@gmail.com"),
			Password:  "password",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.register).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected code %d, received %d", http.StatusCreated, rr.Code)
		}

		created, err := handler.store.GetUserByUsername(strings.ToLower("FirstUser123"))
		if err != nil {
			t.Fatal(err)
		}

		if created == nil {
			t.Error("Expected for the created user to exist, but it's not")
		}
	})

	t.Run("should fail to register due to duplicate username", func(t *testing.T) {
		payload := types_user.RegisterUserPayload{
			FirstName: "FirstUser",
			LastName:  "FirstUser",
			Username:  strings.ToLower("alanturing00"),
			Email:     strings.ToLower("SecondEmailShouldFail@gmail.com"),
			Password:  "password",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.register).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}

		created, err := handler.store.GetUserByEmail(
			strings.ToLower("SecondEmailShouldFail@gmail.com"),
		)
		if created != nil {
			t.Error("Expected for the created user to not be found, but it has been found")
		}
	})

	t.Run("should fail to register due to duplicate email", func(t *testing.T) {
		payload := types_user.RegisterUserPayload{
			FirstName: "FirstUser",
			LastName:  "FirstUser",
			Username:  strings.ToLower("SecondUsernameShouldFail"),
			Email:     strings.ToLower("alanturing0090@gmail.com"),
			Password:  "password",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.register).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}

		created, err := handler.store.GetUserByUsername(strings.ToLower("SecondUsernameShouldFail"))
		if created != nil {
			t.Error("Expected for the created user to not be found, but it has been found")
		}
	})

	t.Run("should fail to register because of invalid data", func(t *testing.T) {
		payload := types_user.RegisterUserPayload{
			FirstName: "FirstUser",
			LastName:  "FirstUser",
			Username:  strings.ToLower("newErrorUser"),
			Email:     strings.ToLower("gmail"),
			Password:  "password",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.register).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}

		created, err := handler.store.GetUserByUsername(strings.ToLower("newErrorUser"))
		if created != nil {
			t.Error("Expected for the created user to not be found, but it has been found")
		}
	})

	t.Run("should fail to register because of short password", func(t *testing.T) {
		payload := types_user.RegisterUserPayload{
			FirstName: "FirstUser",
			LastName:  "FirstUser",
			Username:  strings.ToLower("NewUserCreated"),
			Email:     strings.ToLower("NewUser@gmail.com"),
			Password:  "1",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.register).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}

		created, err := handler.store.GetUserByUsername(strings.ToLower("NewUserCreated"))
		if created != nil {
			t.Error("Expected for the created user to not be found, but it has been found")
		}
	})

	t.Run("should fail to register because of long password", func(t *testing.T) {
		payload := types_user.RegisterUserPayload{
			FirstName: "FirstUser",
			LastName:  "FirstUser",
			Username:  strings.ToLower("NewUserCreated1"),
			Email:     strings.ToLower("NewUser1@gmail.com"),
			Password:  strings.Repeat("1", 5000000),
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.register).Methods("POST")

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected code %d, received %d", http.StatusBadRequest, rr.Code)
		}

		created, err := handler.store.GetUserByUsername(strings.ToLower("NewUserCreated1"))
		if created != nil {
			t.Error("Expected for the created user to not be found, but it has been found")
		}
	})
}

type MockUserStore struct {
	DefaultUsers []types_user.User
}

type MockGetUsersResult struct {
	Result map[string][]types_user.User
}

func (m *MockUserStore) GetUserById(id string) (*types_user.User, error) {
	for i := range m.DefaultUsers {
		u := m.DefaultUsers[i]

		if u.Id == id {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("Cannot find user")
}

func (m *MockUserStore) GetUserByUsername(username string) (*types_user.User, error) {
	for i := range m.DefaultUsers {
		u := m.DefaultUsers[i]

		if u.Username == username {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("Cannot find user")
}

func (m *MockUserStore) GetUserByEmail(email string) (*types_user.User, error) {
	for i := range m.DefaultUsers {
		u := m.DefaultUsers[i]

		if u.Email == email {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("Cannot find user")
}

func (m *MockUserStore) GetUserByUsernameOrEmail(
	username string,
	email string,
) (*types_user.User, error) {
	for i := range m.DefaultUsers {
		u := m.DefaultUsers[i]

		if u.Email == email || u.Username == username {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("Cannot find user")
}

func (m *MockUserStore) CreateUser(u types_user.RegisterUserPayload) (*types_user.User, error) {
	created := types_user.User{
		Id:        strconv.Itoa(rand.Int()),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: time.Now(),
	}

	m.DefaultUsers = append(m.DefaultUsers, created)

	return &created, nil
}

func (m *MockUserStore) GetUsers(
	query types_user.SearchUserQuery,
) ([]types_user.User, error) {
	var res []types_user.User

	for i := range m.DefaultUsers {
		u := m.DefaultUsers[i]

		if len(query.Username) > 0 {
			if strings.Contains(u.Username, query.Username) {
				res = append(res, u)
			}
		} else {
			res = append(res, u)
		}
	}

	return res, nil
}

func (m *MockUserStore) DeleteUserById(
	id string,
) error {
	var res []types_user.User

	for i := range m.DefaultUsers {
		u := m.DefaultUsers[i]

		if u.Id != id {
			res = append(res, u)
		}
	}

	if len(res) == len(m.DefaultUsers) {
		return fmt.Errorf("User not found to delete")
	}

	m.DefaultUsers = res

	return nil
}

func (m *MockUserStore) DeleteUserByUsername(
	username string,
) error {
	var res []types_user.User

	for i := range m.DefaultUsers {
		u := m.DefaultUsers[i]

		if u.Username != username {
			res = append(res, u)
		}
	}

	if len(res) == len(m.DefaultUsers) {
		return fmt.Errorf("User not found to delete")
	}

	m.DefaultUsers = res

	return nil
}
