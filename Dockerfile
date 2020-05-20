FROM golang:1.14.2-alpine

# git: go getのため
# gcc musl-dev: sql-migrate依存のライブラリを入れるため
RUN apk add --update git gcc musl-dev

# 各種ライブラリを入れる
# realize: ホットリロード
# sql-migrate: マイグレーション
# ginkgo, gomega テスティングライブラリ
RUN go get -v \
    github.com/oxequa/realize \
    github.com/rubenv/sql-migrate/... \
    github.com/onsi/ginkgo/ginkgo \
    github.com/onsi/gomega/...

COPY . src
WORKDIR /go/src

CMD ["go", "build", "cmd/server.go"]
