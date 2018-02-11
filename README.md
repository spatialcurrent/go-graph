# go-graph

# Description

**go-graph** contains framework code for building graph databases and APIs in [Go](http://golang.org).

# Building

```
go get github.com/spatialcurrent/go-graph
```

# Usage

**Server**

```
backend, err := LoadBackendFromPlugin("/path/to/custom/graph/plugin.so", "NewGraphBackend")
```

**Plugin**

```
package main

import (
  "net/http"
)

import (
  "github.com/spatialcurrent/go-graph/backend"
)

func NewGraphBackend() backend.Backend {
  return &MyCustomGraphBackend{}
}

func (a *GraphBackend) Type() string {
  return "My Custom Graph Backend"
}
```

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-graph/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
