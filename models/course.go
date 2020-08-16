package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Course struct {
	CourseID uuid.UUID `json:"course_id,omitempty"`
	Name     string    `json:"course_name,omitempty"`
	Duration time.Time `db:"date" json:"date" time_format:"unixNano"`
	CourseBy string    `json:"course_instructor,omitempty"`
}
