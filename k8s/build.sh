#!/bin/bash

# Centos X

## Install go
# To download the Go binary
wget https://dl.google.com/go/go1.17.3.linux-amd64.tar.gz
# Extract the archive you downloaded into /usr/local, creating a Go tree in /usr/local/go
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
# Add /usr/local/go/bin to the PATH environment variable
export PATH=$PATH:/usr/local/go/bin
# Verify that you've installed Go
go version

## Install kubectl
# Download the latest release with the command
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
# Download the kubectl checksum file:
curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
# Validate the kubectl binary against the checksum file
echo "$(<kubectl.sha256) kubectl" | sha256sum --check
# Install
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
# Test to ensure the version you installed is up-to-date
kubectl version --client

## Install minikube
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
# Start your cluster
minikube start

# Show minikube docker images
eval $(minikube docker-env)
docker images

# Build all services
go build -o ../authentication/authsvc authentication/main.go
go build -o ../crawler/crawlsvc crawler/main.go
go build -o ../categorizer/catzesvc categorizer/main.go
go build -o ../categorization/catsvc categorization/main.go
go build -o ../api/apisvc api/main.go

# Copy builted services
cp ../authentication/authsvc docker/
cp ../crawler/crawlsvc docker/
cp ../categorizer/catzesvc docker/
cp ../categorization/catsvc docker/
cp ../api/apisvc docker/

# Build docker file
docker build -t lescatit:v1 docker/
docker inspect lescatit:v1

# Show docker images
docker images

# Apply a configuration for mongodb
kubectl apply -f mongodb/

# Generate certificate
cd ../certs
bash generate.sh

# Create secret for secure communication between services
kubectl create secret generic cert-secret \ 
    --from-file=ca-cert.pem \
    --from-file=server-cert.pem \
    --from-file=server-key.pem

# Create secret for ingress
kubectl create secret tls ingress-secret \
    --key ingress-key.pem \
    --cert ingress-cert.pem

# Go kubernetes direcotory
cd ../k8s

# Enable ingress on minikube
minikube addons enable ingress

# Apply a configuration for services
kubectl apply -f services/

# Get all information
kubectl get all