package server

import (
	"github.com/gin-gonic/gin"
	"rest-api-test/controller"
	"net/http"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/index", controller.IndexAction)

	return router
}

func GetRouterHandle() {
	http.HandleFunc("/index", controller.HandleIndexAction)
}