package NCBISeqRef

import (
	"ncbiasn/NCBIPub"
	"ncbiasn/NCBISeqalign"
)

type Pubdesc struct {
	Pub        NCBIPub.PubEquiv `xml:"pub,omitempty" json:"pub,omitempty"`
	Name       string           `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Fig        string           `xml:"fig,omitempty" json:"fig,omitempty" asn1:"optional"`
	Num        *Numbering       `xml:"num,omitempty" json:"num,omitempty" asn1:"optional"`
	Numexc     bool             `xml:"numexc,omitempty" json:"numexc,omitempty" asn1:"optional"`
	PolyA      bool             `xml:"poly-a,omitempty" json:"poly_a,omitempty" asn1:"optional"`
	Maploc     string           `xml:"maploc,omitempty" json:"maploc,omitempty" asn1:"optional"`
	SeqRaw     *string          `xml:"seq-raw,omitempty" json:"seq_raw,omitempty" asn1:"optional"`
	AlignGroup int64            `xml:"align-group,omitempty" json:"align_group,omitempty" asn1:"optional"`
	Comment    string           `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Reftype    int              `xml:"reftype" json:"reftype"` //Reftype,IntegerEnum:seq(0),sites(1),feats(2),no-target(3)
}

//Numbering,ChoiceOption
type Numbering struct {
	Cont *NumCont `xml:"cont,omitempty" json:"cont,omitempty"`
	Enum *NumEnum `xml:"enum,omitempty" json:"enum,omitempty"`
	Ref  *NumRef  `xml:"ref,omitempty" json:"ref,omitempty"`
	Real *NumReal `xml:"real,omitempty" json:"real,omitempty"`
}

type NumCont struct {
	Refnum    int64 `xml:"refnum" json:"refnum" asn1:"default:1"`
	HasZero   bool  `xml:"has-zero" json:"has_zero"`
	Ascending bool  `xml:"ascending" json:"ascending"`
}
type NumEnum struct {
	Num   int64    `xml:"num" json:"num"`
	Names []string `xml:"names,omitempty" json:"names,omitempty"`
}
type NumRef struct {
	Type   string                 `xml:"type" json:"type"` //Type,EnumList:not-set(0),sources(1),aligns(2)
	Aligns *NCBISeqalign.SeqAlign `xml:"aligns,omitempty" json:"aligns,omitempty" asn1:"optional"`
}
type NumReal struct {
	A     float64 `xml:"a" json:"a"`
	B     float64 `xml:"b" json:"b"`
	Units string  `xml:"units,omitempty" json:"units,omitempty" asn1:"optional"`
}
