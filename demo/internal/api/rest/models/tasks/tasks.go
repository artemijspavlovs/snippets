package tasks

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"

	"github.com/artemijspavlovs/snippets/demo/internal/utils"
)


type StringArray []string

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

type TaskHandlers struct {
	db *sqlx.DB
}

func NewHandlers(db *sqlx.DB) *TaskHandlers {
	return &TaskHandlers{db}
}

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

func (h *TaskHandlers) UpdateTask(c echo.Context) error {
	id := map[string]any{"id": c.Param("id")}

	query := `
		UPDATE tasks 
		SET title = 'new title'
		WHERE id = :id
	`
	result, err := h.db.NamedExec(query, id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "Failed to update the task"},
		)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "Failed to update the task"},
		)
	}

	if affected == 0 {
		return c.JSON(http.StatusNotFound,
			map[string]string{"message": "Task with the requested ID does not exist"})
	}

	return c.JSON(http.StatusOK, id) // return the updated task
}
