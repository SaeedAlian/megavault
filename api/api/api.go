package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SaeedAlian/megavault/api/config"
	"github.com/SaeedAlian/megavault/api/services/blog"
	"github.com/SaeedAlian/megavault/api/services/user"
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

	blogMdFileUploadDir := fmt.Sprintf("%s/blogs/mds", config.Env.UploadsRootDir)
	blogImageUploadDir := fmt.Sprintf("%s/blogs/images", config.Env.UploadsRootDir)

	userStore := user.NewStore(s.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(userSubrouter)

	blogStore := blog.NewStore(s.db)
	blogService := blog.NewHandler(blogStore, userStore, blogMdFileUploadDir, blogImageUploadDir)
	blogService.RegisterRoutes(blogSubrouter)

	log.Println("API Listening on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
