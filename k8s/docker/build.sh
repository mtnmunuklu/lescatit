#!/bin/bash

cp ../../authentication/authsvc .
cp ../../api/apisvc .

docker build -t cws:v1 .
docker inspect cws:v1