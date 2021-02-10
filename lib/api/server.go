package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pytyagi/wisdom/lib/wisdom"
)

// Server struct with config
type Server struct {
	*Config
	router    *echo.Echo
	dispenser wisdom.Dispenser
}

// NewServer return instance of a server
func NewServer(cfg *Config, dispenser wisdom.Dispenser) *Server {
	server := &Server{
		Config:    cfg,
		router:    echo.New(),
		dispenser: dispenser,
	}

	server.router.GET(path.Join(cfg.APIPath, "/ping"), server.Ping)
	server.router.GET(path.Join(cfg.APIPath, "/version"), server.Version)
	server.router.GET(path.Join(cfg.APIPath, "/quote"), server.Quote)
	server.router.Use(middleware.Logger())
	server.router.HideBanner = true
	server.router.HidePort = true
	return server
}

// Start start the server with given config
func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	log.Printf("listening on %s \n", address)

	return s.router.Start(address)
}

// Stop Shutdown the server
func (s *Server) Stop(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}

// Ping is a health check for the server
func (s *Server) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"ping": "pong"})

}

// Version is seek the current version info
func (s *Server) Version(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"Version": wisdom.Version})
}

// Quote isa random quote
func (s *Server) Quote(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"Quote": s.dispenser.Random()})
}
