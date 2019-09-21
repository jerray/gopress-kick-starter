package starter

import "time"

type User struct {
	ID        uint64 `gorm:"column:id"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}

type UserService interface {
	CreateUser(user *User) error
	GetUserByID(id uint64) (*User, error)
}
