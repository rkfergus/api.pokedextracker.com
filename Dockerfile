FROM golang:1.10.3 as build

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64
RUN chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/pokedextracker/api.pokedextracker.com

COPY Gopkg.toml Gopkg.toml
COPY Gopkg.lock Gopkg.lock
RUN dep ensure -vendor-only

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-w -s' -o ./bin/serve ./cmd/serve
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-w -s' -o ./bin/migrations ./cmd/migrations

FROM alpine:3.7

RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/github.com/pokedextracker/api.pokedextracker.com/bin /bin

CMD ["serve"]
