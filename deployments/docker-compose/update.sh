#!/bin/bash

path="$(cd "$(dirname "$0")" && pwd)"
cd $path
docker load < ha-server.tar.gz
docker load < ha-frontend.tar.gz
#重启docker compose
mv -rf /home/app/ha/docker-compose.yml /home/app/ha/docker-compose.yml.bak
cp -rf $path/docker-compose.yml /home/app/ha/docker-compose.yml
docker-compose -f /home/app/ha/docker-compose.yml up -d
echo "finish"
