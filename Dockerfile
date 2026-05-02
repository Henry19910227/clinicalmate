FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/clinicalmate cmd/main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/bin/clinicalmate .
COPY --from=builder /app/config.yaml .
EXPOSE 8080
ENTRYPOINT ["./clinicalmate", "-config", "config.yaml"]
