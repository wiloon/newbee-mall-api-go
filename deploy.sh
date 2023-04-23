#!/bin/bash
package_name="mall-api"
GOOS=linux GOARCH=amd64 go build -o ${package_name} main.go

ansible -i 'wiloon.com,' all  -m shell -a 'systemctl stop mall' -u root
scp config-aliyun.yaml aliyun:/data/mall/config.yaml
scp ${package_name} aliyun:/data/mall
ansible -i 'wiloon.com,' all  -m shell -a 'systemctl start mall' -u root
