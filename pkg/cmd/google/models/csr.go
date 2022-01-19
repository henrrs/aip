package models

import (
	"fmt"

	"aip/pkg/utils"

	"aip/pkg/services/google/sourcerepo"
)

type CSR struct {
	Name string
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

func (cfg CSRConfig) NewCloudSourceRepository(sourcerepoResources sourcerepo.ServiceResources) {

	req, err := sourcerepo.FindByName(sourcerepoResources)

	if err != nil {

		_, err = sourcerepo.AddRepository(sourcerepoResources)

		if err != nil {
			fmt.Println("Error while creating the repository.")
		} else {
			fmt.Println("The repository was created sucessfully.")
		}

	} else {
		fmt.Println(req, err)
	}
}
