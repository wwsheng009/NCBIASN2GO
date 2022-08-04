package NCBIMedline

import (
	"ncbiasn/NCBIBiblio"
	"ncbiasn/NCBIGeneral"
)

type MedlineEntry struct {
	Uid       int64                `xml:"uid,omitempty" json:"uid,omitempty" asn1:"optional"`
	Em        *NCBIGeneral.Date    `xml:"em,omitempty" json:"em,omitempty"`
	Cit       *NCBIBiblio.CitArt   `xml:"cit,omitempty" json:"cit,omitempty"`
	Abstract  string               `xml:"abstract,omitempty" json:"abstract,omitempty" asn1:"optional"`
	Mesh      []MedlineMesh        `xml:"mesh,omitempty" json:"mesh,omitempty" asn1:"optional"`
	Substance []MedlineRn          `xml:"substance,omitempty" json:"substance,omitempty" asn1:"optional"`
	Xref      []MedlineSi          `xml:"xref,omitempty" json:"xref,omitempty" asn1:"optional"`
	Idnum     []string             `xml:"idnum,omitempty" json:"idnum,omitempty" asn1:"optional"`
	Gene      []string             `xml:"gene,omitempty" json:"gene,omitempty" asn1:"optional"`
	Pmid      *NCBIBiblio.PubMedId `xml:"pmid,omitempty" json:"pmid,omitempty" asn1:"optional"`
	PubType   []string             `xml:"pub-type,omitempty" json:"pub_type,omitempty" asn1:"optional"`
	Mlfield   []MedlineField       `xml:"mlfield,omitempty" json:"mlfield,omitempty" asn1:"optional"`
	Status    int                  `xml:"status" json:"status"` //Status,IntegerEnum:publisher(1),premedline(2),medline(3)
}
type MedlineMesh struct {
	Mp   bool          `xml:"mp" json:"mp"`
	Term string        `xml:"term" json:"term"`
	Qual []MedlineQual `xml:"qual,omitempty" json:"qual,omitempty" asn1:"optional"`
}
type MedlineQual struct {
	Mp   bool   `xml:"mp" json:"mp"`
	Subh string `xml:"subh" json:"subh"`
}
type MedlineRn struct {
	Type string `xml:"type" json:"type"` //Type,EnumList:nameonly(0),cas(1),ec(2)
	Cit  string `xml:"cit,omitempty" json:"cit,omitempty" asn1:"optional"`
	Name string `xml:"name" json:"name"`
}
type MedlineSi struct {
	Type string `xml:"type" json:"type"` //Type,EnumList:ddbj(1),carbbank(2),embl(3),hdb(4),genbank(5),hgml(6),mim(7),msd(8),pdb(9),pir(10),prfseqdb(11),psd(12),swissprot(13),gdb(14)
	Cit  string `xml:"cit,omitempty" json:"cit,omitempty" asn1:"optional"`
}
type MedlineField struct {
	Type int      `xml:"type" json:"type"` //Type,IntegerEnum:other(0),comment(1),erratum(2)
	Str  string   `xml:"str" json:"str"`
	Ids  []DocRef `xml:"ids,omitempty" json:"ids,omitempty" asn1:"optional"`
}
type DocRef struct {
	Type int   `xml:"type" json:"type"` //Type,IntegerEnum:medline(1),pubmed(2),ncbigi(3)
	Uid  int64 `xml:"uid" json:"uid"`
}
