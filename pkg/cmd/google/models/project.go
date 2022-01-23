package models

import (
	"fmt"

	"aip/pkg/utils"
)

type project struct {
	Id     string
	Number string
}

type Project struct {
	Project project
}

func NewProject(projectId string) Project {

	p := new(project)
	p.SetId(projectId)

	return Project{
		Project: *p,
	}
}

func (p project) GetId() string {
	return p.Id
}

func (p *project) SetId(id string) {
	p.Id = id
}

func (p project) GetNumber() string {
	return p.Number
}

func (p *project) SetNumber() {

	n := p.describeProjectNumber()

	p.Number = n
}

func (p Project) GetProject() project {
	return p.Project
}

func (p *Project) SetProject(project project) {
	p.Project = project
}

func (p project) describeProjectNumber() string {

	cmd := "gcloud projects describe " + p.Id + " --format \"value(projectNumber)\""

	r, err := utils.ExecCmdWithOutput(cmd)

	if err != nil {
		fmt.Println(err)
	}

	return r
}
