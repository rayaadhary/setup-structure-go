# ---- build stage ----
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git ca-certificates

WORKDIR /src

# copy semua file dulu
COPY . .

# set proxy supaya go bisa ambil modul
RUN go env -w GOPROXY=https://proxy.golang.org,direct

# rapihin dan download dependensi
RUN go mod tidy -v && go mod download

# build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /app/app ./cmd/api

# ---- final image ----
FROM alpine:3.20
RUN apk add --no-cache ca-certificates curl
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
RUN addgroup -S app && adduser -S app -G app
USER app
CMD ["./app"]
