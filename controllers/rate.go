package controllers

import (
	"course_feedback/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Ratings struct {
	CommentID  uuid.UUID `db:"comment_id" json:"comment_id,omitempty"`
	CourseID   uuid.UUID `db:"course_id" json:"course_id,omitempty"`
	UserID     uuid.UUID `db:"user_id" json:"user_id,omitempty"`
	Rating     int       `json:"rating" check(max(rating_id)):"5"`
	Comment    string    `json:"comment"`
	IsComplete bool      `db:"is_complete" json:"is_complete"`
}

//Rate a course and comment
func RateCourse(c *gin.Context) {
	// Validate input
	var input Ratings
	var feedback models.Feedback

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Fetch userid and courseid
	//courseid := c.Param("course_id")
	//userid := c.Param("user_id")
	rating := input.Rating
	//state := c

	//fmt.Println("courseid", courseid, "userid:", userid, "rating:", rating, "state:", state)
	//Check if course is complete for a particular user
	if c.Param("is_complete") == "true" {
		fmt.Println("Checking ratings")
		//Check if rating is in between 1-5
		if rating < 1 || rating > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Rating needs to be from 1-5"})
			return
		}
		if err := models.DB.Where("user_id = ? AND course_id = ? AND is_complete = true ", c.Param("user_id"), c.Param("course_id")).First(&feedback).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Record not found": err})
			return
		}
		if err := models.DB.Model(&feedback).Update("rating", rating).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error while adding ratings": err})
			return

		}
		c.JSON(http.StatusOK, gin.H{"data": feedback})
	}
}

//Calculate & display average rating
func AverageRating(c *gin.Context) {
	var feedback *[]models.Feedback
	var input Ratings
	fmt.Println(c.Param("course_id"))

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//select avg(rating) from feedback where course_id = ''
	if err := models.DB.Select("AVG(rating)").Where("course_id = ?", c.Param("course_id")).First(&feedback).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while fetching record": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": feedback})

}

//Display rating with feedback
func DisplayRatingandFeedback(c *gin.Context) {
	var feedback *[]models.Feedback
	var input Ratings
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Check for ratings only if comments are present
	if err := models.DB.Select("rating, feedback").Where("course_id = ? AND comment is not NULL", input.CourseID).Find(&feedback).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while fetching record": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": feedback})

}
