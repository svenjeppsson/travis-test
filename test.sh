#!/usr/bin/env bash
./start-db-container.sh 3307 testdb secret
docker run -ti --link testdb:mysql -- $(id -u):$(id -g) -e "DBCON=root:secret@tcp(mysql:3307)/TEST" -v $(pwd):/go/src/app xthinker/go-builder:latest
