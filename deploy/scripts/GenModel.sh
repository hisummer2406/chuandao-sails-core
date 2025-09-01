#!/usr/bin/env bash

#目标表名
tableName=$1
#目标目录
targetDir=$2

#数据库配置
host=$3
port=$4
user=$5
password=$6
dbname=$7

echo "开始生成库：$dbname 的表 $tableName"
# 生成模型代码
goctl model mysql datasource -url="${user}:${password}@tcp(${host}:${port})/${dbname}" -table="${tableName}"  -dir="${targetDir}" -cache=true -style=gozero
