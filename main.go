package main

import (
	"rest-api-test/server"
	// "rest-api-test/controller"
	"net/http"
)

func main() {
	// r := server.GetRouter()
	// r.Run(":8080")
	
	server.GetRouterHandle()
	http.ListenAndServe(":8080", nil)
}