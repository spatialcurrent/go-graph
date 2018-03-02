package exp

type Not struct {
  *UnaryOperator
}

func (n Not) Sgol() string {
  return "not "+n.Node.Sgol()
}

func (n Not) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "not",
		"node": n.Node.Map(),
	}
}
