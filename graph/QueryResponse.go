package graph

import (
  "fmt"
  "encoding/json"
  "net/http"
  "strings"
)

import (
  "gopkg.in/yaml.v2"
)

type QueryResponse struct {
  Success bool `json:"success" bson:"success" yaml:"success" hcl:"success"`
  Message string `json:"message" bson:"message" yaml:"message" hcl:"message"`
  Query string `json:"query" bson:"query" yaml:"query" hcl:"query"`
  Results QueryResults `json:"results" bson:"results" yaml:"results" hcl:"results"`
}

func (qr *QueryResponse) Json() (string, error) {
  b, err := json.Marshal(qr)
  if err != nil {
    return "", err
  }
  return string(b), nil
}

func (qr *QueryResponse) WriteToResponse(w http.ResponseWriter, ext string) {

  contentTypes := map[string]string{
    "css":  "text/css",
    "geojson": "application/json",
    "json": "application/json",
    "jsonl": "plain/text",
    "js":   "application/javascript; charset=utf-8",
    "png":  "image/png",
    "jpg":  "image/jpeg",
    "jpeg": "image/jpeg",
    "jp2":  "image/jpeg",
    "txt": "plain/text",
    "yaml": "plain/text",
  }

  contentType := "text/plain"
  if x, ok := contentTypes[ext]; ok {
    contentType = x
  }

  if ext == "geojson" {

    fc, err := qr.Results.FeatureCollection()
    if err != nil {
      w.Header().Set("Content-Type", "plain/text")
      fmt.Fprintf(w, "Could not convert query results into a feature collection.")
      w.WriteHeader(500)
      return
    }
    output_bytes, err := json.Marshal(fc)
    if err != nil {
      w.Header().Set("Content-Type", "plain/text")
      fmt.Fprintf(w, "Could not marshal query response as geojson")
      w.WriteHeader(500)
      return
    }
    w.Header().Set("Content-Type", contentType)
    fmt.Fprintf(w, string(output_bytes))

  } else if ext == "yaml" || ext == "yml" {

    output_bytes, err := yaml.Marshal(qr)
    if err != nil {
      w.Header().Set("Content-Type", "plain/text")
      fmt.Fprintf(w, "Could not marshal query response as yaml")
      w.WriteHeader(500)
      return
    }
    w.Header().Set("Content-Type", contentType)
    fmt.Fprintf(w, string(output_bytes))

  } else if ext == "json" {

    output_text, err := qr.Json()
    if err != nil {
      w.Header().Set("Content-Type", "plain/text")
      fmt.Fprintf(w, "Could not marshal query response as json")
      w.WriteHeader(500)
      return
    }
    w.Header().Set("Content-Type", contentType)
    fmt.Fprintf(w, output_text)

  } else if ext == "jsonl" {

    lines := []string{}
    for _, entity := range qr.Results.Entities {
      result_bytes, err := json.Marshal(entity)
      if err != nil {
        w.Header().Set("Content-Type", "plain/text")
        fmt.Fprintf(w, "Could not marshal query response as json")
        w.WriteHeader(500)
        return
      }
      lines = append(lines, string(result_bytes))
    }
    for _, edge := range qr.Results.Edges {
      result_bytes, err := json.Marshal(edge)
      if err != nil {
        w.Header().Set("Content-Type", "plain/text")
        fmt.Fprintf(w, "Could not marshal query response as json")
        w.WriteHeader(500)
        return
      }
      lines = append(lines, string(result_bytes))
    }
    w.Header().Set("Content-Type", contentType)
    fmt.Fprintf(w, strings.Join(lines, "\n")+"\n")

  } else {

    w.Header().Set("Content-Type", "plain/text")
    fmt.Fprintf(w, "Unknown file extension")
    w.WriteHeader(500)

  }
}

func NewQueryResponse(success bool, message string, query string) QueryResponse {
  return QueryResponse{
    Success: success,
    Message: message,
    Query: query,
  }
}
