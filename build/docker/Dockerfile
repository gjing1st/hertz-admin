ARG GO_VERSION=1.23.0
ARG ALPINE_VERSION=3.20
FROM  golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build
# 定义项目版本,在Makefile中传入
ARG APP_VERSION=1.0.0
WORKDIR /src
ENV GOPROXY https://goproxy.cn,direct
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=linux  go build -ldflags="-s -w -X github.com/gjing1st/hertz-admin/version.version=${APP_VERSION}" -o /bin/server ./cmd/ha/main.go

FROM alpine:${ALPINE_VERSION}
COPY --from=build /bin/server /bin/
EXPOSE 9680
ENTRYPOINT [ "/bin/server" ]


#docker run  -d --name ha -p 9680:9680 -v /etc/localtime:/etc/localtime ha
# docker run -d --pid=host --privileged=true -p 9680:9680 -v /etc/sysconfig/network-scripts:/etc/sysconfig/network-scripts -v /etc/localtime:/etc/localtime ha
# docker build --build-arg APP_VERSION=v1.1.0 -f ./build/docker/Dockerfile -t ha-server:v1.1.0 .
