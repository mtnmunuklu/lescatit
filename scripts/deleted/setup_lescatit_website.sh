#!/bin/bash

# Download lescatit-site from github release and extract to the /var/www directory
echo -e "${GREEN}Download lescatit-site from github release${ENDCOLOR}"
LATEST_RELEASE=$(curl -s https://api.github.com/repos/mtnmunuklu/lescatit-site/releases/latest \
| grep browser_download_url \
| grep tar.gz \
| cut -d '"' -f 4)
curl -LJO $LATEST_RELEASE

echo -e "${GREEN}extract to the /var/www directory${ENDCOLOR}"
FILE_NAME=$(echo $LATEST_RELEASE | cut -d '"' -f 9)
tar -xvf $FILE_NAME -C /var/www/

# Install Hugo
echo -e "${GREEN}Install Hugo${ENDCOLOR}"
echo '[daftaupe-hugo]
name=Copr repo for hugo owned by daftaupe
baseurl=https://copr-be.cloud.fedoraproject.org/results/daftaupe/hugo/epel-7-$basearch/
type=rpm-md
skip_if_unavailable=True
gpgcheck=1
gpgkey=https://copr-be.cloud.fedoraproject.org/results/daftaupe/hugo/pubkey.gpg
repo_gpgcheck=0
enabled=1
enabled_metadata=1' | sudo tee -a /etc/yum.repos.d/CentOS-hugo.repo > /dev/null
yum install hugo

# Run hugo command and create public file
echo -e "${GREEN}Run hugo command and create public file${ENDCOLOR}"
cd /var/www/lescatit-site
hugo