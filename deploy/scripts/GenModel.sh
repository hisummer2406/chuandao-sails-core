#!/usr/bin/env bash

#目标表名
tableName=city_region_mapping
#目标目录
targetDir=./model

#数据库配置
host=localhost
port=3306
user=root
password=pwd123
dbname=cdts_platform_db
templateHome=../goctl/v1.8.5

echo "开始生成库：$dbname 的表 $tableName"
# 生成模型代码
goctl model mysql datasource \
-url="${user}:${password}@tcp(${host}:${port})/${dbname}" \
-table="${tableName}"  \
--dir="${targetDir}" \
--cache=true \
--style=gozero \
--home="${templateHome}"
