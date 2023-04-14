#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o newbee-mall-api main.go
scp newbee-mall-api aliyun:/data/newbee
scp config-aliyun.yaml aliyun:/data/newbee/config.yaml
