package model

import (
	"time"

	"gorm.io/gorm"
)

type DBUser struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DBPost struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string
	Content   string
	AuthorID  uint
	Author    DBUser         `gorm:"foreignKey:AuthorID"`
	Comments  []DBComment    `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DBComment struct {
	ID        uint           `gorm:"primaryKey"`
	Text      string
	PostID    uint
	Post      DBPost         `gorm:"foreignKey:PostID"`
	AuthorID  uint
	Author    DBUser         `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
