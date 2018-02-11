package graph

type OperationChain interface {
  GetName() string
  GetOperations() []Operation
	GetLastOperation() (Operation, error)
	GetLimit() int
	GetOutputType() string
  Hash() (string, error)
  Sgol() (string, error)
}
