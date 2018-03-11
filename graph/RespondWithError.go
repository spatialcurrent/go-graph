package graph

import (
  "encoding/json"
  "fmt"
  "net/http"
)

import (
  "gopkg.in/yaml.v2"
)

func RespondWithError(w http.ResponseWriter, ext string, message string) {

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

  data := map[string]interface{}{"success": false, "message": message}

  if ext == "yaml" || ext == "yml" {

    output_bytes, err := yaml.Marshal(data)
    if err != nil {
      w.Header().Set("Content-Type", "plain/text")
      fmt.Fprintf(w, "Could not marshal query response as yaml")
      w.WriteHeader(500)
      return
    }
    w.Header().Set("Content-Type", contentType)
    fmt.Fprintf(w, string(output_bytes))

  } else if ext == "json" || ext == "geojson" || ext == "jsonl" || ext == "geojsonl" {

    output_bytes, err := json.Marshal(data)
    if err != nil {
      w.Header().Set("Content-Type", "plain/text")
      fmt.Fprintf(w, "Could not marshal query response as "+ext+".")
      w.WriteHeader(500)
      return
    }
    w.Header().Set("Content-Type", contentType)
    fmt.Fprintf(w, string(output_bytes))

  } else {

    w.Header().Set("Content-Type", "plain/text")
    fmt.Fprintf(w, "Unknown file extension")

  }

  w.WriteHeader(500)
}
