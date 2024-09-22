package apiserver

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// APIserver...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *echo.Echo
}

// New...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: echo.New(),
	}
}

// Start...
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Infof("Starting API server on %s", s.config.ServerAddr)

	return s.router.Start(s.config.ServerAddr)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIserver) configureRouter() {
	s.router.GET("/hello", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
}
