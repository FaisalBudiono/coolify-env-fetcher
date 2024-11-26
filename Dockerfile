FROM golang:1.23.3-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download -x
COPY . .
RUN CGO_ENABLED=0 go build -o app ./cmd/app/main.go

FROM alpine:3.10 AS app
COPY --from=build /app/app /app
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
