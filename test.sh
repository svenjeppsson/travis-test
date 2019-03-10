#!/usr/bin/env bash
export DOCKER_PASSWORD=Ty?JrvEfB#2DKE+qiCK3
export DOCKER_USERNAME=xthinker
export DB_CONTAINER=testdb
export DB_USER=root
export GO_BUILDER_VERSION=0.0.8
export DB_PASSWORD=secret

docker run --rm  -v $(pwd)/schema.sql:/docker-entrypoint-initdb.d/schema.sql --name $DB_CONTAINER \
 -e "MYSQL_ROOT_PASSWORD=secret" \
 --health-cmd='mysqladmin ping --silent' -d mariadb:10

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker pull xthinker/go-builder:$GO_BUILDER_VERSION
./waitForContainer.sh $DB_CONTAINER
docker run --rm -ti --link $DB_CONTAINER:mysql \
 -e "DBCON=root:secret@tcp(mysql:3306)/TEST" \
 -e "UID=$(id -u)" \
 -e "GID=$(id -u)" \
 -v $(pwd):/go/src/app xthinker/go-builder:$GO_BUILDER_VERSION
docker stop $DB_CONTAINER
docker container prune -f

