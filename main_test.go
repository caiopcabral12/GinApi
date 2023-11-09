package main

import (
	ct "GinGoApi/controllers"
	db "GinGoApi/database"
	md "GinGoApi/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	req, _ := http.NewRequest("GET", "/students/cpf/12345678900", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	fmt.Println(response.Body)
}

func TestIDUsers(t *testing.T) {
	db.DbConnect()
	MockStudent()
	defer MockStudentDelete()
	r := RoutesTest()
	r.GET("/students/:id", ct.FindStudent)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var mockStudent md.Student
	json.Unmarshal(response.Body.Bytes(), &mockStudent)

	assert.Equal(t, "TestName", mockStudent.Name)
	assert.Equal(t, "TestCollege", mockStudent.Degree)
	assert.Equal(t, "12345678900", mockStudent.Document)
	assert.Equal(t, 32, mockStudent.Age)
	fmt.Println(mockStudent.Name)
	fmt.Println(mockStudent.Degree)
	fmt.Println(mockStudent.Document)
	fmt.Println(mockStudent.Age)

}

func TestDeleteUsers(t *testing.T) {
	db.DbConnect()
	MockStudent()
	r := RoutesTest()
	r.DELETE("/students/:id", ct.DeleteStudent)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	fmt.Println(response.Body)
}

func TestUpdateUsers(t *testing.T) {
	db.DbConnect()
	MockStudent()
	defer MockStudentDelete()
	r := RoutesTest()
	r.PATCH("/students/:id", ct.UpdateStudent)
	student := md.Student{Name: "TestName2", Degree: "TestCollege2", Document: "22345678900", Age: 22}
	jsonValue, _ := json.Marshal(student)
	pathEdit := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathEdit, bytes.NewBuffer(jsonValue))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var mockStudentEdit md.Student
	json.Unmarshal(response.Body.Bytes(), &mockStudentEdit)
	assert.Equal(t, "TestName2", mockStudentEdit.Name)
	assert.Equal(t, "TestCollege2", mockStudentEdit.Degree)
	assert.Equal(t, "22345678900", mockStudentEdit.Document)
	assert.Equal(t, 22, mockStudentEdit.Age)

}
