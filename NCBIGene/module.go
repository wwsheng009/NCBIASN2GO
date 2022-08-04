package NCBIGene

import "ncbiasn/NCBIGeneral"

type GeneRef struct {
	Locus      string              `xml:"locus,omitempty" json:"locus,omitempty" asn1:"optional"`
	Allele     string              `xml:"allele,omitempty" json:"allele,omitempty" asn1:"optional"`
	Desc       string              `xml:"desc,omitempty" json:"desc,omitempty" asn1:"optional"`
	Maploc     string              `xml:"maploc,omitempty" json:"maploc,omitempty" asn1:"optional"`
	Pseudo     bool                `xml:"pseudo" json:"pseudo"`
	Db         []NCBIGeneral.Dbtag `xml:"db,omitempty" json:"db,omitempty" asn1:"optional"`
	Syn        []string            `xml:"syn,omitempty" json:"syn,omitempty" asn1:"optional"`
	LocusTag   string              `xml:"locus-tag,omitempty" json:"locus_tag,omitempty" asn1:"optional"`
	FormalName *GeneNomenclature   `xml:"formal-name,omitempty" json:"formal_name,omitempty" asn1:"optional"`
}
type GeneNomenclature struct {
	Status string             `xml:"status" json:"status"` //Status,EnumList:unknown(0),official(1),interim(2)
	Symbol string             `xml:"symbol,omitempty" json:"symbol,omitempty" asn1:"optional"`
	Name   string             `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Source *NCBIGeneral.Dbtag `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"`
}
