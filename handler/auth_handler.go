package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/dropon/gonzi/database"
)

type AuthService struct {
	repo *database.Repo
}

func (svc AuthService) Show(c *echo.Context) error {
	return c.String(http.StatusOK, "hello, world")
}
