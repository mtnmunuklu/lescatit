<p align="center">
  <img width="300" height="300" src="images/logo.svg">
</p>

<p align="center">
<a href="https://pkg.go.dev/github.com/mtnmunuklu/lescatit"><img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-informational.svg" alt="Go Doc"></a> <a href="https://goreportcard.com/report/github.com/mtnmunuklu/lescatit"><img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A+-success.svg" alt="Go Report"></a> <a href="https://travis-ci.com/"><img src="https://img.shields.io/badge/%E2%9A%99%20build-X-success.svg" alt="Build Status"></a> <a href="https://lescatit.com/"><img src="https://img.shields.io/badge/%F0%9F%93%BD%20demo-online-red.svg" alt="Live Demo"></a>
</p>

# Lescatit <sub><small><small>(Let's categorized it)</small></small></sub>

Lescatit is a project developed in **go**, **mongo**, **docker**, and **kubernetes** technologies, providing URL crawling and categorization functionality.

## Table of Contents

* [Features](#features)
* [Setup](#setup)
* [Usage](#usage)
* [License](#license)

## Features

Lescatit offers the following features:

* Getting user(s) information
* Deleting a user
* Changing user roles
* Updating user passwords
* Updating user email addresses
* Updating usernames
* Getting content of URL(s)
* Crawling URL(s)
* Categorizing URL(s)
* Generating a classification model
* Getting the classification model
* Updating the classification model
* Deleting classification model(s)
* Listing all classification models
* Getting URL categories
* Updating URL categories
* Reporting miscategorization
* Adding URL addresses
* Deleting URL(s)
* Listing all URLs

## Setup

To set up Lescatit, follow these steps:

1. Download the latest version:

    ```
    LATEST_VERSION=$(wget -qO - https://api.github.com/repos/mtnmunuklu/Lescatit/releases/latest \
    | grep tag_name \
    | cut -d  '"' -f 4)

    curl -LJO https://github.com/mtnmunuklu/Lescatit/archive/refs/tags/$LATEST_VERSION.tar.gz
    ```

2. Extract the downloaded file:

    ```
    FILE_NAME=Lescatit-$(echo $LATEST_VERSION | cut -d 'v' -f 2)
    tar -xvf $FILE_NAME.tar.gz
    ```

3. Execute the setup scripts:

    ```
    cd $FILE_NAME/scripts
    # Execute on worker and control plane servers.
    bash setup_tools.sh
    bash setup_k8s.sh
    # Execute only on the first control plane server.
    # It will create setup_k8s_control_plane.sh and setup_k8s_worker.sh files.
    # Control plane and worker scripts are for joining the Kubernetes cluster.
    # You can use these scripts on new nodes when you add new nodes as control plane or worker.
    bash setup_k8s_first_control_plane.sh
    ```

## Usage

Lescatit consists of 5 different services: [authentication](https://github.com/mtnmunuklu/Lescatit/blob/main/authentication), [crawler](https://github.com/mtnmunuklu/Lescatit/blob/main/crawler), [categorizer](https://github.com/mtnmunuklu/Lescatit/blob/main/categorizer), [categorization](https://github.com/mtnmunuklu/Lescatit/blob/main/categorization), and [api](https://github.com/mtnmunuklu/Lescatit/blob/main/api). All incoming requests are first forwarded to the API service. Afterwards, the API service decides to which service the incoming request will be forwarded. The requested URL plays a role in the decision-making process.

To understand the features of each service, the available endpoints, how to make requests, and the expected responses, refer to the [api.pdf](https://github.com/mtnmunuklu/Lescatit/blob/main/docs/api/api.pdf) file under the `docs` folder.

You can also access the documents describing the software structure of each service under the `docs` folder.

## License

This project is licensed under the terms of the **MIT** license.
>You can check out the full license [here](https://github.com/mtnmunuklu/Lescatit/blob/main/LICENSE)

## Buy Me a Coffee

If you find this project useful or interesting, please consider supporting me by buying me a coffee. It would help me to invest more time in open source projects like this one.

<a href="https://www.buymeacoffee.com/mtnmunuklu" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>
