package graph

type Schema interface{
	GetGroups() ([]string, error)
}
