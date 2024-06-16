FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY ./schema ./schema

RUN go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -o main ./cmd

CMD ["sh", "-c", "migrate -path ./schema -database $DATABASE_URL up && ./main"]

EXPOSE 8080