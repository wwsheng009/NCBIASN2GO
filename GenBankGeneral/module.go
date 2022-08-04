package GenBankGeneral

import "ncbiasn/NCBIGeneral"

type GBBlock struct {
	ExtraAccessions []string          `xml:"extra-accessions,omitempty" json:"extra_accessions,omitempty" asn1:"optional"`
	Source          string            `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"`
	Keywords        []string          `xml:"keywords,omitempty" json:"keywords,omitempty" asn1:"optional"`
	Origin          string            `xml:"origin,omitempty" json:"origin,omitempty" asn1:"optional"`
	Date            string            `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
	EntryDate       *NCBIGeneral.Date `xml:"entry-date,omitempty" json:"entry_date,omitempty" asn1:"optional"`
	Div             string            `xml:"div,omitempty" json:"div,omitempty" asn1:"optional"`
	Taxonomy        string            `xml:"taxonomy,omitempty" json:"taxonomy,omitempty" asn1:"optional"`
}
