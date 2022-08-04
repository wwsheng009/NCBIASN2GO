package NCBIPCAssay

import (
	"ncbiasn/NCBIBioSource"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIPCSubstance"
	"ncbiasn/NCBIPub"
)

type PCAssayContainer []PCAssaySubmit
type PCAssaySubmit struct {
	Assay struct {
		Aid       int64                     `xml:"aid,omitempty" json:"aid,omitempty"`
		AidSource *NCBIPCSubstance.PCSource `xml:"aid-source,omitempty" json:"aid_source,omitempty"`
		Descr     *PCAssayDescription       `xml:"descr,omitempty" json:"descr,omitempty"`
		Aidver    *NCBIPCSubstance.PC_ID    `xml:"aidver,omitempty" json:"aidver,omitempty"`
	} `xml:"assay" json:"assay"` //Assay,ChoiceOption
	Data   []PCAssayResults `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
	Revoke []int64          `xml:"revoke,omitempty" json:"revoke,omitempty" asn1:"optional"`
}
type PCAssayResults struct {
	Sid       int64                     `xml:"sid" json:"sid"`
	SidSource *NCBIPCSubstance.PCSource `xml:"sid-source,omitempty" json:"sid_source,omitempty" asn1:"optional"`
	Version   int64                     `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
	Comment   string                    `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Outcome   int                       `xml:"outcome" json:"outcome"` //Outcome,IntegerEnum:inactive(1),active(2),inconclusive(3),unspecified(4),probe(5)
	Rank      int64                     `xml:"rank,omitempty" json:"rank,omitempty" asn1:"optional"`
	Data      []PCAssayData             `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
	Url       string                    `xml:"url,omitempty" json:"url,omitempty" asn1:"optional"`
	Xref      []PCAnnotatedXRef         `xml:"xref,omitempty" json:"xref,omitempty" asn1:"optional"`
	Date      *NCBIGeneral.Date         `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
}
type PCAssayData struct {
	Tid   int64 `xml:"tid" json:"tid"`
	Value struct {
		Ival int64   `xml:"ival,omitempty" json:"ival,omitempty"`
		Fval float64 `xml:"fval,omitempty" json:"fval,omitempty"`
		Bval bool    `xml:"bval,omitempty" json:"bval,omitempty"`
		Sval string  `xml:"sval,omitempty" json:"sval,omitempty"`
	} `xml:"value" json:"value"` //Value,ChoiceOption
}
type PCAssayDescription struct {
	Aid                   *NCBIPCSubstance.PC_ID    `xml:"aid,omitempty" json:"aid,omitempty"`
	AidSource             *NCBIPCSubstance.PCSource `xml:"aid-source,omitempty" json:"aid_source,omitempty" asn1:"optional"`
	Name                  string                    `xml:"name" json:"name"`
	Description           []string                  `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Protocol              []string                  `xml:"protocol,omitempty" json:"protocol,omitempty" asn1:"optional"`
	Comment               []string                  `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Xref                  []PCAnnotatedXRef         `xml:"xref,omitempty" json:"xref,omitempty" asn1:"optional"`
	Results               []PCResultType            `xml:"results,omitempty" json:"results,omitempty" asn1:"optional"`
	Pub                   []NCBIPub.Pub             `xml:"pub,omitempty" json:"pub,omitempty" asn1:"optional"`
	Revision              int64                     `xml:"revision,omitempty" json:"revision,omitempty" asn1:"optional"`
	Target                []PCAssayTargetInfo       `xml:"target,omitempty" json:"target,omitempty" asn1:"optional"`
	ActivityOutcomeMethod int                       `xml:"activity-outcome-method,omitempty" json:"activity_outcome_method,omitempty" asn1:"optional"` //ActivityOutcomeMethod,IntegerEnum:other(0),screening(1),confirmatory(2),summary(3)
	Dr                    []PCAssayDRAttr           `xml:"dr,omitempty" json:"dr,omitempty" asn1:"optional"`
	SubstanceType         int                       `xml:"substance-type,omitempty" json:"substance_type,omitempty" asn1:"optional"` //SubstanceType,IntegerEnum:small-molecule(1),nucleotide(2),other(255)
	GrantNumber           []string                  `xml:"grant-number,omitempty" json:"grant_number,omitempty" asn1:"optional"`
	ProjectCategory       int                       `xml:"project-category,omitempty" json:"project_category,omitempty" asn1:"optional"` //ProjectCategory,IntegerEnum:mlscn(1),mlpcn(2),mlscn-ap(3),mlpcn-ap(4),journal-article(5),assay-vendor(6),literature-extracted(7),literature-author(8),literature-publisher(9),rnaigi(10),other(255)
	IsPanel               bool                      `xml:"is-panel,omitempty" json:"is_panel,omitempty" asn1:"optional"`
	AssayGroup            []string                  `xml:"assay-group,omitempty" json:"assay_group,omitempty" asn1:"optional"`
	PanelInfo             *PCAssayPanel             `xml:"panel-info,omitempty" json:"panel_info,omitempty" asn1:"optional"`
	IsMlpLateStage        bool                      `xml:"is-mlp-late-stage,omitempty" json:"is_mlp_late_stage,omitempty" asn1:"optional"`
	CategorizedComment    []PCCategorizedComment    `xml:"categorized-comment,omitempty" json:"categorized_comment,omitempty" asn1:"optional"`
}
type PCAssayPanel struct {
	Name   string               `xml:"name" json:"name"`
	Descr  string               `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Member []PCAssayPanelMember `xml:"member,omitempty" json:"member,omitempty" asn1:"optional"`
}
type PCAssayPanelMember struct {
	Mid                   int64                  `xml:"mid" json:"mid"`
	Name                  string                 `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Description           string                 `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Protocol              []string               `xml:"protocol,omitempty" json:"protocol,omitempty" asn1:"optional"`
	Comment               []string               `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Target                []PCAssayTargetInfo    `xml:"target,omitempty" json:"target,omitempty" asn1:"optional"`
	Xref                  []PCAnnotatedXRef      `xml:"xref,omitempty" json:"xref,omitempty" asn1:"optional"`
	ActivityOutcomeMethod int                    `xml:"activity-outcome-method,omitempty" json:"activity_outcome_method,omitempty" asn1:"optional"` //ActivityOutcomeMethod,IntegerEnum:other(0),screening(1),confirmatory(2),summary(3)
	Dr                    []PCAssayDRAttr        `xml:"dr,omitempty" json:"dr,omitempty" asn1:"optional"`
	CategorizedComment    []PCCategorizedComment `xml:"categorized-comment,omitempty" json:"categorized_comment,omitempty" asn1:"optional"`
}
type PCCategorizedComment struct {
	Title   string   `xml:"title" json:"title"`
	Comment []string `xml:"comment,omitempty" json:"comment,omitempty"`
}
type PCAssayDRAttr struct {
	Id    int64  `xml:"id" json:"id"`
	Descr string `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Dn    string `xml:"dn,omitempty" json:"dn,omitempty" asn1:"optional"`
	Rn    string `xml:"rn,omitempty" json:"rn,omitempty" asn1:"optional"`
	Type  int    `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"` //Type,IntegerEnum:experimental(0),calculated(1)
}
type PCAssayTargetInfo struct {
	Name         string                   `xml:"name" json:"name"`
	MolId        int64                    `xml:"mol-id" json:"mol_id"`
	MoleculeType int                      `xml:"molecule-type" json:"molecule_type"` //MoleculeType,IntegerEnum:protein(1),dna(2),rna(3),gene(4),biosystem(5),other(255)
	Organism     *NCBIBioSource.BioSource `xml:"organism,omitempty" json:"organism,omitempty" asn1:"optional"`
	Descr        string                   `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Comment      []string                 `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
}
type PCAnnotatedXRef struct {
	Xref    *NCBIPCSubstance.PCXRefData `xml:"xref,omitempty" json:"xref,omitempty"`
	Comment string                      `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Type    int                         `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"` //Type,IntegerEnum:pcit(1),pgene(2)
}
type PCResultType struct {
	Tid         int64    `xml:"tid" json:"tid"`
	Name        string   `xml:"name" json:"name"`
	Description []string `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Type        int      `xml:"type" json:"type"` //Type,IntegerEnum:float(1),int(2),bool(3),string(4)
	Constraints struct {
		Fset   []float64        `xml:"fset,omitempty" json:"fset,omitempty"`
		Fmin   float64          `xml:"fmin,omitempty" json:"fmin,omitempty"`
		Fmax   float64          `xml:"fmax,omitempty" json:"fmax,omitempty"`
		Frange *PCRealMinMax    `xml:"frange,omitempty" json:"frange,omitempty"`
		Iset   []int64          `xml:"iset,omitempty" json:"iset,omitempty"`
		Imin   int64            `xml:"imin,omitempty" json:"imin,omitempty"`
		Imax   int64            `xml:"imax,omitempty" json:"imax,omitempty"`
		Irange *PCIntegerMinMax `xml:"irange,omitempty" json:"irange,omitempty"`
		Sset   []string         `xml:"sset,omitempty" json:"sset,omitempty"`
	} `xml:"constraints,omitempty" json:"constraints,omitempty" asn1:"optional"` //Constraints,ChoiceOption
	Unit       int                     `xml:"unit,omitempty" json:"unit,omitempty" asn1:"optional"` //Unit,IntegerEnum:ppt(1),ppm(2),ppb(3),mm(4),um(5),nm(6),pm(7),fm(8),mgml(9),ugml(10),ngml(11),pgml(12),fgml(13),m(14),percent(15),ratio(16),sec(17),rsec(18),min(19),rmin(20),day(21),rday(22),ml-min-kg(23),l-kg(24),hr-ng-ml(25),cm-sec(26),mg-kg(27),none(254),unspecified(255)
	Sunit      string                  `xml:"sunit,omitempty" json:"sunit,omitempty" asn1:"optional"`
	Transform  int                     `xml:"transform,omitempty" json:"transform,omitempty" asn1:"optional"` //Transform,IntegerEnum:linear(1),ln(2),log(3),reciprocal(4),negative(5),nlog(6),nln(7)
	Stransform string                  `xml:"stransform,omitempty" json:"stransform,omitempty" asn1:"optional"`
	Tc         *PCConcentrationAttr    `xml:"tc,omitempty" json:"tc,omitempty" asn1:"optional"`
	Ac         bool                    `xml:"ac,omitempty" json:"ac,omitempty" asn1:"optional"`
	PanelInfo  *PCAssayPanelTestResult `xml:"panel-info,omitempty" json:"panel_info,omitempty" asn1:"optional"`
	Annot      int                     `xml:"annot,omitempty" json:"annot,omitempty" asn1:"optional"` //Annot,IntegerEnum:pmid(1),mmdb(2),url(3),protein-gi(4),nucleotide-gi(5),taxonomy(6),mim(7),gene(8),probe(9),aid(10),sid(11),cid(12),protein-target-gi(13),biosystems-target-id(14),target-name(15),target-descr(16),target-tax-id(17),gene-target-id(18),dna-nucleotide-target-gi(19),rna-nucleotide-target-gi(20)
}
type PCAssayPanelTestResult struct {
	Mid          int64 `xml:"mid" json:"mid"`
	ReadoutAnnot int   `xml:"readout-annot" json:"readout_annot"` //ReadoutAnnot,IntegerEnum:regular(1),outcome(2),score(3),ac(4)
}
type PCConcentrationAttr struct {
	Concentration float64 `xml:"concentration" json:"concentration"`
	Unit          int     `xml:"unit" json:"unit"` //Unit,IntegerEnum:um(5)
	DrId          int64   `xml:"dr-id,omitempty" json:"dr_id,omitempty" asn1:"optional"`
}
type PCIntegerMinMax struct {
	Min int64 `xml:"min" json:"min"`
	Max int64 `xml:"max" json:"max"`
}
type PCRealMinMax struct {
	Min float64 `xml:"min" json:"min"`
	Max float64 `xml:"max" json:"max"`
}
