package exp

type GreaterThanOrEqual struct {
	*BinaryOperator
}

func (gte GreaterThanOrEqual) Sgol() string {
  return "("+gte.Left.Sgol() + " > "+ gte.Right.Sgol()+")"
}

func (gte GreaterThanOrEqual) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": ">=",
		"left": gte.Left.Map(),
		"right": gte.Right.Map(),
	}
}
