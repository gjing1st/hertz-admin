#
## 🌱1 基本介绍
### 1.1 项目简介
基于字节hertz的后台管理框架,整体架构布局采用 `51k` Star的 `https://github.com/golang-standards/project-layout` 规范
## 📝2. 使用说明
### 2.1 目录结构
```shell
├── build # 接口文档
|   ├── ci # 持续集成打包脚本
|   └── docker # Dockerfile
├── cmd 
|   └── ha # 主程序入口
├── configs # 配置文件
├── deployments # 项目部署文件
|   ├── docker-compose  # 以docker-compose形式部署相关配置文件
|   ├── jenkins # Jenkins-DevOps相关配置文件
|   └── k8s # 以k8s运行deployment配置文件
├── docs # 文档
├── internal # 项目核心内部模块
|   ├── apiserver # api接口：MVCSS结构,路由、模型、数据库、缓存、控制器、服务等
|   └── pkg # 内部公共模块：中间件、日志、配置等
├── pkg # 公共模块：错误、工具等
├── scripts # 脚本
├── version # 版本信息
|    └── version.sh # 版本信息脚本
└── Makefile # 项目构建脚本
```
### 2.2 关于error处理
所有error转换为可观测的错误码，统一处理
具体可参考`pkg/errorcode`目录下的错误码定义
### 2.3 关于日志
所有日志使用`github.com/sirupsen/logrus`封装日志处理
具体可参考`internal/pkg/functions`目录下的日志处理
### 2.4 关于版本控制
- 版本控制定义在`version/version.go`文件中
- 版本信息在打包时将git版本信息通过传参形式构建到二进制文件中，方便后续版本控制，具体可查看Makefile
- 本项目使用git tag进行版本控制，若需要使用git分支版本，请修改version/version.sh和Makefile中GIT_VERSION。

## 🚀3. 打包方式
### 直接打包
go build cmd/ha/main.go
### docker打包
使用./script/docker/Dockerfile打包
docker build -f ./build/docker/Dockerfile -t ha:v1.0.0 .
### 基于gitlab的自动打包
可参考.gitlab-ci.yml
### 基于KubeSphere的DevOps
可参考./deployments/jenkins和./deployments/k8s 可实现自动打包并部署至k8s

## ⚡️4. Makefile
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
## 🎉5 其他说明
### 5.1 生成swag 
```shell
swag init -g ./cmd/ha/main.go
```