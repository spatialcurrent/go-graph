package graph

type Element interface{
	GetGroup() string
	GetProperties() map[string]interface{}
}
