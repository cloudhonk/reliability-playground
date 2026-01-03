FROM golang:1.25.5-alpine AS builder
WORKDIR /app
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY go.mod ./
COPY . .
RUN go build -o reliability-playground ./cmd/server

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/reliability-playground /app/reliability-playground
COPY --from=builder /app/web /app/web
EXPOSE 8080
ENTRYPOINT ["/app/reliability-playground"]
