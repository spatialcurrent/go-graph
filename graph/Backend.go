package graph

import (
	"github.com/spatialcurrent/go-graph/graph/operations"
	"github.com/spatialcurrent/go-graph/graph/elements"
)

type Backend interface {
	Type() string
	Connect(options map[string]string) error
	Ready(options map[string]string) (bool, error)
	Schema(options map[string]string) (Schema, error)
	OperationDefinitions(options map[string]string) ([]operations.OperationDefinition, error)
	FilterFunctions(options map[string]string) ([]FilterFunctionDefinition, error)
	Validate(chain OperationChain, options map[string]string) error
	Execute(chain OperationChain, options map[string]string) (QueryResponse, error)
	NewEntity(group string, vertex string, properties map[string]interface{}) *elements.Entity
	NewEdge(group string, source string, destination string, directed bool, properties map[string]interface{}) *elements.Edge
}
