package graph

type Schema interface{
	GetEntityGroupNames() []string
	GetEdgeGroupNames() []string
	GetTypeNames() []string
	ContainsGroups(groups []string) bool
	ContainsEntities(entities []string) bool
	ContainsEdges(edges []string) bool
	GetEntityPropertyType(entity string, propertyName string) string
	GetEdgePropertyType(edge string, propertyName string) string
	Json() (string, error)
	Yaml() (string, error)
}
