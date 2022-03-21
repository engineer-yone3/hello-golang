package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
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

func HandleIndexAction(w http.ResponseWriter, r *http.Request) {
	sample := Sample {
		Id: 100,
		Name: "テストなまえその２",
		Comment: "この人はクレーマーです（う",
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(&sample)
	fmt.Fprintf(w, buf.String())
}