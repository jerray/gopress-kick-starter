package starter

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

type PostService interface {
	CreatePost(post *Post) error
	GetPostByID(id uint64) (*Post, error)
	ListPostsByUser(user *User) ([]*Post, error)
}
