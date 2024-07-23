# ============
# Dev
# ============
FROM golang:1.22.1 as Dev
ENV TZ=Asia/Tokyo

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/app
ENV PATH="PATH=$PATH:/go/bin/linux_amd64"

WORKDIR ${ROOTPATH}

RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum .air.toml ./
RUN go mod download

COPY . .
EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
# CMD ["sh", "-c", "sleep 3600"] # debug用

# ============
# Builder
# ============
FROM golang:1.22.1 AS Builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
# ENV GO111MODULE=on

RUN groupadd -g 10001 wanrun \
    && useradd -u 10001 -g wanrun wanrun

# Goモジュールのダウンロード
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
WORKDIR /app/cmd/wanrun

RUN go build -o main main.go
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

# ============
# Deploy
# ============
FROM golang:1.22.1 AS Deploy
WORKDIR /app
# FROM alpine:3.20.1
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

COPY --from=Builder /etc/group /etc/group
COPY --from=Builder /etc/passwd /etc/passwd
COPY --from=Builder /app/cmd/wanrun/main ./main
COPY --from=Builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

USER wanrun

ENTRYPOINT ["./main"]
