package graph

import (
  "encoding/json"
)

import (
	"gopkg.in/yaml.v2"
)

type FilterFunctionDefinition struct {
  Class string `json:"class" bson:"class" yaml:"class" hcl:"class"`
}

func (ff FilterFunctionDefinition) Json() (string, error) {
	b, err := json.Marshal(ff)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (ff FilterFunctionDefinition) Yaml() (string, error) {
	b, err := yaml.Marshal(ff)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
