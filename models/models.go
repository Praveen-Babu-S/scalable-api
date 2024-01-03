package models

import "gorm.io/gorm"

// Note represents a note model.
type Note struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

// User represents a user model.
type User struct {
	gorm.Model
	Username string `json:"user"`
	Password string `json:"password"`
}

// SharedNote represents a shared note model.
type SharedNote struct {
	gorm.Model
	NoteID uint `json:"note_id"`
	UserID uint `json:"user_id"`
}
