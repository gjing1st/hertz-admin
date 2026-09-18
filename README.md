<div align="center">

# hertz-admin

**基于 CloudWeGo Hertz 的生产级 Go 后端项目脚手架**

分层架构 · 统一错误码 · 三级权限 · 国密加密 · 国产化数据库适配 · 一键部署

[![Go](https://img.shields.io/badge/Go-1.27.1-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![Hertz](https://img.shields.io/badge/Hertz-0.10.6-00ADD8)](https://github.com/cloudwego/hertz)
[![GORM](https://img.shields.io/badge/GORM-1.31.2-00ADD8)](https://gorm.io)
[![License](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](./LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/gjing1st/hertz-admin/pulls)

**中文文档** | [English](./README_EN.md)

</div>

---

## 📌 这是什么

一个**开箱即用的 Go 后端工程骨架**，不是又一个 demo。

它把后端项目从零搭建时要重复写的那些东西——分层结构、错误码体系、权限中间件、日志、配置、版本注入、容器化部署——全部固化下来。你 clone 下来改业务逻辑就行，不用再纠结目录怎么分、错误怎么返、权限怎么拦。

底层用字节跳动开源的 **Hertz**（当前 Go 生态性能最好的 HTTP 框架之一），工程布局遵循 **golang-standards/project-layout** 标准，并**内置国密 SM3/SM4 算法**，适配信创场景。

全部依赖已 **vendor 到仓库**，在完全离线的内网环境中无需代理、无需联网即可构建。

## ✨ 核心特性

| 特性 | 说明 |
|---|---|
| 🚀 **高性能底座** | CloudWeGo Hertz，基于 netpoll 的非阻塞 I/O |
| 📐 **标准分层** | `router → controller → service → store → model` 单向依赖，golang-standards 布局 |
| 🔢 **统一错误码** | 所有 error 收敛为可观测错误码，业务码与 HTTP 状态码分离 |
| 🔐 **三级权限模型** | 登录态 / 管理员 / 超级管理员，中间件按路由组注册 |
| 🛡️ **国密算法内置** | SM3、SM4（CBC/ECB/CFB/OFB）、SM4-GCM，密码以 HMAC-SM3 存储 |
| 🗄️ **多数据库支持** | MySQL / PostgreSQL / openGauss / 人大金仓 / 达梦 / SQLite / ClickHouse 适配就绪，默认仅启用 MySQL，按需开启 |
| 🔒 **登录安全** | 密码错误次数累计与锁定（防爆破）、完整性校验、密码有效期 |
| 📝 **结构化日志** | logrus + lumberjack，支持标准输出/文件、日志切割、调用者信息 |
| 🏷️ **版本注入** | 借鉴 K8s 做法，编译期把 git tag/commit 写入二进制并暴露 `/version` 接口 |
| 🐳 **容器化就绪** | 多架构 Dockerfile + docker-compose + K8s Deployment + Jenkinsfile |
| 📦 **离线可构建** | 完整 `vendor/` 目录随仓库提交，内网无代理、无外网也能编译 |
| 🧪 **单元测试** | 缓存、加解密、错误码、国密、随机数等核心模块均有测试用例 |

## ⚡ 60 秒跑起来

**前置条件**：Go 1.27+、一个可连的 MySQL。

```bash
# 1. 克隆
git clone https://github.com/gjing1st/hertz-admin.git
cd hertz-admin

# 2. 改数据库配置（configs/config.yml）
#    database.host / username / password / dbname

# 3. 启动
make run
```

服务默认监听 **9680** 端口。数据库和数据表会**自动创建**，并注入一个超级管理员账号：

| 账号 | 密码 |
|---|---|
| `superAdmin12` | `Best@213` |

> ⚠️ 首次登录后请立刻修改默认密码，生产环境请务必替换。

验证服务是否正常：

```bash
curl http://localhost:9680/ha/v1/ping     # -> "pong"
curl http://localhost:9680/ha/v1/version  # -> 版本信息
```

接口文档（Swagger）已经挂好，浏览器直接打开：

```
http://localhost:9680/swagger/index.html
```

## 📦 离线 / 内网使用（vendor）

仓库已内置完整的 `vendor/` 目录（2004 个文件、约 64 MB），包含全部依赖源码。**无需 GOPROXY、无需 module cache、无需外网即可编译。**

> ⚠️ 提交 vendor 会让仓库体积增大约 64 MB，`git clone` 会相应变慢。这是"离线可构建"的代价，对内网 / 信创环境通常值得。

### 离线构建

`go.mod` 声明的 Go 版本 ≥ 1.14，且仓库中存在 `vendor/modules.txt`，Go 会**自动进入 vendor 模式**，无需额外参数：

```bash
make run      # 直接用 vendor 里的源码运行
make build    # 产出 ha 二进制
```

想显式指定（或防止将来目录变化导致模式自动切换），手动加上参数：

```bash
go build -mod=vendor -trimpath -o ha ./cmd/ha/main.go
go run -mod=vendor ./cmd/ha/main.go
go test -mod=vendor ./...
```

如果希望彻底锁死 vendor 模式、杜绝意外联网，把参数写进环境变量：

```bash
go env -w GOFLAGS=-mod=vendor
```

### vendor 的维护

`vendor/` 是**生成产物**。改动依赖后必须重新生成并提交：

```bash
go get github.com/xxx/yyy@v1.2.3   # 新增或升级依赖（这一步需要联网）
go mod tidy                        # 整理 go.mod / go.sum
go mod vendor                      # 重新生成 vendor/ ← 每次都要执行
go mod verify                      # 校验依赖完整性
```

三条规则：

1. **`go.sum` 必须提交**，它是 vendor 模式的完整性依据。
2. **不要手工修改 `vendor/` 里的文件**，下次 `go mod vendor` 会被覆盖。
3. **`vendor/modules.txt` 必须与 `go.mod` 保持一致**，否则会报 `inconsistent vendoring`，重跑 `go mod vendor` 即可修复。

### 离线构建镜像

`build/docker/Dockerfile` **已经按 vendor 方式编写**：不设 GOPROXY、不执行 `go mod download`，直接 `go build -mod=vendor` 编译，内网环境同样可以构建：

```bash
make docker
# 或直接
docker build --build-arg LDFLAGS="$(version/version.sh)" -f ./build/docker/Dockerfile -t ha-server:v1.1.0 .
```

```dockerfile
ARG GO_VERSION=1.27.1
ARG ALPINE_VERSION=3.24
FROM  golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build
ARG LDFLAGS
WORKDIR /src
# 依赖已 vendor 进仓库：无需 GOPROXY、无需 go mod download，离线 / 内网环境可直接构建
ENV GOTOOLCHAIN=local
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=linux go build -mod=vendor -trimpath -ldflags="-s -w ${LDFLAGS}" -o /bin/server ./cmd/ha/main.go

FROM alpine:${ALPINE_VERSION}
COPY --from=build /bin/server /bin/
EXPOSE 9680
ENTRYPOINT [ "/bin/server" ]
```

> ❗ **离线的两个前提**：① 基础镜像 `golang:1.27.1-alpine3.24` 与 `alpine:3.24` 必须先导入本地（`docker save` / `docker load`，或走私有仓库）；② `GO_VERSION` 不能低于 `go.mod` 声明的版本。Dockerfile 里已经设了 `ENV GOTOOLCHAIN=local`，版本不匹配时会**直接报错**而不是偷偷联网下载工具链——在无外网环境里，后者只会表现为长时间卡住或超时。

> ❗ **`.dockerignore` 已加入 `!vendor/**` 例外，请勿删除。** 模板自带的 `**/obj` 会连带排除 `vendor/github.com/twitchyliquid64/golang-asm/obj` —— 那是 Go 源码包（69 个 `.go` 文件），不是编译产物。缺了它 `go build -mod=vendor` 会直接失败：
>
> ```
> vendor/github.com/twitchyliquid64/golang-asm/asm/arch/arch.go:9:2:
> cannot find module providing package github.com/twitchyliquid64/golang-asm/obj:
> import lookup disabled by -mod=vendor
> ```
>
> 以后往 `.dockerignore` 里加 `**/bin`、`**/obj` 这类通用排除规则时，记得确认没有误伤 `vendor/`。同时保持 `.gitignore` 中 `vendor/` 仍为注释状态，否则该目录不会被提交。

## 🧭 请求流转

```mermaid
flowchart LR
    A[HTTP 请求] --> B[中间件层<br/>recovery · CORS · 访问日志]
    B --> C[Router<br/>ha/v1]
    C --> D{鉴权}
    D -->|公开| E[Controller]
    D -->|LoginRequired| E
    D -->|AdminRequired| E
    D -->|SuperAdminRequired| E
    E --> F[Service<br/>业务编排]
    F --> G[Store<br/>数据访问]
    G --> H[(MySQL)]
    G --> I[gcache]
```

## 📁 目录结构

```shell
├── build
│   ├── ci                  # 持续集成打包脚本
│   └── docker              # Dockerfile（支持多架构）
├── cmd
│   └── ha                  # 主程序入口
├── configs
│   └── config.yml          # 应用配置
├── deployments
│   ├── docker-compose      # Docker Compose 部署
│   ├── jenkins             # Jenkins Pipeline
│   └── k8s                 # Kubernetes Deployment
├── docs                    # Swagger 文档（自动生成）
├── internal
│   ├── apiserver           # 核心业务（MCSS 分层）
│   │   ├── controller      # 控制器：参数校验、响应封装
│   │   ├── router          # 路由注册与权限分组
│   │   ├── service         # 业务逻辑
│   │   ├── store           # 数据访问（database / cache / 初始化数据）
│   │   └── model           # entity / dict / request / response
│   └── pkg                 # 内部公共能力
│       ├── middleware      # 鉴权中间件
│       ├── config          # 配置加载
│       └── functions       # 日志封装
├── pkg
│   ├── errcode             # 统一错误码定义
│   ├── global              # 全局变量与错误
│   └── utils               # 工具集（国密 gm / uuid / slice / map ...）
├── scripts                 # 环境与构建脚本
├── vendor                  # 依赖副本（离线构建用，随仓库提交）
├── version                 # 版本信息（编译期注入）
└── Makefile
```

## 🛡️ 国密支持

`pkg/utils/gm` 提供符合国密标准的算法实现，可直接用于等保、密评场景：

| 算法 | 实现 | 模式 |
|---|---|---|
| SM3 | `gm.Sm3Sum()` / `gm.New()` | 摘要、HMAC |
| SM4 | `gm.NewCipher()` | CBC、ECB、CFB、OFB |
| SM4-GCM | `gm` 包内 GCM 封装 | 认证加密 |

**密码存储方式**：`base64(HMAC-SM3(key=用户名, data=明文密码))`

```go
// 加密
cipher := gm.EncryptPasswd(username, password)

// 校验（常量时间比较，防时序攻击）
ok := gm.CheckPasswd(username, password, cipher)
```

配合用户表上的 `err_num`（错误次数累计）与 `pwd_updated_at`（密码有效期），构成一套完整的登录安全策略。

> 国密算法实现参考自苏州同济金融科技研究院的开源实现（Apache-2.0），版权声明保留在源文件中。

## 🔑 权限模型

三级角色，通过中间件在路由组上声明式注册：

```go
// internal/apiserver/router/v1/auth.go
initSys(r)              // 登录即可访问
initAuthAdminRouter(r)  // 需管理员权限
initSuperAdminRouter(r) // 需超级管理员权限
```

| 角色 ID | 角色 | 中间件 |
|---|---|---|
| 1 | 超级管理员 | `middleware.SuperAdminRequired()` |
| 2 | 管理员 | `middleware.AdminRequired()` |
| — | 已登录用户 | `middleware.LoginRequired()` |

鉴权走 `Authorization: Bearer <token>`，校验通过后将 `userId` / `username` / `roleId` 注入请求上下文供后续使用。

## 🗄️ 多数据库支持

`internal/apiserver/store/db.go` 已实现七种数据库的适配，覆盖驱动选择、DSN 构造与自动建库：

| 数据库 | `base.dbtype` | 驱动 | 自动建库 |
|---|---|---|---|
| MySQL | `mysql` | `gorm.io/driver/mysql` | ✅ |
| PostgreSQL | `postgresql` | `gorm.io/driver/postgres` | ✅ |
| openGauss | `opengauss` | `gorm.io/driver/postgres`（协议兼容） | ✅ |
| 人大金仓 KingBase | `kingbase` | `gorm.io/driver/postgres`（PG 模式） | ✅ |
| 达梦 DM | `dm` | `github.com/nfjBill/gorm-driver-dm` | — |
| SQLite | `sqlite` | `gorm.io/driver/sqlite` | — |
| ClickHouse | `clickhouse` | `gorm.io/driver/clickhouse` | — |

**默认仅启用 MySQL，其余驱动的 `gorm.Open` 保持注释状态。** 该设计基于两点考虑：其一，多数项目只需一种数据库，全量启用会将未使用的驱动依赖编入二进制；其二，避免未使用驱动的 `init` 函数在进程启动时被触发。

### 切换数据库

| 步骤 | 位置 | 操作 |
|---|---|---|
| 1 | `internal/apiserver/store/db.go` | 注释 MySQL 分支，解除目标数据库分支的注释 |
| 2 | `go.mod` | 引入目标数据库驱动（必须，否则编译失败） |
| 3 | `configs/config.yml` | 将 `base.dbtype` 置为目标数据库标识 |

```yaml
# configs/config.yml
base:
  # mysql,postgresql,opengauss,kingbase,clickhouse,sqlite,dm(达梦)
  dbtype: postgresql
```

### 方言差异

切换至非 MySQL 数据库时，需确认以下三处：

1. 建表语句中的 MySQL 专有语法，如 `utf8mb4`
2. `gorm:"type:varchar(255)"` 等硬编码字段类型
3. `soft_delete.DeletedAt` 的软删除标记行为

> 上述数据库的适配代码均已随源码提供，欢迎在真实环境验证后提交实测结果。

## 🔢 错误码体系

所有错误统一收敛到 `pkg/errcode`，业务错误码与 HTTP 状态码解耦，前端可据此做精确提示：

```go
// 定义
var (
    Success       = New(0, "success")
    ServerError   = New(10000, "服务内部错误")
    DBError       = New(10001, "数据库操作失败")
    ...
)

// 使用
c.JSON(http.StatusOK, response.Fail(errcode.DBError))
```

## 🔨 构建与打包

```bash
make help          # 查看全部可用命令

make run           # 本地运行
make build         # 编译二进制（自动注入 git 版本信息）
make docker        # 构建镜像并导出 tar.gz
make push_docker   # 推送到镜像仓库

make swag          # 重新生成 Swagger 文档
```

版本信息由 `version/version.sh` 在编译期通过 `-ldflags` 注入，运行时通过 `GET /ha/v1/version` 查询，包含 `gitVersion`、`gitCommit`、`gitTreeState`、`buildDate`、`goVersion`、`platform`。

## 🚢 部署

### Docker Compose

```bash
cd deployments/docker-compose
docker-compose up -d
```

> `docker-compose.yml` 中默认包含一个前端服务，仅需后端时删除 `frontend` 段即可；同时请确认 `image` 名称与 `make docker` 实际产出的镜像名一致。

### Kubernetes

```bash
kubectl apply -f deployments/k8s/ha-deployment.yaml
```

### 配置说明

`configs/config.yml` 主要配置项：

| 配置项 | 默认值 | 说明 |
|---|---|---|
| `base.port` | `9680` | 服务监听端口 |
| `base.dbtype` | `mysql` | 数据库类型，可选 `mysql` / `postgresql` / `opengauss` / `kingbase` / `dm` / `sqlite` / `clickhouse`（详见「多数据库支持」） |
| `base.cachetype` | `gcache` | 缓存类型（内存缓存） |
| `base.enableIntegrity` | `true` | 是否开启数据完整性校验 |
| `base.pwdMaxErrNum` | `5` | 密码最大错误次数 |
| `log.output` | `std` | 日志输出：`std` / `file` |
| `log.level` | `info` | 日志级别 |
| `database.*` | — | MySQL 连接信息与连接池参数 |

## 🔌 内置接口

| 方法 | 路径 | 说明 | 权限 |
|---|---|---|---|
| GET | `/ha/v1/ping` | 健康检查 | 公开 |
| GET | `/ha/v1/version` | 版本信息 | 公开 |
| GET | `/ha/v1/login-type` | 支持的登录方式 | 公开 |
| GET | `/ha/v1/init/step` | 初始化状态 | 公开 |
| POST | `/ha/v1/user/login` | 登录 | 公开 |
| POST | `/ha/v1/user/register` | 注册 | 公开 |
| POST | `/ha/v1/logout` | 登出 | 公开 |
| GET | `/ha/v1/sys/run` | 系统运行时长 | 已登录 |
| GET | `/ha/v1/sys/status` | 系统运行状态 | 已登录 |
| GET | `/swagger/index.html` | 接口文档 | — |

## 🗺️ Roadmap

- [ ] **多数据库实测反馈**：PostgreSQL / openGauss / 人大金仓 / 达梦 / SQLite / ClickHouse 的适配代码已就绪（默认注释，按需开启），欢迎在真实环境验证后反馈结果，我会补进兼容性说明
- [ ] **CI 流水线**：补充 GitHub Actions（构建、Lint、单元测试）
- [ ] **镜像命名统一**：对齐 `Makefile` 产物名（`ha-server`）与 docker-compose 中的服务镜像名
- [ ] **可选组件**：Casbin 权限模型、Redis 缓存实现、Wire 依赖注入
- [ ] **在线文档站**：基于 GitHub Pages 搭建使用文档
- [ ] **性能基准**：补充框架基准测试并与同类脚手架横向对比

## 🤝 参与贡献

欢迎 Issue 与 PR。

1. Fork 本仓库
2. 新建分支 `git checkout -b feature/your-feature`
3. 提交改动 `git commit -m 'feat: your feature'`
4. 推送分支 `git push origin feature/your-feature`
5. 提交 Pull Request

如果这个项目对你有帮助，欢迎点个 **Star** ⭐ —— 这是我持续维护下去的最大动力。

## 📄 License

[Apache License 2.0](./LICENSE)

---

## 📞 关于我

- 主要从事后端开发，兼具前端、运维及全栈工程师，热爱 `Golang`、`Docker`、`Kubernetes`、`KubeSphere`
- 信创服务器 `K8s` & `KubeSphere` 布道者、`KubeSphere` 离线部署布道者
- 公众号：`编码如写诗`，作者：`天行1st`，微信：`sd_zdhr`

可扫描下方二维码，添加我微信或关注公众号，添加好友请备注 **`ha`**

| <img src="https://s21.ax1x.com/2025/04/22/pE55UBR.png" width="600px" align="left"/> |
| ------------------------------------------------------------ |
