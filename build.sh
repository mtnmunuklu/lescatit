#!/bin/bash
go clean --cache && go test -v -cover CWS/...
go build -o authentication/authsvc authentication/main.go
go build -o api/apisvc api/main.go
go build -o categorization/catsvc categorization/main.go