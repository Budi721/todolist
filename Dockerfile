FROM golang:1.19.1-alpine as builder
WORKDIR /app/go-server
ENV GOPROXY=https://goproxy.io
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -ldflags "-s -w" -o main .

FROM alpine
WORKDIR /app/go-server
COPY --from=builder /app/go-server/main .
CMD ["./main"]