package models

import (
	"aip/pkg/cmd/google/services/sourcerepo"
	"aip/pkg/utils"
	"fmt"
)

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

func (cfg *CSRConfig) SetCsr(csr csr) {
	cfg.CSR = csr
}

func (cfg *CSRConfig) SetProject(project project) {
	cfg.Project = project
}

func (cfg CSRConfig) GetName() string {
	return cfg.CSR.GetName()
}

func (cfg CSRConfig) GetBranch() string {
	return cfg.CSR.GetBranch()
}

func (cfg CSRConfig) GetTeam() []string {
	return cfg.CSR.GetTeam()
}

func (cfg CSRConfig) GetProject() project {
	return cfg.Project
}

func (cfg CSRConfig) HasTeam() bool {
	return cfg.CSR.Team != nil
}

func (cfg CSRConfig) InitCSR() error {
	csrName := cfg.GetName()
	projectId := cfg.GetProject().GetId()

	return cfg.CSR.InitCSR(csrName, projectId)
}

func (cfg *CSRConfig) UpdateTeam() {
	cfg.CSR.UpdateTeam()
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
