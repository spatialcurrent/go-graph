package exp

type Like struct {
	*BinaryOperator
}

func (l Like) Sgol() string {
  return "("+l.Left.Sgol() + " like "+ l.Right.Sgol()+")"
}

func (l Like) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "like",
		"left": l.Left.Map(),
		"right": l.Right.Map(),
	}
}
