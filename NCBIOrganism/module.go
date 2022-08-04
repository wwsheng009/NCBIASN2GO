package NCBIOrganism

import "ncbiasn/NCBIGeneral"

type OrgRef struct {
	Taxname string              `xml:"taxname,omitempty" json:"taxname,omitempty" asn1:"optional"`
	Common  string              `xml:"common,omitempty" json:"common,omitempty" asn1:"optional"`
	Mod     []string            `xml:"mod,omitempty" json:"mod,omitempty" asn1:"optional"`
	Db      []NCBIGeneral.Dbtag `xml:"db,omitempty" json:"db,omitempty" asn1:"optional"`
	Syn     []string            `xml:"syn,omitempty" json:"syn,omitempty" asn1:"optional"`
	Orgname *OrgName            `xml:"orgname,omitempty" json:"orgname,omitempty" asn1:"optional"`
}
type OrgName struct {
	Name struct {
		Binomial    *BinomialOrgName `xml:"binomial,omitempty" json:"binomial,omitempty"`
		Virus       string           `xml:"virus,omitempty" json:"virus,omitempty"`
		Hybrid      MultiOrgName     `xml:"hybrid,omitempty" json:"hybrid,omitempty"`
		Namedhybrid *BinomialOrgName `xml:"namedhybrid,omitempty" json:"namedhybrid,omitempty"`
		Partial     PartialOrgName   `xml:"partial,omitempty" json:"partial,omitempty"`
	} `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"` //Name,ChoiceOption
	Attrib  string   `xml:"attrib,omitempty" json:"attrib,omitempty" asn1:"optional"`
	Mod     []OrgMod `xml:"mod,omitempty" json:"mod,omitempty" asn1:"optional"`
	Lineage string   `xml:"lineage,omitempty" json:"lineage,omitempty" asn1:"optional"`
	Gcode   int64    `xml:"gcode,omitempty" json:"gcode,omitempty" asn1:"optional"`
	Mgcode  int64    `xml:"mgcode,omitempty" json:"mgcode,omitempty" asn1:"optional"`
	Div     string   `xml:"div,omitempty" json:"div,omitempty" asn1:"optional"`
	Pgcode  int64    `xml:"pgcode,omitempty" json:"pgcode,omitempty" asn1:"optional"`
}
type OrgMod struct {
	Subtype int    `xml:"subtype" json:"subtype"` //Subtype,IntegerEnum:strain(2),substrain(3),type(4),subtype(5),variety(6),serotype(7),serogroup(8),serovar(9),cultivar(10),pathovar(11),chemovar(12),biovar(13),biotype(14),group(15),subgroup(16),isolate(17),common(18),acronym(19),dosage(20),nat-host(21),sub-species(22),specimen-voucher(23),authority(24),forma(25),forma-specialis(26),ecotype(27),synonym(28),anamorph(29),teleomorph(30),breed(31),gb-acronym(32),gb-anamorph(33),gb-synonym(34),culture-collection(35),bio-material(36),metagenome-source(37),type-material(38),nomenclature(39),old-lineage(253),old-name(254),other(255)
	Subname string `xml:"subname" json:"subname"`
	Attrib  string `xml:"attrib,omitempty" json:"attrib,omitempty" asn1:"optional"`
}
type BinomialOrgName struct {
	Genus      string `xml:"genus" json:"genus"`
	Species    string `xml:"species,omitempty" json:"species,omitempty" asn1:"optional"`
	Subspecies string `xml:"subspecies,omitempty" json:"subspecies,omitempty" asn1:"optional"`
}
type MultiOrgName []OrgName
type PartialOrgName []TaxElement
type TaxElement struct {
	FixedLevel int    `xml:"fixed-level" json:"fixed_level"` //FixedLevel,IntegerEnum:other(0),family(1),order(2),class(3)
	Level      string `xml:"level,omitempty" json:"level,omitempty" asn1:"optional"`
	Name       string `xml:"name" json:"name"`
}
