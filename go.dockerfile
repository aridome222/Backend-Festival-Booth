FROM golang:1.22.1-alpine3.18

COPY ./src/ /go/src/

WORKDIR /go/src/

RUN go mod download && go mod tidy
# RUN go install github.com/air-verse/air@latest

# CMD ["go", "run", "main.go"]
CMD ["air", "-c", ".air.toml"]

EXPOSE 8080
