package NCBIPCSubstance

import (
	"encoding/asn1"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIPub"
)

type PCSubstance struct {
	Sid      *PC_ID        `xml:"sid,omitempty" json:"sid,omitempty"`
	Source   *PCSource     `xml:"source,omitempty" json:"source,omitempty"`
	Pub      []NCBIPub.Pub `xml:"pub,omitempty" json:"pub,omitempty" asn1:"optional"`
	Synonyms []string      `xml:"synonyms,omitempty" json:"synonyms,omitempty" asn1:"optional"`
	Comment  []string      `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Xref     []PCXRefData  `xml:"xref,omitempty" json:"xref,omitempty" asn1:"optional"`
	Compound PCCompounds   `xml:"compound,omitempty" json:"compound,omitempty" asn1:"optional"`
}
type PCSubstances []PCSubstance
type PC_ID struct {
	Id      int64 `xml:"id" json:"id"`
	Version int64 `xml:"version" json:"version"`
}

//PCSource,ChoiceOption
type PCSource struct {
	Individual *NCBIPub.Pub  `xml:"individual,omitempty" json:"individual,omitempty"`
	Db         *PCDBTracking `xml:"db,omitempty" json:"db,omitempty"`
	Mmdb       *PCMMDBSource `xml:"mmdb,omitempty" json:"mmdb,omitempty"`
}

type PCDBTracking struct {
	Name        string                `xml:"name" json:"name"`
	SourceId    *NCBIGeneral.ObjectId `xml:"source-id,omitempty" json:"source_id,omitempty"`
	Date        *NCBIGeneral.Date     `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
	Description string                `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Pub         *NCBIPub.Pub          `xml:"pub,omitempty" json:"pub,omitempty" asn1:"optional"`
}
type PCMMDBSource struct {
	MmdbId       int64    `xml:"mmdb-id" json:"mmdb_id"`
	MoleculeId   int64    `xml:"molecule-id" json:"molecule_id"`
	MoleculeName []string `xml:"molecule-name,omitempty" json:"molecule_name,omitempty"`
	ResidueId    int64    `xml:"residue-id,omitempty" json:"residue_id,omitempty" asn1:"optional"`
	ResidueName  string   `xml:"residue-name,omitempty" json:"residue_name,omitempty" asn1:"optional"`
	AtomId       int64    `xml:"atom-id,omitempty" json:"atom_id,omitempty" asn1:"optional"`
	AtomName     string   `xml:"atom-name,omitempty" json:"atom_name,omitempty" asn1:"optional"`
}
type PCXRefData struct {
	Regid               string `xml:"regid,omitempty" json:"regid,omitempty"`
	Rn                  string `xml:"rn,omitempty" json:"rn,omitempty"`
	Mesh                string `xml:"mesh,omitempty" json:"mesh,omitempty"`
	Pmid                int64  `xml:"pmid,omitempty" json:"pmid,omitempty"`
	Gi                  int64  `xml:"gi,omitempty" json:"gi,omitempty"`
	Mmdb                int64  `xml:"mmdb,omitempty" json:"mmdb,omitempty"`
	Sid                 int64  `xml:"sid,omitempty" json:"sid,omitempty"`
	Cid                 int64  `xml:"cid,omitempty" json:"cid,omitempty"`
	Dburl               string `xml:"dburl,omitempty" json:"dburl,omitempty"`
	Sburl               string `xml:"sburl,omitempty" json:"sburl,omitempty"`
	Asurl               string `xml:"asurl,omitempty" json:"asurl,omitempty"`
	ProteinGi           int64  `xml:"protein-gi,omitempty" json:"protein_gi,omitempty"`
	NucleotideGi        int64  `xml:"nucleotide-gi,omitempty" json:"nucleotide_gi,omitempty"`
	Taxonomy            int64  `xml:"taxonomy,omitempty" json:"taxonomy,omitempty"`
	Aid                 int64  `xml:"aid,omitempty" json:"aid,omitempty"`
	Mim                 int64  `xml:"mim,omitempty" json:"mim,omitempty"`
	Gene                int64  `xml:"gene,omitempty" json:"gene,omitempty"`
	Probe               int64  `xml:"probe,omitempty" json:"probe,omitempty"`
	Biosystem           int64  `xml:"biosystem,omitempty" json:"biosystem,omitempty"`
	Geogse              int64  `xml:"geogse,omitempty" json:"geogse,omitempty"`
	Geogsm              int64  `xml:"geogsm,omitempty" json:"geogsm,omitempty"`
	Patent              string `xml:"patent,omitempty" json:"patent,omitempty"`
	ProteinAccession    string `xml:"protein-accession,omitempty" json:"protein_accession,omitempty"`
	NucleotideAccession string `xml:"nucleotide-accession,omitempty" json:"nucleotide_accession,omitempty"`
	Doi                 string `xml:"doi,omitempty" json:"doi,omitempty"`
	Citation            string `xml:"citation,omitempty" json:"citation,omitempty"`
}

//PCXRefData,ChoiceOption
type PCCompound struct {
	Id           *PCCompoundType  `xml:"id,omitempty" json:"id,omitempty"`
	Atoms        *PCAtoms         `xml:"atoms,omitempty" json:"atoms,omitempty" asn1:"optional"`
	Bonds        *PCBonds         `xml:"bonds,omitempty" json:"bonds,omitempty" asn1:"optional"`
	Stereo       []PCStereoCenter `xml:"stereo,omitempty" json:"stereo,omitempty" asn1:"optional"`
	Coords       []PCCoordinates  `xml:"coords,omitempty" json:"coords,omitempty" asn1:"optional"`
	Charge       int64            `xml:"charge,omitempty" json:"charge,omitempty" asn1:"optional"`
	Props        []PCInfoData     `xml:"props,omitempty" json:"props,omitempty" asn1:"optional"`
	Stereogroups []PCStereoGroup  `xml:"stereogroups,omitempty" json:"stereogroups,omitempty" asn1:"optional"`
	Count        *PCCount         `xml:"count,omitempty" json:"count,omitempty" asn1:"optional"`
	Vbalt        PCCompounds      `xml:"vbalt,omitempty" json:"vbalt,omitempty" asn1:"optional"`
	Groups       []PCGroup        `xml:"groups,omitempty" json:"groups,omitempty" asn1:"optional"`
}
type PCCompounds []PCCompound
type PCCompoundType struct {
	Type int `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"` //Type,IntegerEnum:deposited(0),standardized(1),component(2),neutralized(3),mixture(4),tautomer(5),pka-state(6),unknown(255)
	Id   struct {
		Cid int64 `xml:"cid,omitempty" json:"cid,omitempty"`
		Sid int64 `xml:"sid,omitempty" json:"sid,omitempty"`
		Xid int64 `xml:"xid,omitempty" json:"xid,omitempty"`
	} `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"` //Id,ChoiceOption
}
type PCGroup struct {
	Atoms        []int64 `xml:"atoms,omitempty" json:"atoms,omitempty"`
	Type         int     `xml:"type" json:"type"`                                                     //Type,IntegerEnum:sup(1),mul(2),sru(3),mon(4),mer(5),cop(6),cro(7),mod(8),gra(9),com(10),mix(11),for(12),dat(13),any(14),gen(15),unknown(255)
	Subtype      int     `xml:"subtype,omitempty" json:"subtype,omitempty" asn1:"optional"`           //Subtype,IntegerEnum:alt(1),ran(2),blo(3),unknown(255)
	Connectivity int     `xml:"connectivity,omitempty" json:"connectivity,omitempty" asn1:"optional"` //Connectivity,IntegerEnum:hh(1),ht(2),eu(3),unknown(255)
	Label        int64   `xml:"label,omitempty" json:"label,omitempty" asn1:"optional"`
	Subscript    string  `xml:"subscript,omitempty" json:"subscript,omitempty" asn1:"optional"`
	RepeatCount  struct {
		Exact int64 `xml:"exact,omitempty" json:"exact,omitempty"`
		Range struct {
			Lower int64 `xml:"lower" json:"lower"`
			Upper int64 `xml:"upper" json:"upper"`
		} `xml:"range,omitempty" json:"range,omitempty"`
	} `xml:"repeat-count,omitempty" json:"repeat_count,omitempty" asn1:"optional"` //RepeatCount,ChoiceOption
	Bonds struct {
		From []int64 `xml:"from,omitempty" json:"from,omitempty"`
		To   []int64 `xml:"to,omitempty" json:"to,omitempty"`
	} `xml:"bonds,omitempty" json:"bonds,omitempty" asn1:"optional"`
	Brackets struct {
		Left  *PCBracket `xml:"left,omitempty" json:"left,omitempty"`
		Right *PCBracket `xml:"right,omitempty" json:"right,omitempty"`
	} `xml:"brackets,omitempty" json:"brackets,omitempty" asn1:"optional"`
}
type PCBracket struct {
	X1 float64 `xml:"x1" json:"x1"`
	Y1 float64 `xml:"y1" json:"y1"`
	X2 float64 `xml:"x2" json:"x2"`
	Y2 float64 `xml:"y2" json:"y2"`
}
type PCCount struct {
	HeavyAtom       int64 `xml:"heavy-atom" json:"heavy_atom"`
	AtomChiral      int64 `xml:"atom-chiral" json:"atom_chiral"`
	AtomChiralDef   int64 `xml:"atom-chiral-def" json:"atom_chiral_def"`
	AtomChiralUndef int64 `xml:"atom-chiral-undef" json:"atom_chiral_undef"`
	BondChiral      int64 `xml:"bond-chiral" json:"bond_chiral"`
	BondChiralDef   int64 `xml:"bond-chiral-def" json:"bond_chiral_def"`
	BondChiralUndef int64 `xml:"bond-chiral-undef" json:"bond_chiral_undef"`
	IsotopeAtom     int64 `xml:"isotope-atom" json:"isotope_atom"`
	CovalentUnit    int64 `xml:"covalent-unit" json:"covalent_unit"`
	Tautomers       int64 `xml:"tautomers" json:"tautomers"`
}
type PCStereoGroup struct {
	Type int     `xml:"type" json:"type"` //Type,IntegerEnum:absolute(1),or(2),and(3),unknown(255)
	Aid  []int64 `xml:"aid,omitempty" json:"aid,omitempty"`
}
type PCInfoData struct {
	Urn   *PCUrn `xml:"urn,omitempty" json:"urn,omitempty"`
	Value struct {
		Bval    bool              `xml:"bval,omitempty" json:"bval,omitempty"`
		Bvec    []bool            `xml:"bvec,omitempty" json:"bvec,omitempty"`
		Ival    int64             `xml:"ival,omitempty" json:"ival,omitempty"`
		Ivec    []int64           `xml:"ivec,omitempty" json:"ivec,omitempty"`
		Fval    float64           `xml:"fval,omitempty" json:"fval,omitempty"`
		Fvec    []float64         `xml:"fvec,omitempty" json:"fvec,omitempty"`
		Sval    string            `xml:"sval,omitempty" json:"sval,omitempty"`
		Slist   []string          `xml:"slist,omitempty" json:"slist,omitempty"`
		Date    *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty"`
		Binary  []byte            `xml:"binary,omitempty" json:"binary,omitempty"`
		Bitlist asn1.BitString    `xml:"bitlist,omitempty" json:"bitlist,omitempty"`
	} `xml:"value" json:"value"` //Value,ChoiceOption
}
type PCUrn struct {
	Label          string         `xml:"label" json:"label"`
	Name           string         `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Datatype       *PCUrnDataType `xml:"datatype,omitempty" json:"datatype,omitempty" asn1:"optional"`
	Parameters     string         `xml:"parameters,omitempty" json:"parameters,omitempty" asn1:"optional"`
	Implementation string         `xml:"implementation,omitempty" json:"implementation,omitempty" asn1:"optional"`
	Version        string         `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
	Software       string         `xml:"software,omitempty" json:"software,omitempty" asn1:"optional"`
	Source         string         `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"`
	Release        string         `xml:"release,omitempty" json:"release,omitempty" asn1:"optional"`
}
type PCUrnDataType int

//PCUrnDataType,IntegerEnum:string(1),stringlist(2),int(3),intvec(4),uint(5),uintvec(6),double(7),doublevec(8),bool(9),boolvec(10),uint64(11),binary(12),url(13),unicode(14),date(15),fingerprint(16),unknown(255)
type PCCoordinates struct {
	Type       []PCCoordinateType `xml:"type,omitempty" json:"type,omitempty"`
	Aid        []int64            `xml:"aid,omitempty" json:"aid,omitempty"`
	Conformers []PCConformer      `xml:"conformers,omitempty" json:"conformers,omitempty" asn1:"optional"`
	Atomlabels []PCAtomString     `xml:"atomlabels,omitempty" json:"atomlabels,omitempty" asn1:"optional"`
	Data       []PCInfoData       `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
}
type PCConformer struct {
	X     []float64          `xml:"x,omitempty" json:"x,omitempty"`
	Y     []float64          `xml:"y,omitempty" json:"y,omitempty"`
	Z     []float64          `xml:"z,omitempty" json:"z,omitempty" asn1:"optional"`
	Style *PCDrawAnnotations `xml:"style,omitempty" json:"style,omitempty" asn1:"optional"`
	Data  []PCInfoData       `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
}
type PCConformers []PCConformer
type PCCoordinateType int

//PCCoordinateType,IntegerEnum:twod(1),threed(2),submitted(3),experimental(4),computed(5),standardized(6),augmented(7),aligned(8),compact(9),units-angstroms(10),units-nanometers(11),units-pixel(12),units-points(13),units-stdbonds(14),units-unknown(255)
type PCDrawAnnotations struct {
	Annotation []PCBondAnnotation `xml:"annotation,omitempty" json:"annotation,omitempty"`
	Aid1       []int64            `xml:"aid1,omitempty" json:"aid1,omitempty"`
	Aid2       []int64            `xml:"aid2,omitempty" json:"aid2,omitempty"`
}
type PCBondAnnotation int

//PCBondAnnotation,IntegerEnum:crossed(1),dashed(2),wavy(3),dotted(4),wedge-up(5),wedge-down(6),arrow(7),aromatic(8),resonance(9),bold(10),fischer(11),closeContact(12),unknown(255)
type PCAtoms struct {
	Aid     []int64         `xml:"aid,omitempty" json:"aid,omitempty"`
	Element []PCElement     `xml:"element,omitempty" json:"element,omitempty"`
	Label   []PCAtomString  `xml:"label,omitempty" json:"label,omitempty" asn1:"optional"`
	Isotope []PCAtomInt     `xml:"isotope,omitempty" json:"isotope,omitempty" asn1:"optional"`
	Charge  []PCAtomInt     `xml:"charge,omitempty" json:"charge,omitempty" asn1:"optional"`
	Radical []PCAtomRadical `xml:"radical,omitempty" json:"radical,omitempty" asn1:"optional"`
	Source  []PCAtomSource  `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"`
	Comment []PCAtomString  `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
}
type PCAtomSource struct {
	Aid    int64         `xml:"aid" json:"aid"`
	Source *PCMMDBSource `xml:"source,omitempty" json:"source,omitempty"`
}
type PCAtomInt struct {
	Aid   int64 `xml:"aid" json:"aid"`
	Value int64 `xml:"value" json:"value"`
}
type PCAtomString struct {
	Aid   int64  `xml:"aid" json:"aid"`
	Value string `xml:"value" json:"value"`
}
type PCAtomRadical struct {
	Aid  int64 `xml:"aid" json:"aid"`
	Type int   `xml:"type" json:"type"` //Type,IntegerEnum:singlet(1),doublet(2),triplet(3),quartet(4),quintet(5),hextet(6),heptet(7),octet(8),none(255)
}
type PCElement int

//PCElement,IntegerEnum:a(255),d(254),r(253),lp(252),h(1),he(2),li(3),be(4),b(5),c(6),n(7),o(8),f(9),ne(10),na(11),mg(12),al(13),si(14),p(15),s(16),cl(17),ar(18),k(19),ca(20),sc(21),ti(22),v(23),cr(24),mn(25),fe(26),co(27),ni(28),cu(29),zn(30),ga(31),ge(32),as(33),se(34),br(35),kr(36),rb(37),sr(38),y(39),zr(40),nb(41),mo(42),tc(43),ru(44),rh(45),pd(46),ag(47),cd(48),in(49),sn(50),sb(51),te(52),i(53),xe(54),cs(55),ba(56),la(57),ce(58),pr(59),nd(60),pm(61),sm(62),eu(63),gd(64),tb(65),dy(66),ho(67),er(68),tm(69),yb(70),lu(71),hf(72),ta(73),w(74),re(75),os(76),ir(77),pt(78),au(79),hg(80),tl(81),pb(82),bi(83),po(84),at(85),rn(86),fr(87),ra(88),ac(89),th(90),pa(91),u(92),np(93),pu(94),am(95),cm(96),bk(97),cf(98),es(99),fm(100),md(101),no(102),lr(103),rf(104),db(105),sg(106),bh(107),hs(108),mt(109),ds(110),rg(111),cn(112),nh(113),fl(114),mc(115),lv(116),ts(117),og(118)
type PCBonds struct {
	Aid1  []int64      `xml:"aid1,omitempty" json:"aid1,omitempty"`
	Aid2  []int64      `xml:"aid2,omitempty" json:"aid2,omitempty"`
	Order []PCBondType `xml:"order,omitempty" json:"order,omitempty"`
}
type PCBondType int

//PCBondType,IntegerEnum:single(1),double(2),triple(3),quadruple(4),dative(5),complex(6),ionic(7),unknown(255)
type PCStereoCenter struct {
	Tetrahedral  *PCStereoTetrahedral         `xml:"tetrahedral,omitempty" json:"tetrahedral,omitempty"`
	Planar       *PCStereoPlanar              `xml:"planar,omitempty" json:"planar,omitempty"`
	Squareplanar *PCStereoSquarePlanar        `xml:"squareplanar,omitempty" json:"squareplanar,omitempty"`
	Octahedral   *PCStereoOctahedral          `xml:"octahedral,omitempty" json:"octahedral,omitempty"`
	Bipyramid    *PCStereoTrigonalBiPyramid   `xml:"bipyramid,omitempty" json:"bipyramid,omitempty"`
	Tshape       *PCStereoTShape              `xml:"tshape,omitempty" json:"tshape,omitempty"`
	Pentagonal   *PCStereoPentagonalBiPyramid `xml:"pentagonal,omitempty" json:"pentagonal,omitempty"`
}

//PCStereoCenter,ChoiceOption
type PCStereoTetrahedral struct {
	Center int64 `xml:"center" json:"center"`
	Above  int64 `xml:"above" json:"above"`
	Top    int64 `xml:"top" json:"top"`
	Bottom int64 `xml:"bottom" json:"bottom"`
	Below  int64 `xml:"below" json:"below"`
	Parity int   `xml:"parity,omitempty" json:"parity,omitempty" asn1:"optional"` //Parity,IntegerEnum:clockwise(1),counterclockwise(2),any(3),unknown(255)
	Type   int   `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"`     //Type,IntegerEnum:tetrahedral(1),cumulenic(2),biaryl(3)
}
type PCStereoPlanar struct {
	Left    int64 `xml:"left" json:"left"`
	Ltop    int64 `xml:"ltop" json:"ltop"`
	Lbottom int64 `xml:"lbottom" json:"lbottom"`
	Right   int64 `xml:"right" json:"right"`
	Rtop    int64 `xml:"rtop" json:"rtop"`
	Rbottom int64 `xml:"rbottom" json:"rbottom"`
	Parity  int   `xml:"parity,omitempty" json:"parity,omitempty" asn1:"optional"` //Parity,IntegerEnum:same(1),opposite(2),any(3),unknown(255)
	Type    int   `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"`     //Type,IntegerEnum:planar(1),cumulenic(2)
}
type PCStereoSquarePlanar struct {
	Center int64 `xml:"center" json:"center"`
	Lbelow int64 `xml:"lbelow" json:"lbelow"`
	Rbelow int64 `xml:"rbelow" json:"rbelow"`
	Labove int64 `xml:"labove" json:"labove"`
	Rabove int64 `xml:"rabove" json:"rabove"`
	Parity int   `xml:"parity,omitempty" json:"parity,omitempty" asn1:"optional"` //Parity,IntegerEnum:u-shape(1),z-shape(2),x-shape(3),any(4),unknown(255)
}
type PCStereoOctahedral struct {
	Center int64 `xml:"center" json:"center"`
	Top    int64 `xml:"top" json:"top"`
	Bottom int64 `xml:"bottom" json:"bottom"`
	Labove int64 `xml:"labove" json:"labove"`
	Lbelow int64 `xml:"lbelow" json:"lbelow"`
	Rabove int64 `xml:"rabove" json:"rabove"`
	Rbelow int64 `xml:"rbelow" json:"rbelow"`
}
type PCStereoTrigonalBiPyramid struct {
	Center int64 `xml:"center" json:"center"`
	Above  int64 `xml:"above" json:"above"`
	Below  int64 `xml:"below" json:"below"`
	Top    int64 `xml:"top" json:"top"`
	Bottom int64 `xml:"bottom" json:"bottom"`
	Right  int64 `xml:"right" json:"right"`
}
type PCStereoTShape struct {
	Center int64 `xml:"center" json:"center"`
	Top    int64 `xml:"top" json:"top"`
	Bottom int64 `xml:"bottom" json:"bottom"`
	Above  int64 `xml:"above" json:"above"`
}
type PCStereoPentagonalBiPyramid struct {
	Center int64 `xml:"center" json:"center"`
	Top    int64 `xml:"top" json:"top"`
	Bottom int64 `xml:"bottom" json:"bottom"`
	Left   int64 `xml:"left" json:"left"`
	Labove int64 `xml:"labove" json:"labove"`
	Lbelow int64 `xml:"lbelow" json:"lbelow"`
	Rabove int64 `xml:"rabove" json:"rabove"`
	Rbelow int64 `xml:"rbelow" json:"rbelow"`
}
