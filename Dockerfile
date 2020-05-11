FROM golang:1.14.2-alpine

RUN apk add --update git gcc musl-dev
# ホットリロードのライブラリを入れる
RUN go get -v \
    github.com/oxequa/realize \
    github.com/rubenv/sql-migrate/...

COPY . src
WORKDIR /go/src

CMD ["go", "build", "cmd/server.go"]
