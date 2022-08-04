package NCBICdd

import (
	"ncbiasn/MMDB"
	"ncbiasn/NCBICn3d"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIOrganism"
	"ncbiasn/NCBIPub"
	"ncbiasn/NCBIScoreMat"
	"ncbiasn/NCBISeqalign"
	"ncbiasn/NCBISeqfeat"
	"ncbiasn/NCBISeqloc"
	"ncbiasn/NCBISequence"
)

type CddOrgRef struct {
	Reference   *NCBIOrganism.OrgRef `xml:"reference,omitempty" json:"reference,omitempty"`
	Active      bool                 `xml:"active" json:"active"`
	ParentTaxId int64                `xml:"parent-tax-id,omitempty" json:"parent_tax_id,omitempty" asn1:"optional"`
	Rank        string               `xml:"rank,omitempty" json:"rank,omitempty" asn1:"optional"`
}
type CddOrgRefSet []CddOrgRef

//CddPrefNodeDescr,ChoiceOption
type CddPrefNodeDescr struct {
	CreateDate  *NCBIGeneral.Date `xml:"create-date,omitempty" json:"create_date,omitempty"`
	Description string            `xml:"description,omitempty" json:"description,omitempty"`
}

type CddPrefNodeDescrSet []CddPrefNodeDescr
type CddPrefNodes struct {
	PreferredNodes CddOrgRefSet        `xml:"preferred-nodes,omitempty" json:"preferred_nodes,omitempty"`
	ModelOrganisms CddOrgRefSet        `xml:"model-organisms,omitempty" json:"model_organisms,omitempty" asn1:"optional"`
	OptionalNodes  CddOrgRefSet        `xml:"optional-nodes,omitempty" json:"optional_nodes,omitempty" asn1:"optional"`
	Description    CddPrefNodeDescrSet `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
}
type GlobalId struct {
	Accession string `xml:"accession" json:"accession"`
	Release   string `xml:"release,omitempty" json:"release,omitempty" asn1:"optional"`
	Version   int64  `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
	Database  string `xml:"database,omitempty" json:"database,omitempty" asn1:"optional"`
}

//CddId,ChoiceOption
type CddId struct {
	Uid int64     `xml:"uid,omitempty" json:"uid,omitempty"`
	Gid *GlobalId `xml:"gid,omitempty" json:"gid,omitempty"`
}

type CddIdSet []CddId
type CddRepeat struct {
	Count    int64              `xml:"count" json:"count"`
	Location *NCBISeqloc.SeqLoc `xml:"location,omitempty" json:"location,omitempty" asn1:"optional"`
	Avglen   int64              `xml:"avglen,omitempty" json:"avglen,omitempty" asn1:"optional"`
}
type CddBookRef struct {
	Bookname      string `xml:"bookname" json:"bookname"`
	Textelement   string `xml:"textelement" json:"textelement"` //Textelement,EnumList:unassigned(0),section(1),figgrp(2),table(3),chapter(4),biblist(5),box(6),glossary(7),appendix(8),other(255)
	Elementid     int64  `xml:"elementid,omitempty" json:"elementid,omitempty" asn1:"optional"`
	Subelementid  int64  `xml:"subelementid,omitempty" json:"subelementid,omitempty" asn1:"optional"`
	Celementid    string `xml:"celementid,omitempty" json:"celementid,omitempty" asn1:"optional"`
	Csubelementid string `xml:"csubelementid,omitempty" json:"csubelementid,omitempty" asn1:"optional"`
}

//CddDescr,ChoiceOption
type CddDescr struct {
	Othername      string               `xml:"othername,omitempty" json:"othername,omitempty"`
	Category       string               `xml:"category,omitempty" json:"category,omitempty"`
	Comment        string               `xml:"comment,omitempty" json:"comment,omitempty"`
	Reference      *NCBIPub.Pub         `xml:"reference,omitempty" json:"reference,omitempty"`
	CreateDate     *NCBIGeneral.Date    `xml:"create-date,omitempty" json:"create_date,omitempty"`
	TaxSource      *NCBIOrganism.OrgRef `xml:"tax-source,omitempty" json:"tax_source,omitempty"`
	Source         string               `xml:"source,omitempty" json:"source,omitempty"`
	Status         int                  `xml:"status,omitempty" json:"status,omitempty"` //Status,IntegerEnum:unassigned(0),finished-ok(1),pending-release(2),other-asis(3),matrix-only(4),update-running(5),auto-updated(6),claimed(7),curated-complete(8),other(255)
	UpdateDate     *NCBIGeneral.Date    `xml:"update-date,omitempty" json:"update_date,omitempty"`
	Scrapbook      []string             `xml:"scrapbook,omitempty" json:"scrapbook,omitempty"`
	SourceId       CddIdSet             `xml:"source-id,omitempty" json:"source_id,omitempty"`
	Repeats        *CddRepeat           `xml:"repeats,omitempty" json:"repeats,omitempty"`
	OldRoot        CddIdSet             `xml:"old-root,omitempty" json:"old_root,omitempty"`
	CurationStatus int                  `xml:"curation-status,omitempty" json:"curation_status,omitempty"` //CurationStatus,IntegerEnum:unassigned(0),prein(1),ofc(2),iac(3),ofv1(4),iav1(5),ofv2(6),iav2(7),postin(8),other(255)
	ReadonlyStatus int                  `xml:"readonly-status,omitempty" json:"readonly_status,omitempty"` //ReadonlyStatus,IntegerEnum:unassigned(0),readonly(1),readwrite(2),other(255)
	BookRef        *CddBookRef          `xml:"book-ref,omitempty" json:"book_ref,omitempty"`
	Attribution    *NCBIPub.Pub         `xml:"attribution,omitempty" json:"attribution,omitempty"`
	Title          string               `xml:"title,omitempty" json:"title,omitempty"`
}

type CddDescrSet []CddDescr
type CddTree struct {
	Name        string      `xml:"name" json:"name"`
	Id          CddIdSet    `xml:"id,omitempty" json:"id,omitempty"`
	Description CddDescrSet `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Parent      *CddId      `xml:"parent,omitempty" json:"parent,omitempty" asn1:"optional"`
	Children    CddIdSet    `xml:"children,omitempty" json:"children,omitempty" asn1:"optional"`
	Siblings    CddIdSet    `xml:"siblings,omitempty" json:"siblings,omitempty" asn1:"optional"`
	Neighbors   CddIdSet    `xml:"neighbors,omitempty" json:"neighbors,omitempty" asn1:"optional"`
}
type CddTreeSet []CddTree
type Matrix struct {
	Ncolumns    int64    `xml:"ncolumns" json:"ncolumns"`
	Nrows       int64    `xml:"nrows" json:"nrows"`
	RowLabels   []string `xml:"row-labels,omitempty" json:"row_labels,omitempty" asn1:"optional"`
	ScaleFactor int64    `xml:"scale-factor" json:"scale_factor"`
	Columns     []int64  `xml:"columns,omitempty" json:"columns,omitempty"`
}
type Triangle struct {
	Nelements int64                 `xml:"nelements" json:"nelements"`
	Scores    NCBISeqalign.ScoreSet `xml:"scores,omitempty" json:"scores,omitempty" asn1:"optional"`
	DivRanks  []int64               `xml:"div-ranks,omitempty" json:"div_ranks,omitempty" asn1:"optional"`
}

//UpdateComment,ChoiceOption
type UpdateComment struct {
	Comment   string             `xml:"comment,omitempty" json:"comment,omitempty"`
	Addthis   *NCBISeqloc.SeqLoc `xml:"addthis,omitempty" json:"addthis,omitempty"`
	Replaces  *NCBISeqloc.SeqLoc `xml:"replaces,omitempty" json:"replaces,omitempty"`
	RejectLoc *NCBISeqloc.SeqLoc `xml:"reject-loc,omitempty" json:"reject_loc,omitempty"`
	Reference *NCBIPub.Pub       `xml:"reference,omitempty" json:"reference,omitempty"`
}

type UpdateAlign struct {
	Description []UpdateComment        `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Seqannot    *NCBISequence.SeqAnnot `xml:"seqannot,omitempty" json:"seqannot,omitempty" asn1:"optional"`
	Type        int                    `xml:"type" json:"type"` //Type,IntegerEnum:unassigned(0),update(1),update-3d(2),demoted(51),demoted-3d(52),other(255)
}
type RejectId struct {
	Description []UpdateComment    `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Ids         []NCBISeqloc.SeqId `xml:"ids,omitempty" json:"ids,omitempty"`
}

//FeatureEvidence,ChoiceOption

type FeatureEvidence struct {
	Comment   string                 `xml:"comment,omitempty" json:"comment,omitempty"`
	Reference *NCBIPub.Pub           `xml:"reference,omitempty" json:"reference,omitempty"`
	Bsannot   *MMDB.BiostrucAnnotSet `xml:"bsannot,omitempty" json:"bsannot,omitempty"`
	Seqfeat   *NCBISeqfeat.SeqFeat   `xml:"seqfeat,omitempty" json:"seqfeat,omitempty"`
	BookRef   *CddBookRef            `xml:"book-ref,omitempty" json:"book_ref,omitempty"`
}

type AlignAnnot struct {
	Location    *NCBISeqloc.SeqLoc `xml:"location,omitempty" json:"location,omitempty"`
	Description string             `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Evidence    []FeatureEvidence  `xml:"evidence,omitempty" json:"evidence,omitempty" asn1:"optional"`
	Type        int64              `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"`
	Aliases     []string           `xml:"aliases,omitempty" json:"aliases,omitempty" asn1:"optional"`
	Motif       string             `xml:"motif,omitempty" json:"motif,omitempty" asn1:"optional"`
	Motifuse    int64              `xml:"motifuse,omitempty" json:"motifuse,omitempty" asn1:"optional"`
}
type AlignAnnotSet []AlignAnnot
type DomainParent struct {
	ParentType int                    `xml:"parent-type" json:"parent_type"` //ParentType,IntegerEnum:classical(0),fusion(1),deletion(2),permutation(3),other(255)
	Parentid   *CddId                 `xml:"parentid,omitempty" json:"parentid,omitempty"`
	Seqannot   *NCBISequence.SeqAnnot `xml:"seqannot,omitempty" json:"seqannot,omitempty" asn1:"optional"`
}
type SequenceTree struct {
	CdAccession string         `xml:"cdAccession,omitempty" json:"cdAccession,omitempty" asn1:"optional"`
	Algorithm   *AlgorithmType `xml:"algorithm,omitempty" json:"algorithm,omitempty"`
	IsAnnotated bool           `xml:"isAnnotated" json:"isAnnotated"`
	Root        *SeqTreeNode   `xml:"root,omitempty" json:"root,omitempty"`
}
type SeqTreeNode struct {
	IsAnnotated bool    `xml:"isAnnotated" json:"isAnnotated"`
	Name        string  `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Distance    float64 `xml:"distance,omitempty" json:"distance,omitempty" asn1:"optional"`
	Children    struct {
		Children  []SeqTreeNode `xml:"children,omitempty" json:"children,omitempty"`
		Footprint struct {
			SeqRange *NCBISeqloc.SeqInterval `xml:"seqRange,omitempty" json:"seqRange,omitempty"`
			RowId    int64                   `xml:"rowId,omitempty" json:"rowId,omitempty" asn1:"optional"`
		} `xml:"footprint,omitempty" json:"footprint,omitempty"`
	} `xml:"children" json:"children"` //Children,ChoiceOption
	Annotation *NodeAnnotation `xml:"annotation,omitempty" json:"annotation,omitempty" asn1:"optional"`
}
type AlgorithmType struct {
	ScoringScheme    int   `xml:"scoring-Scheme" json:"scoring_Scheme"`                                 //ScoringScheme,IntegerEnum:unassigned(0),percent-id(1),kimura-corrected(2),aligned-score(3),aligned-score-ext(4),aligned-score-filled(5),blast-footprint(6),blast-full(7),hybrid-aligned-score(8),other(255)
	ClusteringMethod int   `xml:"clustering-Method" json:"clustering_Method"`                           //ClusteringMethod,IntegerEnum:unassigned(0),single-linkage(1),neighbor-joining(2),fast-minimum-evolution(3),other(255)
	ScoreMatrix      int   `xml:"score-Matrix,omitempty" json:"score_Matrix,omitempty" asn1:"optional"` //ScoreMatrix,IntegerEnum:unassigned(0),blosum45(1),blosum62(2),blosum80(3),pam30(4),pam70(5),pam250(6),other(255)
	GapOpen          int64 `xml:"gapOpen,omitempty" json:"gapOpen,omitempty" asn1:"optional"`
	GapExtend        int64 `xml:"gapExtend,omitempty" json:"gapExtend,omitempty" asn1:"optional"`
	GapScaleFactor   int64 `xml:"gapScaleFactor,omitempty" json:"gapScaleFactor,omitempty" asn1:"optional"`
	NTerminalExt     int64 `xml:"nTerminalExt,omitempty" json:"nTerminalExt,omitempty" asn1:"optional"`
	CTerminalExt     int64 `xml:"cTerminalExt,omitempty" json:"cTerminalExt,omitempty" asn1:"optional"`
	TreeScope        int   `xml:"tree-scope,omitempty" json:"tree_scope,omitempty" asn1:"optional"`         //TreeScope,IntegerEnum:allDescendants(0),immediateChildrenOnly(1),selfOnly(2),other(255)
	ColoringScope    int   `xml:"coloring-scope,omitempty" json:"coloring_scope,omitempty" asn1:"optional"` //ColoringScope,IntegerEnum:allDescendants(0),immediateChildrenOnly(1),other(255)
}
type NodeAnnotation struct {
	PresentInChildCD string `xml:"presentInChildCD,omitempty" json:"presentInChildCD,omitempty" asn1:"optional"`
	Note             string `xml:"note,omitempty" json:"note,omitempty" asn1:"optional"`
}
type Cdd struct {
	Name            string                           `xml:"name" json:"name"`
	Id              CddIdSet                         `xml:"id,omitempty" json:"id,omitempty"`
	Description     CddDescrSet                      `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Seqannot        []NCBISequence.SeqAnnot          `xml:"seqannot,omitempty" json:"seqannot,omitempty" asn1:"optional"`
	Features        *MMDB.BiostrucAnnotSet           `xml:"features,omitempty" json:"features,omitempty" asn1:"optional"`
	Sequences       *NCBISequence.SeqEntry           `xml:"sequences,omitempty" json:"sequences,omitempty" asn1:"optional"`
	ProfileRange    *NCBISeqloc.SeqInterval          `xml:"profile-range,omitempty" json:"profile_range,omitempty" asn1:"optional"`
	TruncMaster     *NCBISequence.Bioseq             `xml:"trunc-master,omitempty" json:"trunc_master,omitempty" asn1:"optional"`
	Posfreq         *Matrix                          `xml:"posfreq,omitempty" json:"posfreq,omitempty" asn1:"optional"`
	Scoremat        *Matrix                          `xml:"scoremat,omitempty" json:"scoremat,omitempty" asn1:"optional"`
	Distance        *Triangle                        `xml:"distance,omitempty" json:"distance,omitempty" asn1:"optional"`
	Parent          *CddId                           `xml:"parent,omitempty" json:"parent,omitempty" asn1:"optional"`
	Children        CddIdSet                         `xml:"children,omitempty" json:"children,omitempty" asn1:"optional"`
	Siblings        CddIdSet                         `xml:"siblings,omitempty" json:"siblings,omitempty" asn1:"optional"`
	Neighbors       CddIdSet                         `xml:"neighbors,omitempty" json:"neighbors,omitempty" asn1:"optional"`
	Pending         []UpdateAlign                    `xml:"pending,omitempty" json:"pending,omitempty" asn1:"optional"`
	Rejects         []RejectId                       `xml:"rejects,omitempty" json:"rejects,omitempty" asn1:"optional"`
	Master3d        []NCBISeqloc.SeqId               `xml:"master3d,omitempty" json:"master3d,omitempty" asn1:"optional"`
	Alignannot      AlignAnnotSet                    `xml:"alignannot,omitempty" json:"alignannot,omitempty" asn1:"optional"`
	StyleDictionary *NCBICn3d.Cn3dStyleDictionary    `xml:"style-dictionary,omitempty" json:"style_dictionary,omitempty" asn1:"optional"`
	UserAnnotations *NCBICn3d.Cn3dUserAnnotations    `xml:"user-annotations,omitempty" json:"user_annotations,omitempty" asn1:"optional"`
	Ancestors       []DomainParent                   `xml:"ancestors,omitempty" json:"ancestors,omitempty" asn1:"optional"`
	Scoreparams     *NCBIScoreMat.PssmWithParameters `xml:"scoreparams,omitempty" json:"scoreparams,omitempty" asn1:"optional"`
	Seqtree         *SequenceTree                    `xml:"seqtree,omitempty" json:"seqtree,omitempty" asn1:"optional"`
}
type CddSet []Cdd
type CddViewerRect struct {
	Top    int64 `xml:"top" json:"top"`
	Left   int64 `xml:"left" json:"left"`
	Width  int64 `xml:"width" json:"width"`
	Height int64 `xml:"height" json:"height"`
}
type CddViewer struct {
	Ctrl       int            `xml:"ctrl" json:"ctrl"` //Ctrl,IntegerEnum:unassigned(0),cd-info(1),align-annot(2),seq-list(3),seq-tree(4),merge-preview(5),cross-hits(6),notes(7),tax-tree(8),dart(9),dart-selected-rows(10),other(255)
	Rect       *CddViewerRect `xml:"rect,omitempty" json:"rect,omitempty" asn1:"optional"`
	Accessions []string       `xml:"accessions,omitempty" json:"accessions,omitempty"`
}
type CddScript struct {
	Type     int    `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"` //Type,IntegerEnum:unassigned(0),user-recorded(1),server-generated(2),other(255)
	Name     string `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Commands string `xml:"commands" json:"commands"`
}
type CddProject struct {
	Cds        []Cdd             `xml:"cds,omitempty" json:"cds,omitempty"`
	Cdcolor    []int64           `xml:"cdcolor,omitempty" json:"cdcolor,omitempty"`
	Viewers    []CddViewer       `xml:"viewers,omitempty" json:"viewers,omitempty"`
	Log        string            `xml:"log" json:"log"`
	Scripts    []CddScript       `xml:"scripts,omitempty" json:"scripts,omitempty" asn1:"optional"`
	Id         CddIdSet          `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Rids       []string          `xml:"rids,omitempty" json:"rids,omitempty" asn1:"optional"`
	CreateDate *NCBIGeneral.Date `xml:"create-date,omitempty" json:"create_date,omitempty" asn1:"optional"`
	UpdateDate *NCBIGeneral.Date `xml:"update-date,omitempty" json:"update_date,omitempty" asn1:"optional"`
	ProjectId  int64             `xml:"project-id,omitempty" json:"project_id,omitempty" asn1:"optional"`
}
