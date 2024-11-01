package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"megavault/api/services/blog"
	"megavault/api/services/user"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userSubrouter := subrouter.PathPrefix("/user").Subrouter()
	blogSubrouter := subrouter.PathPrefix("/blog").Subrouter()

	userStore := user.NewStore(s.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(userSubrouter)

	blogStore := blog.NewStore(s.db)
	blogService := blog.NewHandler(blogStore, userStore)
	blogService.RegisterRoutes(blogSubrouter)

	log.Println("API Listening on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
