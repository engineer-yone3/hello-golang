package entity

import (
	"gorm.io/gorm"
)

type SampleEntity struct {
	gorm.Model
	Title string
	Content string
}