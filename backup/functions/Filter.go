package functions

import (
  "encoding/json"
  "strings"
)

import (
	"gopkg.in/yaml.v2"
)

type Filter struct {
  Selection []string `json:"selection" bson:"selection" yaml:"selection" hcl:"selection"`
  Predicate Predicate `json:"predicate" bson:"predicate" yaml:"predicate" hcl:"predicate"`
}

func (f Filter) SetSelection(selection []string) {
  f.Selection = selection
}

func (f Filter) SetPredicate(predicate Predicate) {
  f.Predicate = predicate
}

func (f Filter) Map(name2class map[string]string) (map[string]interface{}, error) {
  m := map[string]interface{}{}
  m["selection"] = f.Selection
  p, err := f.Predicate.Map(name2class)
  if err != nil {
    return m, err
  }
  m["predicate"] = p
  return m, nil
}

func (f Filter) Sgol() string {
  args := make([]string, 0)
  args = append(args, f.Selection...)
  for _, name := range f.Predicate.Fields {
    args = append(args, f.Predicate.Values[name])
  }
  return f.Predicate.Name+"(" + strings.Join(args, ",") + ")"
}

func (f Filter) Json(name2class map[string]string) (string, error) {
  m, err := f.Map(name2class)
  if err != nil {
    return "", err
  }
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (f Filter) Yaml(name2class map[string]string) (string, error) {
  m, err := f.Map(name2class)
  if err != nil {
    return "", err
  }
	b, err := yaml.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func NewFilter(selection []string, predicate_class string, fields []string, values map[string]string) Filter {
  return Filter{
    Selection: selection,
    Predicate: NewPredicate(predicate_class, fields, values),
  }
}

func NewInverseFilter(selection []string, predicate_class string, fields []string, values map[string]string) Filter {
  return Filter{
    Selection: selection,
    Predicate: NewPredicate(predicate_class, fields, values),
  }
}
