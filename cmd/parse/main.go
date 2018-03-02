package main

import (
  "os"
  "io/ioutil"
  "fmt"
  "strings"
)

import (
  "github.com/spatialcurrent/go-graph/graph/exp"
)

import (
  "gopkg.in/yaml.v2"
)


func main() {
  in, err := ioutil.ReadAll(os.Stdin)
  if err != nil {
    panic(err)
  }
  text := strings.TrimSpace(string(in))
  fmt.Println("Input:", text)
  root, err := exp.Parse(text)
  if err != nil {
    panic(err)
  }
  fmt.Println("Output:", root.Sgol())

  data, err := yaml.Marshal(root.Map())
  if err != nil {
    panic(err)
  }
  fmt.Println("\n"+string(data))
}
