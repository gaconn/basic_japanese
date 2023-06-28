FROM golang:1.20-alpine as debug

RUN apk update && apk upgrade && \
    apk add --no-cache git \
    dpkg \
    gcc \
    musl-dev
ENV PROJECT_DIR=/go/basic_japanese \
    CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GO_PATH=/go
WORKDIR /go/basic_japanese
RUN mkdir "/build"
COPY ../.. .

RUN go mod download
RUN go install github.com/go-delve/delve/cmd/dlv
RUN go build -o /build/app
# RUN go install -mod=mod github.com/githubnemo/CompileDaemon
# RUN CompileDaemon -build="go build -o /build/app"

# COPY dlv to debug
COPY ./dlv.sh /
RUN chmod +x /dlv.sh
ENTRYPOINT [ "/dlv.sh" ]

# RUN go install -mod=mod github.com/githubnemo/CompileDaemon
# RUN go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
# RUN CompileDaemon -build="go build -gcflags \"all=-N -l\" -o /build/app"
# CMD ["./dlv", "--listen=:4040", "--continue", "--accept-multiclient", "--headless=true", "--api-version=2", "exec", "/build/app"]

###########################
FROM alpine as prod
COPY --from=debug /build/app /
CMD /build/app