package NCBIProject

import (
	"ncbiasn/NCBIBiblio"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIPubMed"
	"ncbiasn/NCBISeqRef"
	"ncbiasn/NCBISeqloc"
	"ncbiasn/NCBISequence"
)

type Project struct {
	Descr *ProjectDescr `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Data  *ProjectItem  `xml:"data,omitempty" json:"data,omitempty"`
}

//ProjectItem,ChoiceOption
type ProjectItem struct {
	Pmuid     []int64                  `xml:"pmuid,omitempty" json:"pmuid,omitempty"`
	Protuid   []int64                  `xml:"protuid,omitempty" json:"protuid,omitempty"`
	Nucuid    []int64                  `xml:"nucuid,omitempty" json:"nucuid,omitempty"`
	Sequid    []int64                  `xml:"sequid,omitempty" json:"sequid,omitempty"`
	Genomeuid []int64                  `xml:"genomeuid,omitempty" json:"genomeuid,omitempty"`
	Structuid []int64                  `xml:"structuid,omitempty" json:"structuid,omitempty"`
	Pmid      []NCBIBiblio.PubMedId    `xml:"pmid,omitempty" json:"pmid,omitempty"`
	Protid    []NCBISeqloc.SeqId       `xml:"protid,omitempty" json:"protid,omitempty"`
	Nucid     []NCBISeqloc.SeqId       `xml:"nucid,omitempty" json:"nucid,omitempty"`
	Seqid     []NCBISeqloc.SeqId       `xml:"seqid,omitempty" json:"seqid,omitempty"`
	Genomeid  []NCBISeqloc.SeqId       `xml:"genomeid,omitempty" json:"genomeid,omitempty"`
	Structid  interface{}              `xml:"-" json:"-"` //Structid,NullType
	Pment     []NCBIPubMed.PubmedEntry `xml:"pment,omitempty" json:"pment,omitempty"`
	Protent   []NCBISequence.SeqEntry  `xml:"protent,omitempty" json:"protent,omitempty"`
	Nucent    []NCBISequence.SeqEntry  `xml:"nucent,omitempty" json:"nucent,omitempty"`
	Seqent    []NCBISequence.SeqEntry  `xml:"seqent,omitempty" json:"seqent,omitempty"`
	Genomeent []NCBISequence.SeqEntry  `xml:"genomeent,omitempty" json:"genomeent,omitempty"`
	Structent interface{}              `xml:"-" json:"-"` //Structent,NullType
	Seqannot  []NCBISequence.SeqAnnot  `xml:"seqannot,omitempty" json:"seqannot,omitempty"`
	Loc       []NCBISeqloc.SeqLoc      `xml:"loc,omitempty" json:"loc,omitempty"`
	Proj      []Project                `xml:"proj,omitempty" json:"proj,omitempty"`
}

type ProjectDescr struct {
	Id    []ProjectId `xml:"id,omitempty" json:"id,omitempty"`
	Name  string      `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Descr []Projdesc  `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
}

//Projdesc,ChoiceOption
type Projdesc struct {
	Pub     *NCBISeqRef.Pubdesc `xml:"pub,omitempty" json:"pub,omitempty"`
	Date    *NCBIGeneral.Date   `xml:"date,omitempty" json:"date,omitempty"`
	Comment string              `xml:"comment,omitempty" json:"comment,omitempty"`
	Title   string              `xml:"title,omitempty" json:"title,omitempty"`
}

type ProjectId string
