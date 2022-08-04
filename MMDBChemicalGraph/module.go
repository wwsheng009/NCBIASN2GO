package MMDBChemicalGraph

import (
	"ncbiasn/MMDBCommon"
	"ncbiasn/NCBIBioSource"
	"ncbiasn/NCBIPub"
	"ncbiasn/NCBISeqloc"
)

type BiostrucGraph struct {
	Descr              []BiomolDescr      `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	MoleculeGraphs     []MoleculeGraph    `xml:"molecule-graphs,omitempty" json:"molecule_graphs,omitempty"`
	InterMoleculeBonds []InterResidueBond `xml:"inter-molecule-bonds,omitempty" json:"inter_molecule_bonds,omitempty" asn1:"optional"`
	ResidueGraphs      []ResidueGraph     `xml:"residue-graphs,omitempty" json:"residue_graphs,omitempty" asn1:"optional"`
}

//BiomolDescr,ChoiceOption
type BiomolDescr struct {
	Name         string                   `xml:"name,omitempty" json:"name,omitempty"`
	PdbClass     string                   `xml:"pdb-class,omitempty" json:"pdb_class,omitempty"`
	PdbSource    string                   `xml:"pdb-source,omitempty" json:"pdb_source,omitempty"`
	PdbComment   string                   `xml:"pdb-comment,omitempty" json:"pdb_comment,omitempty"`
	OtherComment string                   `xml:"other-comment,omitempty" json:"other_comment,omitempty"`
	Organism     *NCBIBioSource.BioSource `xml:"organism,omitempty" json:"organism,omitempty"`
	Attribution  *NCBIPub.Pub             `xml:"attribution,omitempty" json:"attribution,omitempty"`
	AssemblyType int                      `xml:"assembly-type,omitempty" json:"assembly_type,omitempty"` //AssemblyType,IntegerEnum:physiological-form(1),crystallographic-cell(2),other(255)
	MoleculeType int                      `xml:"molecule-type,omitempty" json:"molecule_type,omitempty"` //MoleculeType,IntegerEnum:dna(1),rna(2),protein(3),other-biopolymer(4),solvent(5),other-nonpolymer(6),other(255)
}

type MoleculeGraph struct {
	Id                *MoleculeId        `xml:"id,omitempty" json:"id,omitempty"`
	Descr             []BiomolDescr      `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	SeqId             *NCBISeqloc.SeqId  `xml:"seq-id,omitempty" json:"seq_id,omitempty" asn1:"optional"`
	ResidueSequence   []Residue          `xml:"residue-sequence,omitempty" json:"residue_sequence,omitempty"`
	InterResidueBonds []InterResidueBond `xml:"inter-residue-bonds,omitempty" json:"inter_residue_bonds,omitempty" asn1:"optional"`
	Sid               *PCSubstanceId     `xml:"sid,omitempty" json:"sid,omitempty" asn1:"optional"`
}
type MoleculeId int64
type PCSubstanceId int64
type Residue struct {
	Id           *ResidueId        `xml:"id,omitempty" json:"id,omitempty"`
	Name         string            `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	ResidueGraph *ResidueGraphPntr `xml:"residue-graph,omitempty" json:"residue_graph,omitempty"`
}
type ResidueId int64

//ResidueGraphPntr,ChoiceOption
type ResidueGraphPntr struct {
	Local    *ResidueGraphId              `xml:"local,omitempty" json:"local,omitempty"`
	Biostruc *BiostrucGraphPntr           `xml:"biostruc,omitempty" json:"biostruc,omitempty"`
	Standard *BiostrucResidueGraphSetPntr `xml:"standard,omitempty" json:"standard,omitempty"`
}
type BiostrucGraphPntr struct {
	BiostrucId     *MMDBCommon.BiostrucId `xml:"biostruc-id,omitempty" json:"biostruc_id,omitempty"`
	ResidueGraphId *ResidueGraphId        `xml:"residue-graph-id,omitempty" json:"residue_graph_id,omitempty"`
}
type BiostrucResidueGraphSetPntr struct {
	BiostrucResidueGraphSetId *MMDBCommon.BiostrucId `xml:"biostruc-residue-graph-set-id,omitempty" json:"biostruc_residue_graph_set_id,omitempty"`
	ResidueGraphId            *ResidueGraphId        `xml:"residue-graph-id,omitempty" json:"residue_graph_id,omitempty"`
}
type ResidueGraph struct {
	Id            *ResidueGraphId    `xml:"id,omitempty" json:"id,omitempty"`
	Descr         []BiomolDescr      `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	ResidueType   int                `xml:"residue-type,omitempty" json:"residue_type,omitempty" asn1:"optional"` //ResidueType,IntegerEnum:deoxyribonucleotide(1),ribonucleotide(2),amino-acid(3),other(255)
	IupacCode     []string           `xml:"iupac-code,omitempty" json:"iupac_code,omitempty" asn1:"optional"`
	Atoms         []Atom             `xml:"atoms,omitempty" json:"atoms,omitempty"`
	Bonds         []IntraResidueBond `xml:"bonds,omitempty" json:"bonds,omitempty"`
	ChiralCenters []ChiralCenter     `xml:"chiral-centers,omitempty" json:"chiral_centers,omitempty" asn1:"optional"`
}
type ResidueGraphId int64
type Atom struct {
	Id              *AtomId  `xml:"id,omitempty" json:"id,omitempty"`
	Name            string   `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	IupacCode       []string `xml:"iupac-code,omitempty" json:"iupac_code,omitempty" asn1:"optional"`
	Element         string   `xml:"element" json:"element"`                                                       //Element,EnumList:h(1),he(2),li(3),be(4),b(5),c(6),n(7),o(8),f(9),ne(10),na(11),mg(12),al(13),si(14),p(15),s(16),cl(17),ar(18),k(19),ca(20),sc(21),ti(22),v(23),cr(24),mn(25),fe(26),co(27),ni(28),cu(29),zn(30),ga(31),ge(32),as(33),se(34),br(35),kr(36),rb(37),sr(38),y(39),zr(40),nb(41),mo(42),tc(43),ru(44),rh(45),pd(46),ag(47),cd(48),in(49),sn(50),sb(51),te(52),i(53),xe(54),cs(55),ba(56),la(57),ce(58),pr(59),nd(60),pm(61),sm(62),eu(63),gd(64),tb(65),dy(66),ho(67),er(68),tm(69),yb(70),lu(71),hf(72),ta(73),w(74),re(75),os(76),ir(77),pt(78),au(79),hg(80),tl(81),pb(82),bi(83),po(84),at(85),rn(86),fr(87),ra(88),ac(89),th(90),pa(91),u(92),np(93),pu(94),am(95),cm(96),bk(97),cf(98),es(99),fm(100),md(101),no(102),lr(103),other(254),unknown(255)
	IonizableProton string   `xml:"ionizable-proton,omitempty" json:"ionizable_proton,omitempty" asn1:"optional"` //IonizableProton,EnumList:true(1),false(2),unknown(255)
}
type AtomId int64
type IntraResidueBond struct {
	AtomId1   *AtomId `xml:"atom-id-1,omitempty" json:"atom_id_1,omitempty"`
	AtomId2   *AtomId `xml:"atom-id-2,omitempty" json:"atom_id_2,omitempty"`
	BondOrder int     `xml:"bond-order,omitempty" json:"bond_order,omitempty" asn1:"optional"` //BondOrder,IntegerEnum:single(1),partial-double(2),aromatic(3),double(4),triple(5),other(6),unknown(255)
}
type ChiralCenter struct {
	C    *AtomId `xml:"c,omitempty" json:"c,omitempty"`
	N1   *AtomId `xml:"n1,omitempty" json:"n1,omitempty"`
	N2   *AtomId `xml:"n2,omitempty" json:"n2,omitempty"`
	N3   *AtomId `xml:"n3,omitempty" json:"n3,omitempty"`
	Sign string  `xml:"sign" json:"sign"` //Sign,EnumList:positive(1),negative(2)
}
type InterResidueBond struct {
	AtomId1   *AtomPntr `xml:"atom-id-1,omitempty" json:"atom_id_1,omitempty"`
	AtomId2   *AtomPntr `xml:"atom-id-2,omitempty" json:"atom_id_2,omitempty"`
	BondOrder int       `xml:"bond-order,omitempty" json:"bond_order,omitempty" asn1:"optional"` //BondOrder,IntegerEnum:single(1),partial-double(2),aromatic(3),double(4),triple(5),other(6),unknown(255)
}
type AtomPntr struct {
	MoleculeId *MoleculeId `xml:"molecule-id,omitempty" json:"molecule_id,omitempty"`
	ResidueId  *ResidueId  `xml:"residue-id,omitempty" json:"residue_id,omitempty"`
	AtomId     *AtomId     `xml:"atom-id,omitempty" json:"atom_id,omitempty"`
}
type AtomPntrSet []AtomPntr
