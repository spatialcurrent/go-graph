package operations

import (
  "encoding/json"
)

import (
	"gopkg.in/yaml.v2"
)

type Field struct {
  Name string `json:"name" bson:"name" yaml:"name" hcl:"name"`
  Required bool `json:"required" bson:"required" yaml:"required" hcl:"required"`
}

type OperationDefinition struct {
  Name string `json:"name" bson:"name" yaml:"name" hcl:"name"`
  Fields []Field `json:"fields" bson:"fields" yaml:"fields" hcl:"fields"`
  Next []string `json:"next" bson:"next" yaml:"next" hcl:"next"`
}

func (op OperationDefinition) Json() (string, error) {
	b, err := json.Marshal(op)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (op OperationDefinition) Yaml() (string, error) {
	b, err := yaml.Marshal(op)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
