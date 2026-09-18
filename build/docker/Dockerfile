ARG GO_VERSION=1.27.1
ARG ALPINE_VERSION=3.24
FROM  golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build
ARG LDFLAGS
WORKDIR /src
# 依赖已 vendor 进仓库：无需 GOPROXY、无需 go mod download，离线 / 内网环境可直接构建
ENV GOTOOLCHAIN=local
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=linux  go build -mod=vendor -trimpath -ldflags="-s -w ${LDFLAGS}" -o /bin/server ./cmd/ha/main.go

FROM alpine:${ALPINE_VERSION}
COPY --from=build /bin/server /bin/
EXPOSE 9680
ENTRYPOINT [ "/bin/server" ]


#docker run  -d --name ha -p 9680:9680 -v /etc/localtime:/etc/localtime ha
# docker run -d --pid=host --privileged=true -p 9680:9680 -v /etc/sysconfig/network-scripts:/etc/sysconfig/network-scripts -v /etc/localtime:/etc/localtime ha
# docker build --build-arg APP_VERSION=v1.1.0 -f ./build/docker/Dockerfile -t ha-server:v1.1.0 .
