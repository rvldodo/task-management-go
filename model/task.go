package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title string `json:"title"`
	Body string `json:"body"`
	Done bool `json:"done"`
}