
FROM golang:1.18
WORKDIR /app-server
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth
EXPOSE 3000

CMD ["/auth"]
