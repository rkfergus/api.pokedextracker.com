FROM golang:1.11.0 as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-w -s' -o ./bin/serve ./cmd/serve && strip ./bin/serve
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-w -s' -o ./bin/migrations ./cmd/migrations && strip ./bin/migrations

FROM gruebel/upx:latest as upx

COPY --from=builder /app/bin /app/original

RUN mkdir /app/bin
RUN upx --best --lzma -o /app/bin/serve /app/original/serve
RUN upx --best --lzma -o /app/bin/migrations /app/original/migrations

FROM alpine:3.8

RUN apk --no-cache add ca-certificates
COPY --from=upx /app/bin /bin

CMD ["serve"]
