package models

import (
	"fmt"

	"aip/pkg/utils"
)

func GetProjectNumber(projectId string) string {

	cmd := "gcloud projects describe " + projectId + " --format \"value(projectNumber)\""

	r, err := utils.ExecCmdWithOutput(cmd)

	if err != nil {
		fmt.Println(err)
	}

	return r
}
