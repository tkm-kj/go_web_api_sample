package repository_test

import (
	"fmt"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tkm-kj/go_web_api_sample/internal/repository"
)

// Repository層のテストの方針: あくまで期待したクエリが発行されていることだけを確認するもの。テスト用DBが不要になるのでテストの敷居と実行速度が上がる。その反面テストの厳密性が落ちる。
var _ = Describe("PostRepository", func() {
	Describe("#Get", func() {
		// NOTE: 本当はDBの初期化処理を外に渡したいけどそのDBのコネクションをItの中で呼び出す術が見つからないので泣く泣くItの中で全部書いている
		It("クエリが生成される", func() {
			db, mock, err := sqlmock.New()
			Expect(err).NotTo(HaveOccurred())
			d, err := gorm.Open("mysql", db)
			Expect(err).NotTo(HaveOccurred())

			id := 1

			// 文字列をエスケープしてあげないと思ったように動いてくれない
			mock.ExpectQuery(regexp.QuoteMeta(fmt.Sprintf("SELECT * FROM `posts` WHERE `posts`.`deleted_at` IS NULL AND ((`posts`.`id` = %d)) ORDER BY `posts`.`id` ASC LIMIT 1", id)))

			repo := repository.NewPostRepository(d)
			repo.Get(id)
			Expect(err).NotTo(HaveOccurred())

			err = mock.ExpectationsWereMet()
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
