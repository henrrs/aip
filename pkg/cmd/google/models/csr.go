package models

import (
	"fmt"

	"aip/pkg/utils"

	"aip/pkg/services/google/sourcerepo"
)

// CSR means Cloud Source Repositories

type CSR struct {
	Name string
	Team []string
}

type CSRConfig struct {
	CSR       CSR
	ProjectId string
}

func NewCSRConfig(fileName string) *CSRConfig {
	c := new(CSRConfig)
	c = utils.ReadFile(fileName, c).(*CSRConfig)

	return c
}

func (cfg *CSRConfig) NewCSR(sourcerepo sourcerepo.SourceRepoResources) error {

	req, err := sourcerepo.FindByName()

	if err != nil {

		_, err = sourcerepo.AddRepository()

		if err != nil {
			fmt.Println("Error while creating the repository.")

			return err

		} else {
			fmt.Println("The repository was created sucessfully.")
		}

	} else {
		fmt.Println(req, err)
	}

	return nil
}

func (cfg *CSRConfig) InitCSR() error {

	p, err := utils.GetCurrentDir()

	if err != nil {
		fmt.Println(err)
	}

	repoPath := p + "/" + cfg.CSR.Name

	s1 := utils.NewScript("",
		"gcloud auth activate-service-account --key-file=$GOOGLE_APPLICATION_CREDENTIALS",
		"gcloud source repos clone "+cfg.CSR.Name+" --project=\""+cfg.ProjectId+"\" ")

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
