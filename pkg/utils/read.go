package utils

import (
	"fmt"
	"strings"
	"reflect"
	"errors"
	"io/ioutil"
	"encoding/json"
	"sigs.k8s.io/yaml"
)

const (
	YAML = "yaml"
	JSON = "json"
)

func ReadFile(fileName string, i interface{}) interface{} {

	iType := reflect.TypeOf(i).Elem()
	t := reflect.New(iType).Interface()

	content, err := readFileContent(fileName)
	ext, err := fileExtension(fileName)

	switch opc := ext; opc {

	case JSON:
		err = json.Unmarshal(content, &t)

	case YAML:
		err = yaml.Unmarshal(content, &t)

	default:
		fmt.Println("Extension not supported.")
	}

	if err != nil {
		fmt.Println(err)
	}

	return t
}

func fileExtension(fileName string) (string, error) {

	p := strings.Split(fileName, ".")

	s := len(p)

	if s == 2 {
		return p[1], nil
	} 

	return "", errors.New("File name is out of pattern.")
}

func readFileContent(fileName string) ([]byte, error) {

	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	return content, err
}