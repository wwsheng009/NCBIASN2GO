package NCBIBioSource

import "ncbiasn/NCBIOrganism"

type BioSource struct {
	Genome     int                  `xml:"genome" json:"genome"` //Genome,IntegerEnum:unknown(0),genomic(1),chloroplast(2),chromoplast(3),kinetoplast(4),mitochondrion(5),plastid(6),macronuclear(7),extrachrom(8),plasmid(9),transposon(10),insertion-seq(11),cyanelle(12),proviral(13),virion(14),nucleomorph(15),apicoplast(16),leucoplast(17),proplastid(18),endogenous-virus(19),hydrogenosome(20),chromosome(21),chromatophore(22),plasmid-in-mitochondrion(23),plasmid-in-plastid(24)
	Origin     int                  `xml:"origin" json:"origin"` //Origin,IntegerEnum:unknown(0),natural(1),natmut(2),mut(3),artificial(4),synthetic(5),other(255)
	Org        *NCBIOrganism.OrgRef `xml:"org,omitempty" json:"org,omitempty"`
	Subtype    []SubSource          `xml:"subtype,omitempty" json:"subtype,omitempty" asn1:"optional"`
	IsFocus    interface{}          `xml:"-" json:"-" asn1:"optional"` //IsFocus,NullType
	PcrPrimers PCRReactionSet       `xml:"pcr-primers,omitempty" json:"pcr_primers,omitempty" asn1:"optional"`
}
type PCRReactionSet []PCRReaction
type PCRReaction struct {
	Forward PCRPrimerSet `xml:"forward,omitempty" json:"forward,omitempty" asn1:"optional"`
	Reverse PCRPrimerSet `xml:"reverse,omitempty" json:"reverse,omitempty" asn1:"optional"`
}
type PCRPrimerSet []PCRPrimer
type PCRPrimer struct {
	Seq  PCRPrimerSeq  `xml:"seq,omitempty" json:"seq,omitempty" asn1:"optional"`
	Name PCRPrimerName `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
}
type PCRPrimerSeq string
type PCRPrimerName string
type SubSource struct {
	Subtype int    `xml:"subtype" json:"subtype"` //Subtype,IntegerEnum:chromosome(1),map(2),clone(3),subclone(4),haplotype(5),genotype(6),sex(7),cell-line(8),cell-type(9),tissue-type(10),clone-lib(11),dev-stage(12),frequency(13),germline(14),rearranged(15),lab-host(16),pop-variant(17),tissue-lib(18),plasmid-name(19),transposon-name(20),insertion-seq-name(21),plastid-name(22),country(23),segment(24),endogenous-virus-name(25),transgenic(26),environmental-sample(27),isolation-source(28),lat-lon(29),collection-date(30),collected-by(31),identified-by(32),fwd-primer-seq(33),rev-primer-seq(34),fwd-primer-name(35),rev-primer-name(36),metagenomic(37),mating-type(38),linkage-group(39),haplogroup(40),whole-replicon(41),phenotype(42),altitude(43),other(255)
	Name    string `xml:"name" json:"name"`
	Attrib  string `xml:"attrib,omitempty" json:"attrib,omitempty" asn1:"optional"`
}
