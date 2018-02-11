package graph

type FeatureCollection struct {
  Type string `json:"type" bson:"type" yaml:"type" hcl:"type"`
  TotalFeatures int `json:"totalFeatures" bson:"totalFeatures" yaml:"totalFeatures" hcl:"totalFeatures"`
  Features []Feature `json:"features" bson:"features" yaml:"features" hcl:"features"`
}

func NewFeatureCollection(features []Feature) (FeatureCollection) {
  return FeatureCollection{
    Type: "FeatureCollection",
    TotalFeatures: len(features),
    Features: features,
  }
}
