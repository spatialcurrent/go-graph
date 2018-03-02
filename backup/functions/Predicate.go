package functions

import (
  "strings"
)

import (
  "github.com/pkg/errors"
  //"github.com/imdario/mergo"
)

type Predicate struct {
  Name string
  Fields []string
  Values map[string]string
  Predicate *Predicate
}

func (p Predicate) Map(name2class map[string]string) (map[string]interface{}, error) {
  m := map[string]interface{}{}
  if c, ok := name2class[strings.ToLower(p.Name)]; ok {
    m["class"] = c
    for k, v := range p.Values {
      m[k] = v
    }
    if p.Predicate != nil {
      m["predicate"] = p.predicate.Map(name2class)
    }
  } else {
    return m, errors.New("Could not find matching class for predicate with name "+p.Name+".")
  }
  return m, nil
}

func NewPredicate(name string, fields []string, values map[string]string) Predicate {
  return Predicate{Name: name, Fields: fields, Values: values}
}
