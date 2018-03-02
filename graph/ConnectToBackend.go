package graph

func ConnectToBackend(pluginPath string, symbol string, options map[string]string) (Backend, error) {
  backend, err := LoadBackendFromPlugin(pluginPath, symbol)
  if err != nil {
    return backend, err
  }
  err = backend.Connect(options)
  return backend, err
}
