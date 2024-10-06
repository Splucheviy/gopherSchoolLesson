package apiserver

import (
	"github.com/Splucheviy/gopherSchoolLesson/internal/app/store"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *echo.Echo
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: echo.New(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.GET("/hello", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	s.router.POST("/users", handleUsersCreate)
}

// Function to handle creating a user
func handleUsersCreate(c echo.Context) error {
	return nil
}
