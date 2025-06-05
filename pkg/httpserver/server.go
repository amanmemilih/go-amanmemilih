package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server *http.Server
	notify chan error
	Router *gin.Engine

	shutdownTimeout time.Duration
}

func New(opts ...Option) *Server {
	router := gin.Default()

	srv := &http.Server{
		Handler: router.Handler(),
	}

	for _, opt := range opts {
		opt(srv)
	}

	engine := srv.Handler.(*gin.Engine)

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return &Server{
		server:          srv,
		notify:          make(chan error, 1),
		Router:          engine,
		shutdownTimeout: _defaultShutdownTimeout,
	}
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
