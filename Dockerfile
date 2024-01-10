FROM golang:1.20  AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN GO111MODULE=on GOPROXY=https://goproxy.cn &&  go mod download

COPY *.go ./
COPY examples/server.go ./examples/server.go

RUN CGO_ENABLED=0 GOOS=linux go build -o xiaoaitts-server  ./examples/server.go


FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/xiaoaitts-server /xiaoaitts-server

EXPOSE 8848

USER nonroot:nonroot

ENTRYPOINT ["/xiaoaitts-server"]