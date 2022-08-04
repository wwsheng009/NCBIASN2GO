package PDBGeneral

import "ncbiasn/NCBIGeneral"

type PDBBlock struct {
	Deposition *NCBIGeneral.Date `xml:"deposition,omitempty" json:"deposition,omitempty"`
	Class      string            `xml:"class" json:"class"`
	Compound   []string          `xml:"compound,omitempty" json:"compound,omitempty"`
	Source     []string          `xml:"source,omitempty" json:"source,omitempty"`
	ExpMethod  string            `xml:"exp-method,omitempty" json:"exp_method,omitempty" asn1:"optional"`
	Replace    *PDBReplace       `xml:"replace,omitempty" json:"replace,omitempty" asn1:"optional"`
}
type PDBReplace struct {
	Date *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty"`
	Ids  []string          `xml:"ids,omitempty" json:"ids,omitempty"`
}
