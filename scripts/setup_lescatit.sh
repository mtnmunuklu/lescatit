#!/bin/bash

# Centos X

RED="\e[31m"
GREEN="\e[32m"
ENDCOLOR="\e[0m"

# Apply traefik configuration
echo -e "${GREEN}Apply traefik configuration${ENDCOLOR}"
kubectl apply -f ../k8s/traefik/00-role.yml \
              -f ../k8s/traefik/00-account.yml \
              -f ../k8s/traefik/01-role-binding.yml \
              -f ../k8s/traefik/02-traefik.yml \
              -f ../k8s/traefik/02-traefik-services.yml

# Apply common configurations
echo -e "${GREEN}Apply common configurations${ENDCOLOR}"
kubectl apply -f ../k8s/common/00-scripts.yml \
              -f ../k8s/common/01-ds.yml
              -f ../k8s/common/02-stc.yml
              -f ../k8s/common/03-pv.yml

# Apply mongodb configuration
echo -e "${GREEN}Apply mongodb configuration${ENDCOLOR}"
kubectl apply -f ../k8s/mongodb/00-stc.yml \
              -f ../k8s/mongodb/01-pv.yml \
              -f ../k8s/mongodb/02-pvc.yml \
              -f ../k8s/mongodb/03-configs.yml \
              -f ../k8s/mongodb/04-scripts.yml \
              -f ../k8s/mongodb/05-secrets.yml \
              -f ../k8s/mongodb/06-sts.yml

# Generate certificates for docker registry, services and ingress
echo -e "${GREEN}Generate certificates for docker registry, services and ingress${ENDCOLOR}"
bash certs/generate_certificates.sh

# Create secret for secure communication between services
echo -e "${GREEN}Create secret for secure communication between services${ENDCOLOR}"
kubectl create secret generic service-certs --from-file=../certs/services/ca-cert.pem --from-file=../certs/services/server-cert.pem --from-file=../certs/services/server-key.pem

# Create secret for ingress
echo -e "${GREEN}Create secret for ingress${ENDCOLOR}"
kubectl create secret tls ingress-certs --key ../certs/api/lescatit-key.pem --cert ../certs/api/lescatit-cert.pem

# Apply service configuration
echo -e "${GREEN}Apply service configuration${ENDCOLOR}"
kubectl apply -f ../k8s/services/00-configs.yml \
              -f ../k8s/services/01-secrets.yml \
              -f ../k8s/services/02-authentication.yml \
              -f ../k8s/services/03-categorization.yml \
              -f ../k8s/services/04-categorizer.yml \
              -f ../k8s/services/05-crawler.yml \
              -f ../k8s/services/06-api.yml \
              -f ../k8s/services/07-web.yml \
              -f ../k8s/services/08-tls_ingress.yml

# Get all information
echo -e "${GREEN}Get all information${ENDCOLOR}"
kubectl get all