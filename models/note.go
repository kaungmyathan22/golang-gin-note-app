package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	UserID    uint64 `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NotesAll() *[]Note {
	var notes []Note
	DB.Where("deleted_at is NULL").Order("updated_at desc").Find(&notes)
	return &notes
}

func NoteCreate(name, content string) *Note {
	entry := Note{Name: name, Content: content}
	DB.Create(&entry)
	fmt.Println(entry)
	return &entry
}

func NotesFind(id uint64) *Note {
	// func NotesFind(user *User, id uint64) *Note {
	var note Note
	DB.Where("id = ?", id).First(&note)
	return &note
}
