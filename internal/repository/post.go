package repository

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

type PostRepository interface {
	Get(id int) (*Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Get(id int) (*Post, error) {
	var post Post
	r.db.First(&post, id)
	if post.ID == 0 {
		return nil, &RecordNotFoundError{}
	}
	return &post, nil
}
