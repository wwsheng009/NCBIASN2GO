package NCBIBioTree

import "ncbiasn/NCBIGeneral"

type BioTreeContainer struct {
	Treetype string                  `xml:"treetype,omitempty" json:"treetype,omitempty" asn1:"optional"`
	Fdict    FeatureDictSet          `xml:"fdict,omitempty" json:"fdict,omitempty"`
	Nodes    NodeSet                 `xml:"nodes,omitempty" json:"nodes,omitempty"`
	Label    string                  `xml:"label,omitempty" json:"label,omitempty" asn1:"optional"`
	User     *NCBIGeneral.UserObject `xml:"user,omitempty" json:"user,omitempty" asn1:"optional"`
}
type NodeSet []Node
type Node struct {
	Id       int64          `xml:"id" json:"id"`
	Parent   int64          `xml:"parent,omitempty" json:"parent,omitempty" asn1:"optional"`
	Features NodeFeatureSet `xml:"features,omitempty" json:"features,omitempty" asn1:"optional"`
}
type NodeFeatureSet []NodeFeature
type NodeFeature struct {
	Featureid int64  `xml:"featureid" json:"featureid"`
	Value     string `xml:"value" json:"value" asn1:"utf8"`
}
type FeatureDictSet []FeatureDescr
type FeatureDescr struct {
	Id   int64  `xml:"id" json:"id"`
	Name string `xml:"name" json:"name"`
}
type DistanceMatrix struct {
	Labels    []string  `xml:"labels,omitempty" json:"labels,omitempty"`
	Distances []float64 `xml:"distances,omitempty" json:"distances,omitempty"`
}
