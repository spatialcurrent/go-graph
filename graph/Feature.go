package graph

type Feature struct {
  Id interface{} `json:"id" bson:"id" yaml:"id" hcl:"id"`
  Properties map[string]interface{} `json:"properties" bson:"properties" yaml:"properties" hcl:"properties"`
  GeometryName string `json:"geometry_name" bson:"geometry_name" yaml:"geometry_name" hcl:"geometry_name"`
  Geometry Geometry `json:"geometry" bson:"geometry" yaml:"geometry" hcl:"geometry"`
}

func NewFeature(id string, properties map[string]interface{}, geom Geometry) Feature {
  return Feature{
    Id: id,
    Properties: properties,
    GeometryName: "the_geom",
    Geometry: geom,
  }
}
