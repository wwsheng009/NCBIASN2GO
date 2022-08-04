package HomoloGene

import (
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBISeqalign"
	"ncbiasn/NCBISeqloc"
)

type HGEntrySet struct {
	Entries []HGEntry `xml:"entries,omitempty" json:"entries,omitempty"`
}
type HGEntry struct {
	HgId         int64             `xml:"hg-id" json:"hg_id"`
	Version      int64             `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
	Title        string            `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	Caption      string            `xml:"caption,omitempty" json:"caption,omitempty" asn1:"optional"`
	Taxid        int64             `xml:"taxid,omitempty" json:"taxid,omitempty" asn1:"optional"`
	Genes        []HGGene          `xml:"genes,omitempty" json:"genes,omitempty" asn1:"optional"`
	CrDate       *NCBIGeneral.Date `xml:"cr-date,omitempty" json:"cr_date,omitempty" asn1:"optional"`
	UpDate       *NCBIGeneral.Date `xml:"up-date,omitempty" json:"up_date,omitempty" asn1:"optional"`
	Distances    []HGStats         `xml:"distances,omitempty" json:"distances,omitempty" asn1:"optional"`
	Commentaries []HGCommentarySet `xml:"commentaries,omitempty" json:"commentaries,omitempty" asn1:"optional"`
	Warnings     []string          `xml:"warnings,omitempty" json:"warnings,omitempty" asn1:"optional"`
	Node         *HGNode           `xml:"node,omitempty" json:"node,omitempty" asn1:"optional"`
}
type HGGene struct {
	Geneid    int64              `xml:"geneid" json:"geneid"`
	Otherid   int64              `xml:"otherid,omitempty" json:"otherid,omitempty" asn1:"optional"`
	Symbol    string             `xml:"symbol,omitempty" json:"symbol,omitempty" asn1:"optional"`
	Aliases   []string           `xml:"aliases,omitempty" json:"aliases,omitempty" asn1:"optional"`
	Title     string             `xml:"title" json:"title"`
	Taxid     int64              `xml:"taxid" json:"taxid"`
	ProtGi    int64              `xml:"prot-gi,omitempty" json:"prot_gi,omitempty" asn1:"optional"`
	ProtAcc   string             `xml:"prot-acc,omitempty" json:"prot_acc,omitempty" asn1:"optional"`
	ProtLen   int64              `xml:"prot-len,omitempty" json:"prot_len,omitempty" asn1:"optional"`
	NucGi     int64              `xml:"nuc-gi,omitempty" json:"nuc_gi,omitempty" asn1:"optional"`
	NucAcc    string             `xml:"nuc-acc,omitempty" json:"nuc_acc,omitempty" asn1:"optional"`
	GeneLinks []HGLink           `xml:"gene-links,omitempty" json:"gene_links,omitempty" asn1:"optional"`
	ProtLinks []HGLink           `xml:"prot-links,omitempty" json:"prot_links,omitempty" asn1:"optional"`
	Domains   []HGDomain         `xml:"domains,omitempty" json:"domains,omitempty" asn1:"optional"`
	Chr       string             `xml:"chr,omitempty" json:"chr,omitempty" asn1:"optional"`
	Location  *NCBISeqloc.SeqLoc `xml:"location,omitempty" json:"location,omitempty" asn1:"optional"`
	LocusTag  string             `xml:"locus-tag,omitempty" json:"locus_tag,omitempty" asn1:"optional"`
}
type HGStats struct {
	Gi1         int64   `xml:"gi1" json:"gi1"`
	Gi2         int64   `xml:"gi2" json:"gi2"`
	NucChange   float64 `xml:"nuc-change" json:"nuc_change"`
	NucChangeJc float64 `xml:"nuc-change-jc" json:"nuc_change_jc"`
	ProtChange  float64 `xml:"prot-change" json:"prot_change"`
	Ka          float64 `xml:"ka" json:"ka"`
	Ks          float64 `xml:"ks" json:"ks"`
	Knr         float64 `xml:"knr" json:"knr"`
	Knc         float64 `xml:"knc" json:"knc"`
	RecipBest   bool    `xml:"recip-best,omitempty" json:"recip_best,omitempty" asn1:"optional"`
}
type HGCommentary struct {
	Link              *HGLink        `xml:"link,omitempty" json:"link,omitempty"`
	Description       string         `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Caption           string         `xml:"caption,omitempty" json:"caption,omitempty" asn1:"optional"`
	Provider          string         `xml:"provider,omitempty" json:"provider,omitempty" asn1:"optional"`
	OtherLinks        []HGLink       `xml:"other-links,omitempty" json:"other_links,omitempty" asn1:"optional"`
	OtherCommentaries []HGCommentary `xml:"other-commentaries,omitempty" json:"other_commentaries,omitempty" asn1:"optional"`
	Taxid             int64          `xml:"taxid,omitempty" json:"taxid,omitempty" asn1:"optional"`
	Geneid            int64          `xml:"geneid,omitempty" json:"geneid,omitempty" asn1:"optional"`
}
type HGCommentarySet struct {
	HgId         int64          `xml:"hg-id,omitempty" json:"hg_id,omitempty" asn1:"optional"`
	Title        string         `xml:"title" json:"title"`
	Commentaries []HGCommentary `xml:"commentaries,omitempty" json:"commentaries,omitempty"`
}
type HGCommentaryContainer []HGCommentarySet
type HGLink struct {
	Hypertext string `xml:"hypertext" json:"hypertext"`
	Url       string `xml:"url,omitempty" json:"url,omitempty" asn1:"optional"`
}
type HGDomain struct {
	Begin   int64  `xml:"begin" json:"begin"`
	End     int64  `xml:"end" json:"end"`
	PssmId  int64  `xml:"pssm-id,omitempty" json:"pssm_id,omitempty" asn1:"optional"`
	CddId   string `xml:"cdd-id,omitempty" json:"cdd_id,omitempty" asn1:"optional"`
	CddName string `xml:"cdd-name,omitempty" json:"cdd_name,omitempty" asn1:"optional"`
}
type HGNode struct {
	Type        string    `xml:"type" json:"type"` //Type,EnumList:family(0),ortholog(1),paralog(2),leaf(3)
	Id          *HGNodeId `xml:"id,omitempty" json:"id,omitempty"`
	Caption     string    `xml:"caption,omitempty" json:"caption,omitempty" asn1:"optional"`
	CurrentNode bool      `xml:"current-node" json:"current_node"`
	Children    []HGNode  `xml:"children,omitempty" json:"children,omitempty" asn1:"optional"`
	BranchLen   int64     `xml:"branch-len,omitempty" json:"branch_len,omitempty" asn1:"optional"`
}
type HGNodeId struct {
	Id     int64  `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	IdType string `xml:"id-type" json:"id_type"` //IdType,EnumList:none(0),geneid(1),hid(2)
}
type HGAlignment struct {
	HgId      int64                  `xml:"hg-id" json:"hg_id"`
	Alignment *NCBISeqalign.SeqAlign `xml:"alignment,omitempty" json:"alignment,omitempty"`
}
type HGAlignmentSet []HGAlignment
