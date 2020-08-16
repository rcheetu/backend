package controllers

import (
	"course_feedback/models"
	"course_feedback/utils"
	"course_feedback/utils/paginations"
	"fmt"
	"net/http"
	"strconv"

	//"time"
	"github.com/biezhi/gorm-paginator/pagination"

	//"github.com/PrinceNorin/monga/utils/paginations"
	//"course_feedback/utils/paginations/pagging"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CreateFeedbackInput struct {
	CommentID  uuid.UUID `db:"comment_id" json:"comment_id,omitempty"`
	CourseID   uuid.UUID `db:"course_id" json:"course_id,omitempty"`
	UserID     uuid.UUID `db:"user_id" json:"user_id,omitempty"`
	Rating     int       `json:"rating" check(max(rating_id)):"5"`
	Comment    string    `json:"comment"`
	IsComplete bool      `json:"is_complete"`
}



type UpdateFeedbackInput struct {
	CommentID  uuid.UUID `json:"comment_id"`
	CourseID   uuid.UUID `json:"course_id"`
	UserID     uuid.UUID `json:"user_id"`
	Comment    string    `json:"comment"`
	IsComplete bool      `json:"is_complete"`
}

// Get list of comments
func ListComments(c *gin.Context) {
	var feedback []models.Feedback

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
	}, &feedback)
	c.JSON(200, paginator)

}

//Find list of courses
func FindAll(page, limit int, orderBy []string) (*paginations.Pagination, error) {
	var feedback []models.Feedback
	p := &paginations.Param{
		DB:      models.DB.Where("course_id > ?", 0),
		Page:    page,
		Limit:   limit,
		OrderBy: orderBy,
	}
	return paginations.Pagging(p, &feedback)
}

//Find a single course
func Find(course_name string) (*models.Feedback, error) {
	var feedback *[]models.Feedback
	if err := models.DB.Find(&feedback, course_name).Error; err != nil {
		return nil, err
	}
	//Check here
	return nil, nil

}

//Get order and page params
func IndexHandler(c *gin.Context) {
	orderBy := utils.GetOrderParam(c)
	page, limit := utils.GetPageParam(c)

	res, err := FindAll(page, limit, orderBy)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func ShowHandler(c *gin.Context) {
	var input CreateFeedbackInput
	course_name := c.Param(input.Comment)
	//fmt.Println("val from Param ", val)
	//course_id := utils.GetIntParam("course_id", c)
	//fmt.Println("course_id from GetIntParam ", course_id)
	res, err := Find(course_name)
	fmt.Println("res from Find ", res)
	if err != nil {
		fmt.Println("ERROR FROM Find")
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create a comment
func CreateComment(c *gin.Context) {
	// Validate input
	var input CreateFeedbackInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment_id := uuid.NewV4()
	course_id := uuid.NewV4()
	user_id := uuid.NewV4()

	if input.Rating < 1 || input.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating needs to be from 1-5"})
		return
	}

	// Add comment
	feedback := models.Feedback{Rating: input.Rating, Comment: input.Comment, CommentID: comment_id, CourseID: course_id, UserID: user_id, IsComplete: true}
	models.DB.Create(&feedback)

	c.JSON(http.StatusOK, gin.H{"data": feedback})
}

// Edit a comment
func EditComment(c *gin.Context) {

	// Validate input
	var input UpdateFeedbackInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var feedback models.Feedback
	//Change only 1 record
	if err := models.DB.Where("user_id = ? AND course_id = ? AND comment_id = ? ", c.Param("user_id"), c.Param("course_id"), c.Param("comment_id")).First(&feedback).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from Edit": "Record not found!"})
		return
	}

	models.DB.Model(&feedback).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": feedback})
}

// Delete a comment
func DeleteComment(c *gin.Context) {
	// Get model if exist
	//fmt.Println(c.Param("comment_id"))
	var feedback models.Feedback
	if err := models.DB.Where("user_id = ? AND course_id = ? AND comment_id = ?", c.Param("user_id"), c.Param("course_id"), c.Param("comment_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from Delete": "Record not found!"})
		return
	}

	models.DB.Delete(&feedback)

	c.JSON(http.StatusOK, gin.H{"data": true})
}




