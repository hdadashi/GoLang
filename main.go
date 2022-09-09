package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// request body template
type Object struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// request body
var Objects = []Object{
	{ID: "1", Item: "TEST Item 1", Completed: false},
	{ID: "2", Item: "TEST Item 2", Completed: false},
	{ID: "3", Item: "TEST Item 3", Completed: true},
}

// make Objects a JSON file
func getObjects(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Objects)
}

// make newObject to JSON file
func addObjects(context *gin.Context) {
	var newObject Object
	err := context.BindJSON(&newObject)
	if err != nil {
		return
	}
	Objects = append(Objects, newObject)
	context.IndentedJSON(http.StatusCreated, newObject)
}

func SearchByID(id string) (*Object, error, int) {
	for i, t := range Objects {
		if t.ID == id {
			return &Objects[i], nil, i
		}
	}
	return nil, errors.New("error is not nil"), 0
}

func getByID(context *gin.Context) {
	id := context.Param("id")
	Object, err, _ := SearchByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found!"})
		return
	}
	context.IndentedJSON(http.StatusOK, Object)
}

func patchByID(context *gin.Context) {
	var patchByID Object
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
	if i == len(Objects)-1 {
		Objects = Objects[:len(Objects)-1] //Truncate slice
	} else {
		Objects[i] = Objects[len(Objects)-1] //Copy last element to index i
		Objects = Objects[:len(Objects)-1]   //Truncate slice
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Record by ID= " + id + " has been REMOVED!"})
}

func main() {
	router := gin.Default()
	router.GET("/", getObjects)
	router.GET("/:id", getByID)
	router.PATCH("/:id", patchByID)
	router.POST("/", addObjects)
	router.DELETE("/:id", deleteBYID)
	router.Run("localhost:8080")
}
