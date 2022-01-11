<p align="center">
  <img width="250" height="300" src="images/logo.png">
</p>

<p align="center">
<a href="https://pkg.go.dev/"><img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-informational.svg" alt="Go Doc"></a> <a href="https://goreportcard.com/"><img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-X+-success.svg" alt="Go Report"></a> <a href="https://gocover.io/"><img src="https://img.shields.io/badge/%F0%9F%94%8E%20gocover-X%25-success.svg" alt="Coverage Status"></a> <a href="https://travis-ci.com/"><img src="https://img.shields.io/badge/%E2%9A%99%20build-X-success.svg" alt="Build Status"></a> <a href="https://lescatit.com/"><img src="https://img.shields.io/badge/%F0%9F%93%BD%20demo-online-red.svg" alt="Live Demo"></a>
</p>

# Lescatit <sub><small><small>(Let's categorized it)</small></small></sub>

It provides to crawl and categorize URL addresses. It is developed with go, mongo, docker and kubernetes technologies.

## Table of contents
* [Features](#features)
* [Setup](#setup)
* [Usage](#usage)
* [License](#license)

## Features
This project has the following features:
* Getting the user(s) information
* Delete a user
* Changing the user role
* Update user passsword
* Update user email address
* Update username
* Getting content of the URL(s)
* Crawling the URL(s)
* Categorizing the URL(s)
* Generating a classification model
* Getting the classification model
* Update the classification model
* Deleting the classification model(s)
* List all classification models
* Getting URL category
* Update URL category
* Report miscategorization
* Add URL address
* Delete the URL(s)
* List all URLs
	
## Setup

The following steps are applied for setup:

* Download the latest version:

  ```
  LATEST_VERSION=$(wget -qO - https://api.github.com/repos/mtnmunuklu/Lescatit/releases/latest \
  | grep tag_name \
  | cut -d  '"' -f 4)

  curl -LJO https://github.com/mtnmunuklu/Lescatit/archive/refs/tags/$LATEST_VERSION.tar.gz
  ```

* Extract the downloaded file:

  ```
  FILE_NAME=Lescatit-$(echo $LATEST_VERSION | cut -d 'v' -f 2)
  tar -xvf $FILE_NAME.tar.gz
  ```

* Execute the build script:

  ```
  cd $FILE_NAME/k8s
  bash build.sh
  ```

## Usage

Lescatit consists of 5 different services: **auth**, **crawl**, **catze**, **cat** and **api**. All incoming requests are first forwarded to the API service. Afterwards, the API service decides to which service the incoming request will be forwarded. The address requested is important in the decision-making process.

What features each service has and which addresses can be requested, how to make the relevant requests and which responses are returned for these requests are explained in the [swagger.yaml](https://github.com/mtnmunuklu/Lescatit/blob/main/docs/swagger.yml) file under the docs folder.

You can also access the documents describing the **code structure** of each service under the [docs](https://github.com/mtnmunuklu/Lescatit/tree/main/docs) folder.

## License

This project is licensed under the terms of the **MIT** license.
>You can check out the full license [here](https://github.com/mtnmunuklu/Lescatit/blob/main/LICENSE)

## Buy me a coffee

Whether you use this project, learn from it or like it, please consider supporting me with a coffee so I can spend more time on open source projects like this.

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/mtnmunuklu)
