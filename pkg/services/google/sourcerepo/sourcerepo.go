package sourcerepo

import (
	"context"
	"fmt"

	"google.golang.org/api/sourcerepo/v1"
)

type SourceRepoResources struct {
	Service    *sourcerepo.Service
	Project    string
	Repository string
	Role       string
}

func NewSourceRepoResources(projectId, repoName string) SourceRepoResources {

	ctx := context.Background()
	s, err := sourcerepo.NewService(ctx)

	project := "projects/" + projectId
	repos := project + "/repos/" + repoName
	policy := "roles/source.writer"

	if err != nil {
		fmt.Println("Error creating a new sourcerepo service.")
		fmt.Println(err)
	}

	return SourceRepoResources{
		Service:    s,
		Role:       policy,
		Project:    project,
		Repository: repos,
	}
}

func NewRepo(reponame string) *sourcerepo.Repo {
	return &sourcerepo.Repo{
		Name: reponame,
	}
}

func NewBinding(role string, team []string) []*sourcerepo.Binding {

	b := &sourcerepo.Binding{
		Members: team,
		Role:    role,
	}

	bindings := []*sourcerepo.Binding{}
	bindings = append(bindings, b)

	return bindings
}

func NewPolicy(bindings []*sourcerepo.Binding) *sourcerepo.Policy {
	return &sourcerepo.Policy{
		Bindings: bindings,
	}
}

func NewIamPolicy(policy *sourcerepo.Policy) *sourcerepo.SetIamPolicyRequest {
	return &sourcerepo.SetIamPolicyRequest{
		Policy: policy,
	}
}

func (resources SourceRepoResources) FindAll(projectId string) (*sourcerepo.ListReposResponse, error) {

	req, err := resources.Service.Projects.Repos.List(resources.Project).Do()

	return req, err
}

func (resources SourceRepoResources) FindByName() (*sourcerepo.Repo, error) {

	req, err := resources.Service.Projects.Repos.Get(resources.Repository).Do()

	return req, err
}

func (resources SourceRepoResources) AddRepository() (*sourcerepo.Repo, error) {

	repo := NewRepo(resources.Repository)

	req, err := resources.Service.Projects.Repos.Create(resources.Project, repo).Do()

	return req, err
}

func (resources SourceRepoResources) AddDevelopers(team []string) (*sourcerepo.Policy, error) {

	binding := NewBinding(resources.Role, team)
	policy := NewPolicy(binding)
	iampolicy := NewIamPolicy(policy)

	req, err := resources.Service.Projects.Repos.SetIamPolicy(resources.Repository, iampolicy).Do()

	return req, err
}
