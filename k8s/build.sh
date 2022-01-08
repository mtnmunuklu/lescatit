#!/bin/bash

# Centos X

RED="\e[31m"
GREEN="\e[32m"
ENDCOLOR="\e[0m"

## Install tools
echo -e "${GREEN}Install tools${ENDCOLOR}"
#sudo yum update -y
#sudo yum upgrade -y
sudo yum install -y wget

## Install go
echo -e "${GREEN}Install go${ENDCOLOR}"
# To download the Go binary
echo -e "${GREEN}Download the Go binary${ENDCOLOR}"
wget https://dl.google.com/go/go1.17.3.linux-amd64.tar.gz --no-check-certificate
# Extract the archive you downloaded into /usr/local, creating a Go tree in /usr/local/go
echo -e "${GREEN}Extract the archive you downloaded into /usr/local, creating a Go tree in /usr/local/go${ENDCOLOR}"
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
# Add /usr/local/go/bin to the PATH environment variable
echo -e "${GREEN}Add /usr/local/go/bin to the PATH environment variable${ENDCOLOR}"
export PATH=$PATH:/usr/local/go/bin
# Verify that you've installed Go
go version

## Install docker
echo -e "${GREEN}Install docker${ENDCOLOR}"
# To uninstall old versions
echo -e "${GREEN}Uninstall old versions${ENDCOLOR}"
sudo yum remove docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-engine
# Install the yum-utils and set up the stable repository
echo -e "${GREEN}Install the yum-utils and set up the stable repository${ENDCOLOR}"
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
# Install docker engine
echo -e "${GREEN}Install docker engine${ENDCOLOR}"
sudo yum install -y docker-ce docker-ce-cli containerd.io
# Start Docker
echo -e "${GREEN}Start Docker${ENDCOLOR}"
sudo systemctl start docker

# Build all services
echo -e "${GREEN}Build all services${ENDCOLOR}"
go build -o ../authentication/authsvc ../authentication/main.go
go build -o ../crawler/crawlsvc ../crawler/main.go
go build -o ../categorizer/catzesvc ../categorizer/main.go
go build -o ../categorization/catsvc ../categorization/main.go
go build -o ../api/apisvc ../api/main.go

# Copy builted services to docker directory
echo -e "${GREEN}Copy builted services to docker directory${ENDCOLOR}"
cp ../authentication/authsvc docker/
cp ../crawler/crawlsvc docker/
cp ../categorizer/catzesvc docker/
cp ../categorization/catsvc docker/
cp ../api/apisvc docker/

# Build docker file
echo -e "${GREEN}Build docker file${ENDCOLOR}"
docker build -t lescatit:v1 docker/
docker inspect lescatit:v1

# Show docker images
echo -e "${GREEN}Show docker images${ENDCOLOR}"
docker images

## Install kubectl
echo -e "${GREEN}Install kubectl${ENDCOLOR}"
# Download the latest release with the command
echo -e "${GREEN}Download the latest release with the command${ENDCOLOR}"
curl -kLO "https://dl.k8s.io/release/$(curl -k -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
# Download the kubectl checksum file
echo -e "${GREEN}Download the kubectl checksum file${ENDCOLOR}"
curl -kLO "https://dl.k8s.io/$(curl -k -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
# Validate the kubectl binary against the checksum file
echo -e "${GREEN}Validate the kubectl binary against the checksum file${ENDCOLOR}"
echo "$(<kubectl.sha256) kubectl" | sha256sum --check
# Install
echo -e "${GREEN}Install${ENDCOLOR}"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
# Test to ensure the version you installed is up-to-date
echo -e "${GREEN}Test to ensure the version you installed is up-to-date${ENDCOLOR}"
kubectl version --client

## Install minikube
echo -e "${GREEN}Install minikube${ENDCOLOR}"
curl -kLO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
# Start your cluster
echo -e "${GREEN}Start your cluster${ENDCOLOR}"
minikube start --force --driver=docker

# Load lescatit image to minikube
echo -e "${GREEN}Load lescatit image to minikube${ENDCOLOR}"
minikube image load lescatit:v1

# Show minikube docker images
echo -e "${GREEN}Show minikube docker images${ENDCOLOR}"
eval $(minikube docker-env)
docker images

# Apply a configuration for mongodb
echo -e "${GREEN}Apply a configuration for mongodb${ENDCOLOR}"
kubectl apply -f mongodb/

# Generate certificate
echo -e "${GREEN}Generate certificate${ENDCOLOR}"
cd ../certs
bash generate.sh

# Create secret for secure communication between services
echo -e "${GREEN}Create secret for secure communication between services${ENDCOLOR}"
kubectl create secret generic cert-secret --from-file=services/ca-cert.pem --from-file=services/server-cert.pem --from-file=services/server-key.pem

# Create secret for ingress
echo -e "${GREEN}Create secret for ingress${ENDCOLOR}"
kubectl create secret tls ingress-secret --key api/ingress-key.pem --cert api/ingress-cert.pem

# Go kubernetes direcotory
echo -e "${GREEN}Go kubernetes direcotory${ENDCOLOR}"
cd ../k8s

# Enable ingress on minikube
echo -e "${GREEN}Enable ingress on minikube${ENDCOLOR}"
minikube addons enable ingress
# Delete ingrees nginx admission
echo -e "${GREEN}Delete ingrees nginx admission${ENDCOLOR}"
kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission

# Apply a configuration for services
echo -e "${GREEN}Apply a configuration for services${ENDCOLOR}"
kubectl apply -f services/

# Save Minikube IP to api.lescatit.com
echo -e "${GREEN}Save Minikube IP to api.lescatit.com${ENDCOLOR}"
echo "`minikube ip` api.lescatit.com" | sudo tee -a /etc/hosts > /dev/null

# Get all information
echo -e "${GREEN}Get all information${ENDCOLOR}"
kubectl get all