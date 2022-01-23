package cloudbuild

import (
	"context"
	"fmt"

	"aip/pkg/utils"

	"google.golang.org/api/cloudbuild/v1"
)

type CloudBuildTriggerResources struct {
	Service        *cloudbuild.Service
	Name           string
	Description    string
	ServiceAccount string
	FileName       string
	BranchName     string
	RepoName       string
	ProjectId      string
	ProjectNumber  string
}

func NewCloudBuildTriggerResources(triggerName, triggerDescription, branchName, repoName, projectId, projectNumber, pipeline string) CloudBuildTriggerResources {

	ctx := context.Background()
	s, err := cloudbuild.NewService(ctx)

	if err != nil {
		fmt.Println("Error creating a new cloudbuild service.")
		fmt.Println(err)
	}

	serviceAccount := projectNumber + "@cloudbuild.gserviceaccount.com"

	return CloudBuildTriggerResources{
		Service:        s,
		Name:           triggerName,
		Description:    triggerDescription,
		ServiceAccount: serviceAccount,
		FileName:       pipeline,
		BranchName:     branchName,
		RepoName:       repoName,
		ProjectId:      projectId,
		ProjectNumber:  projectNumber,
	}
}

func NewRepoSource(branchName, repoName, projectId string) *cloudbuild.RepoSource {
	return &cloudbuild.RepoSource{
		BranchName: branchName,
		RepoName:   repoName,
		ProjectId:  projectId,
	}
}

func NewBuild() *cloudbuild.Build {
	return &cloudbuild.Build{}
}

func NewBuildTrigger(description, name string, b *cloudbuild.Build, rs *cloudbuild.RepoSource) *cloudbuild.BuildTrigger {
	return &cloudbuild.BuildTrigger{
		Build:           b,
		Description:     description,
		Name:            name,
		TriggerTemplate: rs,
	}
}

func (resources CloudBuildTriggerResources) AddTrigger() (*cloudbuild.BuildTrigger, error) {

	build := NewBuild()
	reposource := NewRepoSource(resources.BranchName, resources.RepoName, resources.ProjectId)

	if resources.FileName != "" {
		build = utils.ReadFile(resources.FileName, build).(*cloudbuild.Build)
	}

	buildtrigger := NewBuildTrigger(resources.Description, resources.Name, build, reposource)

	req, err := resources.Service.Projects.Triggers.Create(resources.ProjectId, buildtrigger).Do()

	return req, err
}

func (resources CloudBuildTriggerResources) AuthorizeCloudBuildServiceAccount() error {

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

		cmd := "gcloud projects add-iam-policy-binding " + resources.ProjectId + " --member \"serviceAccount:" + resources.ServiceAccount + "\" --role " + r

		s.Cmds = append(s.Cmds, cmd)
	}

	err := utils.Run(s)

	return err
}
