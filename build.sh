#!/bin/bash
#go clean --cache && go test -v -cover ./security/... ./authentication/... ./crawler/... ./categorizer/... ./categorization/...
go build -o authentication/authsvc authentication/main.go
go build -o crawler/crawlsvc crawler/main.go
go build -o categorizer/catzesvc categorizer/main.go
go build -o categorization/catsvc categorization/main.go
go build -o api/apisvc api/main.go