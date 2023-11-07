package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name     string `json:"NAME"`
	Degree   string `json:"DEGREE"`
	Document string `json:"CPF"`
	Age      int    `json:"AGE"`
}
