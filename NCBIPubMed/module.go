package NCBIPubMed

import (
	"ncbiasn/NCBIBiblio"
	"ncbiasn/NCBIMedline"
)

type PubmedEntry struct {
	Pmid      *NCBIBiblio.PubMedId      `xml:"pmid,omitempty" json:"pmid,omitempty"`
	Medent    *NCBIMedline.MedlineEntry `xml:"medent,omitempty" json:"medent,omitempty" asn1:"optional"`
	Publisher string                    `xml:"publisher,omitempty" json:"publisher,omitempty" asn1:"optional"`
	Urls      []PubmedUrl               `xml:"urls,omitempty" json:"urls,omitempty" asn1:"optional"`
	Pubid     string                    `xml:"pubid,omitempty" json:"pubid,omitempty" asn1:"optional"`
}
type PubmedUrl struct {
	Location string `xml:"location,omitempty" json:"location,omitempty" asn1:"optional"`
	Url      string `xml:"url" json:"url"`
}
