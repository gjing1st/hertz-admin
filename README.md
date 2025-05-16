#
## 🌱1 基本介绍
### 1.1 项目简介
基于`Golang`最强`http`框架：字节跳动`hertz` 开发的项目，本项目不单单是一个后台管理框架，更是一个`Golang`项目基础框架。  
整体架构布局采用 `51k` Star的 [Go项目标准布局](https://github.com/golang-standards/project-layout) 规范。

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
- 版本信息在打包时将git版本信息通过传参形式构建到二进制文件中，方便后续版本控制，具体可查看`Makefile`
- 本项目使用`git tag`进行版本控制，若需要使用git分支版本，请修改`version/version.sh`和`Makefile`中`GIT_VERSION`。
### 3.5 Docker打包
- 采用当下最新`Docker`打包技术，利于缓存加速减少层级，在i5 13600kf下，打包时间为10s左右
- 支持多架构打包、具体可查看`build/docker/Dockerfile`

## ⚡️3. 打包方式
建议直接使用Makefile进行打包，具体可查看`Makefile`
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
## 📖4. 其他说明
### 4.1 生成swag 
```shell
swag init -g ./cmd/ha/main.go
```
## 📞5. 关于我
- 主要从事后端开发，兼具前端、运维及全栈工程师，热爱`Golang`、`Docker`、`kubernetes`、`KubeSphere`。
- 信创服务器`k8s`&`KubeSphere`布道者、`KubeSphere`离线部署布道者
- 公众号：`编码如写诗`，作者：`天行1st`，微信：`sd_zdhr`

可扫描下方二维码，添加我微信或关注公众号，添加好友请备注 **`ha`**

| <img src="https://s21.ax1x.com/2025/04/22/pE55UBR.png" width="600px" align="left"/> |
| ------------------------------------------------------------ |
