FROM golang:1.17
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/app/app.go
CMD ["./app -upload"]
EXPOSE 4444
