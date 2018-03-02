package exp

type LessThanOrEqual struct {
	*BinaryOperator
}

func (lte LessThanOrEqual) Sgol() string {
  return "("+lte.Left.Sgol() + " <= "+ lte.Right.Sgol()+")"
}

func (lte LessThanOrEqual) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "<=",
		"left": lte.Left.Map(),
		"right": lte.Right.Map(),
	}
}
