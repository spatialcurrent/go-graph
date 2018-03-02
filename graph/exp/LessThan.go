package exp

type LessThan struct {
	*BinaryOperator
}

func (lt LessThan) Sgol() string {
  return "("+lt.Left.Sgol() + " < "+ lt.Right.Sgol()+")"
}

func (lt LessThan) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "<",
		"left": lt.Left.Map(),
		"right": lt.Right.Map(),
	}
}
