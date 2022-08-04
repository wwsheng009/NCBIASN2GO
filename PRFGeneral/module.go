package PRFGeneral

type PRFBlock struct {
	ExtraSrc *PRFExtraSrc `xml:"extra-src,omitempty" json:"extra_src,omitempty" asn1:"optional"`
	Keywords []string     `xml:"keywords,omitempty" json:"keywords,omitempty" asn1:"optional"`
}
type PRFExtraSrc struct {
	Host   string `xml:"host,omitempty" json:"host,omitempty" asn1:"optional"`
	Part   string `xml:"part,omitempty" json:"part,omitempty" asn1:"optional"`
	State  string `xml:"state,omitempty" json:"state,omitempty" asn1:"optional"`
	Strain string `xml:"strain,omitempty" json:"strain,omitempty" asn1:"optional"`
	Taxon  string `xml:"taxon,omitempty" json:"taxon,omitempty" asn1:"optional"`
}
