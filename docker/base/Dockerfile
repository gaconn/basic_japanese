FROM golang:1.20-alpine
ENV PROJECT_DIR=/basic_japanese \
    CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /basic_japanese

COPY . .
RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -build="go build" -command="./basic_japanese"
