package exp

import (
  "strings"
)

type Function struct {
  Name string
  Arguments []string
}

func (f Function) Sgol() string {
  out := f.Name + "("
  for i, arg := range f.Arguments {
    if i > 0 {
      out += ", "
    }
    if strings.Contains(arg, " ") {
      out += "\"" + arg + "\""
    } else {
      out += arg
    }
  }
  out += ")"
  return out
}


func (f Function) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "function",
		"name": f.Name,
		"arguments": f.Arguments,
	}
}
