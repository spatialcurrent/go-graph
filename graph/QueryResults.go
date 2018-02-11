package graph

import (
  "encoding/json"
)

type QueryResults struct {
  Entities []Entity `json:"entities" bson:"entities" yaml:"entities" hcl:"entities"`
  Edges []Edge `json:"edges" bson:"edges" yaml:"edges" hcl"edges`
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

    geom_lon, err := e.GetPropertyAsFloat64("geom_lon")
    if err != nil {
      return fc, err
    }
    geom_lat, err := e.GetPropertyAsFloat64("geom_lat")
    if err != nil {
      return fc, err
    }
    f := NewFeature(e.GetVertex(), e.GetProperties(), NewPoint(geom_lon, geom_lat))
    features = append(features, f)
  }

  fc = NewFeatureCollection(features)
  return fc, nil
}
