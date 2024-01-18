package users

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-faker/faker/v4"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// User struct    represents a basic user for demo, no login information will be necessary
type User struct {
	ID       string `json:"id"       db:"id"`
	Username string `json:"username" db:"username"`
}

// Generate method    is a helper method that generates a dummy representation of a User struct with pre-filled
// fields using fake/random data
func (u *User) Generate() User {
	u.ID = faker.UUIDDigit()
	u.Username = faker.Word()

	return *u
}

// provides a proxy for reusing the instance of sqlx.DB in task specific http handlers
type UserHandlers struct {
	db *sqlx.DB
}

// NewHandlers function    constructor function for UserHandlers
func NewHandlers(db *sqlx.DB) *UserHandlers {
	return &UserHandlers{db}
}

// @Summary Create a new user
// @Description create a new task
// @Accept  json
// @Produce  json
// @Success 200 {object} User "Task successfully retrieved"
// @Router /users [POST]
func (h *UserHandlers) CreateUser(c echo.Context) error {
	u := User{}
	u.Generate()

	query := `INSERT INTO users (id, username) VALUES (:id, :username)`

	_, err := h.db.NamedExec(query, u)
	if err != nil {
		log.Panicln(err)
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to create the user"},
		)
	}

	fmt.Println(c.Request().Body)
	return c.JSON(http.StatusOK, "ok")
}

// @Summary Get a user
// @Description create a new task
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} User "User successfully retrieved"
// @Router /users/{id} [GET]
func (h *UserHandlers) GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
