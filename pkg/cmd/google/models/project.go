package models

import (
	"fmt"

	"aip/pkg/utils"
)

type Project struct {
	Id     string
	Number string
}

func NewProject(projectId string) Project {
	return Project{
		Id: projectId,
	}
}

func (p *Project) SetProjectNumber() {

	n := p.describeProjectNumber()

	p.Number = n
}

func (p Project) describeProjectNumber() string {

	cmd := "gcloud projects describe " + p.Id + " --format \"value(projectNumber)\""

	r, err := utils.ExecCmdWithOutput(cmd)

	if err != nil {
		fmt.Println(err)
	}

	return r
}
