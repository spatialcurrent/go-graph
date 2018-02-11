package graph

type Operation interface {
  GetTypeName() string
  Sgol() (string, error)
}
