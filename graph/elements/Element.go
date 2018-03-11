package elements

type Element interface{
	GetClass() string
	GetGroup() string
	GetProperties() map[string]interface{}
	SetProperty(name string, value interface{})
}
