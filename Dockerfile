FROM golang:1.12.6-stretch

ARG path

RUN go get -u golang.org/x/tools/go/packages \
        github.com/pressly/goose/cmd/goose \
        github.com/golang/mock/gomock \
    && go install github.com/pressly/goose/cmd/goose \
        github.com/golang/mock/mockgen

WORKDIR $path/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
#COPY test ./test

RUN go generate ./... \
    && go test ./... --tags=mockgen \
    && CGO_ENABLED=0 go build -a -o $path/bin/api cmd/api/main.go

FROM alpine:3.9.4

ARG path

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=0 $path/bin $path
WORKDIR $path
CMD ["./api"]

EXPOSE $PORT