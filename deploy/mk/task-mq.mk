VERSION=latest

SERVER_NAME=task
SERVER_TYPE=mq

# 环境配置
# docker的镜像发布地址
DOCKER_REPO=crpi-mnb3psxzsxgtr7db.cn-shanghai.personal.cr.aliyuncs.com/we_talk/${SERVER_NAME}-${SERVER_TYPE}
# 版本
VERSION=$(VERSION)
# 编译的程序名称
APP_NAME=wetalk-${SERVER_NAME}-${SERVER_TYPE}

# 编译文件
DOCKER_FILE=./deploy/dockerfile/Dockerfile_${SERVER_NAME}_${SERVER_TYPE}

# 环境的编译发布
build:

	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${SERVER_NAME}-${SERVER_TYPE} ./apps/${SERVER_NAME}/${SERVER_TYPE}/${SERVER_NAME}.go
	docker build . -f ${DOCKER_FILE} --no-cache -t ${APP_NAME}

# 镜像标签
tag:

	@echo 'create tag ${VERSION}'
	docker tag ${APP_NAME} ${DOCKER_REPO}:${VERSION}

publish:

	@echo 'publish ${VERSION} to ${DOCKER_REPO}'
	docker push $(DOCKER_REPO):${VERSION}

release-test: build tag publish