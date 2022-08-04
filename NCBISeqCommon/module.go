package NCBISeqCommon

import "ncbiasn/NCBIGeneral"

//FeatId,ChoiceOption
type FeatId struct {
	Gibb    int64                 `xml:"gibb" json:"gibb"`
	Giim    *GiimportId           `xml:"giim,omitempty" json:"giim,omitempty"`
	Local   *NCBIGeneral.ObjectId `xml:"local,omitempty" json:"local,omitempty"`
	General *NCBIGeneral.Dbtag    `xml:"general,omitempty" json:"general,omitempty"`
}

type GiimportId struct {
	Id      int64  `xml:"id" json:"id"`
	Db      string `xml:"db,omitempty" json:"db,omitempty" asn1:"optional"`
	Release string `xml:"release,omitempty" json:"release,omitempty" asn1:"optional"`
}

type Heterogen string

type SeqLiteral struct {
	Length  int64                `xml:"length" json:"length"`
	Fuzz    *NCBIGeneral.IntFuzz `xml:"fuzz,omitempty" json:"fuzz,omitempty" asn1:"optional"`
	SeqData *SeqData             `xml:"seq-data,omitempty" json:"seq_data,omitempty" asn1:"optional"`
}

//SeqData,ChoiceOption
type SeqData struct {
	Iupacna      IUPACna    `xml:"iupacna,omitempty" json:"iupacna,omitempty"`
	Iupacaa      IUPACaa    `xml:"iupacaa,omitempty" json:"iupacaa,omitempty"`
	Ncbi2na      *NCBI2na   `xml:"-" json:"-"`
	Ncbi2naRaw   string     `xml:"ncbi2na,omitempty" json:"ncbi2na,omitempty"`
	Ncbi4na      *NCBI4na   `xml:"ncbi4na,omitempty" json:"ncbi4na,omitempty"`
	Ncbi8na      *NCBI8na   `xml:"ncbi8na,omitempty" json:"ncbi8na,omitempty"`
	Ncbipna      *NCBIpna   `xml:"ncbipna,omitempty" json:"ncbipna,omitempty"`
	Ncbi8aa      *NCBI8aa   `xml:"ncbi8aa,omitempty" json:"ncbi8aa,omitempty"`
	Ncbieaa      NCBIeaa    `xml:"ncbieaa,omitempty" json:"ncbieaa,omitempty"`
	Ncbipaa      *NCBIpaa   `xml:"ncbipaa,omitempty" json:"ncbipaa,omitempty"`
	Ncbistdaa    *NCBIstdaa `xml:"-" json:"-"`
	NcbistdaaRaw string     `xml:"ncbistdaa,omitempty" json:"ncbistdaa,omitempty"` //ncbi使用的是hex编码，而不是base64编码
	Gap          *SeqGap    `xml:"gap,omitempty" json:"gap,omitempty"`
}

type IUPACna string
type IUPACaa string
type NCBI2na []byte
type NCBI4na []byte
type NCBI8na []byte
type NCBIpna []byte
type NCBI8aa []byte
type NCBIeaa string
type NCBIpaa []byte
type NCBIstdaa []byte
type SeqGap struct {
	Type            int               `xml:"type" json:"type"`                                           //Type,IntegerEnum:unknown(0),fragment(1),clone(2),short-arm(3),heterochromatin(4),centromere(5),telomere(6),repeat(7),contig(8),scaffold(9),contamination(10),other(255)
	Linkage         int               `xml:"linkage,omitempty" json:"linkage,omitempty" asn1:"optional"` //Linkage,IntegerEnum:unlinked(0),linked(1),other(255)
	LinkageEvidence []LinkageEvidence `xml:"linkage-evidence,omitempty" json:"linkage_evidence,omitempty" asn1:"optional"`
}
type LinkageEvidence struct {
	Type int64 `xml:"type" json:"type"`
}
