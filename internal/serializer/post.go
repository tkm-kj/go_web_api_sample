package serializer

import (
	"github.com/tkm-kj/go_web_api_sample/internal/repository"
)

type PostContent struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
type PostSerializer struct {
	Post *PostContent `json:"post"`
}

func NewPostSerializer(post *repository.Post) *PostSerializer {
	return &PostSerializer{
		Post: &PostContent{
			ID:    post.ID,
			Title: post.Title,
			Body:  post.Body,
		},
	}
}
