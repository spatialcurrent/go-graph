package graph

type Backend interface {
	Type() string
	Connect(map[string]string) error
	GetSchema() (Schema, error)
	Validate(chain OperationChain) error
	Execute(chain OperationChain, options map[string]string) (QueryResponse, error)
}
