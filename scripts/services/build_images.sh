#!/bin/bash

# Copy builted services to docker directory
echo -e "${GREEN}Copy builted services to docker directory${ENDCOLOR}"
cp ../authentication/authsvc ../k8s/docker
cp ../crawler/crawlsvc ../k8s/docker
cp ../categorizer/catzesvc ../k8s/docker
cp ../categorization/catsvc ../k8s/docker
cp ../api/apisvc ../k8s/docker

# Build app docker file
echo -e "${GREEN}Build docker file${ENDCOLOR}"
docker build -t mtnmunuklu/lescatit:v0.1.0 ../k8s/docker/

# Push app image to local registry
echo -e "${GREEN}Push app image to local registry${ENDCOLOR}"
docker push mtnmunuklu/lescatit:v0.1.0

# Show docker images
echo -e "${GREEN}Show docker images${ENDCOLOR}"
docker images
