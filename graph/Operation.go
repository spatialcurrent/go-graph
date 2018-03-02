package graph

type Operation interface {
  GetTypeName() string
  Validate(schema Schema) error
  Sgol() (string, error)
}
