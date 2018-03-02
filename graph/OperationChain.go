package graph

type OperationChain interface {
  GetName() string
  Validate(schema Schema) error
  GetOperations() []Operation
  GetFirstOperation() (Operation, error)
	GetLastOperation() (Operation, error)
	GetLimit() int
	GetOutputType() string
  Hash() (string, error)
  Sgol() (string, error)
}
