package models

import (
	v2 "gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string `json:"NAME" validate:"nonzero"`
	Degree   string `json:"DEGREE" validate:"nonzero"`
	Document string `json:"CPF" validate:"len=11, regexp=^[0-9]*$"`
	Age      int    `json:"AGE" validate:"nonzero"`
}

func Validate(student *Student) error {
	if err := v2.Validate(student); err != nil {
		return err
	}

	return nil
}
