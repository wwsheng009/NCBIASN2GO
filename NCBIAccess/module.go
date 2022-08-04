package NCBIAccess

type LinkSet struct {
	Num     int64   `xml:"num" json:"num"`
	Uids    []int64 `xml:"uids,omitempty" json:"uids,omitempty" asn1:"optional"`
	Weights []int64 `xml:"weights,omitempty" json:"weights,omitempty" asn1:"optional"`
}
