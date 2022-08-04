package PIRGeneral

import "ncbiasn/NCBISeqloc"

type PIRBlock struct {
	HadPunct       bool               `xml:"had-punct,omitempty" json:"had_punct,omitempty" asn1:"optional"`
	Host           string             `xml:"host,omitempty" json:"host,omitempty" asn1:"optional"`
	Source         string             `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"`
	Summary        string             `xml:"summary,omitempty" json:"summary,omitempty" asn1:"optional"`
	Genetic        string             `xml:"genetic,omitempty" json:"genetic,omitempty" asn1:"optional"`
	Includes       string             `xml:"includes,omitempty" json:"includes,omitempty" asn1:"optional"`
	Placement      string             `xml:"placement,omitempty" json:"placement,omitempty" asn1:"optional"`
	Superfamily    string             `xml:"superfamily,omitempty" json:"superfamily,omitempty" asn1:"optional"`
	Keywords       []string           `xml:"keywords,omitempty" json:"keywords,omitempty" asn1:"optional"`
	CrossReference string             `xml:"cross-reference,omitempty" json:"cross_reference,omitempty" asn1:"optional"`
	Date           string             `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
	SeqRaw         string             `xml:"seq-raw,omitempty" json:"seq_raw,omitempty" asn1:"optional"`
	Seqref         []NCBISeqloc.SeqId `xml:"seqref,omitempty" json:"seqref,omitempty" asn1:"optional"`
}
