package controllers

import (
	"course_feedback/models"
	"course_feedback/utils"
	"course_feedback/utils/paginations"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Ratings struct {
	CommentID  uuid.UUID `db:"comment_id" json:"comment_id,omitempty"`
	CourseID   uuid.UUID `db:"id" json:"course_id,omitempty"`
	UserID     uuid.UUID `db:"user_id" json:"user_id,omitempty"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment" gorm:"size:255"`
	IsComplete bool      `db:"is_complete" json:"is_complete"`
}

//Rate a course and comment
func RateCourse(c *gin.Context) {
	// Validate input
	var input Ratings
	//var feedback models.Feedback

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rating := input.Rating
	is_complete := input.IsComplete
	if is_complete {
		fmt.Println("Checking ratings")
		//Check if rating is in between 1-5
		if rating < 1 || rating > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Rating needs to be from 1-5"})
			return
		}
		if err := models.DB.Table("feedbacks").Where("user_id = ? AND course_id = ? AND is_complete = true ", c.Param("user_id"), c.Param("course_id")).Update("rating", rating).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Record not found": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "true"})
	}
}

//Calculate & display average rating for a particular course_id
func AverageRating(c *gin.Context) {
	var avg_rating float64

	course_id := c.Param("course_id")
	fmt.Println(course_id)

	//SELECT avg(rating)::float, comment FROM "feedbacks"  WHERE (course_id = '') GROUP BY comment
	row := models.DB.Table("feedbacks").Select("avg(rating)").Where("course_id = ?", course_id).Row()
	row.Scan(&avg_rating)

	c.JSON(http.StatusOK, gin.H{"avg_rating": avg_rating})

}

//Display ratings and feedback for a course selected
func DisplayRatingandFeedback(c *gin.Context) {
	var feedback []models.Feedback
	course_id := c.Param("course_id")
	fmt.Println(course_id)

	//Fetch details corresponding to course_id
	if err := models.DB.Select("rating,comment").Where("course_id = ? ", course_id).Find(&feedback).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Record not found": err})
		return
	}
	for _, data := range feedback {
		//Return Rating and Feedback
		c.JSON(http.StatusOK, gin.H{"Rating": data.Rating, "Feedback": data.Comment})
	}

}

//Get order and page params
func SortByHighestRating(c *gin.Context) {
	orderBy := utils.GetOrderParams(c)
	page, limit := utils.GetPageParam(c)

	res, err := SortRatings(page, limit, orderBy)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func SortByDate(c *gin.Context) {
	orderBy := utils.GetOrderParamsByDate(c)
	page, limit := utils.GetPageParam(c)

	res, err := SortRatings(page, limit, orderBy)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func SortByLowestRating(c *gin.Context) {
	orderBy := utils.GetOrderParamsByLowest(c)
	page, limit := utils.GetPageParam(c)

	res, err := SortRatings(page, limit, orderBy)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

//Find list of comments
func SortRatings(page, limit int, orderBy []string) (*paginations.Pagination, error) {
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
