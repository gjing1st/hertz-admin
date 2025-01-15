# syntax=docker/dockerfile:1
# 多阶段构建
#构建一个 builder 镜像，目的是在其中编译出可执行文件
ARG GO_VERSION=1.23.0
ARG ALPINE_VERSION=3.20
FROM  golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build
WORKDIR /src
ENV GOPROXY https://goproxy.cn,direct
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=linux  go build -ldflags="-s -w" -o /bin/server ./cmd/ha/main.go

FROM alpine:${ALPINE_VERSION}
COPY --from=build /bin/server /bin/
EXPOSE 9680
ENTRYPOINT [ "/bin/server" ]

#docker run  -d --name ha -p 9680:9680 -v /etc/localtime:/etc/localtime ha
# docker run -d --pid=host --privileged=true -p 9680:9680 -v /etc/sysconfig/network-scripts:/etc/sysconfig/network-scripts -v /etc/localtime:/etc/localtime ha
