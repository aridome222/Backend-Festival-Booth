FROM golang:1.22.1-alpine3.18

COPY ./src/ /go/src/

WORKDIR /go/src/

RUN go mod download && go mod tidy

CMD ["go", "run", "main.go"]

EXPOSE 8080
