FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /migrate ./cmd/migrate.go
RUN go build -o /start ./cmd/start.go

EXPOSE 8888

CMD /migrate && /start
