package models

type trigger struct {
	Name          string
	Description   string
	Substitutions []string
}

func (t trigger) GetName() string {
	return t.Name
}

func (t trigger) GetDescription() string {
	return t.Description
}

func (t trigger) GetSubstitutions() []string {
	return t.Substitutions
}
