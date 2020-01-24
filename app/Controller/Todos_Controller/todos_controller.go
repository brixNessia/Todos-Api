package todos_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, "GET")
}

func PostTodos(c *gin.Context) {
	c.JSON(http.StatusOK, "POST")
}