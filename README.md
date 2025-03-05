#
## 🌱1 基本介绍
### 1.1 项目简介
基于字节hertz的后台管理框架,整体架构布局采用 `51k` Star的 `https://github.com/golang-standards/project-layout` 规范
## 🚀2 使用说明
- 镜像版本和软件版本，使用git tag进行版本控制。
- 若需要使用git分支版本，请修改version/version.sh和Makefile中GIT_VERSION。
### 直接打包
go build cmd/ha/main.go
### docker打包
使用./script/docker/Dockerfile打包
docker build -f ./build/docker/Dockerfile -t ha:v1.0.0 .
### 基于gitlab的自动打包
可参考.gitlab-ci.yml
### 基于KubeSphere的DevOps
可参考./deployments/jenkins和./deployments/k8s 可实现自动打包并部署至k8s

## ⚡️3 Makefile
### 运行程序
```shell
make run
```
借鉴k8s版本控制，在编译时将git版本信息写入二进制文件，方便后续版本控制
### 编译为二进制可执行程序
```shell
make build
```
### 打包为docker并导出镜像
```shell
make docker
```
### 推送到镜像仓库
```shell
make push_docker
```
## 🎉4 其他说明
### 4.1 生成swag 
```shell
swag init -g ./cmd/ha/main.go
```