package models

import (
	"fmt"

	"aip/pkg/utils"
)

type project struct {
	Id     string
	Number string
}

func NewProject(projectId string) project {
	return project{
		Id: projectId,
	}
}

func (p project) GetId() string {
	return p.Id
}

func (p project) GetNumber() string {
	return p.Number
}

func (p *project) SetProjectNumber() {

	n := p.describeProjectNumber()

	p.Number = n
}

func (p project) describeProjectNumber() string {

	cmd := "gcloud projects describe " + p.Id + " --format \"value(projectNumber)\""

	r, err := utils.ExecCmdWithOutput(cmd)

	if err != nil {
		fmt.Println(err)
	}

	return r
}
