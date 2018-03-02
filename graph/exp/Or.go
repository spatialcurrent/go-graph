package exp

type Or struct {
	*BinaryOperator
}

func (o Or) Sgol() string {
  return "("+o.Left.Sgol() + " or "+ o.Right.Sgol()+")"
}

func (o Or) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "or",
		"left": o.Left.Map(),
		"right": o.Right.Map(),
	}
}
