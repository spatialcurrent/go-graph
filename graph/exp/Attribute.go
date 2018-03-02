package exp

type Attribute struct {
	Name string
}

func (a Attribute) Sgol() string {
	return "@"+a.Name
}

func (a Attribute) Map() map[string]interface{} {
	return map[string]interface{}{
	  "attribute": a.Name,
	}
}
