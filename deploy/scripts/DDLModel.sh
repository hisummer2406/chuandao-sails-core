#!/usr/bin/env bash

templateHome=../goctl/v1.8.5

echo "开始根据sql文件生成"

goctl model mysql ddl \
  -src ./sql/order_status_log.sql \
  -dir ./model \
  -c \
  --home="${templateHome}"