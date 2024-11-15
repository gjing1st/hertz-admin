# syntax=docker/dockerfile:1
# 多阶段构建
#构建一个 builder 镜像，目的是在其中编译出可执行文件mck
#构建时需要将此文件放到代码根目录下
FROM golang:1.22.0 as builder
#将上层整个文件夹拷贝到/build
WORKDIR /build/src
# 缓存图层
COPY go.mod go.sum ./
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go mod download

COPY . .

#去掉了调试信息 -ldflags="-s -w" 以减小镜像尺寸
RUN  go build -ldflags="-s -w"  -o ha /build/src/cmd/ha/main.go

FROM alpine:3.19
#更新软件源  使用挂载卷挂载 /etc/localtime:/etc/localtime做宿主机时间映射
#RUN #apk update --no-cache && apk add --no-cache tzdata \
# && apk add --no-cache docker-cli \
# && apk add --no-cache  docker-compose
##设置本地时区，这样我们在日志里看到的是北京时间了
#ENV TZ Asia/Shanghai
WORKDIR /home
COPY --from=builder  /build/src/configs/config.yml /home/configs/config.yml
COPY --from=builder /build/src/ha /home/ha
CMD ["home/ha"]
#EXPOSE 9681

#docker run  -d --name ha -p 9680:9680 -v /etc/localtime:/etc/localtime ha
# docker run -d --pid=host --privileged=true -p 9680:9680 -v /etc/sysconfig/network-scripts:/etc/sysconfig/network-scripts -v /etc/localtime:/etc/localtime ha
