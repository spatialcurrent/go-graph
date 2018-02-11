package graph

type Edge struct {
  group string
  source string
  destination string
  directed bool
  properties map[string]interface{}
}

func (e *Edge) GetGroup() string {
  return e.group
}

func (e *Edge) GetSource() string {
  return e.source
}

func (e *Edge) GetDestination() string {
  return e.destination
}

func (e *Edge) IsDirected() bool {
  return e.directed
}

func (e *Edge) GetProperties() map[string]interface{} {
  return e.properties
}
