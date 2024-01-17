package users

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       string `json:"id"       db:"id"`
	Username string `json:"username" db:"username"`
}

type UserHandlers struct {
	db *sqlx.DB
}

func NewHandlers(db *sqlx.DB) *UserHandlers {
	return &UserHandlers{db}
}

func (h *UserHandlers) CreateUser(c echo.Context) error {
	fmt.Println(c.Request().Body)
	return c.JSON(http.StatusOK, "ok")
}

func (h *UserHandlers) GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}