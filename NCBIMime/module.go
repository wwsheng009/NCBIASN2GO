package NCBIMime

import (
	"ncbiasn/MMDB"
	"ncbiasn/NCBICdd"
	"ncbiasn/NCBICn3d"
	"ncbiasn/NCBIMedline"
	"ncbiasn/NCBISequence"
)

//NcbiMimeAsn1,ChoiceOption
type NcbiMimeAsn1 struct {
	Entrez     *EntrezGeneral         `xml:"entrez,omitempty" json:"entrez,omitempty"`
	Alignstruc *BiostrucAlign         `xml:"alignstruc,omitempty" json:"alignstruc,omitempty"`
	Alignseq   *BiostrucAlignSeq      `xml:"alignseq,omitempty" json:"alignseq,omitempty"`
	Strucseq   *BiostrucSeq           `xml:"strucseq,omitempty" json:"strucseq,omitempty"`
	Strucseqs  *BiostrucSeqs          `xml:"strucseqs,omitempty" json:"strucseqs,omitempty"`
	General    *BiostrucSeqsAlignsCdd `xml:"general,omitempty" json:"general,omitempty"`
}

type BundleSeqsAligns struct {
	Sequences       []NCBISequence.SeqEntry       `xml:"sequences,omitempty" json:"sequences,omitempty" asn1:"optional"`
	Seqaligns       []NCBISequence.SeqAnnot       `xml:"seqaligns,omitempty" json:"seqaligns,omitempty" asn1:"optional"`
	Strucaligns     *MMDB.BiostrucAnnotSet        `xml:"strucaligns,omitempty" json:"strucaligns,omitempty" asn1:"optional"`
	Imports         []NCBISequence.SeqAnnot       `xml:"imports,omitempty" json:"imports,omitempty" asn1:"optional"`
	StyleDictionary *NCBICn3d.Cn3dStyleDictionary `xml:"style-dictionary,omitempty" json:"style_dictionary,omitempty" asn1:"optional"`
	UserAnnotations *NCBICn3d.Cn3dUserAnnotations `xml:"user-annotations,omitempty" json:"user_annotations,omitempty" asn1:"optional"`
}
type BiostrucSeqsAlignsCdd struct {
	SeqAlignData struct {
		Bundle *BundleSeqsAligns `xml:"bundle,omitempty" json:"bundle,omitempty"`
		Cdd    *NCBICdd.Cdd      `xml:"cdd,omitempty" json:"cdd,omitempty"`
	} `xml:"seq-align-data" json:"seq_align_data"` //SeqAlignData,ChoiceOption
	Structures    []MMDB.Biostruc `xml:"structures,omitempty" json:"structures,omitempty" asn1:"optional"`
	StructureType string          `xml:"structure-type,omitempty" json:"structure_type,omitempty" asn1:"optional"` //StructureType,EnumList:ncbi-backbone(2),ncbi-all-atom(3),pdb-model(4)
}
type BiostrucAlign struct {
	Master          *MMDB.Biostruc                `xml:"master,omitempty" json:"master,omitempty"`
	Slaves          []MMDB.Biostruc               `xml:"slaves,omitempty" json:"slaves,omitempty"`
	Alignments      *MMDB.BiostrucAnnotSet        `xml:"alignments,omitempty" json:"alignments,omitempty"`
	Sequences       []NCBISequence.SeqEntry       `xml:"sequences,omitempty" json:"sequences,omitempty"`
	Seqalign        []NCBISequence.SeqAnnot       `xml:"seqalign,omitempty" json:"seqalign,omitempty"`
	StyleDictionary *NCBICn3d.Cn3dStyleDictionary `xml:"style-dictionary,omitempty" json:"style_dictionary,omitempty" asn1:"optional"`
	UserAnnotations *NCBICn3d.Cn3dUserAnnotations `xml:"user-annotations,omitempty" json:"user_annotations,omitempty" asn1:"optional"`
}
type BiostrucAlignSeq struct {
	Sequences       []NCBISequence.SeqEntry       `xml:"sequences,omitempty" json:"sequences,omitempty"`
	Seqalign        []NCBISequence.SeqAnnot       `xml:"seqalign,omitempty" json:"seqalign,omitempty"`
	StyleDictionary *NCBICn3d.Cn3dStyleDictionary `xml:"style-dictionary,omitempty" json:"style_dictionary,omitempty" asn1:"optional"`
	UserAnnotations *NCBICn3d.Cn3dUserAnnotations `xml:"user-annotations,omitempty" json:"user_annotations,omitempty" asn1:"optional"`
}
type BiostrucSeq struct {
	Structure       *MMDB.Biostruc                `xml:"structure,omitempty" json:"structure,omitempty"`
	Sequences       []NCBISequence.SeqEntry       `xml:"sequences,omitempty" json:"sequences,omitempty"`
	StyleDictionary *NCBICn3d.Cn3dStyleDictionary `xml:"style-dictionary,omitempty" json:"style_dictionary,omitempty" asn1:"optional"`
	UserAnnotations *NCBICn3d.Cn3dUserAnnotations `xml:"user-annotations,omitempty" json:"user_annotations,omitempty" asn1:"optional"`
}
type BiostrucSeqs struct {
	Structure       *MMDB.Biostruc                `xml:"structure,omitempty" json:"structure,omitempty"`
	Sequences       []NCBISequence.SeqEntry       `xml:"sequences,omitempty" json:"sequences,omitempty"`
	Seqalign        []NCBISequence.SeqAnnot       `xml:"seqalign,omitempty" json:"seqalign,omitempty"`
	StyleDictionary *NCBICn3d.Cn3dStyleDictionary `xml:"style-dictionary,omitempty" json:"style_dictionary,omitempty" asn1:"optional"`
	UserAnnotations *NCBICn3d.Cn3dUserAnnotations `xml:"user-annotations,omitempty" json:"user_annotations,omitempty" asn1:"optional"`
}

//EntrezStyle,EnumList:docsum(1),genbank(2),genpept(3),fasta(4),asn1(5),graphic(6),alignment(7),globalview(8),report(9),medlars(10),embl(11),pdb(12),kinemage(13)
type EntrezStyle string

type EntrezGeneral struct {
	Title string `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	Data  struct {
		Ml         *NCBIMedline.MedlineEntry `xml:"ml,omitempty" json:"ml,omitempty"`
		Prot       *NCBISequence.SeqEntry    `xml:"prot,omitempty" json:"prot,omitempty"`
		Nuc        *NCBISequence.SeqEntry    `xml:"nuc,omitempty" json:"nuc,omitempty"`
		Genome     *NCBISequence.SeqEntry    `xml:"genome,omitempty" json:"genome,omitempty"`
		Structure  *MMDB.Biostruc            `xml:"structure,omitempty" json:"structure,omitempty"`
		StrucAnnot *MMDB.BiostrucAnnotSet    `xml:"strucAnnot,omitempty" json:"strucAnnot,omitempty"`
	} `xml:"data" json:"data"` //Data,ChoiceOption
	Style    *EntrezStyle `xml:"style,omitempty" json:"style,omitempty"`
	Location string       `xml:"location,omitempty" json:"location,omitempty" asn1:"optional"`
}
