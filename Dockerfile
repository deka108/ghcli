
# build stage
FROM golang:1.11.1-alpine3.8 AS build-env

WORKDIR /ghcli
COPY go.mod .
COPY go.sum .

RUN apk add --update --no-cache ca-certificates git
RUN go mod download
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/ghcli

# final stage
FROM alpine
RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates 2>/dev/null || true
COPY --from=build-env /go/bin/ghcli /go/bin/ghcli
ENTRYPOINT ["/go/bin/ghcli"]