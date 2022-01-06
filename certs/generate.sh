rm -f api/*.pem
rm -f services/*.pem
rm -f services/*.srl

# Creation of certificates for secure communication between services
# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes \
    -keyout services/ca-key.pem -out services/ca-cert.pem \
    -subj "/C=TR/ST=Istanbul/L=DavutPasa/O=Lescatit/OU=Software/CN=*.lescatit.com/emailAddress=lescatit@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in services/ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout services/server-key.pem \
    -out services/server-req.pem \
    -subj "/C=TR/ST=Istanbul/L=DavutPasa/O=Lescatit/OU=Software/CN=*.lescatit.com/emailAddress=lescatit@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in services/server-req.pem  \
    -days 60 -CA services/ca-cert.pem \
    -CAkey services/ca-key.pem -CAcreateserial -out services/server-cert.pem -extfile services/server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in services/server-cert.pem -noout -text

# Creation of certificates for ingress
openssl req -x509 -nodes -days 365 -newkey rsa:4096 \
    -out api/ingress-cert.pem \
    -keyout api/ingress-key.pem \
    -subj "/C=TR/ST=Istanbul/L=DavutPasa/O=Lescatit/OU=Software/CN=*.lescatit.com/emailAddress=lescatit@gmail.com"