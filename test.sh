#!/usr/bin/env bash
export DBPORT=3306
export DBCON="root:secret@tcp(localhost:${DBPORT})/TEST"
./start-db-container.sh
dep ensure
golangci-lint run
go vet .
go test -short -coverprofile=cov.out
