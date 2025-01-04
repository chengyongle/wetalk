#!/bin/bash
reso_addr='crpi-mnb3psxzsxgtr7db.cn-shanghai.personal.cr.aliyuncs.com/we_talk/im-api'
tag='latest'

container_name="we_talk-im-api"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-im -v /easy-im/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 8882:8882  --name=${container_name} -d ${reso_addr}:${tag}