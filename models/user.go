package models

import "time"

type User struct {
	ID               string    `json:"user_id,omitempty"`
	Name             string    `json:"name,omitempty"`
	CourseEnrolledID string    `json:"course_id,omitempty"`
	EnrolledOn       time.Time `db:"enrolled_date" json:"date"`
	IsComplete       bool      `db:"is_complete" json:"is_complete"`
}
