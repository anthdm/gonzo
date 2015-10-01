package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/dropon/gonzi/database"
	"github.com/dropon/gonzi/model"
)

type UserService struct {
	Repo *database.Repo
}

func (svc UserService) Show(c *echo.Context) error {
	var user model.User
	svc.Repo.First(&user, c.Param("id"))
	return encodeJSON(c, http.StatusOK, &user)
}

func (svc UserService) List(c *echo.Context) error {
	var users []*model.User
	svc.Repo.Find(&users)
	return encodeJSON(c, http.StatusOK, users)
}

func (svc UserService) Create(c *echo.Context) error {
	var user model.User
	if err := decode(c.Request(), &user); err != nil {
		return err
	}

	if errs := user.Validate(); errs != nil {
		return errs[0]
	}
	svc.Repo.Create(&user)

	return c.JSON(http.StatusCreated, &user)
}

func (svc UserService) Update(c *echo.Context) error {
	return c.String(http.StatusOK, "hello, v1 world")
}

func (svc UserService) Delete(c *echo.Context) error {
	return c.String(http.StatusOK, "hello, v1 world")
}
