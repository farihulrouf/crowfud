package handler

import (
	"crowdgo/todo"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"fmt"
)

type todoHandler struct {
	todoService todo.Service
}

func NewTodoHandler(todoService todo.Service) *todoHandler {
	return &todoHandler{todoService}
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	var inputTodo todo.InputTodo

	err := c.ShouldBindJSON(&inputTodo)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			erroMessage := fmt.Sprintf("Error on filed %s, condiiton: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, erroMessage)
			//fmt.Println(err)
			return
		}
		
	}

	//h.todoService.CreateTodo(inputTodo)

	c.JSON(http.StatusOK, gin.H {
		"name": inputTodo.Name,
		//"SubTitle": bookInput.SubTitle,
	})
	
}