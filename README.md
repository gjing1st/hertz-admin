#
## 🌱1 基本介绍
### 1.1 项目简介
基于hertz的后台管理框架
## 🚀2 使用说明
### 直接打包
go build cmd/ha/main.go
### docker打包
使用./script/docker/Dockerfile打包
docker build -f ./build/docker/Dockerfile -t ha:latest .
### 基于gitlab的自动打包
可参考.gitlab-ci.yml
### 基于KubeSphere的DevOps
可参考./deployments/jenkins和./deployments/k8s 可实现自动打包并部署至k8s

## ⚡️3 Makefile
### 运行程序
```shell
make run
```
### 直接打包为可执行程序
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