# Install the container's OS
FROM ubuntu:latest as builder

# ENVIREMENT
ENV WEBSITE=/website
ENV APP=/app

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
WORKDIR ${WEBSITE}
RUN LATEST_RELEASE=$(curl -s https://api.github.com/repos/mtnmunuklu/LescatitSite/releases/latest \
 | grep tarball_url \
 | cut -d '"' -f 4) \
 && curl -LkSs $LATEST_RELEASE | tar -xz -C . --strip-components=1

RUN mkdir -p themes && cd themes && git clone https://github.com/jeblister/kube.git

RUN hugo -v --source=${WEBSITE} \
    --destination=${WEBSITE}/public/

# Set up Services
# Create a directory for services
WORKDIR ${APP}
# Copy all services to app directory
COPY authsvc authsvc
COPY crawlsvc crawlsvc
COPY catzesvc catzesvc
COPY catsvc catsvc
COPY apisvc apisvc

# Install NGINX and deactivate NGINX's default index.html file
# Move the static site files to NGINX's html directory
# This directory is where the static site files will be served from by NGINX
FROM nginx:stable-alpine
RUN apk add --no-cache libc6-compat 
RUN mv /usr/share/nginx/html/index.html /usr/share/nginx/html/old-index.html
COPY --from=builder /website/public /usr/share/nginx/html
COPY --from=builder /app /app

# Expose port 80 and ports between 9000-9004
EXPOSE 80
EXPOSE 9001
EXPOSE 9003
EXPOSE 9004
EXPOSE 9002
EXPOSE 9000