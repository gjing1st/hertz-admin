#!/bin/bash
# 该文件为hss升级脚本，需改写。
#2.7(){
# echo '开始升级至2.7版本'
# echo '注意：2.7版本需要手动更新助手服务'
# systemctl stop hssbackup
# chmod 777 hssbackup
# /bin/cp -rf hssbackup /home/app/hss/hssbackup
# systemctl start hssbackup
# echo '备份服务升级完成'
#
# #升级更高版本
# 2.8
#}
#2.8(){
# #2.8升级2.9需要新增的配置文件
# echo '开始升级至2.8版本'
# 2.9
#}
#2.9(){
# echo '开始升级至2.9版本'
#}
## 通过backend版本信息接口，获取当前版本
#res=$(curl -s http://127.0.0.1:9628/hss/v1/sys/version/info)
#v=$(echo $res | grep -Po '"version":(.+?),' | grep -Po '\d+\.\d+'  )
##获取版本信息，去除版本中首字母v，最后结果如：2.8
#v=${v#v}
##echo $v
##执行2.8版本升级2.9要做的操作，版本依次升级。
## 每个版本升级中，不重启docker-compose，只升级版本中需要特殊变化的内容。如添加了某项配置，增加了某个服务
#${v}
#导入镜像
path="$(cd "$(dirname "$0")" && pwd)"
cd $path
docker load < admin-backend.tar.gz
docker load < admin-frontend.tar.gz
#docker load < admin-fulight.tar.gz
#重启docker compose
#cd /home/app/hss
docker-compose -f /home/app/admin/docker-compose.yml down
#5.复制新的docker-compose文件并重启
mv -rf /home/app/admin/docker-compose.yml /home/app/admin/docker-compose.yml.bak
cp -rf $path/docker-compose.yml /home/app/admin/docker-compose.yml

docker-compose -f /home/app/admin/docker-compose.yml up -d
echo "finish"
