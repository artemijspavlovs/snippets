package tasks

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/artemijspavlovs/snippets/demo/internal/utils"
)

type TaskStatus string

// const (
//     StatusPending   TaskStatus = "pending"
//     StatusInProgress TaskStatus = "in_progress"
//     StatusCompleted TaskStatus = "completed"
//	   // etc.
// )

type StringArray []string

// parsePostgreSQLArray function    transforms the default array type retrieved from a PostgreSQL database
// ( {value, value, value} ) into a go string array []string{value, value, value}
func parsePostgreSQLArray(s string) ([]string, error) {
	trimmed := strings.Trim(s, "{}")
	if trimmed == "" {
		return StringArray{}, nil
	}
	list := strings.Split(trimmed, ",")

	return list, nil
}

// Scan method    extends the StringArray type with the ability to natively retrieve PostgreSQL rows that
// include string arrays as their column values
func (sa *StringArray) Scan(src any) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("source is not a string")
	}

	a, err := parsePostgreSQLArray(str)
	if err != nil {
		return err
	}

	*sa = a
	return nil
}

// Task struct   represents an instance of task and maps the struct fields to json and database property names.
// Additionally, it introduces standardization through custom times where applicable. It applies omitempty tag
// to the fields that can be either left out or will be populated automatically during creation. for custom type
// information, refer to the respectful docstrings
type Task struct {
	ID          string      `json:"id"                    db:"id"`
	Title       string      `json:"title"                 db:"title"`
	Description string      `json:"description,omitempty" db:"description"`
	Status      TaskStatus  `json:"status"                db:"status"`
	CreatedAt   time.Time   `json:"created_at"            db:"created_at"`
	CreatedBy   string      `json:"created_by"            db:"created_by"`
	UpdatedAt   time.Time   `json:"updated_at"            db:"updated_at"`
	UpdatedBy   string      `json:"updated_by"            db:"updated_by"`
	DueDate     time.Time   `json:"due_date,omitempty"    db:"due_date"`
	Tags        StringArray `json:"tags,omitempty"        db:"tags"`
}

// Generate method    is a helper method that generates a dummy representation of a Task struct with pre-filled
// fields using fake/random data
func (t *Task) Generate() Task {
	t.ID = faker.UUIDDigit()
	t.Title = faker.Sentence()
	t.Description = faker.Sentence()
	t.Status = TaskStatus("todo") // change to valid states
	t.Tags = []string{faker.Word()}

	t.CreatedAt = time.Unix(time.Now().Unix(), 0)
	t.CreatedBy = faker.Word()
	t.UpdatedAt = utils.GenerateRandomTimestampFromUnix()
	t.UpdatedBy = faker.Word()
	t.DueDate = utils.GenerateRandomTimestampFromUnix()

	j, _ := json.Marshal(t)
	fmt.Printf("%v\n\n", string(j))
	return *t
}

// TaskHandlers struct    provides a proxy for reusing the instance of sqlx.DB in task specific http handlers
type TaskHandlers struct {
	db *sqlx.DB
}

// NewHandlers function    constructor function for TaskHandlers
func NewHandlers(db *sqlx.DB) *TaskHandlers {
	return &TaskHandlers{db}
}

// @Summary List tasks
// @Description get tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} Task
// @Router /tasks [GET]
func (h *TaskHandlers) GetAllTasks(c echo.Context) error {
	tasks := []Task{}
	query := `SELECT * FROM tasks`

	err := h.db.Select(&tasks, query)
	if err != nil {
		fmt.Println(err)
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to retrieve tasks"},
		)
	}

	return c.JSON(http.StatusOK, tasks)
}

// @Summary Create task
// @Description create a new task
// @Accept  json
// @Produce  json
// @Success 200 {object} Task "Task successfully created"
// @Router /tasks [POST]
func (h *TaskHandlers) CreateTask(c echo.Context) error {
	t := Task{}
	t.Generate()

	query := `INSERT INTO tasks (id, title, description, status, created_at, created_by, updated_at, updated_by, due_date, tags)
	VALUES (:id, :title, :description, :status, :created_at, :created_by, :updated_at, :updated_by, :due_date, :tags)`

	_, err := h.db.NamedExec(query, t)
	if err != nil {
		log.Println(err)
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to create the task"},
		)
	}

	return c.JSON(http.StatusCreated, t)
}

// @Summary Create task
// @Description create a new task
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {object} Task "Task successfully retrieved"
// @Router /tasks/{id} [GET]
func (h *TaskHandlers) GetTaskByID(c echo.Context) error {
	id := c.Param("id")
	var task Task
	query := `SELECT * FROM tasks WHERE id = $1`

	err := h.db.Get(&task, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				map[string]string{"message": "Task with the requested ID does not exist"})
		}

		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to retrieve task with this ID"},
		)
	}

	return c.JSON(http.StatusOK, task)
}

// @Summary Create task
// @Description create a new task
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {object} Task "Task successfully retrieved"
// @Router /tasks/{id} [PUT]
func (h *TaskHandlers) UpdateTask(c echo.Context) error {
	// TODO: handle the update and get queries in a single transaction
	id := map[string]any{"id": c.Param("id")}

	query := `
		UPDATE tasks
		SET title = 'new title'
		WHERE id = :id
	`
	result, err := h.db.NamedExec(query, id)
	if err != nil {
		fmt.Println(err)
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "Failed to update the task"},
		)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "Failed to update the task"},
		)
	}

	if affected == 0 {
		return c.JSON(http.StatusNotFound,
			map[string]string{"message": "Task with the requested ID does not exist"})
	}

	var task Task
	query = `SELECT * FROM tasks WHERE id = $1`

	err = h.db.Get(&task, query, id["id"])
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				map[string]string{"message": "Task with the requested ID does not exist"})
		}

		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to retrieve task with this ID"},
		)
	}

	return c.JSON(http.StatusOK, task) // return the updated task
}
