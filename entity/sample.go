package entity

import (
	"gorm.io/gorm"
)

// 「gorm.Model」はid, created_at, updated_at, deleted_atに変わる
type SampleEntity struct {
	Title string
	Content string
	gorm.Model
}