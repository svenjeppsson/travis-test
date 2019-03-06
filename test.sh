#!/usr/bin/env bash
export DBPORT=3306
export DBCON="root:secret@tcp(localhost:${DBPORT})/TEST"
./start-db-container.sh
dep ensure
go test -short -coverprofile=cov.out `go list./..|grep -v vendor/`
docker stop testdb