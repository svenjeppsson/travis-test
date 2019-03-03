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

docker run --rm  -p localhost:3307:3306 -v $(pwd)/schema.sql:/docker-entrypoint-initdb.d/schema.sql --name testdb -e MYSQL_ROOT_PASSWORD=secret --health-cmd='mysqladmin ping --silent' -d mariadb:10
waitContainer testdb
export DBCON="root:secret@tcp(localhost:3307)/test"
