package routes

import (
	ct "GinGoApi/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("assets", "./assets")
	r.GET("/students", ct.ShowStudents)
	r.GET("/index", ct.IndexPage)
	r.GET("/students/:id", ct.FindStudent)
	r.GET("/:name", ct.Greetings)
	r.POST("/students", ct.NewStudent)
	r.DELETE("/students/:id", ct.DeleteStudent)
	r.PATCH("/students/:id", ct.UpdateStudent)
	r.GET("/students/cpf/:cpf", ct.CPFSearchStudent)
	r.NoRoute(ct.UnknownRoute)

	r.Run(":3333")

}
