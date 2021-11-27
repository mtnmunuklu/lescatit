#!/bin/bash

cp ../authentication/authsvc docker/
cp ../crawler/crawlsvc docker/
cp ../categorizer/catzesvc docker/
cp ../categorization/catsvc docker/
cp ../api/apisvc docker/

docker build -t cws:v1 docker/
docker inspect cws:v1

kubectl apply -f mongodb/

kubectl create secret generic certs --from-file=../certs/ca-cert.pem --from-file=../certs/server-cert.pem --from-file=../certs/server-key.pem

kubectl apply -f services/