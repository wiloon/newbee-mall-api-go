#!/bin/bash
package_name="mall-api"
echo "building"
GOOS=linux GOARCH=amd64 go build -o ${package_name} main.go
echo "stop service"
ansible -i 'wiloon.com,' all  -m shell -a 'systemctl stop mall' -u root
scp ${package_name} aliyun:/data/mall
scp config-aliyun.yaml aliyun:/data/mall/config.yaml
cp ${package_name} /Users/wiloon/tmp
echo "start service"
ansible -i 'wiloon.com,' all  -m shell -a 'systemctl start mall' -u root

