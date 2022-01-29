package utils

import (
	"os"
	"os/exec"
)

type Script struct {
	Dir  string
	Cmds []string
}

func NewScript(dir string, cmds ...string) Script {
	return Script{
		Dir:  dir,
		Cmds: cmds,
	}
}

func Run(script ...Script) error {

	for _, s := range script {

		if s.Dir != "" {
			os.Chdir(s.Dir)
		}

		for _, c := range s.Cmds {

			err := ExecCmdWithoutOutput(c)

			if err != nil {
				return err
			}

		}
	}

	return nil
}

func ExecCmdWithOutput(cmd string) (string, error) {
	res, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		return "", err
	}

	r := ByteToString(res)

	return r, nil
}

func ExecCmdWithoutOutput(command string) error {
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()

	return err
}

func GetCurrentDir() (string, error) {
	p, err := exec.Command("pwd").Output()

	if err != nil {
		return "", err
	}

	path := ByteToString(p)

	return path, nil
}
