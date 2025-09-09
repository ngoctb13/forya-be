package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/config"
)

type Server struct {
	httpServer *http.Server
	router     *gin.Engine
	cfg        *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	router := gin.New()
	return &Server{
		router: router,
		cfg:    cfg,
	}
}

func (s *Server) ListenHTTP() error {
	_ = os.Setenv("PORT", "8090")
	listen, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Printf("Error listening on port %v", os.Getenv("PORT"))
		panic(err)
	}

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	s.httpServer = &http.Server{
		Addr:    address,
		Handler: s.router,
	}

	log.Printf("Starting http server on port %v ...", os.Getenv("PORT"))

	return s.httpServer.Serve(listen)
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Printf("Shutting down server gracefully...")
	return s.httpServer.Shutdown(ctx)
}
