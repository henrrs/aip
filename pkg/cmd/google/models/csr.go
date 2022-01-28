package models

import (
	"fmt"

	"aip/pkg/utils"
)

// CSR means Cloud Source Repositories

type csr struct {
	Name   string
	Branch string
	Team   []string
}

func (csr csr) GetName() string {
	return csr.Name
}

func (csr csr) GetBranch() string {
	return csr.Branch
}

func (csr csr) GetTeam() []string {
	return csr.Team
}

func (csr csr) HasTeam() bool {
	return csr.Team != nil
}

func (csr csr) InitCSR(csrName, projectId string) error {

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

func (csr *csr) UpdateTeam() {

	for i := 0; i < len(csr.Team); i++ {
		csr.Team[i] = "user:" + csr.Team[i]
	}
}
