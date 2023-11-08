package controllers

import (
	db "GinGoApi/database"
	md "GinGoApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) {
	var student []md.Student
	db.DB.Find(&student)

	c.JSON(200, student)
}

func FindStudent(c *gin.Context) {
	var student md.Student
	id := c.Params.ByName("id")
	db.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found!"})
		return
	}
	c.JSON(200, student)
}

func CPFSearchStudent(c *gin.Context) {
	var student md.Student
	cpf := c.Param("cpf")
	db.DB.Where(&md.Student{Document: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "CPF not found!"})
		return
	}
	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	var student md.Student
	id := c.Params.ByName("id")
	db.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"Success": "Student deleted!"})
}

func UpdateStudent(c *gin.Context) {
	var student md.Student
	id := c.Params.ByName("id")
	db.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := md.Validate(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Student updated!"})
}

func Greetings(c *gin.Context) {
	//name := c.Params.ByName("NAME")
	c.JSON(200, gin.H{
		"Greetings from Gin!": "Mr Caio Welcome to this API",
	})

}

func NewStudent(c *gin.Context) {
	var student md.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := md.Validate(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
