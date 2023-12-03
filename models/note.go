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

func (note *Note) Update(name string, content string) {
	note.Name = name
	note.Content = content
	DB.Save(note)
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

func NotesMarkDelete(id uint64) {
	// UPDATE notes SET deleted_at=<Current Time> WHERE id = <id> and user_id = <user_id>
	DB.Where("id = ?", id).Delete(&Note{})
}
