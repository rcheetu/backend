package main

import (
	"course_feedback/controllers"
	"course_feedback/models"

	"github.com/gin-gonic/gin"
)

/*func init() {
	r := gin.Default()
	//r := router.Get()
	//g := r.Group("/api/v1")
	{

	}
}*/

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	r.GET("/api/v1/", controllers.IndexHandler)

	//Sort by rating
	//r.GET("/api/v1/:course_name", controllers.ShowHandler)

	//Comments
	r.GET("/api/v1/listallfeedback", controllers.ListComments)
	r.POST("/api/v1/addcomment", controllers.CreateComment)
	r.PATCH("/api/v1/edit/:userid/:courseid/:commentid", controllers.EditComment)
	r.DELETE("/api/v1/delete/:userid/:courseid/:commentid", controllers.DeleteComment)

	//Courses
	r.GET("/api/v1/listcourse", controllers.ListCourse)
	r.POST("/api/v1/addcourse", controllers.CreateCourse)

	//RatingsandFeedback
	r.POST("/api/v1/ratecourse/:is_complete/:user_id/:course_id", controllers.RateCourse)
	r.GET("/api/v1/avgrating/:courseid", controllers.AverageRating)
	r.GET("/api/v1/listfeedback/:courseid", controllers.DisplayRatingandFeedback)

	r.Run()
}
