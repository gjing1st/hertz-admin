MODULE = gin-admin-frame

# 运行
.PHONY: run
run:
	go run cmd/ha/main.go

.PHONY: build
build:
	go build -o ha cmd/ha/main.go

# 打包成docker镜像
.PHONY: docker
docker:
	docker build -f ./build/docker/Dockerfile -t ha:latest .
	docker save -o ha.tar ha:latest
	gzip ha.tar
	docker image prune -f
	#docker system prune -f
