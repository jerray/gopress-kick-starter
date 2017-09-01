package models

import "time"

type Post struct {
	ID        uint64 `gorm:"column:id"`
	UserID    uint64 `gorm:"column:user_id"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Post) TableName() string {
	return "posts"
}
