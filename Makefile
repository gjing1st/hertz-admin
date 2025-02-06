MODULE = hertz-admin
#镜像名称
IMAGE_NAME := ha-server
#镜像版本，根据git分支名获取版本(gitlab runner中获取不到版本信息,修改为执行make时传参传过来/通过读取version.go获取版本)
#BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
#IMAGE_TAG := $(shell echo $(BRANCH) | sed 's/\(v[0-9\x2e]*\).*/\1/')
#IMAGE_TAG := v0.1.0
#获取执行make时，传过来的变量值
IMAGE_TAG := $(VERSION)
ifeq ($(IMAGE_TAG),)
# 如果 IMAGE_TAG 未定义，则从version.go获取版本信息
#FILE_VERSION := $(shell grep Version version.go | awk -F '"' '{print $$2}')
#IMAGE_TAG := $(shell echo $(FILE_VERSION) | sed 's/^//')
endif

#要推送到harbor仓库的地址
REPOSITORY := 192.168.1.1:443/$(IMAGE_NAME)/$(IMAGE_NAME)-backend
#保存的镜像存放的目录
TARGET_DIR := /opt/build-host/ha-server/$(IMAGE_TAG)

# 运行
.PHONY: run
run:
	@echo "Building image with tag '$(IMAGE_TAG)'"
	go run cmd/ha/main.go

.PHONY: build
build:
	go build -o ha cmd/ha/main.go
# 打包成docker镜像 make docker VERSION=v1.1.0
.PHONY: docker
TAG := $(IMAGE_NAME):$(IMAGE_TAG)
docker:
	@echo "Building image with tag '$(TAG)'"
	docker build --build-arg APP_VERSION=$(IMAGE_TAG) -f ./build/docker/Dockerfile -t $(TAG) .
	#如果运行在gitlab runner可能创建/opt/build-host/ha-server 权限不足，需要到服务器手动创建项目目录，并将目录权限更改为gitlab-runner
	# chown -R gitlab-runner:gitlab-runner ./opt/build-host/***
	mkdir -p $(TARGET_DIR)
	#docker save -o $(TARGET_DIR)/$(IMAGE_NAME)-$(IMAGE_TAG).tar $(TAG)
	docker save -o $(TARGET_DIR)/$(IMAGE_NAME).tar $(TAG)
	gzip -f $(TARGET_DIR)/$(IMAGE_NAME).tar
	docker image prune -f

# 推送到镜像仓库
.PHONY: push_docker
push_docker:
	@echo "Pushing image with tag '$(IMAGE_NAME):$(IMAGE_TAG)' to repository '$(REPOSITORY):$(IMAGE_TAG)'"
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(REPOSITORY):$(IMAGE_TAG)
	docker login -u admin -p 123456 192.168.1.1:443
	docker push $(REPOSITORY):$(IMAGE_TAG)

#生成swag接口文档
.PHONY: swag
swag:
	swag init -g ./cmd/ha/main.go