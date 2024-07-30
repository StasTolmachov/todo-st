package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()
	router.GET("/todos", GetTodos)

	router.Run(":8080")
}
