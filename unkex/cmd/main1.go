package main


import (
	"fmt"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExRate struct {
	Uah float32 `json:"uah"`
	Usd float32 `json:"usd"`
	Id  string  `json:"id"`
}


var todos = []ExRate{
	{Uah: 24.24, Usd: 23.23},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo ExRate

	fmt.Println("jdlkfjsldkjf");

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	fmt.Println(newTodo);

	todos = append(todos, newTodo);
	context.IndentedJSON(http.StatusCreated, newTodo);
}

func getTodo(context *gin.Context) {
	id := context.Param("id");

	item, err := getTodoItem(id);

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"});
		return;
	}

	context.IndentedJSON(http.StatusOK, item);
}

func getTodoItem(id string) (*ExRate, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil;
		}
	}
	return nil, errors.New("Could not find item by given Id");
}


func main() {
	router := gin.Default()
	router.GET("/todos", getTodos);
	router.GET("/todos/:id", getTodo);
	router.POST("/todos", addTodo);

	fmt.Println("Hey world");

	router.Run("localhost:9090");
}