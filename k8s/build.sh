#!/bin/bash

# Copy builted services
cp ../authentication/authsvc docker/
cp ../crawler/crawlsvc docker/
cp ../categorizer/catzesvc docker/
cp ../categorization/catsvc docker/
cp ../api/apisvc docker/

# Build docker file
docker build -t lescatit:v1 docker/
docker inspect lescatit:v1

# Apply a configuration for mongodb
kubectl apply -f mongodb/

# Create secret for secure communication between services
kubectl create secret generic cert-secret \ 
    --from-file=../certs/ca-cert.pem \
    --from-file=../certs/server-cert.pem \
    --from-file=../certs/server-key.pem

# Create secret for ingress
kubectl create secret tls ingress-secret \
    --key ingress-key.pem \
    --cert ingress-cert.pem

# Apply a configuration for services
kubectl apply -f services/