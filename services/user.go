package services

import starter "github.com/jerray/gopress-kick-starter"

type UserService struct {
	db *DatabaseService
}

func NewUserService(db *DatabaseService) *UserService {
	return &UserService{db}
}

func (u *UserService) CreateUser(user *starter.User) error {
	return u.db.Create(user).Error
}

func (u *UserService) GetUserByID(id uint64) (*starter.User, error) {
	user := new(starter.User)
	err := u.db.Model(user).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
