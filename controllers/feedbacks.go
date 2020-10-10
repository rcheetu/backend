package controllers

import (
	"course_feedback/models"
	"course_feedback/utils"
	"course_feedback/utils/paginations"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/biezhi/gorm-paginator/pagination"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CreateFeedbackInput struct {
	CommentID  uuid.UUID `db:"comment_id" json:"comment_id,omitempty"`
	CourseID   uuid.UUID `db:"course_id" json:"course_id,omitempty"`
	UserID     uuid.UUID `db:"user_id" json:"user_id,omitempty"`
	Rating     int       `json:"rating" check(max(rating_id)):"5"`
	Comment    string    `json:"comment" gorm:"size:500"`
	IsComplete bool      `json:"is_complete"`
}

type UpdateFeedbackInput struct {
	CommentID  uuid.UUID `db:"comment_id" json:"comment_id,omitempty"`
	CourseID   uuid.UUID `db:"course_id" json:"course_id,omitempty"`
	UserID     uuid.UUID `db:"user_id" json:"user_id,omitempty"`
	Comment    string    `json:"comment" gorm:"size:500"`
	IsComplete bool      `json:"is_complete"`
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

//Find list of comments
func FindAll(page, limit int, orderBy []string) (*paginations.Pagination, error) {
	var feedback []models.Feedback
	p := &paginations.Param{
		DB: models.DB,
		//.Where("course_id > ?", 0),
		Page:    page,
		Limit:   limit,
		OrderBy: orderBy,
	}
	return paginations.Pagging(p, &feedback)
}

// Get list of comments
func ListComments(c *gin.Context) {

	var feedback []models.Feedback

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	course_id := c.Param("course_id")
	fmt.Println(course_id)

	// Validate input
	//var input CreateFeedbackInput
	//if err := c.ShouldBindJSON(&input); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//Pagination functionality
	paginator := pagination.Paging(&pagination.Param{
		//Check number of courses
		DB: models.DB.Where("course_id = ?", course_id),
		//DB: models.DB.Where()
		Page:  page,
		Limit: limit,
		//OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &feedback)
	c.JSON(200, paginator)

}

// Create a comment
func CreateComment(c *gin.Context) {
	// Validate input
	var input CreateFeedbackInput
	var feedback models.Feedback
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Convert course_id to uuid
	course_id := c.Param("course_id")
	courseid, _ := uuid.FromString(course_id)

	comment_id := uuid.NewV4()
	//course_id := uuid.NewV4()
	user_id := uuid.NewV4()

	if len(input.Comment) < 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comments should be of minimum 20 characters"})
		return
	}

	if len(input.Comment) > 500 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comments should be of maximum 500 characters"})
		return
	}

	if input.Rating < 1 || input.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating needs to be from 1-5"})
		return
	}

	//Check for course_id
	//if err := models.DB.Where("course_id = ? ", c.Param("course_id")).First(&feedback).Error; err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error while inserting comment": "Course not found!"})
	//	return
	//}
	//models.DB.Model(&feedback)

	//Add comment
	feedback = models.Feedback{Rating: input.Rating, CourseID: courseid, Comment: input.Comment, CommentID: comment_id, UserID: user_id, Date: time.Now(), IsComplete: true}
	models.DB.Create(&feedback)

	c.JSON(http.StatusOK, gin.H{"data": feedback})
}

// Edit a comment
func EditComment(c *gin.Context) {

	// Validate input
	var input UpdateFeedbackInput
	//var feedback models.Feedback

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := input.Comment
	user_id := input.UserID
	course_id := input.CourseID
	comment_id := input.CommentID

	fmt.Println(comment, user_id, course_id, comment_id)

	if len(input.Comment) < 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comments should be of minimum 20 characters"})
		return
	}

	if len(input.Comment) > 500 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comments should be of maximum 500 characters"})
		return
	}

	//Update record
	if err := models.DB.Table("feedbacks").Where("user_id = ? AND course_id = ? AND comment_id = ? ", user_id, course_id, comment_id).Update("comment", comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from Edit": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "true", "comment": comment})
}

// Delete a comment
func DeleteComment(c *gin.Context) {

	var input UpdateFeedbackInput
	var feedback []models.Feedback

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commentid := input.CommentID
	userid := input.UserID
	courseid := input.CourseID

	if err := models.DB.Where("user_id = ? AND course_id = ? AND comment_id = ?", userid, courseid, commentid).Delete(&feedback).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from Delete": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
