package context

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go-ms-session/config"
	"go-ms-session/logger"
	"go-ms-session/repository"
	"os"
	"strconv"
)

type Server struct {
	Serve           *echo.Echo
	ServerName      string
	Config          config.ServerConfig
	MongoRepository *repository.MongoRepository
}

func CreateServer(serveName string, c config.Config, skipper middleware.Skipper) *Server {
	s := &Server{}
	s.ServerName = serveName
	s.Config = config.LoadConfig(
		c.ParenPath,
		c.Resource,
		c.Env,
		c.Application)
	s.Serve = echo.New()

	logger.Logger = logrus.New()
	s.Serve.Logger = logger.GetEchoLogger()
	s.Serve.Use(logger.Hook())

	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", skipper)
	p.Use(s.Serve)

	return s
}

func (s *Server) GetPort(portConfig int) string {
	var port = os.Getenv("PORT") // ----> (A)
	if port == "" {
		port = strconv.Itoa(portConfig)
	}
	return ":" + port // ----> (B)
}
