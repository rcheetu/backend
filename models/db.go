package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	user := getEnv("PG_USER", "postgres")
	password := getEnv("PG_PASSWORD", "")
	host := getEnv("PG_HOST", "postgres")
	port := getEnv("PG_PORT", "5432")
	database := getEnv("PG_DB", "tasks")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	DB, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("Failed to connect to database", err)
		panic(err)
	}
	fmt.Println("Database connected")
	if !DB.HasTable(&Feedback{}) {
		//Creating tables
		DB.CreateTable(&Feedback{})
		DB.AutoMigrate(&Course{})
		DB.AutoMigrate(&User{})

	}

	// Enable Logger, show detailed log
	DB.LogMode(true)
}

/*
//Connecting to sqlite mock database for now
func ConnectDataBase() {

	db.LogMode(true)
	db.AutoMigrate(&Feedback{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&User{})

	//Set contraints
	db.Debug().Model(&Course{}).AddUniqueIndex("course_pr_key", "id")
	db.Debug().Model(&User{}).AddUniqueIndex("user_pr_key", "id")
	db.Debug().Model(&Feedback{}).AddUniqueIndex("feedback_pr_key", "comment_id")
	//db.Debug().Model(&Course{}).AddUniqueIndex("feedback_pr_key", "course_id")
	//db.Debug().Model(&Course{}).AddForeignKey("course_id", "feedbacks(course_id)", "CASCADE", "CASCADE")

	DB = db
*/
// Enable Logger, show detailed log
//db.LogMode(true)

//}
