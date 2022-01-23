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

1. [Go](https://golang.org/doc/install)

2. [Cloud SDK](https://cloud.google.com/sdk)

3. [Git](https://git-scm.com/about)

## :white_check_mark: Features

- [x] Creation of Cloud Source Repositories (CSR) on Google Cloud Platform (GCP)
- [x] Creation of Cloud Build Triggers (CBT) on Google Cloud Platform (GCP)
- [x] Creation of Continuous Integration and Continuous Deployment (CI/CD) Pipeline on Google Cloud Platform (GCP)

## :white_check_mark: Pendency

- [ ] Implement tests


## :white_check_mark: How it works

For GCP features you need to generate a [service account key](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) with the necessary roles for the tool authentication with the provider. 

1. Having the service account key you need to set it's path in the setup file, on line 3, and then execute the follow:

```bash
. setup.sh
```

&nbsp;&nbsp;&nbsp;&nbsp;Please be careful about Identity Access Management (IAM) least privilege recomendations for your service account key. Be sure that you're only giving the needed permission for it. If you have any doubt about least privilege principle, take a look [here](https://cloud.google.com/iam/docs/recommender-overview).

2. For development purposes you should also follow the instructions shown [here](https://golang.org/doc/gopath_code) in order to be able to install and test the CLI accordingly. 

## :white_check_mark: Usage example

1. Once you clone the repository you can run the CLI in two ways, the first one:

```bash
go run main.go google create ci-cd-pipeline -c="your-file.yaml" -p="your-file.yaml"
```

Or if you configure correctly go environment variables, you can install the CLI as follow:

```bash
go install aip
```

And run as:

```bash
aip google create ci-cd-pipeline -c="your-file.yaml" -p="your-file.yaml"
```

The provided files can be in yaml or json format.

If you have any doubt about how a command must be executed you can just use "-h" or "--help" flag, like that:


```bash
aip --help
```

```bash
aip google -h
```

```bash
aip google create ci-cd-pipeline --help
```
On the confs folder you find an template for write your files.

## :white_check_mark: Project Structure

The follow structure was adopted it. We have the cmd folder where we have our Cobra commands implemented it for CLI function. In the services/google we have the implementation of GCP services. The utils folder is used for functions that can be used in multiple contexts like for example, file reading. The structure for others providers is still not decided it.

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
