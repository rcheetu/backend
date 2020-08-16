package controllers

import (
	"course_feedback/models"
	"net/http"
	"strconv"

	//"time"
	"github.com/biezhi/gorm-paginator/pagination"

	//"github.com/PrinceNorin/monga/utils/paginations"
	//"course_feedback/utils/paginations/pagging"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CreateCourseInput struct {
	CourseID uuid.UUID `db:"course_id" json:"course_id,omitempty"`
	Name     string    `db:"course_name" json:"course_name,omitempty"`
	Duration string    `db:"duration" json:"duration,omitempty"`
	CourseBy string    `json:"course_by" json:"course_by"`
}

// Create a course
func CreateCourse(c *gin.Context) {
	// Validate input
	var input CreateCourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course_id := uuid.NewV4()

	// Add course
	course := models.Course{CourseBy: input.CourseBy, CourseID: course_id, Name: input.Name}
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
		//Check number of courses
		DB:    models.DB.Where("course_id > ?", 0),
		Page:  page,
		Limit: limit,
		//OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &course)
	c.JSON(200, paginator)

}
