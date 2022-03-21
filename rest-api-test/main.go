package main

import (
	"rest-api-test/server"
)

func main() {
	r := server.GetRouter()
	r.Run(":8080")
}