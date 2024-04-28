#!/bin/bash
TOPDIR=$(dirname $0)

if [ "$TOPDIR" != "." ]; then
	cd $TOPDIR
fi

set -e
echo "install ..."
#创建prometheus存储映射目录,注意修改产品线名称
mkdir -p /home/app/admin
cp -r ./* /home/app/admin
mkdir -p /home/app/admin/opengauss
chmod -R 755 /home/app/admin/*
cat > /usr/lib/systemd/system/assist.service << _DEVINFO
[Unit]
Description=assist
After=network.target
[Service]
Type=simple
ExecStart= /home/app/admin/assist
ExecStop=/bin/kill $MAINPID
Restart=on-failure
[Install]
WantedBy=multi-user.target
_DEVINFO
systemctl enable assist.service
systemctl start assist.service
cd /home/app/admin
rpm -ivh selftest-1.0.0-2.x86_64.rpm --force --nodeps
echo "Loading docker images ..."
docker load < admin-backend.tar.gz
docker load < admin-frontend.tar.gz
docker load < opengauss.tar.gz
echo "Starting containers by docker-compose ..."
docker-compose down
docker-compose up -d

echo "Done"
