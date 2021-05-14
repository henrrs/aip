# AIP

## About
Automating Infrastructure Provisioning (AIP) it's a CLI tool been developet it in Go for automating infrastructure provisioning in different cloud providers. In the first moment I'll only be developing features to automate Google Cloud Platform (GCP) infrastructure but the project was structure so anyone can contribute developing features to others providers as your need.

| <!-- --> | <!-- --> | 
--------------- |  ---------------
First Launch:   | **2021-05-16**    
Last Revision:  | **2021-05-16**    
Version:        | **1.0**

## Requirements

1. [Cloud SDK (gcloud)](https://cloud.google.com/sdk)

```bash
sudo apt-get update
wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-307.0.0-linux-x86_64.tar.gz
tar –xvzf google-cloud-sdk-307.0.0-linux-x86_64.tar.gz
cd google-cloud-sdk
./install.sh

```

2. [Git](https://git-scm.com/about)

## Features

- [x] Continuous Integration and Continuous Deployment (CI/CD) Pipeline creation (GCP)


## How it works

For GCP features you need to generate a [service account key](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) with the necessary roles for the tool authentication with the provider. 

1. Having the service account key you need to set it's path in the GOOGLE_APPLICATION_CREDENTIALS enviroment variable with the follow command:

```bash
export GOOGLE_APPLICATION_CREDENTIALS=$PWD/your-service-key.json
```
