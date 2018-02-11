package graph

type Geometry struct {
  Type string `json:"type" bson:"type" yaml:"type" hcl:"type"`
  Coordinates []interface{} `json:"coordinates" bson:"coordinates" yaml:"coordinates" hcl:"coordinates"`
}

func NewPoint(lon float64, lat float64) Geometry {
  return Geometry{
    Type: "Point",
    Coordinates: []interface{}{lon, lat},
  }
}
