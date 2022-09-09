package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// request body template
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// request body
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: true},
}

// make todos a JSON file
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// make newTodo to JSON file
func addTodos(context *gin.Context) {
	var newTodo todo
	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func SearchByID(id string) (*todo, error, int) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil, i
		}
	}
	return nil, errors.New("error is not nil"), 0
}

func getByID(context *gin.Context) {
	id := context.Param("id")
	todo, err, _ := SearchByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found!"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func patchByID(context *gin.Context) {
	var patchByID todo
	err1 := context.BindJSON(&patchByID)
	if err1 != nil {
		return
	}
	id := context.Param("id")
	data, err, _ := SearchByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found!"})
		return
	}
	data.Completed = patchByID.Completed
	data.Item = patchByID.Item
	patchByID.ID = id
	context.IndentedJSON(http.StatusOK, patchByID)
}

func deleteBYID(context *gin.Context) {
	id := context.Param("id")
	_, err, i := SearchByID(id)
	fmt.Print(i)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found!"})
		return
	}
	if i == len(todos)-1 {
		todos = todos[:len(todos)-1] //Truncate slice
	} else {
		todos[i] = todos[len(todos)-1] //Copy last element to index i
		todos = todos[:len(todos)-1]   //Truncate slice
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Record by ID= " + id + " has been REMOVED!"})
}

func main() {
	router := gin.Default()
	router.GET("/", getTodos)
	router.GET("/:id", getByID)
	router.PATCH("/:id", patchByID)
	router.POST("/", addTodos)
	router.DELETE("/:id", deleteBYID)
	router.Run("localhost:8080")
}
