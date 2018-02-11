package graph

import (
	"plugin"
)

import (
	"github.com/mitchellh/go-homedir"
)

func LoadBackendFromPlugin(path string, symbol string) (Backend, error) {

	path_expanded, err := homedir.Expand(path)
	if err != nil {
		return nil, err
	}

	p, err := plugin.Open(path_expanded)
	if err != nil {
		return nil, err
	}

	f, err := p.Lookup(symbol)
	if err != nil {
		return nil, err
	}

	graph_backend := f.(func() Backend)().(Backend)

	return graph_backend, nil
}
