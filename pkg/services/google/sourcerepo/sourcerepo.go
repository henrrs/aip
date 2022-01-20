package sourcerepo

import (
	"context"
	"fmt"

	"aip/pkg/utils"

	"google.golang.org/api/sourcerepo/v1"
)

type ServiceResources struct {
	Service    *sourcerepo.Service
	Project    string
	Repository string
	Role       string
}

func NewServiceResources(service *sourcerepo.Service, resources ...string) ServiceResources {
	return ServiceResources{
		Service:    service,
		Project:    resources[0],
		Repository: resources[1],
		Role:       resources[2],
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

func NewSourceRepoService(projectId, repoName string) ServiceResources {

	ctx := context.Background()
	s, err := sourcerepo.NewService(ctx)

	if err != nil {
		fmt.Println("Error creating a new sourcerepo service.")
		fmt.Println(err)
	}

	project := "projects/" + projectId
	repos := project + "/repos/" + repoName
	policy := "roles/source.writer"

	serviceresources := NewServiceResources(s, project, repos, policy)

	return serviceresources
}

func POCNewSourceRepoService(projectId, repoName string) ServiceResources {

	ctx := context.Background()
	s, err := sourcerepo.NewService(ctx)

	project := "projects/" + projectId
	repos := project + "/repos/" + repoName
	policy := "roles/source.writer"

	if err != nil {
		fmt.Println("Error creating a new sourcerepo service.")
		fmt.Println(err)
	}

	return ServiceResources{
		Service:    s,
		Role:       policy,
		Project:    project,
		Repository: repos,
	}
}

func FindAll(serviceresources ServiceResources, projectId string) (*sourcerepo.ListReposResponse, error) {

	req, err := serviceresources.Service.Projects.Repos.List(serviceresources.Project).Do()

	return req, err
}

func FindByName(serviceresources ServiceResources) (*sourcerepo.Repo, error) {

	req, err := serviceresources.Service.Projects.Repos.Get(serviceresources.Repository).Do()

	return req, err
}

func AddRepository(serviceresources ServiceResources) (*sourcerepo.Repo, error) {

	repo := NewRepo(serviceresources.Repository)

	req, err := serviceresources.Service.Projects.Repos.Create(serviceresources.Project, repo).Do()

	return req, err
}

func AddDevelopers(serviceresources ServiceResources, team []string) (*sourcerepo.Policy, error) {

	binding := NewBinding(serviceresources.Role, team)
	policy := NewPolicy(binding)
	iampolicy := NewIamPolicy(policy)

	req, err := serviceresources.Service.Projects.Repos.SetIamPolicy(serviceresources.Repository, iampolicy).Do()

	return req, err
}

func (resources ServiceResources) POCAddDevelopers(team []string) (*sourcerepo.Policy, error) {

	binding := NewBinding(resources.Role, team)
	policy := NewPolicy(binding)
	iampolicy := NewIamPolicy(policy)

	req, err := resources.Service.Projects.Repos.SetIamPolicy(resources.Repository, iampolicy).Do()

	return req, err
}

func InitRepo(projectId, reponame string) error {

	p, err := utils.GetCurrentDir()

	if err != nil {
		fmt.Println(err)
	}

	repoPath := p + "/" + reponame

	s1 := utils.NewScript(repoPath,
		"gcloud auth activate-service-account --key-file=$GOOGLE_APPLICATION_CREDENTIALS",
		"gcloud source repos clone "+reponame+" --project=\""+projectId+"\" ")

	s2 := utils.NewScript(p,
		"touch README.MD",
		"git add .",
		"git commit -m \"Initial Commit\"",
		"git push --all",
		"rm -rf "+repoPath)

	err = utils.Run(s1, s2)

	if err != nil {
		return err
	}

	return nil
}
