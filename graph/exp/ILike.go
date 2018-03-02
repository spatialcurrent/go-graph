package exp

type ILike struct {
	*BinaryOperator
}

func (i ILike) Sgol() string {
  return "("+i.Left.Sgol() + " ilike "+ i.Right.Sgol()+")"
}

func (i ILike) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "ilike",
		"left": i.Left.Map(),
		"right": i.Right.Map(),
	}
}
