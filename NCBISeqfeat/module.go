package NCBISeqfeat

import (
	"ncbiasn/NCBIBiblio"
	"ncbiasn/NCBIBioSource"
	"ncbiasn/NCBIGene"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIOrganism"
	"ncbiasn/NCBIProtein"
	"ncbiasn/NCBIPub"
	"ncbiasn/NCBIRsite"
	"ncbiasn/NCBISeqCommon"
	"ncbiasn/NCBISeqRef"
	"ncbiasn/NCBISeqloc"
	"ncbiasn/NCBITxInit"
	"ncbiasn/NCBIVariation"
	"ncbiasn/NCBI_RNA"
)

//FeatId,ChoiceOption
// type FeatId struct {
// 	Gibb	int64			`xml:"gibb,omitempty" json:"gibb,omitempty"`
// 	Giim	*NCBISeqloc.GiimportId	`xml:"giim,omitempty" json:"giim,omitempty"`
// 	Local	*NCBIGeneral.ObjectId	`xml:"local,omitempty" json:"local,omitempty"`
// 	General	*NCBIGeneral.Dbtag	`xml:"general,omitempty" json:"general,omitempty"`
// }

type SeqFeat struct {
	Id         *NCBISeqCommon.FeatId    `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Data       *SeqFeatData             `xml:"data,omitempty" json:"data,omitempty"`
	Partial    bool                     `xml:"partial,omitempty" json:"partial,omitempty" asn1:"optional"`
	Except     bool                     `xml:"except,omitempty" json:"except,omitempty" asn1:"optional"`
	Comment    string                   `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Product    *NCBISeqloc.SeqLoc       `xml:"product,omitempty" json:"product,omitempty" asn1:"optional"`
	Location   *NCBISeqloc.SeqLoc       `xml:"location,omitempty" json:"location,omitempty"`
	Qual       []GbQual                 `xml:"qual,omitempty" json:"qual,omitempty" asn1:"optional"`
	Title      string                   `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	Ext        *NCBIGeneral.UserObject  `xml:"ext,omitempty" json:"ext,omitempty" asn1:"optional"`
	Cit        *NCBIPub.PubSet          `xml:"cit,omitempty" json:"cit,omitempty" asn1:"optional"`
	ExpEv      string                   `xml:"exp-ev,omitempty" json:"exp_ev,omitempty" asn1:"optional"` //ExpEv,EnumList:experimental(1),not-experimental(2)
	Xref       []SeqFeatXref            `xml:"xref,omitempty" json:"xref,omitempty" asn1:"optional"`
	Dbxref     []NCBIGeneral.Dbtag      `xml:"dbxref,omitempty" json:"dbxref,omitempty" asn1:"optional"`
	Pseudo     bool                     `xml:"pseudo,omitempty" json:"pseudo,omitempty" asn1:"optional"`
	ExceptText string                   `xml:"except-text,omitempty" json:"except_text,omitempty" asn1:"optional"`
	Ids        []NCBISeqCommon.FeatId   `xml:"ids,omitempty" json:"ids,omitempty" asn1:"optional"`
	Exts       []NCBIGeneral.UserObject `xml:"exts,omitempty" json:"exts,omitempty" asn1:"optional"`
	Support    *SeqFeatSupport          `xml:"support,omitempty" json:"support,omitempty" asn1:"optional"`
}

//SeqFeatData,ChoiceOption
type SeqFeatData struct {
	Gene          *NCBIGene.GeneRef           `xml:"gene,omitempty" json:"gene,omitempty"`
	Org           *NCBIOrganism.OrgRef        `xml:"org,omitempty" json:"org,omitempty"`
	Cdregion      *Cdregion                   `xml:"cdregion,omitempty" json:"cdregion,omitempty"`
	Prot          *NCBIProtein.ProtRef        `xml:"prot,omitempty" json:"prot,omitempty"`
	Rna           *NCBI_RNA.RNARef            `xml:"rna,omitempty" json:"rna,omitempty"`
	Pub           *NCBISeqRef.Pubdesc         `xml:"pub,omitempty" json:"pub,omitempty"`
	Seq           *NCBISeqloc.SeqLoc          `xml:"seq,omitempty" json:"seq,omitempty"`
	Imp           *ImpFeat                    `xml:"imp,omitempty" json:"imp,omitempty"`
	Region        string                      `xml:"region,omitempty" json:"region,omitempty"`
	Comment       interface{}                 `xml:"-" json:"-"`                           //Comment,NullType
	Bond          string                      `xml:"bond,omitempty" json:"bond,omitempty"` //Bond,EnumList:disulfide(1),thiolester(2),xlink(3),thioether(4),other(255)
	Site          string                      `xml:"site,omitempty" json:"site,omitempty"` //Site,EnumList:active(1),binding(2),cleavage(3),inhibit(4),modified(5),glycosylation(6),myristoylation(7),mutagenized(8),metal-binding(9),phosphorylation(10),acetylation(11),amidation(12),methylation(13),hydroxylation(14),sulfatation(15),oxidative-deamination(16),pyrrolidone-carboxylic-acid(17),gamma-carboxyglutamic-acid(18),blocked(19),lipid-binding(20),np-binding(21),dna-binding(22),signal-peptide(23),transit-peptide(24),transmembrane-region(25),nitrosylation(26),other(255)
	Rsite         *NCBIRsite.RsiteRef         `xml:"rsite,omitempty" json:"rsite,omitempty"`
	User          *NCBIGeneral.UserObject     `xml:"user,omitempty" json:"user,omitempty"`
	Txinit        *NCBITxInit.Txinit          `xml:"txinit,omitempty" json:"txinit,omitempty"`
	Num           *NCBISeqRef.Numbering       `xml:"num,omitempty" json:"num,omitempty"`
	PsecStr       string                      `xml:"psec-str,omitempty" json:"psec_str,omitempty"` //PsecStr,EnumList:helix(1),sheet(2),turn(3)
	NonStdResidue string                      `xml:"non-std-residue,omitempty" json:"non_std_residue,omitempty"`
	Het           NCBISeqCommon.Heterogen     `xml:"het,omitempty" json:"het,omitempty"`
	Biosrc        *NCBIBioSource.BioSource    `xml:"biosrc,omitempty" json:"biosrc,omitempty"`
	Clone         *CloneRef                   `xml:"clone,omitempty" json:"clone,omitempty"`
	Variation     *NCBIVariation.VariationRef `xml:"variation,omitempty" json:"variation,omitempty"`
}

type SeqFeatXref struct {
	Id   *NCBISeqCommon.FeatId `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Data *SeqFeatData          `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
}
type SeqFeatSupport struct {
	Experiment    []ExperimentSupport    `xml:"experiment,omitempty" json:"experiment,omitempty" asn1:"optional"`
	Inference     []InferenceSupport     `xml:"inference,omitempty" json:"inference,omitempty" asn1:"optional"`
	ModelEvidence []ModelEvidenceSupport `xml:"model-evidence,omitempty" json:"model_evidence,omitempty" asn1:"optional"`
}

//EvidenceCategory,IntegerEnum:not-set(0),coordinates(1),description(2),existence(3)
type EvidenceCategory int

type ExperimentSupport struct {
	Category    *EvidenceCategory     `xml:"category,omitempty" json:"category,omitempty" asn1:"optional"`
	Explanation string                `xml:"explanation" json:"explanation"`
	Pmids       []NCBIBiblio.PubMedId `xml:"pmids,omitempty" json:"pmids,omitempty" asn1:"optional"`
	Dois        []NCBIBiblio.DOI      `xml:"dois,omitempty" json:"dois,omitempty" asn1:"optional"`
}
type ProgramId struct {
	Name    string `xml:"name" json:"name"`
	Version string `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
}
type EvidenceBasis struct {
	Programs   []ProgramId        `xml:"programs,omitempty" json:"programs,omitempty" asn1:"optional"`
	Accessions []NCBISeqloc.SeqId `xml:"accessions,omitempty" json:"accessions,omitempty" asn1:"optional"`
}
type InferenceSupport struct {
	Category    *EvidenceCategory     `xml:"category,omitempty" json:"category,omitempty" asn1:"optional"`
	Type        int                   `xml:"type" json:"type"` //Type,IntegerEnum:not-set(0),similar-to-sequence(1),similar-to-aa(2),similar-to-dna(3),similar-to-rna(4),similar-to-mrna(5),similiar-to-est(6),similar-to-other-rna(7),profile(8),nucleotide-motif(9),protein-motif(10),ab-initio-prediction(11),alignment(12),other(255)
	OtherType   string                `xml:"other-type,omitempty" json:"other_type,omitempty" asn1:"optional"`
	SameSpecies bool                  `xml:"same-species" json:"same_species"`
	Basis       *EvidenceBasis        `xml:"basis,omitempty" json:"basis,omitempty"`
	Pmids       []NCBIBiblio.PubMedId `xml:"pmids,omitempty" json:"pmids,omitempty" asn1:"optional"`
	Dois        []NCBIBiblio.DOI      `xml:"dois,omitempty" json:"dois,omitempty" asn1:"optional"`
}
type ModelEvidenceItem struct {
	Id                   *NCBISeqloc.SeqId `xml:"id,omitempty" json:"id,omitempty"`
	ExonCount            int64             `xml:"exon-count,omitempty" json:"exon_count,omitempty" asn1:"optional"`
	ExonLength           int64             `xml:"exon-length,omitempty" json:"exon_length,omitempty" asn1:"optional"`
	FullLength           bool              `xml:"full-length" json:"full_length"`
	SupportsAllExonCombo bool              `xml:"supports-all-exon-combo" json:"supports_all_exon_combo"`
}
type ModelEvidenceSupport struct {
	Method               string              `xml:"method,omitempty" json:"method,omitempty" asn1:"optional"`
	Mrna                 []ModelEvidenceItem `xml:"mrna,omitempty" json:"mrna,omitempty" asn1:"optional"`
	Est                  []ModelEvidenceItem `xml:"est,omitempty" json:"est,omitempty" asn1:"optional"`
	Protein              []ModelEvidenceItem `xml:"protein,omitempty" json:"protein,omitempty" asn1:"optional"`
	Identification       *NCBISeqloc.SeqId   `xml:"identification,omitempty" json:"identification,omitempty" asn1:"optional"`
	Dbxref               []NCBIGeneral.Dbtag `xml:"dbxref,omitempty" json:"dbxref,omitempty" asn1:"optional"`
	ExonCount            int64               `xml:"exon-count,omitempty" json:"exon_count,omitempty" asn1:"optional"`
	ExonLength           int64               `xml:"exon-length,omitempty" json:"exon_length,omitempty" asn1:"optional"`
	FullLength           bool                `xml:"full-length" json:"full_length"`
	SupportsAllExonCombo bool                `xml:"supports-all-exon-combo" json:"supports_all_exon_combo"`
}
type Cdregion struct {
	Orf       bool        `xml:"orf,omitempty" json:"orf,omitempty" asn1:"optional"`
	Frame     string      `xml:"frame" json:"frame"` //Frame,EnumList:not-set(0),one(1),two(2),three(3)
	Conflict  bool        `xml:"conflict,omitempty" json:"conflict,omitempty" asn1:"optional"`
	Gaps      int64       `xml:"gaps,omitempty" json:"gaps,omitempty" asn1:"optional"`
	Mismatch  int64       `xml:"mismatch,omitempty" json:"mismatch,omitempty" asn1:"optional"`
	Code      GeneticCode `xml:"code,omitempty" json:"code,omitempty" asn1:"optional"`
	CodeBreak []CodeBreak `xml:"code-break,omitempty" json:"code_break,omitempty" asn1:"optional"`
	Stops     int64       `xml:"stops,omitempty" json:"stops,omitempty" asn1:"optional"`
}
type GeneticCode []struct {
	Name       string `xml:"name,omitempty" json:"name,omitempty"`
	Id         int64  `xml:"id,omitempty" json:"id,omitempty"`
	Ncbieaa    string `xml:"ncbieaa,omitempty" json:"ncbieaa,omitempty"`
	Ncbi8aa    []byte `xml:"ncbi8aa,omitempty" json:"ncbi8aa,omitempty"`
	Ncbistdaa  []byte `xml:"ncbistdaa,omitempty" json:"ncbistdaa,omitempty"`
	Sncbieaa   string `xml:"sncbieaa,omitempty" json:"sncbieaa,omitempty"`
	Sncbi8aa   []byte `xml:"sncbi8aa,omitempty" json:"sncbi8aa,omitempty"`
	Sncbistdaa []byte `xml:"sncbistdaa,omitempty" json:"sncbistdaa,omitempty"`
}
type CodeBreak struct {
	Loc *NCBISeqloc.SeqLoc `xml:"loc,omitempty" json:"loc,omitempty"`
	Aa  struct {
		Ncbieaa   int64 `xml:"ncbieaa,omitempty" json:"ncbieaa,omitempty"`
		Ncbi8aa   int64 `xml:"ncbi8aa,omitempty" json:"ncbi8aa,omitempty"`
		Ncbistdaa int64 `xml:"ncbistdaa,omitempty" json:"ncbistdaa,omitempty"`
	} `xml:"aa" json:"aa"` //Aa,ChoiceOption
}
type GeneticCodeTable []GeneticCode
type ImpFeat struct {
	Key   string `xml:"key" json:"key"`
	Loc   string `xml:"loc,omitempty" json:"loc,omitempty" asn1:"optional"`
	Descr string `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
}
type GbQual struct {
	Qual string `xml:"qual" json:"qual"`
	Val  string `xml:"val" json:"val"`
}
type CloneRef struct {
	Name            string      `xml:"name" json:"name"`
	Library         string      `xml:"library,omitempty" json:"library,omitempty" asn1:"optional"`
	Concordant      bool        `xml:"concordant" json:"concordant"`
	Unique          bool        `xml:"unique" json:"unique"`
	PlacementMethod int         `xml:"placement-method,omitempty" json:"placement_method,omitempty" asn1:"optional"` //PlacementMethod,IntegerEnum:end-seq(0),insert-alignment(1),sts(2),fish(3),fingerprint(4),end-seq-insert-alignment(5),external(253),curated(254),other(255)
	CloneSeq        CloneSeqSet `xml:"clone-seq,omitempty" json:"clone_seq,omitempty" asn1:"optional"`
}
type CloneSeqSet []CloneSeq
type CloneSeq struct {
	Type       int                `xml:"type" json:"type"`                                                 //Type,IntegerEnum:insert(0),end(1),other(255)
	Confidence int                `xml:"confidence,omitempty" json:"confidence,omitempty" asn1:"optional"` //Confidence,IntegerEnum:multiple(0),na(1),nohit-rep(2),nohitnorep(3),other-chrm(4),unique(5),virtual(6),multiple-rep(7),multiplenorep(8),no-hit(9),other(255)
	Location   *NCBISeqloc.SeqLoc `xml:"location,omitempty" json:"location,omitempty"`
	Seq        *NCBISeqloc.SeqLoc `xml:"seq,omitempty" json:"seq,omitempty" asn1:"optional"`
	AlignId    *NCBIGeneral.Dbtag `xml:"align-id,omitempty" json:"align_id,omitempty" asn1:"optional"`
	Support    int                `xml:"support,omitempty" json:"support,omitempty" asn1:"optional"` //Support,IntegerEnum:prototype(0),supporting(1),supports-other(2),non-supporting(3)
}
