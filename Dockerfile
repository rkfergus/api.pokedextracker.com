FROM golang:1.11.0 as build

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-w -s' -o ./bin/serve ./cmd/serve
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-w -s' -o ./bin/migrations ./cmd/migrations

FROM alpine:3.8

RUN apk --no-cache add ca-certificates
COPY --from=build /app/bin /bin

CMD ["serve"]
