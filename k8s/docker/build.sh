#!/bin/bash

cp ../../authentication/authsvc .
cp ../../crawler/crawlsvc .
cp ../../categorizer/catzesvc
cp ../../categorization/catsvc .
cp ../../api/apisvc .

docker build -t cws:v1 .
docker inspect cws:v1