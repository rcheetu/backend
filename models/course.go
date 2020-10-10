package models

import (
	uuid "github.com/satori/go.uuid"
)

type Course struct {
	ID uuid.UUID `json:"id,omitempty"`
	//CourseID   uuid.UUID `db:"course_id" json:"course_id,omitempty"`
	CourseName string `json:"course_name,omitempty"`
	Duration   int    `db:"duration" json:"duration"`
	CourseBy   string `json:"course_by,omitempty"`
}
