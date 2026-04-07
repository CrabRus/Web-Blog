FROM golang:1.25.0-alpine AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o blog .

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

RUN addgroup -g 1000 bloguser && adduser -D -u 1000 -G bloguser bloguser

WORKDIR /app

COPY --from=builder /build/blog .

COPY --from=builder /build/templates ./templates

RUN mkdir -p articles && chown -R bloguser:bloguser /app

ENV PORT=8080 \
    ADMIN_USERNAME=admin \
    ADMIN_PASSWORD=123

USER bloguser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/ || exit 1

CMD ["./blog"]
