#!/bin/bash

#sudo yum update -y
#sudo yum upgrade -y
sudo yum install -y wget tar

# Install Go
echo -e "${GREEN}Install Go${ENDCOLOR}"
# To download the Go binary
echo -e "${GREEN}Download the Go binary${ENDCOLOR}"
wget https://dl.google.com/go/go1.19.2.linux-amd64.tar.gz --no-check-certificate
# Extract the archive you downloaded into /usr/local, creating a Go tree in /usr/local/go
echo -e "${GREEN}Extract the archive you downloaded into /usr/local, creating a Go tree in /usr/local/go${ENDCOLOR}"
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz
# Add /usr/local/go/bin to the PATH environment variable
echo -e "${GREEN}Add /usr/local/go/bin to the PATH environment variable${ENDCOLOR}"
export PATH=$PATH:/usr/local/go/bin
# Verify that you've installed Go
echo -e "${GREEN}Verify that you've installed Go${ENDCOLOR}"
go version

# Install Docker
echo -e "${GREEN}Install Docker${ENDCOLOR}"
# To uninstall old versions
echo -e "${GREEN}Uninstall old versions${ENDCOLOR}"
sudo yum remove docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-engine
# Install the yum-utils and set up the stable repository
echo -e "${GREEN}Install the yum-utils and set up the stable repository${ENDCOLOR}"
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
# Install docker engine
echo -e "${GREEN}Install docker engine${ENDCOLOR}"
sudo yum install -y docker-ce docker-ce-cli
# Enable Docker
echo -e "${GREEN}Enable Docker${ENDCOLOR}"
sudo systemctl enable docker
# Start Docker
echo -e "${GREEN}Start Docker${ENDCOLOR}"
sudo systemctl start docker