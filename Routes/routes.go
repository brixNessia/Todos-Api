package routes

import (
	"github.com/gin-gonic/gin"
	CONTROLLER "Hello/Controller/Todos_Controller"
)

var router = gin.Default()

func Routes(){
	r := router.Group("/api/todos")
	{
		r.GET("/", CONTROLLER.GetTodos)
		r.POST("/", CONTROLLER.PostTodos)
	}

	router.Run()
}