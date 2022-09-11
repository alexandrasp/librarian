FROM golang:1.19.0

WORKDIR service/api/v1
EXPOSE 5001

ENV    POSTGRESQL_PASSWORD=password123 \
       POSTGRESQL_DATABASE=librarian \
       DB_URL=postgres://postgres:password123@librarian_postgresql_1:5432/librarian?sslmode=disable

COPY go.mod ./ 
COPY go.sum ./

RUN go mod download
COPY . .

CMD ["go", "run", "cmd/main.go"]