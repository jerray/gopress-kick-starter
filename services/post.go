package services

import starter "github.com/jerray/gopress-kick-starter"

type PostService struct {
	db *DatabaseService
}

func NewPostService(db *DatabaseService) *PostService {
	return &PostService{db}
}

func (p *PostService) CreatePost(post *starter.Post) error {
	return p.db.Create(post).Error
}

func (p *PostService) GetPostByID(id uint64) (*starter.Post, error) {
	post := new(starter.Post)
	err := p.db.Model(post).Where("id = ?", id).First(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostService) ListPostsByUser(user *starter.User) ([]*starter.Post, error) {
	posts := []*starter.Post{}
	err := p.db.Where("user_id = ?", user.ID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
