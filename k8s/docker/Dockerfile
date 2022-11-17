# Install the container's OS
FROM ubuntu:latest as HUGOINSTALL

# ENVIREMENT
ENV HUGO_SITE=/hugo-site

# Install Hugo
RUN apt-get update
RUN apt-get -y install apt-utils
RUN apt-get -y install curl
RUN apt-get -y install wget
RUN apt-get -y install dpkg
RUN apt-get -y install tar
RUN apt-get -y install git

RUN curl -s https://api.github.com/repos/gohugoio/hugo/releases/latest \
 | grep  browser_download_url \
 | grep linux-amd64.deb \
 | grep -v extended \
 | cut -d '"' -f 4 \
 | wget -i -
RUN dpkg -i hugo*_linux-amd64.deb

# Install Lescatit web
WORKDIR ${HUGO_SITE}
RUN LATEST_RELEASE=$(curl -s https://api.github.com/repos/mtnmunuklu/LescatitSite/releases/latest \
 | grep tarball_url \
 | cut -d '"' -f 4) \
 && curl -LkSs $LATEST_RELEASE | tar -xz -C . --strip-components=1

RUN mkdir -p themes && cd themes && git clone https://github.com/jeblister/kube.git

RUN hugo -v --source=${HUGO_SITE} \
    --destination=${HUGO_SITE}/public

# Install NGINX and deactivate NGINX's default index.html file
# Move the static site files to NGINX's html directory
# This directory is where the static site files will be served from by NGINX
FROM nginx:stable-alpine
RUN mv /usr/share/nginx/html/index.html /usr/share/nginx/html/old-index.html
COPY --from=HUGOINSTALL ${HUGO_SITE}/public/ /usr/share/nginx/html/

# The container will listen on port 80 using the TCP protocol
EXPOSE 80

# Set up Services
# Create a directory for services
WORKDIR /app
# Copy all services to app directory
COPY authsvc /app/authsvc
COPY crawlsvc /app/crawlsvc
COPY catzesvc /app/catzesvc
COPY catsvc /app/catsvc
COPY apisvc /app/apisvc

# The container will listen on the below ports using the tcp protocol
EXPOSE 9001
EXPOSE 9003
EXPOSE 9004
EXPOSE 9002
EXPOSE 9000

# Execute the all services
CMD ["/app/authsvc", "/app/crawlsvc", "/app/catzesvc", "/app/catsvc", "/app/apisvc"]