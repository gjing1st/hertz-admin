#!/bin/bash
TOPDIR=$(dirname $0)

if [ "$TOPDIR" != "." ]; then
	cd $TOPDIR
fi

set -e
echo "install ..."
#创建prometheus存储映射目录,注意修改产品线名称
mkdir -p /home/app/ha
cp -r ./* /home/app/ha
cd /home/app/ha
echo "Loading docker images ..."
docker load < ha-server.tar.gz
docker load < ha-frontend.tar.gz
echo "Starting containers by docker-compose ..."
docker-compose up -d

echo "Done"
