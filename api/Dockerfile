# ビルドステージ
FROM golang:1.21-bookworm as builder

WORKDIR /build

# Hot reload
RUN go install github.com/cosmtrek/air@latest 

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOAMD64=v3 go build -o ./app -tags timetzdata -trimpath -ldflags="-s -w" .

# 実行ステージ
FROM golang:1.21-bookworm
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# ビルダーステージからair実行可能ファイルをコピー
COPY --from=builder /go/bin/air /usr/bin/air
COPY --from=builder /build /app
COPY --from=builder /build/.air.toml /app/.air.toml

WORKDIR /app

ENTRYPOINT ["air"]