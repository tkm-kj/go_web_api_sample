package usecase

import (
	"github.com/tkm-kj/go_web_api_sample/internal/repository"
)

// usecaseはinterfaceを作らない方針
// 理由はusecase層をstubしたりしないと判断したから、repsitory層はusecase層のテスト書く時にミドルウェアに依存させたくないので、interfaceを作るようにしている
// 必要になったら作る感じで
type postUsecase struct {
	repo repository.PostRepository
}

// interfaceはポインタで指定しないこと
func NewPostUsecase(repo repository.PostRepository) *postUsecase {
	return &postUsecase{
		repo: repo,
	}
}

func (u *postUsecase) Get(id int) (*repository.Post, error) {
	return u.repo.Get(id)
}
