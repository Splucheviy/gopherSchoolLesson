package apiserver

import (
	"net/http"

	"github.com/Splucheviy/gopherSchoolLesson/internal/app/model"
	"github.com/Splucheviy/gopherSchoolLesson/internal/app/store"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	sessionName = "user-session"
)

type server struct {
	router       *echo.Echo
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       echo.New(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	s.router.POST("/users", s.handleUserCreate)
	s.router.POST("/sessions", s.handleSessionsCreate)
}



func (s *server) handleUserCreate(c echo.Context) error {
	type request struct {
		Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
		Password string `param:"password" query:"password" form:"password" json:"password" xml:"password"`
	}

	var req request

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.store.User().Create(u); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	u.Sanitize()
	return c.JSON(http.StatusCreated, u)
}

func (s *server) handleSessionsCreate(c echo.Context) error {
	type request struct {
		Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
		Password string `param:"password" query:"password" form:"password" json:"password" xml:"password"`
	}
	var req request

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	u, err := s.store.User().FindByEmail(req.Email)
	if err != nil || !u.ComparePassword(req.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email address or password"})
	}

	session, err := s.sessionStore.Get(c.Request(), sessionName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	session.Values["user_id"] = u.ID
	if err := s.sessionStore.Save(c.Request(), c.Response(), session); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	return c.JSON(http.StatusOK, nil)
}
	