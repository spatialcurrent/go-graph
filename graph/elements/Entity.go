package elements

import (
  "github.com/pkg/errors"
)

type Entity struct {
  Class string `json:"class,omitempty" bson:"class,omitempty" yaml:"class,omitempty" hcl:"class,omitempty"`
  Group string `json:"group" bson:"group" yaml:"group" hcl:"group"`
  Vertex string `json:"vertex" bson:"vertex" yaml:"vertex" hcl:"group"`
  Properties map[string]interface{} `json:"properties" bson:"properties" yaml:"properties" hcl:"properties"`
}

func (e *Entity) GetClass() string {
  return e.Class
}

func (e *Entity) GetGroup() string {
  return e.Group
}

func (e *Entity) GetVertex() string {
  return e.Vertex
}

func (e *Entity) GetProperties() map[string]interface{} {
  return e.Properties
}

func (e *Entity) GetPropertyAsFloat64(name string) (float64, error) {
  if i, ok := e.Properties[name]; ok {
    switch i.(type) {
    case float64:
      return i.(float64), nil
    case int:
      return float64(i.(int)), nil
    case string:
      return 0.0, errors.New("Error: Property "+name+" is of type string.  Expecting float64.")
    }
    return 0.0, errors.New("Error: Property "+name+" is of an unknown type.")
  } else {
    return 0.0, errors.New("Error: Property "+name+" does not exist.")
  }
}
