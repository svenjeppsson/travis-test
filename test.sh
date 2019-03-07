#!/usr/bin/env bash
export DBPORT=3306
export DBCON="root:secret@tcp(localhost:${DBPORT})/TEST"
./start-db-container.sh
dep ensure
go fmt ./...
golangci-lint run
go vet .
go test -v ./... -run ^TestIntegration -coverprofile=cov.out
