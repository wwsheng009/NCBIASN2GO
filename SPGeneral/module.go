package SPGeneral

import (
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBISeqloc"
)

type SPBlock struct {
	Class    string              `xml:"class" json:"class"` //Class,EnumList:not-set(0),standard(1),prelim(2),other(255)
	ExtraAcc []string            `xml:"extra-acc,omitempty" json:"extra_acc,omitempty" asn1:"optional"`
	Imeth    bool                `xml:"imeth" json:"imeth"`
	Plasnm   []string            `xml:"plasnm,omitempty" json:"plasnm,omitempty" asn1:"optional"`
	Seqref   []NCBISeqloc.SeqId  `xml:"seqref,omitempty" json:"seqref,omitempty" asn1:"optional"`
	Dbref    []NCBIGeneral.Dbtag `xml:"dbref,omitempty" json:"dbref,omitempty" asn1:"optional"`
	Keywords []string            `xml:"keywords,omitempty" json:"keywords,omitempty" asn1:"optional"`
	Created  *NCBIGeneral.Date   `xml:"created,omitempty" json:"created,omitempty" asn1:"optional"`
	Sequpd   *NCBIGeneral.Date   `xml:"sequpd,omitempty" json:"sequpd,omitempty" asn1:"optional"`
	Annotupd *NCBIGeneral.Date   `xml:"annotupd,omitempty" json:"annotupd,omitempty" asn1:"optional"`
}
