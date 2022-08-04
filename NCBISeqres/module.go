package NCBISeqres

import "ncbiasn/NCBISeqloc"

type SeqGraph struct {
	Title   string             `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	Comment string             `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Loc     *NCBISeqloc.SeqLoc `xml:"loc,omitempty" json:"loc,omitempty"`
	TitleX  string             `xml:"title-x,omitempty" json:"title_x,omitempty" asn1:"optional"`
	TitleY  string             `xml:"title-y,omitempty" json:"title_y,omitempty" asn1:"optional"`
	Comp    int64              `xml:"comp,omitempty" json:"comp,omitempty" asn1:"optional"`
	A       float64            `xml:"a,omitempty" json:"a,omitempty" asn1:"optional"`
	B       float64            `xml:"b,omitempty" json:"b,omitempty" asn1:"optional"`
	Numval  int64              `xml:"numval" json:"numval"`
	Graph   struct {
		Real *RealGraph `xml:"real,omitempty" json:"real,omitempty"`
		Int  *IntGraph  `xml:"int,omitempty" json:"int,omitempty"`
		Byte *ByteGraph `xml:"byte,omitempty" json:"byte,omitempty"`
	} `xml:"graph" json:"graph"` //Graph,ChoiceOption
}
type RealGraph struct {
	Max    float64   `xml:"max" json:"max"`
	Min    float64   `xml:"min" json:"min"`
	Axis   float64   `xml:"axis" json:"axis"`
	Values []float64 `xml:"values,omitempty" json:"values,omitempty"`
}
type IntGraph struct {
	Max    int64   `xml:"max" json:"max"`
	Min    int64   `xml:"min" json:"min"`
	Axis   int64   `xml:"axis" json:"axis"`
	Values []int64 `xml:"values,omitempty" json:"values,omitempty"`
}
type ByteGraph struct {
	Max    int64  `xml:"max" json:"max"`
	Min    int64  `xml:"min" json:"min"`
	Axis   int64  `xml:"axis" json:"axis"`
	Values []byte `xml:"values,omitempty" json:"values,omitempty"`
}
