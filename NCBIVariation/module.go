package NCBIVariation

import (
	"ncbiasn/NCBIBioSource"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIPub"
	"ncbiasn/NCBISeqCommon"
	"ncbiasn/NCBISeqloc"
)

type VariantProperties struct {
	Version                  int64   `xml:"version" json:"version"`
	ResourceLink             int     `xml:"resource-link,omitempty" json:"resource_link,omitempty" asn1:"optional"`                           //ResourceLink,IntegerEnum:preserved(1),provisional(2),has3D(4),submitterLinkout(8),clinical(16),genotypeKit(32)
	GeneLocation             int     `xml:"gene-location,omitempty" json:"gene_location,omitempty" asn1:"optional"`                           //GeneLocation,IntegerEnum:in-gene(1),near-gene-5(2),near-gene-3(4),intron(8),donor(16),acceptor(32),utr-5(64),utr-3(128),in-start-codon(256),in-stop-codon(512),intergenic(1024),conserved-noncoding(2048)
	Effect                   int     `xml:"effect,omitempty" json:"effect,omitempty" asn1:"optional"`                                         //Effect,IntegerEnum:no-change(0),synonymous(1),nonsense(2),missense(4),frameshift(8),up-regulator(16),down-regulator(32),methylation(64),stop-gain(128),stop-loss(256)
	Mapping                  int     `xml:"mapping,omitempty" json:"mapping,omitempty" asn1:"optional"`                                       //Mapping,IntegerEnum:has-other-snp(1),has-assembly-conflict(2),is-assembly-specific(4)
	MapWeight                int     `xml:"map-weight,omitempty" json:"map_weight,omitempty" asn1:"optional"`                                 //MapWeight,IntegerEnum:is-uniquely-placed(1),placed-twice-on-same-chrom(2),placed-twice-on-diff-chrom(3),many-placements(10)
	FrequencyBasedValidation int     `xml:"frequency-based-validation,omitempty" json:"frequency_based_validation,omitempty" asn1:"optional"` //FrequencyBasedValidation,IntegerEnum:is-mutation(1),above-5pct-all(2),above-5pct-1plus(4),validated(8),above-1pct-all(16),above-1pct-1plus(32)
	Genotype                 int     `xml:"genotype,omitempty" json:"genotype,omitempty" asn1:"optional"`                                     //Genotype,IntegerEnum:in-haplotype-set(1),has-genotypes(2)
	ProjectData              []int64 `xml:"project-data,omitempty" json:"project_data,omitempty" asn1:"optional"`
	QualityCheck             int     `xml:"quality-check,omitempty" json:"quality_check,omitempty" asn1:"optional"` //QualityCheck,IntegerEnum:contig-allele-missing(1),withdrawn-by-submitter(2),non-overlapping-alleles(4),strain-specific(8),genotype-conflict(16)
	Confidence               int     `xml:"confidence,omitempty" json:"confidence,omitempty" asn1:"optional"`       //Confidence,IntegerEnum:unknown(0),likely-artifact(1),other(255)
	OtherValidation          bool    `xml:"other-validation,omitempty" json:"other_validation,omitempty" asn1:"optional"`
	AlleleOrigin             int     `xml:"allele-origin,omitempty" json:"allele_origin,omitempty" asn1:"optional"` //AlleleOrigin,IntegerEnum:unknown(0),germline(1),somatic(2),inherited(4),paternal(8),maternal(16),de-novo(32),biparental(64),uniparental(128),not-tested(256),tested-inconclusive(512),not-reported(1024),other(1073741824)
	AlleleState              int     `xml:"allele-state,omitempty" json:"allele_state,omitempty" asn1:"optional"`   //AlleleState,IntegerEnum:unknown(0),homozygous(1),heterozygous(2),hemizygous(3),nullizygous(4),other(255)
	AlleleFrequency          float64 `xml:"allele-frequency,omitempty" json:"allele_frequency,omitempty" asn1:"optional"`
	IsAncestralAllele        bool    `xml:"is-ancestral-allele,omitempty" json:"is_ancestral_allele,omitempty" asn1:"optional"`
}
type Phenotype struct {
	Source               string              `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"`
	Term                 string              `xml:"term,omitempty" json:"term,omitempty" asn1:"optional"`
	Xref                 []NCBIGeneral.Dbtag `xml:"xref,omitempty" json:"xref,omitempty" asn1:"optional"`
	ClinicalSignificance int                 `xml:"clinical-significance,omitempty" json:"clinical_significance,omitempty" asn1:"optional"` //ClinicalSignificance,IntegerEnum:unknown(0),untested(1),non-pathogenic(2),probable-non-pathogenic(3),probable-pathogenic(4),pathogenic(5),drug-response(6),histocompatibility(7),other(255)
}
type PopulationData struct {
	Population        string                 `xml:"population" json:"population"`
	GenotypeFrequency float64                `xml:"genotype-frequency,omitempty" json:"genotype_frequency,omitempty" asn1:"optional"`
	ChromosomesTested int64                  `xml:"chromosomes-tested,omitempty" json:"chromosomes_tested,omitempty" asn1:"optional"`
	SampleIds         []NCBIGeneral.ObjectId `xml:"sample-ids,omitempty" json:"sample_ids,omitempty" asn1:"optional"`
	AlleleFrequency   float64                `xml:"allele-frequency,omitempty" json:"allele_frequency,omitempty" asn1:"optional"`
	Flags             int                    `xml:"flags,omitempty" json:"flags,omitempty" asn1:"optional"` //Flags,IntegerEnum:is-default-population(1),is-minor-allele(2),is-rare-allele(4)
}
type ExtLoc struct {
	Id       *NCBIGeneral.ObjectId `xml:"id,omitempty" json:"id,omitempty"`
	Location *NCBISeqloc.SeqLoc    `xml:"location,omitempty" json:"location,omitempty"`
}
type VariationRef struct {
	Id                *NCBIGeneral.Dbtag    `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	ParentId          *NCBIGeneral.Dbtag    `xml:"parent-id,omitempty" json:"parent_id,omitempty" asn1:"optional"`
	SampleId          *NCBIGeneral.ObjectId `xml:"sample-id,omitempty" json:"sample_id,omitempty" asn1:"optional"`
	OtherIds          []NCBIGeneral.Dbtag   `xml:"other-ids,omitempty" json:"other_ids,omitempty" asn1:"optional"`
	Name              string                `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Synonyms          []string              `xml:"synonyms,omitempty" json:"synonyms,omitempty" asn1:"optional"`
	Description       string                `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Phenotype         []Phenotype           `xml:"phenotype,omitempty" json:"phenotype,omitempty" asn1:"optional"`
	Method            []int                 `xml:"method,omitempty" json:"method,omitempty" asn1:"optional"`
	PopulationData    []PopulationData      `xml:"population-data,omitempty" json:"population_data,omitempty" asn1:"optional"`
	VariantProp       *VariantProperties    `xml:"variant-prop,omitempty" json:"variant_prop,omitempty" asn1:"optional"`
	Validated         bool                  `xml:"validated,omitempty" json:"validated,omitempty" asn1:"optional"`
	ClinicalTest      []NCBIGeneral.Dbtag   `xml:"clinical-test,omitempty" json:"clinical_test,omitempty" asn1:"optional"`
	AlleleOrigin      int                   `xml:"allele-origin,omitempty" json:"allele_origin,omitempty" asn1:"optional"` //AlleleOrigin,IntegerEnum:unknown(0),germline(1),somatic(2),inherited(4),paternal(8),maternal(16),de-novo(32),biparental(64),uniparental(128),not-tested(256),tested-inconclusive(512),other(1073741824)
	AlleleState       int                   `xml:"allele-state,omitempty" json:"allele_state,omitempty" asn1:"optional"`   //AlleleState,IntegerEnum:unknown(0),homozygous(1),heterozygous(2),hemizygous(3),nullizygous(4),other(255)
	AlleleFrequency   float64               `xml:"allele-frequency,omitempty" json:"allele_frequency,omitempty" asn1:"optional"`
	IsAncestralAllele bool                  `xml:"is-ancestral-allele,omitempty" json:"is_ancestral_allele,omitempty" asn1:"optional"`
	Pub               *NCBIPub.Pub          `xml:"pub,omitempty" json:"pub,omitempty" asn1:"optional"`
	Data              struct {
		Unknown           interface{}    `xml:"-" json:"-"` //Unknown,NullType
		Note              string         `xml:"note,omitempty" json:"note,omitempty"`
		UniparentalDisomy interface{}    `xml:"-" json:"-"` //UniparentalDisomy,NullType
		Instance          *VariationInst `xml:"instance,omitempty" json:"instance,omitempty"`
		Set               struct {
			Type       int            `xml:"type" json:"type"` //Type,IntegerEnum:unknown(0),compound(1),products(2),haplotype(3),genotype(4),mosaic(5),individual(6),population(7),alleles(8),package(9),other(255)
			Variations []VariationRef `xml:"variations,omitempty" json:"variations,omitempty"`
			Name       string         `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
		} `xml:"set,omitempty" json:"set,omitempty"`
		Complex interface{} `xml:"-" json:"-"` //Complex,NullType
	} `xml:"data" json:"data"` //Data,ChoiceOption
	Consequence []struct {
		Unknown    interface{}   `xml:"-" json:"-"` //Unknown,NullType
		Splicing   interface{}   `xml:"-" json:"-"` //Splicing,NullType
		Note       string        `xml:"note,omitempty" json:"note,omitempty"`
		Variation  *VariationRef `xml:"variation,omitempty" json:"variation,omitempty"`
		Frameshift struct {
			Phase   int64 `xml:"phase,omitempty" json:"phase,omitempty" asn1:"optional"`
			XLength int64 `xml:"x-length,omitempty" json:"x_length,omitempty" asn1:"optional"`
		} `xml:"frameshift,omitempty" json:"frameshift,omitempty"`
		LossOfHeterozygosity struct {
			Reference string `xml:"reference,omitempty" json:"reference,omitempty" asn1:"optional"`
			Test      string `xml:"test,omitempty" json:"test,omitempty" asn1:"optional"`
		} `xml:"loss-of-heterozygosity,omitempty" json:"loss_of_heterozygosity,omitempty"`
	} `xml:"consequence,omitempty" json:"consequence,omitempty" asn1:"optional"`
	Location      *NCBISeqloc.SeqLoc      `xml:"location,omitempty" json:"location,omitempty" asn1:"optional"`
	ExtLocs       []ExtLoc                `xml:"ext-locs,omitempty" json:"ext_locs,omitempty" asn1:"optional"`
	Ext           *NCBIGeneral.UserObject `xml:"ext,omitempty" json:"ext,omitempty" asn1:"optional"`
	SomaticOrigin []struct {
		Source    *NCBIBioSource.SubSource `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"`
		Condition struct {
			Description string              `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
			ObjectId    []NCBIGeneral.Dbtag `xml:"object-id,omitempty" json:"object_id,omitempty" asn1:"optional"`
		} `xml:"condition,omitempty" json:"condition,omitempty" asn1:"optional"`
	} `xml:"somatic-origin,omitempty" json:"somatic_origin,omitempty" asn1:"optional"`
}
type DeltaItem struct {
	Seq struct {
		Literal *NCBISeqCommon.SeqLiteral `xml:"literal,omitempty" json:"literal,omitempty"`
		Loc     *NCBISeqloc.SeqLoc        `xml:"loc,omitempty" json:"loc,omitempty"`
		This    interface{}               `xml:"-" json:"-"` //This,NullType
	} `xml:"seq,omitempty" json:"seq,omitempty" asn1:"optional"` //Seq,ChoiceOption
	Multiplier     int64                `xml:"multiplier,omitempty" json:"multiplier,omitempty" asn1:"optional"`
	MultiplierFuzz *NCBIGeneral.IntFuzz `xml:"multiplier-fuzz,omitempty" json:"multiplier_fuzz,omitempty" asn1:"optional"`
	Action         int                  `xml:"action" json:"action"` //Action,IntegerEnum:morph(0),offset(1),del-at(2),ins-before(3)
}
type VariationInst struct {
	Type        int         `xml:"type" json:"type"` //Type,IntegerEnum:unknown(0),identity(1),inv(2),snv(3),mnp(4),delins(5),del(6),ins(7),microsatellite(8),transposon(9),cnv(10),direct-copy(11),rev-direct-copy(12),inverted-copy(13),everted-copy(14),translocation(15),prot-missense(16),prot-nonsense(17),prot-neutral(18),prot-silent(19),prot-other(20),other(255)
	Delta       []DeltaItem `xml:"delta,omitempty" json:"delta,omitempty"`
	Observation int         `xml:"observation,omitempty" json:"observation,omitempty" asn1:"optional"` //Observation,IntegerEnum:asserted(1),reference(2),variant(4)
}
