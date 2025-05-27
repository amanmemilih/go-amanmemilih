package httpserver

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Option -.
type Option func(*http.Server)

func WithNewGinEngine() Option {
	return func(s *http.Server) {
		s.Handler = gin.New().Handler()
	}
}

// Port -.
func Port(port string) Option {
	return func(s *http.Server) {
		s.Addr = net.JoinHostPort("", port)
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *http.Server) {
		s.ReadTimeout = timeout
		s.WriteTimeout = timeout
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *http.Server) {
		s.ReadTimeout = timeout
	}
}
