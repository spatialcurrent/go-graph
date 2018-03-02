package graph

type Schema interface{
	GetEntityGroupNames() []string
	GetEdgeGroupNames() []string
	GetTypeNames() []string
	ContainsGroups(groups []string) bool
	ContainsEntities(entities []string) bool
	ContainsEdges(edges []string) bool
	Json() (string, error)
	Yaml() (string, error)
}
