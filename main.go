package main

import(
	"github.com/gin-gonic/gin"
	//"net/http"
	//"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/postgres"
	//"log"
	//"strconv"
	"curdTest/src"
)


func main() {
	route := gin.Default()
	route.GET("/getAllStudents", curd.GetAllStudents)
	route.GET("/getStudentById", curd.GetStudentById)
	route.POST("/createStudentRecord", curd.CreateStudentRecord)
	route.PUT("/updateStudentById", curd.UpdateStudentById)
	route.DELETE("/deleteStudentById", curd.DeleteStudentById)
	route.Run(":8080")
}