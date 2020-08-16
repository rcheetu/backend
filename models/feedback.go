package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Feedback struct {
	CommentID  uuid.UUID `db:"comment_id" json:"comment_id" default:"uuid_generate_v4()"`
	CourseID   uuid.UUID `db:"course_id" json:"course_id" default:"uuid_generate_v4()"`
	UserID     uuid.UUID `db:"user_id" json:"user_id" default:"uuid_generate_v4()"`
	Rating     int       `db:"rating_id" json:"rating"`
	Date       time.Time `db:"date" json:"date" time_format:"unixNano"`
	Comment    string    `db:"comment" json:"comment,omitempty"`
	IsComplete bool      `db:"is_complete" json:"is_complete"`
}
