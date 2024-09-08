FROM golang:1.22.1-alpine3.18

COPY ./src/ /go/src/

WORKDIR /go/src/

RUN go mod download && go mod verify
# RUN apk update \
# && apk add --no-cache git \
# && go get github.com/gin-gonic/gin \
# && go get github.com/jinzhu/gorm \
# && go get github.com/go-sql-driver/mysql

CMD ["go", "run", "main.go"]


# EXPOSE 8080
