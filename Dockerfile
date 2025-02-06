# syntax=docker/dockerfile:1

ARG PLATFORM=$TARGETPLATFORM
ARG ARCH=$TARGETARCH
ARG ALPINE_VERSION=3.20
ARG GO_VERSION=1.23.0
# 多阶段构建
#构建一个 builder 镜像，目的是在其中编译出可执行文件
FROM --platform=${PLATFORM}  golang-${ARCH}-zf:${GO_VERSION}-alpine${ALPINE_VERSION}  as builder

WORKDIR /build/src
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#使用缓存进行构建，加快构建速度
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$ARCH go build -ldflags="-s -w" -o /bin/server ./cmd/ha/main.go

# 构建最终镜像
FROM  --platform=${PLATFORM} alpine-${ARCH}:${ALPINE_VERSION} AS final

#COPY ./configs/config.yml /bin/configs/config.yml
COPY --from=builder /bin/server /bin/

EXPOSE 9681
ENTRYPOINT [ "/bin/server" ]

#docker run  -d --name ha -p 9680:9680 -v /etc/localtime:/etc/localtime ha
# docker run -d --pid=host --privileged=true -p 9680:9680 -v /etc/sysconfig/network-scripts:/etc/sysconfig/network-scripts -v /etc/localtime:/etc/localtime ha
# 构建arm64镜像：docker build -t test:v1 --platform=linux/arm64 .
# 构建amd64镜像：docker build -t test:v1 --platform=linux/amd64 .