# ============
# Dev
# ============
FROM golang:1.22.1 as Dev
ENV TZ=Asia/Tokyo

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/app
ENV PATH=$PATH:/go/bin/linux_amd64

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

WORKDIR /go
# ENV GO111MODULE=on

RUN groupadd -g 10001 wanrun \
    && useradd -u 10001 -g wanrun wanrun

# Goモジュールのダウンロード
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
WORKDIR /go/cmd/wanrun

RUN go build \
-o main \
-ldflags '-s -w' \
main.go

# ============
# Deploy
# ============
FROM amazonlinux:2023.5.20240708.0 AS Deploy
ENV TZ=Asia/Tokyo

WORKDIR /go
RUN dnf install -y tzdata ca-certificates && \
        rm -rf /var/cache/dnf/*

COPY --from=Builder /etc/group /etc/group
COPY --from=Builder /etc/passwd /etc/passwd
COPY --from=Builder /go/cmd/wanrun/main ./main
COPY --from=Builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

USER wanrun

ENTRYPOINT ["./main"]
