package main

import (
	ct "GinGoApi/controllers"
	db "GinGoApi/database"
	md "GinGoApi/models"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func RoutesTest() *gin.Engine {
	routes := gin.Default()
	return routes
}

func MockStudent() {
	student := md.Student{Name: "TestName", Degree: "TestCollege", Document: "12345678900", Age: 32}
	db.DB.Create(&student)

	ID = int(student.ID)
}

func MockStudentDelete() {
	var student md.Student
	db.DB.Delete(&student, ID)
}

func TestStatusCode(t *testing.T) {
	r := RoutesTest()
	r.GET("/:name", ct.Greetings)
	req, _ := http.NewRequest("GET", "/test", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)

	responseMock := `{"Greetings from Gin!":"Mr Caio Welcome to this API"}`

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, responseMock, string(responseBody))
}

func TestAllUsers(t *testing.T) {
	db.DbConnect()
	MockStudent()
	defer MockStudentDelete()
	r := RoutesTest()
	r.GET("/students", ct.ShowStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	fmt.Println(response.Body)
}

func TestCPFUsers(t *testing.T) {
	db.DbConnect()
	MockStudent()
	defer MockStudentDelete()
	r := RoutesTest()
	r.GET("/students/cpf/:cpf", ct.CPFSearchStudent)
	req, _ := http.NewRequest("GET", "/students/cpf/12347678900", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	fmt.Println(response.Body)
}
