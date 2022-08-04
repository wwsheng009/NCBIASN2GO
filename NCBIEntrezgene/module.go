package NCBIEntrezgene

import (
	"ncbiasn/NCBIBioSource"
	"ncbiasn/NCBIGene"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIProtein"
	"ncbiasn/NCBIPub"
	"ncbiasn/NCBISeqloc"
	"ncbiasn/NCBI_RNA"
)

type Entrezgene struct {
	TrackInfo      *GeneTrack               `xml:"track-info,omitempty" json:"track_info,omitempty" asn1:"optional"`
	Type           int                      `xml:"type" json:"type"` //Type,IntegerEnum:unknown(0),tRNA(1),rRNA(2),snRNA(3),scRNA(4),snoRNA(5),protein-coding(6),pseudo(7),transposon(8),miscRNA(9),ncRNA(10),biological-region(11),other(255)
	Source         *NCBIBioSource.BioSource `xml:"source,omitempty" json:"source,omitempty"`
	Gene           *NCBIGene.GeneRef        `xml:"gene,omitempty" json:"gene,omitempty"`
	Prot           *NCBIProtein.ProtRef     `xml:"prot,omitempty" json:"prot,omitempty" asn1:"optional"`
	Rna            *NCBI_RNA.RNARef         `xml:"rna,omitempty" json:"rna,omitempty" asn1:"optional"`
	Summary        string                   `xml:"summary,omitempty" json:"summary,omitempty" asn1:"optional"`
	Location       []Maps                   `xml:"location,omitempty" json:"location,omitempty" asn1:"optional"`
	GeneSource     *GeneSource              `xml:"gene-source,omitempty" json:"gene_source,omitempty" asn1:"optional"`
	Locus          []GeneCommentary         `xml:"locus,omitempty" json:"locus,omitempty" asn1:"optional"`
	Properties     []GeneCommentary         `xml:"properties,omitempty" json:"properties,omitempty" asn1:"optional"`
	Refgene        []GeneCommentary         `xml:"refgene,omitempty" json:"refgene,omitempty" asn1:"optional"`
	Homology       []GeneCommentary         `xml:"homology,omitempty" json:"homology,omitempty" asn1:"optional"`
	Comments       []GeneCommentary         `xml:"comments,omitempty" json:"comments,omitempty" asn1:"optional"`
	UniqueKeys     []NCBIGeneral.Dbtag      `xml:"unique-keys,omitempty" json:"unique_keys,omitempty" asn1:"optional"`
	XtraIndexTerms []string                 `xml:"xtra-index-terms,omitempty" json:"xtra_index_terms,omitempty" asn1:"optional"`
	XtraProperties []XtraTerms              `xml:"xtra-properties,omitempty" json:"xtra_properties,omitempty" asn1:"optional"`
	XtraIq         []XtraTerms              `xml:"xtra-iq,omitempty" json:"xtra_iq,omitempty" asn1:"optional"`
	NonUniqueKeys  []NCBIGeneral.Dbtag      `xml:"non-unique-keys,omitempty" json:"non_unique_keys,omitempty" asn1:"optional"`
}
type EntrezgeneSet []Entrezgene
type GeneTrack struct {
	Geneid          int64               `xml:"geneid" json:"geneid"`
	Status          int                 `xml:"status" json:"status"` //Status,IntegerEnum:live(0),secondary(1),discontinued(2)
	CurrentId       []NCBIGeneral.Dbtag `xml:"current-id,omitempty" json:"current_id,omitempty" asn1:"optional"`
	CreateDate      *NCBIGeneral.Date   `xml:"create-date,omitempty" json:"create_date,omitempty"`
	UpdateDate      *NCBIGeneral.Date   `xml:"update-date,omitempty" json:"update_date,omitempty"`
	DiscontinueDate *NCBIGeneral.Date   `xml:"discontinue-date,omitempty" json:"discontinue_date,omitempty" asn1:"optional"`
}
type GeneSource struct {
	Src          string `xml:"src" json:"src"`
	SrcInt       int64  `xml:"src-int,omitempty" json:"src_int,omitempty" asn1:"optional"`
	SrcStr1      string `xml:"src-str1,omitempty" json:"src_str1,omitempty" asn1:"optional"`
	SrcStr2      string `xml:"src-str2,omitempty" json:"src_str2,omitempty" asn1:"optional"`
	GeneDisplay  bool   `xml:"gene-display" json:"gene_display"`
	LocusDisplay bool   `xml:"locus-display" json:"locus_display"`
	ExtraTerms   bool   `xml:"extra-terms" json:"extra_terms"`
}
type GeneCommentary struct {
	Type           int                 `xml:"type" json:"type"` //Type,IntegerEnum:genomic(1),pre-RNA(2),mRNA(3),rRNA(4),tRNA(5),snRNA(6),scRNA(7),peptide(8),other-genetic(9),genomic-mRNA(10),cRNA(11),mature-peptide(12),pre-protein(13),miscRNA(14),snoRNA(15),property(16),reference(17),generif(18),phenotype(19),complex(20),compound(21),ncRNA(22),gene-group(23),assembly(24),assembly-unit(25),c-region(26),d-segment(27),j-segment(28),v-segment(29),comment(254),other(255)
	Heading        string              `xml:"heading,omitempty" json:"heading,omitempty" asn1:"optional"`
	Label          string              `xml:"label,omitempty" json:"label,omitempty" asn1:"optional"`
	Text           string              `xml:"text,omitempty" json:"text,omitempty" asn1:"optional"`
	Accession      string              `xml:"accession,omitempty" json:"accession,omitempty" asn1:"optional"`
	Version        int64               `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
	XtraProperties []XtraTerms         `xml:"xtra-properties,omitempty" json:"xtra_properties,omitempty" asn1:"optional"`
	Refs           []NCBIPub.Pub       `xml:"refs,omitempty" json:"refs,omitempty" asn1:"optional"`
	Source         []OtherSource       `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"`
	GenomicCoords  []NCBISeqloc.SeqLoc `xml:"genomic-coords,omitempty" json:"genomic_coords,omitempty" asn1:"optional"`
	Seqs           []NCBISeqloc.SeqLoc `xml:"seqs,omitempty" json:"seqs,omitempty" asn1:"optional"`
	Products       []GeneCommentary    `xml:"products,omitempty" json:"products,omitempty" asn1:"optional"`
	Properties     []GeneCommentary    `xml:"properties,omitempty" json:"properties,omitempty" asn1:"optional"`
	Comment        []GeneCommentary    `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	CreateDate     *NCBIGeneral.Date   `xml:"create-date,omitempty" json:"create_date,omitempty" asn1:"optional"`
	UpdateDate     *NCBIGeneral.Date   `xml:"update-date,omitempty" json:"update_date,omitempty" asn1:"optional"`
	Rna            *NCBI_RNA.RNARef    `xml:"rna,omitempty" json:"rna,omitempty" asn1:"optional"`
}
type OtherSource struct {
	Src      *NCBIGeneral.Dbtag `xml:"src,omitempty" json:"src,omitempty" asn1:"optional"`
	PreText  string             `xml:"pre-text,omitempty" json:"pre_text,omitempty" asn1:"optional"`
	Anchor   string             `xml:"anchor,omitempty" json:"anchor,omitempty" asn1:"optional"`
	Url      string             `xml:"url,omitempty" json:"url,omitempty" asn1:"optional"`
	PostText string             `xml:"post-text,omitempty" json:"post_text,omitempty" asn1:"optional"`
}
type Maps struct {
	DisplayStr string `xml:"display-str" json:"display_str"`
	Method     struct {
		Proxy   string `xml:"proxy,omitempty" json:"proxy,omitempty"`
		MapType string `xml:"map-type,omitempty" json:"map_type,omitempty"` //MapType,EnumList:cyto(0),bp(1),cM(2),cR(3),min(4)
	} `xml:"method" json:"method"` //Method,ChoiceOption
}
type XtraTerms struct {
	Tag   string `xml:"tag" json:"tag"`
	Value string `xml:"value" json:"value"`
}
