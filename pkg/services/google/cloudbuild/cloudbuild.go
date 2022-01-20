package cloudbuild

import (
	"context"
	"fmt"

	m "aip/pkg/cmd/google/models"
	"aip/pkg/utils"

	"google.golang.org/api/cloudbuild/v1"
)

type ServiceResources struct {
	Service        *cloudbuild.Service
	BranchName     string
	RepoName       string
	ProjectId      string
	ProjectNumber  string
	ServiceAccount string
	Description    string
	Name           string
	FileName       string
}

func NewServiceResources(service *cloudbuild.Service, resources []string) ServiceResources {
	return ServiceResources{
		Service:        service,
		BranchName:     resources[0],
		RepoName:       resources[1],
		ProjectId:      resources[2],
		Description:    resources[3],
		Name:           resources[4],
		FileName:       resources[5],
		ProjectNumber:  resources[6],
		ServiceAccount: resources[7],
	}
}

func NewRepoSource(branchName string, repoName string, projectId string) *cloudbuild.RepoSource {
	return &cloudbuild.RepoSource{
		BranchName: branchName,
		RepoName:   repoName,
		ProjectId:  projectId,
	}
}

func NewBuild(fileName string) *cloudbuild.Build {
	b := &cloudbuild.Build{}

	return utils.ReadFile(fileName, b).(*cloudbuild.Build)
}

func NewBuildTrigger(description string, name string, b *cloudbuild.Build, rs *cloudbuild.RepoSource) *cloudbuild.BuildTrigger {
	return &cloudbuild.BuildTrigger{
		Build:           b,
		Description:     description,
		Name:            name,
		TriggerTemplate: rs,
	}
}

func NewCloudBuildService(resources ...string) ServiceResources {

	ctx := context.Background()
	s, err := cloudbuild.NewService(ctx)

	if err != nil {
		fmt.Println("Error creating a new cloudbuild service.")
		fmt.Println(err)
	}

	projectNumber := m.GetProjectNumber(resources[2])
	serviceAccount := projectNumber + "@cloudbuild.gserviceaccount.com"

	resources = append(resources, projectNumber)
	resources = append(resources, serviceAccount)

	serviceresources := NewServiceResources(s, resources)

	return serviceresources
}

func AddTrigger(serviceresources ServiceResources) (*cloudbuild.BuildTrigger, error) {

	build := NewBuild(serviceresources.FileName)
	reposource := NewRepoSource(serviceresources.BranchName, serviceresources.RepoName, serviceresources.ProjectId)
	buildtrigger := NewBuildTrigger(serviceresources.Description, serviceresources.Name, build, reposource)

	req, err := serviceresources.Service.Projects.Triggers.Create(serviceresources.ProjectId, buildtrigger).Do()

	return req, err
}

func AuthorizeCloudBuildServiceAccount(serviceresources ServiceResources) error {

	s := utils.Script{
		Dir: "",
	}

	roles := []string{
		"roles/run.admin",
		"roles/iam.serviceAccountUser",
		"roles/cloudbuild.builds.builder",
		"roles/appengine.appAdmin",
		"roles/cloudfunctions.developer",
		"roles/compute.instanceAdmin.v1",
		"roles/container.developer",
	}

	for _, r := range roles {

		cmd := "gcloud projects add-iam-policy-binding " + serviceresources.ProjectId + " --member \"serviceAccount:" + serviceresources.ServiceAccount + "\" --role " + r

		s.Cmds = append(s.Cmds, cmd)
	}

	err := utils.Run(s)

	return err
}
