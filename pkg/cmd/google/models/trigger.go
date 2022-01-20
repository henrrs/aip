package models

import "aip/pkg/utils"

type Trigger struct {
	Name          string
	Description   string
	Substitutions []string
}

type TriggerConfig struct {
	Trigger Trigger
	CSR     csr
}

func NewTriggerConfig(fileName string) *TriggerConfig {
	c := new(TriggerConfig)
	c = utils.ReadFile(fileName, c).(*TriggerConfig)

	return c
}
