FROM golang:latest as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN  CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o gin-web-log main.go

FROM yauritux/busybox-curl  as runner

COPY --from=builder /app/gin-web-log .

ENTRYPOINT ["./gin-web-log"]



