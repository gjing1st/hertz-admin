MODULE = hertz-admin


#镜像名称
#IMAGE_NAME := ha
PROJECT_NAME := $(shell grep ProjectName version.go | awk -F '"' '{print $$2}')
IMAGE_NAME := $(shell echo $(PROJECT_NAME) | sed 's/^//')-backend
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
TARGET_DIR := /opt/build-host/$(PROJECT_NAME)/$(IMAGE_TAG)
#TARGET_DIR := ./

# 运行
.PHONY: run
run:
	@echo "Building image with tag '$(IMAGE_TAG)'"
	go run cmd/ha/main.go

.PHONY: build
build:
	go build -o ha cmd/ha/main.go
# 打包成docker镜像
.PHONY: docker
TAG := $(IMAGE_NAME):$(IMAGE_TAG)
docker:
	@echo "Building image with tag '$(TAG)'"
	docker build -f ./build/docker/Dockerfile -t $(TAG) .
	#如果运行在gitlab runner可能创建/opt/build-host/$(PROJECT_NAME)权限不足，需要到135服务器手动创建项目目录，并将目录权限更改为gitlab-runner
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

# 自动打包(根据dbes编写,其他产品线需自行修改)
.PHONY: install_package
#获取执行make时，传过来的变量值,直接使用
#COMMIT_COUNT := $(COMMIT_COUNT)
INSTALL_PATH := /opt/build-host/$(PROJECT_NAME)/$(PROJECT_NAME)_$(IMAGE_TAG).${COMMIT_COUNT}
install_package:
	@echo "COMMIT_COUNT '$(COMMIT_COUNT)' install package install_path '$(INSTALL_PATH)'"
	# 创建/opt/build-host/dbes/dbes_v0.2.0文件夹，用来打包v0.2.0版本安装包
	mkdir -p $(INSTALL_PATH)
	# 将该版本的backend安装包拷贝到打包目录下，此次只拷贝了backend镜像包。该段Makefile可拷贝到其他项目，修改执行打包操作
	#cp -f $(TARGET_DIR)/$(IMAGE_NAME)-$(IMAGE_TAG).tar.gz  $(INSTALL_PATH)/$(IMAGE_NAME).tar.gz
	#将该版本下的文件全部拷贝到打包目录，包含前端后端和其他文件
	cp -rf $(TARGET_DIR)/*  $(INSTALL_PATH)/
	# 每个版本有新的变动，写到./deployments/docker-compose/目录下，该目录下应包含install.sh(安装脚本)和update.sh（升级脚本）和其他项目依赖项
	# install.sh和update.sh需要保持backend中是最新的，每次从该git仓库修改和打包。update.sh需要修改请求backend地址
	cp -rf ./deployments/docker-compose/* $(INSTALL_PATH)/
	# dbes下需要写proxy代理的配置文件，为方便，这里写到了config目录下
	cp -rf ./config $(INSTALL_PATH)/
	echo "{\"version\":\"$(shell echo "$(IMAGE_TAG).${COMMIT_COUNT}" | cut -c2-)\"}" > $(INSTALL_PATH)/config/version.json
	# common文件夹中包含助手，自检服务，config/config.yml配置文件(避免本地调试配置文件覆盖线上标准配置)等
	# 每个产品线，需要单独维护产品线下的common文件夹
	cp -rf $(INSTALL_PATH)/../common/* $(INSTALL_PATH)/
	#转换windows下写的sh为unix，其他sh自行添加
	dos2unix $(INSTALL_PATH)/install.sh
	dos2unix $(INSTALL_PATH)/update.sh
	cd $(INSTALL_PATH)/..
	# 打包成tar.gz安装包
	chmod +x $(INSTALL_PATH)/install.sh
	chmod +x $(INSTALL_PATH)/update.sh
	tar -zcvf $(PROJECT_NAME)_$(IMAGE_TAG).$(COMMIT_COUNT).tar.gz  -C /opt/build-host/$(PROJECT_NAME) $(PROJECT_NAME)_$(IMAGE_TAG).$(COMMIT_COUNT)
	mv -f $(PROJECT_NAME)_$(IMAGE_TAG).$(COMMIT_COUNT).tar.gz $(INSTALL_PATH)/..
	rm -rf $(INSTALL_PATH)


#生成swag接口文档
.PHONY: swag
swag:
	swag init -g ./cmd/ha/main.go