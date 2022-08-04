package NCBIID1Access

import (
	"ncbiasn/NCBISeqloc"
	"ncbiasn/NCBISequence"
)

//ID1serverRequest,ChoiceOption
type ID1serverRequest struct {
	Init            interface{}          `xml:"-" json:"-"` //Init,NullType
	Getgi           *NCBISeqloc.SeqId    `xml:"getgi,omitempty" json:"getgi,omitempty"`
	Getsefromgi     *ID1serverMaxcomplex `xml:"getsefromgi,omitempty" json:"getsefromgi,omitempty"`
	Fini            interface{}          `xml:"-" json:"-"` //Fini,NullType
	Getseqidsfromgi int64                `xml:"getseqidsfromgi,omitempty" json:"getseqidsfromgi,omitempty"`
	Getgihist       int64                `xml:"getgihist,omitempty" json:"getgihist,omitempty"`
	Getgirev        int64                `xml:"getgirev,omitempty" json:"getgirev,omitempty"`
	Getgistate      int64                `xml:"getgistate,omitempty" json:"getgistate,omitempty"`
	Getsewithinfo   *ID1serverMaxcomplex `xml:"getsewithinfo,omitempty" json:"getsewithinfo,omitempty"`
	Getblobinfo     *ID1serverMaxcomplex `xml:"getblobinfo,omitempty" json:"getblobinfo,omitempty"`
}

type ID1serverMaxcomplex struct {
	Maxplex *EntryComplexities `xml:"maxplex,omitempty" json:"maxplex,omitempty"`
	Gi      int64              `xml:"gi" json:"gi"`
	Ent     int64              `xml:"ent,omitempty" json:"ent,omitempty" asn1:"optional"`
	Sat     string             `xml:"sat,omitempty" json:"sat,omitempty" asn1:"optional"`
}

//EntryComplexities,IntegerEnum:entry(0),bioseq(1),bioseq-set(2),nuc-prot(3),pub-set(4)
type EntryComplexities int

type ID1SeqHist struct {
	Hist *NCBISequence.SeqHist `xml:"hist,omitempty" json:"hist,omitempty"`
}

//ID1serverBack,ChoiceOption
type ID1serverBack struct {
	Init            interface{}            `xml:"-" json:"-"` //Init,NullType
	Error           int64                  `xml:"error,omitempty" json:"error,omitempty"`
	Gotgi           int64                  `xml:"gotgi,omitempty" json:"gotgi,omitempty"`
	Gotseqentry     *NCBISequence.SeqEntry `xml:"gotseqentry,omitempty" json:"gotseqentry,omitempty"`
	Gotdeadseqentry *NCBISequence.SeqEntry `xml:"gotdeadseqentry,omitempty" json:"gotdeadseqentry,omitempty"`
	Fini            interface{}            `xml:"-" json:"-"` //Fini,NullType
	Gistate         int64                  `xml:"gistate,omitempty" json:"gistate,omitempty"`
	Ids             []NCBISeqloc.SeqId     `xml:"ids,omitempty" json:"ids,omitempty"`
	Gihist          []ID1SeqHist           `xml:"gihist,omitempty" json:"gihist,omitempty"`
	Girevhist       []ID1SeqHist           `xml:"girevhist,omitempty" json:"girevhist,omitempty"`
	Gotsewithinfo   *ID1SeqEntryInfo       `xml:"gotsewithinfo,omitempty" json:"gotsewithinfo,omitempty"`
	Gotblobinfo     *ID1blobInfo           `xml:"gotblobinfo,omitempty" json:"gotblobinfo,omitempty"`
}

type ID1serverDebug []ID1serverBack
type ID1blobInfo struct {
	Gi           int64  `xml:"gi" json:"gi"`
	Sat          int64  `xml:"sat" json:"sat"`
	SatKey       int64  `xml:"sat-key" json:"sat_key"`
	Satname      string `xml:"satname" json:"satname"`
	Suppress     int64  `xml:"suppress" json:"suppress"`
	Withdrawn    int64  `xml:"withdrawn" json:"withdrawn"`
	Confidential int64  `xml:"confidential" json:"confidential"`
	BlobState    int64  `xml:"blob-state" json:"blob_state"`
	Comment      string `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Extfeatmask  int64  `xml:"extfeatmask,omitempty" json:"extfeatmask,omitempty" asn1:"optional"`
}
type ID1SeqEntryInfo struct {
	BlobInfo *ID1blobInfo           `xml:"blob-info,omitempty" json:"blob_info,omitempty"`
	Blob     *NCBISequence.SeqEntry `xml:"blob,omitempty" json:"blob,omitempty" asn1:"optional"`
}
