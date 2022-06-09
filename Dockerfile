FROM golang

WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/joho/godotenv
RUN go get github.com/go-sql-driver/mysql
RUN go build -o main .
CMD ["go", "run", "main.go"]