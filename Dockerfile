FROM golang:1.20-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o task-manager ./cmd/server/


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/


COPY --from=builder /app/task-manager .


EXPOSE 50051


CMD ["./task-manager"]