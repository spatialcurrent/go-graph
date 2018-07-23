package graph

import (
  "encoding/json"
)

import (
  "github.com/spatialcurrent/go-graph/graph/elements"
)

type QueryResults struct {
  Entities []elements.Entity `json:"entities,omitempty" bson:"entities" yaml:"entities,omitempty" hcl:"entities"`
  Edges []elements.Edge `json:"edges,omitempty" bson:"edges" yaml:"edges,omitempty" hcl:"edges"`
}

func (results *QueryResults) Json() (string, error) {
  b, err := json.Marshal(results)
  if err != nil {
    return "", err
  }
  return string(b), nil
}

func (results *QueryResults) FeatureCollection() (FeatureCollection, error) {

  fc := FeatureCollection{}
  features := make([]Feature, 0)
  for _, e := range results.Entities {

    properties := e.GetProperties()
    if attributes, ok := properties["attributes"]; ok {
      p := map[string]interface{}{}
      for k, v := range attributes.(map[string]string) {
        p[k] = v
      }
      properties = p
    }

    geom_lon, err := e.GetPropertyAsFloat64("geom_lon")
    if err != nil {
      return fc, err
    }
    geom_lat, err := e.GetPropertyAsFloat64("geom_lat")
    if err != nil {
      return fc, err
    }
    f := NewFeature(e.GetVertex(), properties, NewPoint(geom_lon, geom_lat))
    features = append(features, f)
  }

  fc = NewFeatureCollection(features)
  return fc, nil
}
