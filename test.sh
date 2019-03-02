#!/usr/bin/env bash

function getContainerHealth {
  docker inspect --format "{{json .State.Health.Status }}" $1
}

function waitContainer {
  while STATUS=$(getContainerHealth $1); [ $STATUS != "\"healthy\"" ]; do
    if [ $STATUS == "\"unhealthy\"" ]; then
      echo "Failed!"
      exit -1
    fi
    printf .
    lf=$'\n'
    sleep 1
  done
  printf "$lf"
}

docker run --rm -v $(pwd)/schema.sql:/docker-entrypoint-initdb.d/schema.sql --name testdb -e MYSQL_ROOT_PASSWORD=secret --health-cmd='mysqladmin ping --silent' -d mariadb:10
waitContainer testdb
docker build . -t go-alltest
docker run --rm --link testdb:mysql -v "`pwd`/.cache/pkg:/go/pkg" -v "`pwd`:/go/src/app"  go-alltest
docker stop testdb
