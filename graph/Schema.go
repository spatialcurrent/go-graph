package graph

type Schema interface{
	GetEntityGroupNames() []string
	GetEdgeGroupNames() []string
	GetTypeNames() []string
	ContainsGroups(groups []string) bool
	ContainsEntities(entities []string) bool
	ContainsEdges(edges []string) bool
	GetEntityPropertyNames(entity string) []string
	//GetEntityPropertyType(entity string, propertyName string) string
	GetEntityPropertyTypes(group string) map[string]string
	GetEdgePropertyNames(edge string) []string
	GetEdgePropertyTypes(group string) map[string]string
	//GetEdgePropertyType(edge string, propertyName string) string
	Json() (string, error)
	Yaml() (string, error)
}
