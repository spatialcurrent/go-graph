package exp

type GreaterThan struct {
	*BinaryOperator
}

func (gt GreaterThan) Sgol() string {
  return "("+gt.Left.Sgol() + " > "+ gt.Right.Sgol()+")"
}

func (gt GreaterThan) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": ">",
		"left": gt.Left.Map(),
		"right": gt.Right.Map(),
	}
}
