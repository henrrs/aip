# :gear: AIP

<p align="center">
![](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![](https://img.shields.io/badge/Google_Cloud-4285F4?style=for-the-badge&logo=google-cloud&logoColor=white)
</p>

## :white_check_mark: About
Automating Infrastructure Provisioning (AIP) it's a CLI tool that is been developed in Go for automating infrastructure provisioning in different cloud providers. At the first moment I'll only be developing features to automate Google Cloud Platform (GCP) infrastructure but the project was structure so anyone can contribute developing features to others providers as your need.

| <!-- --> | <!-- --> | 
--------------- |  ---------------
First Launch:   | **2021-05-16**    
Last Revision:  | **2021-05-16**    
Version:        | **1.0**

## :white_check_mark: Requirements

1. [Cloud SDK](https://cloud.google.com/sdk)

```bash
sudo apt-get update
wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-307.0.0-linux-x86_64.tar.gz
tar –xvzf google-cloud-sdk-307.0.0-linux-x86_64.tar.gz
cd google-cloud-sdk
./install.sh

```

2. [Git](https://git-scm.com/about)

## :white_check_mark: Features

- [x] Continuous Integration and Continuous Deployment (CI/CD) Pipeline creation (GCP)


## :white_check_mark: How it works

For GCP features you need to generate a [service account key](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) with the necessary roles for the tool authentication with the provider. 

1. Having the service account key you need to set it's path in the GOOGLE_APPLICATION_CREDENTIALS enviroment variable with the follow command:

```bash
export GOOGLE_APPLICATION_CREDENTIALS=$PWD/your-service-key.json
```

## :white_check_mark: Usage example


## :white_check_mark: Project Structure
