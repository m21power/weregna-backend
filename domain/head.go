package domain

import "time"

type Head struct {
	ID               int       `json:"id" db:"id"`
	Email            string    `json:"email" db:"email"`
	Password         string    `json:"password" db:"password"`
	ProfilePic       *string   `json:"profile_pic,omitempty" db:"profile_pic"`
	Name             string    `json:"name" db:"name"`
	Role             string    `json:"role" db:"role"`
	TelegramUsername *string   `json:"telegram_username,omitempty" db:"telegram_username"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}

type HeadRepository interface {
	FindHeadByEmail(email string) (*Head, error)
	FindHeadByID(id int) (*Head, error)
	CreateHead(head *Head) error
	UpdateHead(head *Head) error
	DeleteHead(id int) error
	GetHeads() ([]*Head, error)
	AddMyStudent(headID int, studentID int) error
	GetMyStudents(headID int) ([]*StudentModel, error)
}
