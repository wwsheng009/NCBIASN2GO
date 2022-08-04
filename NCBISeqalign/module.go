package NCBISeqalign

import (
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBISeqloc"
)

type SeqAlignSet []SeqAlign
type SeqAlign struct {
	Type  string  `xml:"type" json:"type"` //Type,EnumList:not-set(0),global(1),diags(2),partial(3),disc(4),other(255)
	Dim   int64   `xml:"dim,omitempty" json:"dim,omitempty" asn1:"optional"`
	Score []Score `xml:"score,omitempty" json:"score,omitempty" asn1:"optional"`
	Segs  struct {
		Dendiag []DenseDiag `xml:"dendiag,omitempty" json:"dendiag,omitempty"`
		Denseg  *DenseSeg   `xml:"denseg,omitempty" json:"denseg,omitempty"`
		Std     []StdSeg    `xml:"std,omitempty" json:"std,omitempty"`
		Packed  *PackedSeg  `xml:"packed,omitempty" json:"packed,omitempty"`
		Disc    SeqAlignSet `xml:"disc,omitempty" json:"disc,omitempty"`
		Spliced *SplicedSeg `xml:"spliced,omitempty" json:"spliced,omitempty"`
		Sparse  *SparseSeg  `xml:"sparse,omitempty" json:"sparse,omitempty"`
	} `xml:"segs" json:"segs"` //Segs,ChoiceOption
	Bounds []NCBISeqloc.SeqLoc      `xml:"bounds,omitempty" json:"bounds,omitempty" asn1:"optional"`
	Id     []NCBIGeneral.ObjectId   `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Ext    []NCBIGeneral.UserObject `xml:"ext,omitempty" json:"ext,omitempty" asn1:"optional"`
}
type DenseDiag struct {
	Dim     int64                 `xml:"dim" json:"dim" asn1:"default:2"`
	Ids     []NCBISeqloc.SeqId    `xml:"ids,omitempty" json:"ids,omitempty"`
	Starts  []int64               `xml:"starts,omitempty" json:"starts,omitempty"`
	Len     int64                 `xml:"len" json:"len"`
	Strands []NCBISeqloc.NaStrand `xml:"strands,omitempty" json:"strands,omitempty" asn1:"optional"`
	Scores  []Score               `xml:"scores,omitempty" json:"scores,omitempty" asn1:"optional"`
}
type DenseSeg struct {
	Dim     int64                 `xml:"dim" json:"dim" asn1:"default:2"`
	Numseg  int64                 `xml:"numseg" json:"numseg"`
	Ids     []NCBISeqloc.SeqId    `xml:"ids,omitempty" json:"ids,omitempty"`
	Starts  []int64               `xml:"starts,omitempty" json:"starts,omitempty"`
	Lens    []int64               `xml:"lens,omitempty" json:"lens,omitempty"`
	Strands []NCBISeqloc.NaStrand `xml:"strands,omitempty" json:"strands,omitempty" asn1:"optional"`
	Scores  []Score               `xml:"scores,omitempty" json:"scores,omitempty" asn1:"optional"`
}
type PackedSeg struct {
	Dim     int64                 `xml:"dim" json:"dim" asn1:"default:2"`
	Numseg  int64                 `xml:"numseg" json:"numseg"`
	Ids     []NCBISeqloc.SeqId    `xml:"ids,omitempty" json:"ids,omitempty"`
	Starts  []int64               `xml:"starts,omitempty" json:"starts,omitempty"`
	Present []byte                `xml:"present,omitempty" json:"present,omitempty"`
	Lens    []int64               `xml:"lens,omitempty" json:"lens,omitempty"`
	Strands []NCBISeqloc.NaStrand `xml:"strands,omitempty" json:"strands,omitempty" asn1:"optional"`
	Scores  []Score               `xml:"scores,omitempty" json:"scores,omitempty" asn1:"optional"`
}
type StdSeg struct {
	Dim    int64               `xml:"dim" json:"dim" asn1:"default:2"`
	Ids    []NCBISeqloc.SeqId  `xml:"ids,omitempty" json:"ids,omitempty" asn1:"optional"`
	Loc    []NCBISeqloc.SeqLoc `xml:"loc,omitempty" json:"loc,omitempty"`
	Scores []Score             `xml:"scores,omitempty" json:"scores,omitempty" asn1:"optional"`
}
type SplicedSeg struct {
	ProductId     *NCBISeqloc.SeqId    `xml:"product-id,omitempty" json:"product_id,omitempty" asn1:"optional"`
	GenomicId     *NCBISeqloc.SeqId    `xml:"genomic-id,omitempty" json:"genomic_id,omitempty" asn1:"optional"`
	ProductStrand *NCBISeqloc.NaStrand `xml:"product-strand,omitempty" json:"product_strand,omitempty" asn1:"optional"`
	GenomicStrand *NCBISeqloc.NaStrand `xml:"genomic-strand,omitempty" json:"genomic_strand,omitempty" asn1:"optional"`
	ProductType   string               `xml:"product-type" json:"product_type"` //ProductType,EnumList:transcript(0),protein(1)
	Exons         []SplicedExon        `xml:"exons,omitempty" json:"exons,omitempty"`
	PolyA         int64                `xml:"poly-a,omitempty" json:"poly_a,omitempty" asn1:"optional"`
	ProductLength int64                `xml:"product-length,omitempty" json:"product_length,omitempty" asn1:"optional"`
	Modifiers     []SplicedSegModifier `xml:"modifiers,omitempty" json:"modifiers,omitempty" asn1:"optional"`
}

//SplicedSegModifier,ChoiceOption
type SplicedSegModifier struct {
	StartCodonFound bool `xml:"start-codon-found,omitempty" json:"start_codon_found,omitempty"`
	StopCodonFound  bool `xml:"stop-codon-found,omitempty" json:"stop_codon_found,omitempty"`
}

type SplicedExon struct {
	ProductStart       *ProductPos              `xml:"product-start,omitempty" json:"product_start,omitempty"`
	ProductEnd         *ProductPos              `xml:"product-end,omitempty" json:"product_end,omitempty"`
	GenomicStart       int64                    `xml:"genomic-start" json:"genomic_start"`
	GenomicEnd         int64                    `xml:"genomic-end" json:"genomic_end"`
	ProductId          *NCBISeqloc.SeqId        `xml:"product-id,omitempty" json:"product_id,omitempty" asn1:"optional"`
	GenomicId          *NCBISeqloc.SeqId        `xml:"genomic-id,omitempty" json:"genomic_id,omitempty" asn1:"optional"`
	ProductStrand      *NCBISeqloc.NaStrand     `xml:"product-strand,omitempty" json:"product_strand,omitempty" asn1:"optional"`
	GenomicStrand      *NCBISeqloc.NaStrand     `xml:"genomic-strand,omitempty" json:"genomic_strand,omitempty" asn1:"optional"`
	Parts              []SplicedExonChunk       `xml:"parts,omitempty" json:"parts,omitempty" asn1:"optional"`
	Scores             ScoreSet                 `xml:"scores,omitempty" json:"scores,omitempty" asn1:"optional"`
	AcceptorBeforeExon *SpliceSite              `xml:"acceptor-before-exon,omitempty" json:"acceptor_before_exon,omitempty" asn1:"optional"`
	DonorAfterExon     *SpliceSite              `xml:"donor-after-exon,omitempty" json:"donor_after_exon,omitempty" asn1:"optional"`
	Partial            bool                     `xml:"partial,omitempty" json:"partial,omitempty" asn1:"optional"`
	Ext                []NCBIGeneral.UserObject `xml:"ext,omitempty" json:"ext,omitempty" asn1:"optional"`
}

//ProductPos,ChoiceOption
type ProductPos struct {
	Nucpos  int64    `xml:"nucpos,omitempty" json:"nucpos,omitempty"`
	Protpos *ProtPos `xml:"protpos,omitempty" json:"protpos,omitempty"`
}

type ProtPos struct {
	Amin  int64 `xml:"amin" json:"amin"`
	Frame int64 `xml:"frame" json:"frame" asn1:"default:0"`
}
type SplicedExonChunk struct {
	Match      int64 `xml:"match,omitempty" json:"match,omitempty"`
	Mismatch   int64 `xml:"mismatch,omitempty" json:"mismatch,omitempty"`
	Diag       int64 `xml:"diag,omitempty" json:"diag,omitempty"`
	ProductIns int64 `xml:"product-ins,omitempty" json:"product_ins,omitempty"`
	GenomicIns int64 `xml:"genomic-ins,omitempty" json:"genomic_ins,omitempty"`
}

//SplicedExonChunk,ChoiceOption
type SpliceSite struct {
	Bases string `xml:"bases" json:"bases"`
}
type SparseSeg struct {
	MasterId  *NCBISeqloc.SeqId `xml:"master-id,omitempty" json:"master_id,omitempty" asn1:"optional"`
	Rows      []SparseAlign     `xml:"rows,omitempty" json:"rows,omitempty"`
	RowScores []Score           `xml:"row-scores,omitempty" json:"row_scores,omitempty" asn1:"optional"`
	Ext       []SparseSegExt    `xml:"ext,omitempty" json:"ext,omitempty" asn1:"optional"`
}
type SparseAlign struct {
	FirstId       *NCBISeqloc.SeqId     `xml:"first-id,omitempty" json:"first_id,omitempty"`
	SecondId      *NCBISeqloc.SeqId     `xml:"second-id,omitempty" json:"second_id,omitempty"`
	Numseg        int64                 `xml:"numseg" json:"numseg"`
	FirstStarts   []int64               `xml:"first-starts,omitempty" json:"first_starts,omitempty"`
	SecondStarts  []int64               `xml:"second-starts,omitempty" json:"second_starts,omitempty"`
	Lens          []int64               `xml:"lens,omitempty" json:"lens,omitempty"`
	SecondStrands []NCBISeqloc.NaStrand `xml:"second-strands,omitempty" json:"second_strands,omitempty" asn1:"optional"`
	SegScores     []Score               `xml:"seg-scores,omitempty" json:"seg_scores,omitempty" asn1:"optional"`
}
type SparseSegExt struct {
	Index int64 `xml:"index" json:"index"`
}
type Score struct {
	Id    *NCBIGeneral.ObjectId `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Value struct {
		Real float64 `xml:"real,omitempty" json:"real,omitempty"`
		Int  int64   `xml:"int,omitempty" json:"int,omitempty"`
	} `xml:"value" json:"value"` //Value,ChoiceOption
}
type ScoreSet []Score
