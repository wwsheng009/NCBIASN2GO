package NCBIMedlars

import "ncbiasn/NCBIBiblio"

type MedlarsEntry struct {
	Pmid *NCBIBiblio.PubMedId `xml:"pmid,omitempty" json:"pmid,omitempty"`
	Muid int64                `xml:"muid,omitempty" json:"muid,omitempty" asn1:"optional"`
	Recs []MedlarsRecord      `xml:"recs,omitempty" json:"recs,omitempty"`
}
type MedlarsRecord struct {
	Code int64  `xml:"code" json:"code"`
	Abbr string `xml:"abbr,omitempty" json:"abbr,omitempty" asn1:"optional"`
	Data string `xml:"data" json:"data"`
}
