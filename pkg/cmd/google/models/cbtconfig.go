package models

import (
	"aip/pkg/cmd/google/services/cloudbuild"
	"aip/pkg/utils"
	"fmt"
)

type CBTConfig struct {
	Trigger trigger
	Project project
	CSR     csr
}

func NewCBTConfig(fileName string) *CBTConfig {
	c := new(CBTConfig)
	c = utils.ReadFile(fileName, c).(*CBTConfig)

	return c
}

func NewCBTConfigWithoutParameters() *CBTConfig {
	c := new(CBTConfig)

	return c
}

func (cfg CBTConfig) NewCBT(cloudbuildResources cloudbuild.CloudBuildTriggerResources) error {

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

func (cbt CBTConfig) GetCsr() csr {
	return cbt.CSR
}

func (cbt CBTConfig) GetProject() project {
	return cbt.Project
}

func (cbt CBTConfig) GetTrigger() trigger {
	return cbt.Trigger
}
