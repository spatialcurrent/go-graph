package graph

import (
  "github.com/pkg/errors"
)

type Entity struct {
  group string
  vertex string
  properties map[string]interface{}
}

func (e *Entity) GetGroup() string {
  return e.group
}

func (e *Entity) GetVertex() string {
  return e.vertex
}

func (e *Entity) GetProperties() map[string]interface{} {
  return e.properties
}

func (e *Entity) GetPropertyAsFloat64(name string) (float64, error) {
  if i, ok := e.properties[name]; ok {
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
