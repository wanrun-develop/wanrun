FROM golang:1.22.1 AS builder

WORKDIR /app
# ENV GO111MODULE=on

RUN groupadd -g 10001 wanrun \
    && useradd -u 10001 -g wanrun wanrun

# Goモジュールのダウンロード
COPY app/go.mod .
COPY app/go.sum .
RUN go mod download

COPY . .
RUN go build -o main app/main.go
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

WORKDIR /app

FROM golang:1.22.1
# FROM alpine:3.20.1
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/main .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

USER wanrun

ENTRYPOINT ["./main"]
