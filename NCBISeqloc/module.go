package NCBISeqloc

import (
	"ncbiasn/NCBIBiblio"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBISeqCommon"
)

//SeqId,ChoiceOption
type SeqId struct {
	Local           *NCBIGeneral.ObjectId     `xml:"local,omitempty" json:"local,omitempty"`
	Gibbsq          int64                     `xml:"gibbsq,omitempty" json:"gibbsq,omitempty"`
	Gibbmt          int64                     `xml:"gibbmt,omitempty" json:"gibbmt,omitempty"`
	Giim            *NCBISeqCommon.GiimportId `xml:"giim,omitempty" json:"giim,omitempty"`
	Genbank         *TextseqId                `xml:"genbank,omitempty" json:"genbank,omitempty"`
	Embl            *TextseqId                `xml:"embl,omitempty" json:"embl,omitempty"`
	Pir             *TextseqId                `xml:"pir,omitempty" json:"pir,omitempty"`
	Swissprot       *TextseqId                `xml:"swissprot,omitempty" json:"swissprot,omitempty"`
	Patent          *PatentSeqId              `xml:"patent,omitempty" json:"patent,omitempty"`
	Other           *TextseqId                `xml:"other,omitempty" json:"other,omitempty"`
	General         *NCBIGeneral.Dbtag        `xml:"general,omitempty" json:"general,omitempty"`
	Gi              int64                     `xml:"gi,omitempty" json:"gi,omitempty"`
	Ddbj            *TextseqId                `xml:"ddbj,omitempty" json:"ddbj,omitempty"`
	Prf             *TextseqId                `xml:"prf,omitempty" json:"prf,omitempty"`
	Pdb             *PDBSeqId                 `xml:"pdb,omitempty" json:"pdb,omitempty"`
	Tpg             *TextseqId                `xml:"tpg,omitempty" json:"tpg,omitempty"`
	Tpe             *TextseqId                `xml:"tpe,omitempty" json:"tpe,omitempty"`
	Tpd             *TextseqId                `xml:"tpd,omitempty" json:"tpd,omitempty"`
	Gpipe           *TextseqId                `xml:"gpipe,omitempty" json:"gpipe,omitempty"`
	NamedAnnotTrack *TextseqId                `xml:"named-annot-track,omitempty" json:"named_annot_track,omitempty"`
}

type SeqIdSet []SeqId
type PatentSeqId struct {
	Seqid int64             `xml:"seqid" json:"seqid"`
	Cit   *NCBIBiblio.IdPat `xml:"cit,omitempty" json:"cit,omitempty"`
}
type TextseqId struct {
	Name      string `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Accession string `xml:"accession,omitempty" json:"accession,omitempty" asn1:"optional"`
	Release   string `xml:"release,omitempty" json:"release,omitempty" asn1:"optional"`
	Version   int64  `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
}

// type GiimportId struct {
// 	Id	int64	`xml:"id" json:"id"`
// 	Db	string	`xml:"db,omitempty" json:"db,omitempty" asn1:"optional"`
// 	Release	string	`xml:"release,omitempty" json:"release,omitempty" asn1:"optional"`
// }
type PDBSeqId struct {
	Mol   PDBMolId          `xml:"mol,omitempty" json:"mol,omitempty"`
	Chain int64             `xml:"chain" json:"chain" asn1:"default:32"`
	Rel   *NCBIGeneral.Date `xml:"rel,omitempty" json:"rel,omitempty" asn1:"optional"`
}
type PDBMolId string
type SeqLoc struct {
	Null      interface{}           `xml:"-" json:"-"` //Null,NullType
	Empty     *SeqId                `xml:"empty,omitempty" json:"empty,omitempty"`
	Whole     *SeqId                `xml:"whole,omitempty" json:"whole,omitempty"`
	Int       *SeqInterval          `xml:"int,omitempty" json:"int,omitempty"`
	PackedInt PackedSeqint          `xml:"packed-int,omitempty" json:"packed_int,omitempty"`
	Pnt       *SeqPoint             `xml:"pnt,omitempty" json:"pnt,omitempty"`
	PackedPnt *PackedSeqpnt         `xml:"packed-pnt,omitempty" json:"packed_pnt,omitempty"`
	Mix       SeqLocMix             `xml:"mix,omitempty" json:"mix,omitempty"`
	Equiv     SeqLocEquiv           `xml:"equiv,omitempty" json:"equiv,omitempty"`
	Bond      *SeqBond              `xml:"bond,omitempty" json:"bond,omitempty"`
	Feat      *NCBISeqCommon.FeatId `xml:"feat,omitempty" json:"feat,omitempty"`
}

//SeqLoc,ChoiceOption
type SeqInterval struct {
	From     int64                `xml:"from" json:"from"`
	To       int64                `xml:"to" json:"to"`
	Strand   *NaStrand            `xml:"strand,omitempty" json:"strand,omitempty" asn1:"optional"`
	Id       *SeqId               `xml:"id,omitempty" json:"id,omitempty"`
	FuzzFrom *NCBIGeneral.IntFuzz `xml:"fuzz-from,omitempty" json:"fuzz_from,omitempty" asn1:"optional"`
	FuzzTo   *NCBIGeneral.IntFuzz `xml:"fuzz-to,omitempty" json:"fuzz_to,omitempty" asn1:"optional"`
}
type PackedSeqint []SeqInterval
type SeqPoint struct {
	Point  int64                `xml:"point" json:"point"`
	Strand *NaStrand            `xml:"strand,omitempty" json:"strand,omitempty" asn1:"optional"`
	Id     *SeqId               `xml:"id,omitempty" json:"id,omitempty"`
	Fuzz   *NCBIGeneral.IntFuzz `xml:"fuzz,omitempty" json:"fuzz,omitempty" asn1:"optional"`
}
type PackedSeqpnt struct {
	Strand *NaStrand            `xml:"strand,omitempty" json:"strand,omitempty" asn1:"optional"`
	Id     *SeqId               `xml:"id,omitempty" json:"id,omitempty"`
	Fuzz   *NCBIGeneral.IntFuzz `xml:"fuzz,omitempty" json:"fuzz,omitempty" asn1:"optional"`
	Points []int64              `xml:"points,omitempty" json:"points,omitempty"`
}
type NaStrand string

//NaStrand,EnumList:unknown(0),plus(1),minus(2),both(3),both-rev(4),other(255)
type SeqBond struct {
	A *SeqPoint `xml:"a,omitempty" json:"a,omitempty"`
	B *SeqPoint `xml:"b,omitempty" json:"b,omitempty" asn1:"optional"`
}
type SeqLocMix []SeqLoc
type SeqLocEquiv []SeqLoc
