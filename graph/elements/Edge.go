package elements

type Edge struct {
  Class string `json:"class,omitempty" bson:"class,omitempty" yaml:"class,omitempty" hcl:"class,omitempty"`
  Group string `json:"group" bson:"group" yaml:"group" hcl:"group"`
  Source string `json:"source" bson:"source" yaml:"source" hcl:"source"`
  Destination string `json:"destination" bson:"destination" yaml:"destination" hcl:"destination"`
  Directed bool `json:"directed" bson:"directed" yaml:"directed" hcl:"directed"`
  Properties map[string]interface{} `json:"properties" bson:"properties" yaml:"properties" hcl:"properties"`
}

func (e *Edge) GetClass() string {
  return e.Class
}

func (e *Edge) GetGroup() string {
  return e.Group
}

func (e *Edge) GetSource() string {
  return e.Source
}

func (e *Edge) GetDestination() string {
  return e.Destination
}

func (e *Edge) IsDirected() bool {
  return e.Directed
}

func (e *Edge) GetProperties() map[string]interface{} {
  return e.Properties
}

func (e *Edge) SetProperty(name string, value interface{}) {
  e.Properties[name] = value
}

func NewEdge(class string, group string, source string, destination string, directed bool, properties map[string]interface{}) *Edge {
  return &Edge{
    Class: class,
    Group: group,
    Source: source,
    Destination: destination,
    Properties: properties,
  }
}
