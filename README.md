# :gear: AIP

<p align="center">
  
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white">
<img src="https://img.shields.io/badge/Google_Cloud-4285F4?style=for-the-badge&logo=google-cloud&logoColor=white">

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

After execute the above commands, it will appear the following question. Just type "N" in your terminal, as shown below:

![alt text](https://vitux.com/wp-content/uploads/word-image-1355.png)

After that you'll be asked to continue with the installation. Just type "Y" in your terminal, as shown below:

![alt text](https://vitux.com/wp-content/uploads/word-image-1356.png)

Here you'll be asked a path for bringing the Google Cloud CLIs into your enviroment. You can type your choosed path, but here we can just proceed with the default path pressing "enter", as shown below:

![alt text](https://vitux.com/wp-content/uploads/word-image-1357.png)

Once it's finished it, your terminal window will display the following output on it:

![alt text](https://vitux.com/wp-content/uploads/word-image-1358.png)

![alt text](https://github.com/[username]/[reponame]/blob/[branch]/image.jpg?raw=true)

2. [Git](https://git-scm.com/about)

## :white_check_mark: Features

- [x] Continuous Integration and Continuous Deployment (CI/CD) Pipeline creation (GCP)

## :white_check_mark: Pendency

- [ ] Implement tests


## :white_check_mark: How it works

For GCP features you need to generate a [service account key](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) with the necessary roles for the tool authentication with the provider. 

1. Having the service account key you need to set it's path in the GOOGLE_APPLICATION_CREDENTIALS enviroment variable with the follow command:

```bash
export GOOGLE_APPLICATION_CREDENTIALS=$PWD/your-service-key.json
```

## :white_check_mark: Usage example


## :white_check_mark: Project Structure

    .
    ├── pkg                     
    │   ├── cmd   
    |   |   └── google
    |   |       └── create
    │   ├── services     
    |   |   └── google
    |   |       ├── cloudbuild
    |   |       ├── sourcerepo
    |   |       └── ...
    │   └── utils                
    └── ...
