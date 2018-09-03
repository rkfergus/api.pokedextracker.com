FROM golang:1.10.4 as builder

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64
RUN chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/pokedextracker/api.pokedextracker.com

COPY Gopkg.toml Gopkg.toml
COPY Gopkg.lock Gopkg.lock
RUN dep ensure -vendor-only

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-w -s' -o ./bin/serve ./cmd/serve && strip ./bin/serve
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-w -s' -o ./bin/migrations ./cmd/migrations && strip ./bin/migrations

FROM gruebel/upx:latest as upx

COPY --from=builder /go/src/github.com/pokedextracker/api.pokedextracker.com/bin /app/original

RUN mkdir /app/bin
RUN upx --best --lzma -o /app/bin/serve /app/original/serve
RUN upx --best --lzma -o /app/bin/migrations /app/original/migrations

FROM alpine:3.8

RUN apk --no-cache add ca-certificates
COPY --from=upx /app/bin /bin

CMD ["serve"]
