package NCBI_RNA

import "ncbiasn/NCBISeqloc"

type RNARef struct {
	Type   string `xml:"type" json:"type"` //Type,EnumList:unknown(0),premsg(1),mRNA(2),tRNA(3),rRNA(4),snRNA(5),scRNA(6),snoRNA(7),ncRNA(8),tmRNA(9),miscRNA(10),other(255)
	Pseudo bool   `xml:"pseudo,omitempty" json:"pseudo,omitempty" asn1:"optional"`
	Ext    struct {
		Name string   `xml:"name,omitempty" json:"name,omitempty"`
		TRNA *TrnaExt `xml:"tRNA,omitempty" json:"tRNA,omitempty"`
		Gen  *RNAGen  `xml:"gen,omitempty" json:"gen,omitempty"`
	} `xml:"ext,omitempty" json:"ext,omitempty" asn1:"optional"` //Ext,ChoiceOption
}
type TrnaExt struct {
	Aa struct {
		Iupacaa   int64 `xml:"iupacaa,omitempty" json:"iupacaa,omitempty"`
		Ncbieaa   int64 `xml:"ncbieaa,omitempty" json:"ncbieaa,omitempty"`
		Ncbi8aa   int64 `xml:"ncbi8aa,omitempty" json:"ncbi8aa,omitempty"`
		Ncbistdaa int64 `xml:"ncbistdaa,omitempty" json:"ncbistdaa,omitempty"`
	} `xml:"aa,omitempty" json:"aa,omitempty" asn1:"optional"` //Aa,ChoiceOption
	Codon     []int64            `xml:"codon,omitempty" json:"codon,omitempty" asn1:"optional"`
	Anticodon *NCBISeqloc.SeqLoc `xml:"anticodon,omitempty" json:"anticodon,omitempty" asn1:"optional"`
}
type RNAGen struct {
	Class   string     `xml:"class,omitempty" json:"class,omitempty" asn1:"optional"`
	Product string     `xml:"product,omitempty" json:"product,omitempty" asn1:"optional"`
	Quals   RNAQualSet `xml:"quals,omitempty" json:"quals,omitempty" asn1:"optional"`
}
type RNAQual struct {
	Qual string `xml:"qual" json:"qual"`
	Val  string `xml:"val" json:"val"`
}
type RNAQualSet []RNAQual
