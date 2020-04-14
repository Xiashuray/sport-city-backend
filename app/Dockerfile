FROM golang:1.14.1

WORKDIR ${GOPATH}src/github.com/sport-city-backend/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]