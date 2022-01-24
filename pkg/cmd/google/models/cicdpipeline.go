package models

import (
	"aip/pkg/utils"
)

type CiCdPipeline struct {
	Project project
	CSR     csr
	Trigger trigger
}

func NewCiCdPipeline(fileName string) *CiCdPipeline {
	c := new(CiCdPipeline)
	c = utils.ReadFile(fileName, c).(*CiCdPipeline)

	return c
}

func (pipeline CiCdPipeline) GetProject() project {
	return pipeline.Project
}

func (pipeline CiCdPipeline) GetCsr() csr {
	return pipeline.CSR
}

func (pipeline CiCdPipeline) GetTrigger() trigger {
	return pipeline.Trigger
}
