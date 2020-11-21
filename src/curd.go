package curd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	//gorm.Model
	ID  int	`json:"id"`
	Grade int `json:"grade"`
	Name  string	`json:"name"`
	Score float64	`json:"score"`
}

var db *gorm.DB
var err error

var db_name string = "db_name"
var db_host string = "host_name"
var db_user string = "user_name"
//var db_type string = "postgres"
//var db_passwd string = "passwd"


func connDB() {
	db, err = gorm.Open("postgres", "host=" + db_host + " user="+ db_user +" dbname="+ db_name +" sslmode=disable" )//password="+db_passwd)
	if err!=nil{
		log.Fatal("DB Connection ERROR!")
	} else{
		log.Println("DB connected")
	}
	db.AutoMigrate(&Student{})
}


func closeDB() {
	if err = db.Close(); err != nil{
		log.Println("Closing DB Error!")
	}
}

func CreateStudentRecord(c *gin.Context){
	connDB()
	defer closeDB()

	var student Student
	c.ShouldBindJSON(&student)

	if student.ID==0 || student.Grade==0 || student.Name=="" || student.Score==0{
		c.String(http.StatusBadRequest, "missing params\n")
		c.JSON(http.StatusBadRequest,student)
		return
	}

	var preStudent Student

	if db.Where("id = ?", student.ID).First(&preStudent); preStudent.ID!=0{
		c.String(http.StatusBadRequest, "Record for this id already existed\n")
		c.JSON(http.StatusBadRequest, preStudent)
		return
	}

	db.Create(&student)

	c.String(http.StatusOK, "record created\n")
	c.JSON(http.StatusOK, student)
}

func GetAllStudents(c *gin.Context) {
	connDB()
	defer closeDB()

	var students []Student
	db.Find(&students)
	c.JSON(200, students)
}


func GetStudentById(c *gin.Context) {
	connDB()
	defer closeDB()

	var id string = c.Query("id")
	var student Student
	db.First(&student, id)
	if err := db.Where("id = ?", id).First(&student).Error; err != nil {
		//c.AbortWithStatus(404)
		c.String(403, "record not found, no record for id = " + id)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}

func UpdateStudentById(c *gin.Context) {
	connDB()
	defer closeDB()

	var student Student
	var preStudent Student
	c.ShouldBindJSON(&student)

	if student.ID==0 || student.Grade==0 || student.Name=="" || student.Score==0{
		c.String(http.StatusBadRequest, "missing params\n")
		c.JSON(http.StatusBadRequest,student)
		return
	}

	if err := db.Where("id = ?", student.ID).First(&preStudent).Error; err != nil {
		//c.AbortWithStatus(404)
		c.String(403, "record not found, no record for id = " + strconv.Itoa(student.ID))
		fmt.Println(err)
		return
	}

	c.ShouldBindJSON(&student)
	db.Save(&student)
	c.String(http.StatusOK, "Updated successfully from: \n")
	c.JSON(http.StatusOK, preStudent)
	c.String(http.StatusOK, "\nto: \n")
	c.JSON(http.StatusOK, student)

}

func DeleteStudentById(c *gin.Context) {
	connDB()
	defer closeDB()

	var student Student
	c.ShouldBindJSON(&student)

	if student.ID==0{
		c.String(http.StatusBadRequest, "missing params\n")
		c.JSON(http.StatusBadRequest,student)
		return
	}

	var preStudent Student

	if db.Where("id = ?", student.ID).First(&preStudent); preStudent.ID==0{
		c.String(http.StatusBadRequest, "Record for id = "+strconv.Itoa(student.ID)+" not existed\n")
		return
	}

	if err := db.Where("id = ?", student.ID).Delete(&student).Error; err != nil {
		//c.AbortWithStatus(404)
		c.String(http.StatusBadRequest, "delete failed\n")
		c.String(http.StatusBadRequest, err.Error())
		fmt.Println(err)
		return
	}

	c.String(http.StatusOK, "Successfully deleted: \n")
	c.JSON(http.StatusOK, preStudent)
}