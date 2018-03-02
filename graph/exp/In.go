package exp

type In struct {
	*BinaryOperator
}

func (i In) Sgol() string {
  return "("+i.Left.Sgol() + " in "+ i.Right.Sgol()+")"
}

func (i In) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "in",
		"left": i.Left.Map(),
		"right": i.Right.Map(),
	}
}
