#!/usr/bin/env bash
docker run --rm  -p $1:3306 -v $(pwd)/schema.sql:/docker-entrypoint-initdb.d/schema.sql --name $2 -e MYSQL_ROOT_PASSWORD=$3 --health-cmd='mysqladmin ping --silent' -d mariadb:10
