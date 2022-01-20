package utils

import "fmt"

func GetProjectNumber(projectId string) string {

	cmd := "gcloud projects describe " + projectId + " --format \"value(projectNumber)\""

	r, err := ExecCmdWithOutput(cmd)

	if err != nil {
		fmt.Println(err)
	}

	return r
}
