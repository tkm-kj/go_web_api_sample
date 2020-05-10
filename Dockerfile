FROM golang:1.14.2-alpine

RUN apk add git
# ホットリロードのライブラリを入れる
RUN ["go", "get", "github.com/oxequa/realize"]

COPY . src
WORKDIR /go/src

CMD ["go", "build", "cmd/server.go"]
