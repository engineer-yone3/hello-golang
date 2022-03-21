package controller

import (
	"github.com/gin-gonic/gin"
)

type Sample struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Comment string `json:"comment"`
}


func IndexAction(c *gin.Context) {
	sample := Sample {
		Id: 100,
		Name: "テストなまえ",
		Comment: "この人はクレーマーです",
	}
	// c.String(200, "Hello, World")
	c.JSON(200, sample)
}