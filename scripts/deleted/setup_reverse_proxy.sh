#!/bin/bash

# Centos X

# Install Apache
echo -e "${GREEN}Install Apache${ENDCOLOR}"
sudo dnf install httpd

# Enable and start Apache
echo -e "${GREEN}Enable and start Apache${ENDCOLOR}"
sudo systemctl enable httpd
sudo systemctl start httpd

# Firewall set up, open up the http and https ports
echo -e "${GREEN}Firewall set up, open up the http and https ports${ENDCOLOR}"
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --reload

# Install mod_ssl
echo -e "${GREEN}Install mod_ssl${ENDCOLOR}"
sudo dnf install mod_ssl

# Restart Apache
echo -e "${GREEN}Restart Apache${ENDCOLOR}"
sudo systemctl restart httpd

# Copy lescatit certificates to the /etc/pki/tls
echo -e "${GREEN}Copy lescatit certificates to the /etc/pki/tls${ENDCOLOR}"
cp ../certs/api/lescatit-cert.pem /etc/pki/tls/certs/
cp ../certs/api/lescatit-key.pem /etc/pki/tls/private/

# Setup lescatit website
echo -e "${GREEN}Setup lescatit website${ENDCOLOR}"
bash setup_lescatit_website.sh

# Assign the permissions to the Document root
echo -e "${GREEN}Assign the permissions to the Document root${ENDCOLOR}"
sudo chown -R apache:apache /var/www/lescatit-site

# Open a new file in the /etc/httpd/conf.d directory
echo -e "${GREEN}Open a new file in the /etc/httpd/conf.d directory${ENDCOLOR}"
echo "<VirtualHost *:443>
  ServerName lescatit.com
  ServerAlias www.lescatit.com

  DocumentRoot /var/www/lescatit-site/public/
  <Directory /var/www/lescatit-site/public/>
    Options -Indexes +FollowSymLinks
    AllowOverride All
  </Directory>

  ErrorLog /var/log/httpd/error.log
  CustomLog /var/log/httpd/access.log combined

  SSLEngine On
  SSLProxyEngine on
  SSLProxyVerify none
  SSLProxyCheckPeerCN off
  SSLProxyCheckPeerName off
  SSLProxyCheckPeerExpire off
  SSLCertificateFile    /etc/pki/tls/certs/lescatit-cert.pem
  SSLCertificateKeyFile  /etc/pki/tls/private/lescatit-key.pem

  <Location /api>
    ProxyPreserveHost On
    ProxyPass https://lescatit.com/api
    ProxyPassReverse https://lescatit.com/api
  </Location>

</VirtualHost>

<VirtualHost *:80>
  ServerName lescatit.com
  ServerAlias www.lescatit.com
  Redirect permanent / https://lescatit.com/
</VirtualHost>" | sudo tee -a /etc/httpd/conf.d/lescatit.conf > /dev/null

# Check your Apache configuration for syntax errors
echo -e "${GREEN}Check your Apache configuration for syntax errors${ENDCOLOR}"
sudo apachectl configtest

# Reload Apache to pick up the configuration changes
echo -e "${GREEN}Reload Apache to pick up the configuration changes${ENDCOLOR}"
sudo systemctl reload httpd
