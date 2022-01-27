package models

import (
	"fmt"

	"aip/pkg/cmd/google/services/sourcerepo"
	"aip/pkg/utils"
)

// CSR means Cloud Source Repositories

type csr struct {
	Name   string
	Branch string
	Team   []string
}

type CSRConfig struct {
	CSR     csr
	Project project
}

func NewCSRConfig(fileName string) *CSRConfig {
	c := new(CSRConfig)
	c = utils.ReadFile(fileName, c).(*CSRConfig)

	return c
}

func NewCSRConfigWithoutParameters() *CSRConfig {
	c := new(CSRConfig)

	return c
}

func (cfg CSRConfig) NewCSR(sourcerepo sourcerepo.SourceRepoResources) error {

	req, err := sourcerepo.FindByName()

	if err != nil {

		_, err = sourcerepo.AddRepository()

		if err != nil {
			fmt.Println("Error while creating the repository.")

		} else {
			fmt.Println("The repository was created sucessfully.")
		}

	} else {
		fmt.Println(req, err)
	}

	return nil
}

func (cfg CSRConfig) GetName() string {
	return cfg.CSR.Name
}

func (csr csr) GetCsrName() string {
	return csr.Name
}

func (cfg CSRConfig) GetBranch() string {
	return cfg.CSR.Branch
}

func (csr csr) GetCsrBranch() string {
	return csr.Branch
}

func (cfg CSRConfig) GetTeam() []string {
	return cfg.CSR.Team
}

func (csr csr) GetCsrTeam() []string {
	return csr.Team
}

func (cfg CSRConfig) GetProject() project {
	return cfg.Project
}

func (cfg CSRConfig) HasTeam() bool {
	return cfg.CSR.Team != nil
}

func (csr csr) CsrHasTeam() bool {
	return csr.Team != nil
}

func (cfg CSRConfig) InitCSR() error {

	p, err := utils.GetCurrentDir()

	if err != nil {
		fmt.Println(err)
	}

	csrName := cfg.GetName()
	csrProject := cfg.GetProject()
	projectId := csrProject.GetId()

	repoPath := p + "/" + csrName

	s1 := utils.NewScript("",
		"gcloud auth activate-service-account --key-file=$GOOGLE_APPLICATION_CREDENTIALS",
		"gcloud source repos clone "+csrName+" --project=\""+projectId+"\" ")

	s2 := utils.NewScript(repoPath,
		"touch README.MD",
		"git add .",
		"git commit -m \"Initial Commit\"",
		"git push --all")

	s3 := utils.NewScript(p,
		"rm -rf "+repoPath)

	err = utils.Run(s1, s2, s3)

	if err != nil {
		return err
	}

	return nil
}

func (csr csr) InitCSRByParameters(csrName, projectId string) error {

	p, err := utils.GetCurrentDir()

	if err != nil {
		fmt.Println(err)
	}

	repoPath := p + "/" + csrName

	s1 := utils.NewScript("",
		"gcloud auth activate-service-account --key-file=$GOOGLE_APPLICATION_CREDENTIALS",
		"gcloud source repos clone "+csrName+" --project=\""+projectId+"\" ")

	s2 := utils.NewScript(repoPath,
		"touch README.MD",
		"git add .",
		"git commit -m \"Initial Commit\"",
		"git push --all")

	s3 := utils.NewScript(p,
		"rm -rf "+repoPath)

	err = utils.Run(s1, s2, s3)

	if err != nil {
		return err
	}

	return nil
}

func (cfg *CSRConfig) UpdateTeam() {

	for i := 0; i < len(cfg.CSR.Team); i++ {
		cfg.CSR.Team[i] = "user:" + cfg.CSR.Team[i]
	}
}

func (csr *csr) UpdateCSRTeam() {

	for i := 0; i < len(csr.Team); i++ {
		csr.Team[i] = "user:" + csr.Team[i]
	}
}

func (cfg CSRConfig) AddTeam(sourcerepo sourcerepo.SourceRepoResources) error {

	team := cfg.GetTeam()

	req, err := sourcerepo.AddDevelopers(team)

	if err != nil {
		fmt.Println(err, req)

		return err
	}

	fmt.Println("The developers were added sucessfully to the repository.")

	return nil
}
