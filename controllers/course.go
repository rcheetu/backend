package controllers

import (
	"course_feedback/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CreateCourseInput struct {
	ID         uuid.UUID `db:"id" json:"id,omitempty"`
	CourseName string    `db:"course_name" json:"course_name,omitempty"`
	Duration   int       `db:"duration" json:"duration,omitempty"`
	CourseBy   string    `json:"course_by" json:"course_by"`
}

// Create a course
func CreateCourse(c *gin.Context) {
	// Validate input
	var input CreateCourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uuid.NewV4()

	// Add course
	course := models.Course{ID: id, CourseBy: input.CourseBy, CourseName: input.CourseName, Duration: input.Duration}
	models.DB.Create(&course)

	c.JSON(http.StatusOK, gin.H{"data": course})
}

// Get list of courses
func ListCourse(c *gin.Context) {
	var course []models.Course

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	//Pagination functionality
	paginator := pagination.Paging(&pagination.Param{
		DB: models.DB,
		//.Where("course_id > ?", 0),
		Page:  page,
		Limit: limit,
		//OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &course)
	c.JSON(200, paginator)

}

func ShowHandler(c *gin.Context) {
	var course models.Course
	course_name := c.Param("course_name")
	fmt.Println(course_name)

	if err := models.DB.Where("course_name = ? ", course_name).First(&course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from ShowHandler": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": course})
}
