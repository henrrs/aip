package utils

import (
	"os"
	"os/exec"
	"strings"
)

type Script struct {
	Cmds 	[]string
	Dir		string
}

func NewScript(dir string, cmds ...string) Script {
	return Script {
		Cmds: 	cmds,
		Dir:	dir,
	}
}

func ExecCommand(cmd string) (string, error) {
	
	res, err := exec.Command( "bash", "-c", cmd ).Output()

	if err != nil {
		return "", err
	}

	r := responseToString(res)

	return r, nil
}

func Run(script ...Script) error {

	for _, s := range script {

		for _, c := range s.Cmds {

			cmd := exec.Command( "bash", "-c", c )
			err := cmd.Run()
	
			if err != nil {
				return err
			}

		}

		if s.Dir != "" {
			os.Chdir(s.Dir)
		}
	}

	return nil
}

func GetCurrentDir() (string, error) {
	p, err := exec.Command("pwd").Output()

	if err != nil {
		return "", err
	}

	path := responseToString(p)

	return path, nil
}

func responseToString(response []byte) string {

	r := string(response)
	r = strings.TrimSuffix(r, "\n")

	return r
}