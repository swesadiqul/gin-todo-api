package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: "1", Title: "Learn Go", Completed: false},
	{Id: "2", Title: "Write API", Completed: true},
}

// GET /todos -get all todos items
func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		// context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoByID(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func updateTodo(context *gin.Context) {
	id := context.Param("id")
	var todo todo
	if err := context.BindJSON(&todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	updatedTodo, err := updateTodoByID(id, todo)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, updatedTodo)
}

func updateTodoByID(id string, updatedTodo todo) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			todos[i] = updatedTodo
			return &updatedTodo, nil
		}
	}
	return nil, fmt.Errorf("Todo not found")
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func main() {
	router := gin.Default()
	// To get all todos items
	router.GET("/todos", getAllTodos)

	// To add a new todo item
	router.POST("/todos", addTodo)

	// To get a todo item
	router.GET("/todos/:id", getTodo)

	// To update a todo item
	router.PUT("/todos/:id", updateTodo)

	// To delete a todo item
	router.DELETE("/todos/:id", deleteTodo)

	router.Run(":8080")

	// To delete a todo item
	router.DELETE("/todos/:id", deleteTodo)

	router.Run(":8080")

	// Start the server
	router.Run("localhost:8080")
}
