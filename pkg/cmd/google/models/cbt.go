package models

import (
	"fmt"

	"aip/pkg/cmd/google/services/cloudbuild"
	"aip/pkg/utils"
)

type trigger struct {
	Name          string
	Description   string
	Substitutions []string
}

type CbtConfig struct {
	Trigger trigger
	Project project
	CSR     csr
}

func NewCbtConfig(fileName string) *CbtConfig {
	c := new(CbtConfig)
	c = utils.ReadFile(fileName, c).(*CbtConfig)

	return c
}

func (cfg CbtConfig) NewCBT(cloudbuildResources cloudbuild.CloudBuildTriggerResources) error {

	req, err := cloudbuildResources.AddTrigger()

	if err != nil {
		fmt.Println(req, err)

		return err

	} else {
		fmt.Println("The trigger was created sucessfully.")
	}

	err = cloudbuildResources.AuthorizeCloudBuildServiceAccount()

	if err != nil {
		fmt.Println(err)

		return err

	} else {
		fmt.Println("Cloud build service account is authorized to trigger deploys.")
	}

	return nil
}

func (cbt CbtConfig) GetCsr() csr {
	return cbt.CSR
}

func (cbt CbtConfig) GetProject() project {
	return cbt.Project
}

func (cbt CbtConfig) GetTrigger() trigger {
	return cbt.Trigger
}

func (t trigger) GetName() string {
	return t.Name
}

func (t trigger) GetDescription() string {
	return t.Description
}
