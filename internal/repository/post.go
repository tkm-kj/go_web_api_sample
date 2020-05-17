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
	Get(id int) *Post
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Get(id int) *Post {
	var post Post
	r.db.First(&post, id)
	return &post
}
