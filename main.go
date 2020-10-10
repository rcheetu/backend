package main

import (
	"course_feedback/controllers"
	"course_feedback/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.Init()

	r.GET("/api/v1", controllers.IndexHandler)

	//Courses
	r.GET("/api/v1/listcourse", controllers.ListCourse)
	r.POST("/api/v1/addcourse", controllers.CreateCourse)

	//Comments
	r.GET("/api/v1/listallfeedback/:course_id", controllers.ListComments)
	r.POST("/api/v1/addcomment/:course_id", controllers.CreateComment)
	r.PATCH("/api/v1/updatecomment", controllers.EditComment)
	r.DELETE("/api/v1/deletecomment", controllers.DeleteComment)

	//RatingsandFeedback
	r.POST("/api/v1/ratecourse/:user_id/:course_id", controllers.RateCourse)
	r.GET("/api/v1/avgrating/:course_id", controllers.AverageRating)
	r.GET("/api/v1/fetchratingfeedback/:course_id", controllers.DisplayRatingandFeedback)

	//Sort Ratings
	r.GET("/api/v1/sortratings/:course_name", controllers.ShowHandler)
	r.GET("/api/v1/SortByHighestRating", controllers.SortByHighestRating)
	r.GET("/api/v1/SortByDate", controllers.SortByDate)
	r.GET("/api/v1/SortByLowestRating", controllers.SortByLowestRating)

	r.Run()
}
