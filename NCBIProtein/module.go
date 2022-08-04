package NCBIProtein

import "ncbiasn/NCBIGeneral"

type ProtRef struct {
	Name      []string            `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Desc      string              `xml:"desc,omitempty" json:"desc,omitempty" asn1:"optional"`
	Ec        []string            `xml:"ec,omitempty" json:"ec,omitempty" asn1:"optional"`
	Activity  []string            `xml:"activity,omitempty" json:"activity,omitempty" asn1:"optional"`
	Db        []NCBIGeneral.Dbtag `xml:"db,omitempty" json:"db,omitempty" asn1:"optional"`
	Processed string              `xml:"processed" json:"processed"` //Processed,EnumList:not-set(0),preprotein(1),mature(2),signal-peptide(3),transit-peptide(4),propeptide(5)
}
