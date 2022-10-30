#!/bin/bash

# Remove old certificates
echo -e "${GREEN}Remove old certificates${ENDCOLOR}"
rm -f ../certs/api/*.pem
rm -f ../certs/services/*.pem
rm -f ../certs/services/*.srl

# Creation of certificates for secure communication between services
# 1. Generate CA's private key and self-signed certificate
echo -e "${GREEN}1. Generate CA's private key and self-signed certificate${ENDCOLOR}"
openssl req -x509 -newkey rsa:4096 -days 365 -nodes \
    -keyout ../certs/services/ca-key.pem -out ../certs/services/ca-cert.pem \
    -addext "subjectAltName = DNS:localhost, DNS:auth-service, DNS:crawl-service, DNS:catze-service, DNS:cat-service, IP:0.0.0.0" \
    -subj "/C=TR/ST=Istanbul/L=DavutPasa/O=Lescatit/OU=Software/CN=*.lescatit.com/emailAddress=lescatit@gmail.com"

# CA's self-signed certificate
echo -e "${GREEN}CA's self-signed certificate${ENDCOLOR}"
openssl x509 -in ../certs/services/ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
echo -e "${GREEN}2. Generate web server's private key and certificate signing request (CSR)${ENDCOLOR}"
openssl req -newkey rsa:4096 -nodes -keyout ../certs/services/server-key.pem \
    -out ../certs/services/server-req.pem \
    -addext "subjectAltName = DNS:localhost, DNS:auth-service, DNS:crawl-service, DNS:catze-service, DNS:cat-service, IP:0.0.0.0" \
    -subj "/C=TR/ST=Istanbul/L=DavutPasa/O=Lescatit/OU=Software/CN=*.lescatit.com/emailAddress=lescatit@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
echo -e "${GREEN}3. Use CA's private key to sign web server's CSR and get back the signed certificate${ENDCOLOR}"
openssl x509 -req -in ../certs/services/server-req.pem  \
    -days 60 -CA ../certs/services/ca-cert.pem \
    -CAkey ../certs/services/ca-key.pem -CAcreateserial -out ../certs/services/server-cert.pem

# Server's signed certificates
echo -e "${GREEN}Server's signed certificates${ENDCOLOR}"
openssl x509 -in ../certs/services/server-cert.pem -noout -text

# Creation of certificates for ingress
echo -e "${GREEN}Creation of certificates for ingress${ENDCOLOR}"
openssl req -x509 -nodes -days 365 -newkey rsa:4096 \
    -out ../certs/api/lescatit-cert.pem \
    -keyout ../certs/api/lescatit-key.pem \
    -addext "subjectAltName = DNS:lescatit.com, DNS:api.lescatit.com, DNS:api-service, DNS:hugo-service" \
    -subj "/C=TR/ST=Istanbul/L=DavutPasa/O=Lescatit/OU=Software/CN=*.lescatit.com/emailAddress=lescatit@gmail.com"

# Creation of certificates for docker registry
echo -e "${GREEN}Creation of certificates for docker registry${ENDCOLOR}"
openssl req -x509 -nodes -days 365 -newkey rsa:4096 \
    -out ../certs/docker/registry-cert.pem \
    -keyout ../certs/docker/registry-key.pem \
    -addext "subjectAltName = DNS:registry-service" \
    -subj "/C=TR/ST=Istanbul/L=DavutPasa/O=Lescatit/OU=Software/CN=*.lescatit.com/emailAddress=lescatit@gmail.com"