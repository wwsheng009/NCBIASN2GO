package NCBISeqSplit

import (
	"ncbiasn/NCBISeqCommon"
	"ncbiasn/NCBISeqalign"
	"ncbiasn/NCBISeqloc"
	"ncbiasn/NCBISequence"
)

type ID2SSplitInfo struct {
	BioseqsInfo []ID2SBioseqsInfo      `xml:"bioseqs-info,omitempty" json:"bioseqs_info,omitempty" asn1:"optional"`
	Chunks      []ID2SChunkInfo        `xml:"chunks,omitempty" json:"chunks,omitempty"`
	Skeleton    *NCBISequence.SeqEntry `xml:"skeleton,omitempty" json:"skeleton,omitempty" asn1:"optional"`
}
type ID2SBioseqsInfo struct {
	Info    *ID2SBioseqInfo `xml:"info,omitempty" json:"info,omitempty"`
	Bioseqs ID2SBioseqIds   `xml:"bioseqs,omitempty" json:"bioseqs,omitempty"`
}
type ID2SBioseqInfo struct {
	GapCount     int64 `xml:"gap-count,omitempty" json:"gap_count,omitempty" asn1:"optional"`
	SeqMapHasRef bool  `xml:"seq-map-has-ref,omitempty" json:"seq_map_has_ref,omitempty" asn1:"optional"`
}
type ID2SChunkInfo struct {
	Id      *ID2SChunkId       `xml:"id,omitempty" json:"id,omitempty"`
	Content []ID2SChunkContent `xml:"content,omitempty" json:"content,omitempty"`
}

//ID2SChunkContent,ChoiceOption
type ID2SChunkContent struct {
	SeqDescr      *ID2SSeqDescrInfo      `xml:"seq-descr,omitempty" json:"seq_descr,omitempty"`
	SeqAnnot      *ID2SSeqAnnotInfo      `xml:"seq-annot,omitempty" json:"seq_annot,omitempty"`
	SeqAssembly   *ID2SSeqAssemblyInfo   `xml:"seq-assembly,omitempty" json:"seq_assembly,omitempty"`
	SeqMap        *ID2SSeqMapInfo        `xml:"seq-map,omitempty" json:"seq_map,omitempty"`
	SeqData       *ID2SSeqDataInfo       `xml:"seq-data,omitempty" json:"seq_data,omitempty"`
	SeqAnnotPlace *ID2SSeqAnnotPlaceInfo `xml:"seq-annot-place,omitempty" json:"seq_annot_place,omitempty"`
	BioseqPlace   []ID2SBioseqPlaceInfo  `xml:"bioseq-place,omitempty" json:"bioseq_place,omitempty"`
	FeatIds       []ID2SSeqFeatIdsInfo   `xml:"feat-ids,omitempty" json:"feat_ids,omitempty"`
}

type ID2SSeqDescrInfo struct {
	TypeMask   int64            `xml:"type-mask" json:"type_mask"`
	Bioseqs    ID2SBioseqIds    `xml:"bioseqs,omitempty" json:"bioseqs,omitempty" asn1:"optional"`
	BioseqSets ID2SBioseqSetIds `xml:"bioseq-sets,omitempty" json:"bioseq_sets,omitempty" asn1:"optional"`
}
type ID2SSeqAnnotInfo struct {
	Name   string             `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Align  interface{}        `xml:"-" json:"-" asn1:"optional"` //Align,NullType
	Graph  interface{}        `xml:"-" json:"-" asn1:"optional"` //Graph,NullType
	Feat   []ID2SFeatTypeInfo `xml:"feat,omitempty" json:"feat,omitempty" asn1:"optional"`
	SeqLoc *ID2SSeqLoc        `xml:"seq-loc,omitempty" json:"seq_loc,omitempty" asn1:"optional"`
}
type ID2SSeqAnnotPlaceInfo struct {
	Name       string           `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Bioseqs    ID2SBioseqIds    `xml:"bioseqs,omitempty" json:"bioseqs,omitempty" asn1:"optional"`
	BioseqSets ID2SBioseqSetIds `xml:"bioseq-sets,omitempty" json:"bioseq_sets,omitempty" asn1:"optional"`
}
type ID2SSeqFeatIdsInfo struct {
	FeatTypes   []ID2SFeatTypeInfo `xml:"feat-types,omitempty" json:"feat_types,omitempty" asn1:"optional"`
	XrefTypes   []ID2SFeatTypeInfo `xml:"xref-types,omitempty" json:"xref_types,omitempty" asn1:"optional"`
	LocalIds    []int64            `xml:"local-ids,omitempty" json:"local_ids,omitempty" asn1:"optional"`
	LocalStrIds []string           `xml:"local-str-ids,omitempty" json:"local_str_ids,omitempty" asn1:"optional"`
}
type ID2SFeatTypeInfo struct {
	Type     int64   `xml:"type" json:"type"`
	Subtypes []int64 `xml:"subtypes,omitempty" json:"subtypes,omitempty" asn1:"optional"`
}
type ID2SSeqAssemblyInfo struct {
	Bioseqs ID2SBioseqIds `xml:"bioseqs,omitempty" json:"bioseqs,omitempty"`
}
type ID2SSeqMapInfo ID2SSeqLoc
type ID2SSeqDataInfo ID2SSeqLoc
type ID2SBioseqPlaceInfo struct {
	BioseqSet int64         `xml:"bioseq-set" json:"bioseq_set"`
	SeqIds    ID2SBioseqIds `xml:"seq-ids,omitempty" json:"seq_ids,omitempty"`
}
type ID2SChunk struct {
	Data []ID2SChunkData `xml:"data,omitempty" json:"data,omitempty"`
}
type ID2SChunkData struct {
	Id struct {
		BioseqSet int64             `xml:"bioseq-set,omitempty" json:"bioseq_set,omitempty"`
		Gi        int64             `xml:"gi,omitempty" json:"gi,omitempty"`
		SeqId     *NCBISeqloc.SeqId `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
	} `xml:"id" json:"id"` //Id,ChoiceOption
	Descr    NCBISequence.SeqDescr   `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Annots   []NCBISequence.SeqAnnot `xml:"annots,omitempty" json:"annots,omitempty" asn1:"optional"`
	Assembly []NCBISeqalign.SeqAlign `xml:"assembly,omitempty" json:"assembly,omitempty" asn1:"optional"`
	SeqMap   []ID2SSequencePiece     `xml:"seq-map,omitempty" json:"seq_map,omitempty" asn1:"optional"`
	SeqData  []ID2SSequencePiece     `xml:"seq-data,omitempty" json:"seq_data,omitempty" asn1:"optional"`
	Bioseqs  []NCBISequence.Bioseq   `xml:"bioseqs,omitempty" json:"bioseqs,omitempty" asn1:"optional"`
}
type ID2SSequencePiece struct {
	Start int64                      `xml:"start" json:"start"`
	Data  []NCBISeqCommon.SeqLiteral `xml:"data,omitempty" json:"data,omitempty"`
}
type ID2SChunkId int64
type ID2SBioseqSetIds []int64
type ID2SBioseqIds []struct {
	Gi      int64             `xml:"gi,omitempty" json:"gi,omitempty"`
	SeqId   *NCBISeqloc.SeqId `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
	GiRange *ID2SGiRange      `xml:"gi-range,omitempty" json:"gi_range,omitempty"`
}
type ID2SGiRange struct {
	Start int64 `xml:"start" json:"start"`
	Count int64 `xml:"count" json:"count" asn1:"default:1"`
}

//ID2SSeqLoc,ChoiceOption
type ID2SSeqLoc struct {
	WholeGi       int64              `xml:"whole-gi,omitempty" json:"whole_gi,omitempty"`
	WholeSeqId    *NCBISeqloc.SeqId  `xml:"whole-seq-id,omitempty" json:"whole_seq_id,omitempty"`
	WholeGiRange  *ID2SGiRange       `xml:"whole-gi-range,omitempty" json:"whole_gi_range,omitempty"`
	GiInterval    *ID2SGiInterval    `xml:"gi-interval,omitempty" json:"gi_interval,omitempty"`
	SeqIdInterval *ID2SSeqIdInterval `xml:"seq-id-interval,omitempty" json:"seq_id_interval,omitempty"`
	GiInts        *ID2SGiInts        `xml:"gi-ints,omitempty" json:"gi_ints,omitempty"`
	SeqIdInts     *ID2SSeqIdInts     `xml:"seq-id-ints,omitempty" json:"seq_id_ints,omitempty"`
	LocSet        []ID2SSeqLoc       `xml:"loc-set,omitempty" json:"loc_set,omitempty"`
}

type ID2SGiInterval struct {
	Gi     int64 `xml:"gi" json:"gi"`
	Start  int64 `xml:"start" json:"start"`
	Length int64 `xml:"length" json:"length" asn1:"default:1"`
}
type ID2SSeqIdInterval struct {
	SeqId  *NCBISeqloc.SeqId `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
	Start  int64             `xml:"start" json:"start"`
	Length int64             `xml:"length" json:"length" asn1:"default:1"`
}
type ID2SInterval struct {
	Start  int64 `xml:"start" json:"start"`
	Length int64 `xml:"length" json:"length" asn1:"default:1"`
}
type ID2SGiInts struct {
	Gi   int64          `xml:"gi" json:"gi"`
	Ints []ID2SInterval `xml:"ints,omitempty" json:"ints,omitempty"`
}
type ID2SSeqIdInts struct {
	SeqId *NCBISeqloc.SeqId `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
	Ints  []ID2SInterval    `xml:"ints,omitempty" json:"ints,omitempty"`
}
