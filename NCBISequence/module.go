package NCBISequence

import (
	"ncbiasn/EMBLGeneral"
	"ncbiasn/GenBankGeneral"
	"ncbiasn/NCBIBioSource"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIOrganism"
	"ncbiasn/NCBISeqCommon"
	"ncbiasn/NCBISeqRef"
	"ncbiasn/NCBISeqTable"
	"ncbiasn/NCBISeqalign"
	"ncbiasn/NCBISeqfeat"
	"ncbiasn/NCBISeqloc"
	"ncbiasn/NCBISeqres"
	"ncbiasn/PDBGeneral"
	"ncbiasn/PIRGeneral"
	"ncbiasn/PRFGeneral"
	"ncbiasn/SPGeneral"
)

type Bioseq struct {
	Id    []NCBISeqloc.SeqId `xml:"id,omitempty" json:"id,omitempty"`
	Descr SeqDescr           `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Inst  *SeqInst           `xml:"inst,omitempty" json:"inst,omitempty"`
	Annot []SeqAnnot         `xml:"annot,omitempty" json:"annot,omitempty" asn1:"optional"`
}

//Seqdesc,ChoiceOption
type SeqDescr []Seqdesc

type Seqdesc struct {
	MolType    *GIBBMol                          `xml:"mol-type,omitempty" json:"mol_type,omitempty"`
	Modif      []GIBBMod                         `xml:"modif,omitempty" json:"modif,omitempty"`
	Method     *GIBBMethod                       `xml:"method,omitempty" json:"method,omitempty"`
	Name       string                            `xml:"name,omitempty" json:"name,omitempty"`
	Title      string                            `xml:"title,omitempty" json:"title,omitempty"`
	Org        *NCBIOrganism.OrgRef              `xml:"org,omitempty" json:"org,omitempty"`
	Comment    string                            `xml:"comment,omitempty" json:"comment,omitempty"`
	Num        *NCBISeqRef.Numbering             `xml:"num,omitempty" json:"num,omitempty"`
	Maploc     *NCBIGeneral.Dbtag                `xml:"maploc,omitempty" json:"maploc,omitempty"`
	Pir        *PIRGeneral.PIRBlock              `xml:"pir,omitempty" json:"pir,omitempty"`
	Genbank    *GenBankGeneral.GBBlock           `xml:"genbank,omitempty" json:"genbank,omitempty"`
	Pub        *NCBISeqRef.Pubdesc               `xml:"pub,omitempty" json:"pub,omitempty"`
	Region     string                            `xml:"region,omitempty" json:"region,omitempty"`
	User       *NCBIGeneral.UserObject           `xml:"user,omitempty" json:"user,omitempty"`
	Sp         *SPGeneral.SPBlock                `xml:"sp,omitempty" json:"sp,omitempty"`
	Dbxref     *NCBIGeneral.Dbtag                `xml:"dbxref,omitempty" json:"dbxref,omitempty"`
	Embl       *EMBLGeneral.EMBLBlock            `xml:"embl,omitempty" json:"embl,omitempty"`
	CreateDate *NCBIGeneral.Date                 `xml:"create-date,omitempty" json:"create_date,omitempty"`
	UpdateDate *NCBIGeneral.Date                 `xml:"update-date,omitempty" json:"update_date,omitempty"`
	Prf        *PRFGeneral.PRFBlock              `xml:"prf,omitempty" json:"prf,omitempty"`
	Pdb        *PDBGeneral.PDBBlock              `xml:"pdb,omitempty" json:"pdb,omitempty"`
	Het        NCBISeqCommon.Heterogen           `xml:"het,omitempty" json:"het,omitempty"`
	Source     *NCBIBioSource.BioSource          `xml:"source,omitempty" json:"source,omitempty"`
	Molinfo    *MolInfo                          `xml:"molinfo,omitempty" json:"molinfo,omitempty"`
	Modelev    *NCBISeqfeat.ModelEvidenceSupport `xml:"modelev,omitempty" json:"modelev,omitempty"`
}

type MolInfo struct {
	Biomol       int    `xml:"biomol" json:"biomol"` //Biomol,IntegerEnum:unknown(0),genomic(1),pre-RNA(2),mRNA(3),rRNA(4),tRNA(5),snRNA(6),scRNA(7),peptide(8),other-genetic(9),genomic-mRNA(10),cRNA(11),snoRNA(12),transcribed-RNA(13),ncRNA(14),tmRNA(15),other(255)
	Tech         int    `xml:"tech" json:"tech"`     //Tech,IntegerEnum:unknown(0),standard(1),est(2),sts(3),survey(4),genemap(5),physmap(6),derived(7),concept-trans(8),seq-pept(9),both(10),seq-pept-overlap(11),seq-pept-homol(12),concept-trans-a(13),htgs-1(14),htgs-2(15),htgs-3(16),fli-cdna(17),htgs-0(18),htc(19),wgs(20),barcode(21),composite-wgs-htgs(22),tsa(23),targeted(24),other(255)
	Techexp      string `xml:"techexp,omitempty" json:"techexp,omitempty" asn1:"optional"`
	Completeness int    `xml:"completeness" json:"completeness"` //Completeness,IntegerEnum:unknown(0),complete(1),partial(2),no-left(3),no-right(4),no-ends(5),has-left(6),has-right(7),other(255)
	Gbmoltype    string `xml:"gbmoltype,omitempty" json:"gbmoltype,omitempty" asn1:"optional"`
}
type GIBBMol string

//GIBBMol,EnumList:unknown(0),genomic(1),pre-mRNA(2),mRNA(3),rRNA(4),tRNA(5),snRNA(6),scRNA(7),peptide(8),other-genetic(9),genomic-mRNA(10),other(255)
type GIBBMod string

//GIBBMod,EnumList:dna(0),rna(1),extrachrom(2),plasmid(3),mitochondrial(4),chloroplast(5),kinetoplast(6),cyanelle(7),synthetic(8),recombinant(9),partial(10),complete(11),mutagen(12),natmut(13),transposon(14),insertion-seq(15),no-left(16),no-right(17),macronuclear(18),proviral(19),est(20),sts(21),survey(22),chromoplast(23),genemap(24),restmap(25),physmap(26),other(255)
type GIBBMethod string

//GIBBMethod,EnumList:concept-trans(1),seq-pept(2),both(3),seq-pept-overlap(4),seq-pept-homol(5),concept-trans-a(6),other(255)
//  部分数据结构迁移到NCBISeqRef作为单独的数据定义，防止结构文件想到引用
// type Numbering struct {
// 	Cont	*NumCont	`xml:"cont,omitempty" json:"cont,omitempty"`
// 	Enum	*NumEnum	`xml:"enum,omitempty" json:"enum,omitempty"`
// 	Ref	*NumRef		`xml:"ref,omitempty" json:"ref,omitempty"`
// 	Real	*NumReal	`xml:"real,omitempty" json:"real,omitempty"`
// }
//Numbering,ChoiceOption
// type NumCont struct {
// 	Refnum		int64	`xml:"refnum" json:"refnum" asn1:"default:1"`
// 	HasZero		bool	`xml:"has-zero" json:"has_zero"`
// 	Ascending	bool	`xml:"ascending" json:"ascending"`
// }
// type NumEnum struct {
// 	Num	int64		`xml:"num" json:"num"`
// 	Names	[]string	`xml:"names,omitempty" json:"names,omitempty"`
// }
// type NumRef struct {
// 	Type	string			`xml:"type" json:"type"`//Type,EnumList:not-set(0),sources(1),aligns(2)
// 	Aligns	*NCBISeqalign.SeqAlign	`xml:"aligns,omitempty" json:"aligns,omitempty" asn1:"optional"`
// }
// type NumReal struct {
// 	A	float64	`xml:"a" json:"a"`
// 	B	float64	`xml:"b" json:"b"`
// 	Units	string	`xml:"units,omitempty" json:"units,omitempty" asn1:"optional"`
// }
// type Pubdesc struct {
// 	Pub		*NCBIPub.PubEquiv	`xml:"pub,omitempty" json:"pub,omitempty"`
// 	Name		string			`xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
// 	Fig		string			`xml:"fig,omitempty" json:"fig,omitempty" asn1:"optional"`
// 	Num		*Numbering		`xml:"num,omitempty" json:"num,omitempty" asn1:"optional"`
// 	Numexc		bool			`xml:"numexc,omitempty" json:"numexc,omitempty" asn1:"optional"`
// 	PolyA		bool			`xml:"poly-a,omitempty" json:"poly_a,omitempty" asn1:"optional"`
// 	Maploc		string			`xml:"maploc,omitempty" json:"maploc,omitempty" asn1:"optional"`
// 	SeqRaw		*string			`xml:"seq-raw,omitempty" json:"seq_raw,omitempty" asn1:"optional"`
// 	AlignGroup	int64			`xml:"align-group,omitempty" json:"align_group,omitempty" asn1:"optional"`
// 	Comment		string			`xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
// 	Reftype		int			`xml:"reftype" json:"reftype"`//Reftype,IntegerEnum:seq(0),sites(1),feats(2),no-target(3)
// }

// type Heterogen string
type SeqInst struct {
	Repr     string                 `xml:"repr" json:"repr"` //Repr,EnumList:not-set(0),virtual(1),raw(2),seg(3),const(4),ref(5),consen(6),map(7),delta(8),other(255)
	Mol      string                 `xml:"mol" json:"mol"`   //Mol,EnumList:not-set(0),dna(1),rna(2),aa(3),na(4),other(255)
	Length   int64                  `xml:"length,omitempty" json:"length,omitempty" asn1:"optional"`
	Fuzz     *NCBIGeneral.IntFuzz   `xml:"fuzz,omitempty" json:"fuzz,omitempty" asn1:"optional"`
	Topology string                 `xml:"topology" json:"topology"`                                 //Topology,EnumList:not-set(0),linear(1),circular(2),tandem(3),other(255)
	Strand   string                 `xml:"strand,omitempty" json:"strand,omitempty" asn1:"optional"` //Strand,EnumList:not-set(0),ss(1),ds(2),mixed(3),other(255)
	SeqData  *NCBISeqCommon.SeqData `xml:"seq-data,omitempty" json:"seq_data,omitempty" asn1:"optional"`
	Ext      *SeqExt                `xml:"ext,omitempty" json:"ext,omitempty" asn1:"optional"`
	Hist     *SeqHist               `xml:"hist,omitempty" json:"hist,omitempty" asn1:"optional"`
}
type SeqExt struct {
	Seg   SegExt   `xml:"seg,omitempty" json:"seg,omitempty"`
	Ref   *RefExt  `xml:"ref,omitempty" json:"ref,omitempty"`
	Map   MapExt   `xml:"map,omitempty" json:"map,omitempty"`
	Delta DeltaExt `xml:"delta,omitempty" json:"delta,omitempty"`
}
type SegExt []NCBISeqloc.SeqLoc
type RefExt NCBISeqloc.SeqLoc
type MapExt []NCBISeqfeat.SeqFeat
type DeltaExt []DeltaSeq

//DeltaSeq,ChoiceOption
type DeltaSeq struct {
	Loc     *NCBISeqloc.SeqLoc        `xml:"loc,omitempty" json:"loc,omitempty"`
	Literal *NCBISeqCommon.SeqLiteral `xml:"literal,omitempty" json:"literal,omitempty"`
}

// move to module NCBISeqCommon
// type SeqLiteral struct {
// 	Length	int64			`xml:"length" json:"length"`
// 	Fuzz	*NCBIGeneral.IntFuzz	`xml:"fuzz,omitempty" json:"fuzz,omitempty" asn1:"optional"`
// 	SeqData	*SeqData		`xml:"seq-data,omitempty" json:"seq_data,omitempty" asn1:"optional"`
// }
type SeqHist struct {
	Assembly   []NCBISeqalign.SeqAlign `xml:"assembly,omitempty" json:"assembly,omitempty" asn1:"optional"`
	Replaces   *SeqHistRec             `xml:"replaces,omitempty" json:"replaces,omitempty" asn1:"optional"`
	ReplacedBy *SeqHistRec             `xml:"replaced-by,omitempty" json:"replaced_by,omitempty" asn1:"optional"`
	Deleted    struct {
		Bool bool              `xml:"bool,omitempty" json:"bool,omitempty"`
		Date *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty"`
	} `xml:"deleted,omitempty" json:"deleted,omitempty" asn1:"optional"` //Deleted,ChoiceOption
}
type SeqHistRec struct {
	Date *NCBIGeneral.Date  `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
	Ids  []NCBISeqloc.SeqId `xml:"ids,omitempty" json:"ids,omitempty"`
}

// move to module NCBISeqCommon
// type SeqData struct {
// 	Iupacna		*IUPACna	`xml:"iupacna,omitempty" json:"iupacna,omitempty"`
// 	Iupacaa		*IUPACaa	`xml:"iupacaa,omitempty" json:"iupacaa,omitempty"`
// 	Ncbi2na		*NCBI2na	`xml:"ncbi2na,omitempty" json:"ncbi2na,omitempty"`
// 	Ncbi4na		*NCBI4na	`xml:"ncbi4na,omitempty" json:"ncbi4na,omitempty"`
// 	Ncbi8na		*NCBI8na	`xml:"ncbi8na,omitempty" json:"ncbi8na,omitempty"`
// 	Ncbipna		*NCBIpna	`xml:"ncbipna,omitempty" json:"ncbipna,omitempty"`
// 	Ncbi8aa		*NCBI8aa	`xml:"ncbi8aa,omitempty" json:"ncbi8aa,omitempty"`
// 	Ncbieaa		*NCBIeaa	`xml:"ncbieaa,omitempty" json:"ncbieaa,omitempty"`
// 	Ncbipaa		*NCBIpaa	`xml:"ncbipaa,omitempty" json:"ncbipaa,omitempty"`
// 	Ncbistdaa	*NCBIstdaa	`xml:"ncbistdaa,omitempty" json:"ncbistdaa,omitempty"`
// 	Gap		*SeqGap		`xml:"gap,omitempty" json:"gap,omitempty"`
// }
//SeqData,ChoiceOption
// type SeqGap struct {
// 	Type		int			`xml:"type" json:"type"`//Type,IntegerEnum:unknown(0),fragment(1),clone(2),short-arm(3),heterochromatin(4),centromere(5),telomere(6),repeat(7),contig(8),scaffold(9),contamination(10),other(255)
// 	Linkage		int			`xml:"linkage,omitempty" json:"linkage,omitempty" asn1:"optional"`//Linkage,IntegerEnum:unlinked(0),linked(1),other(255)
// 	LinkageEvidence	[]LinkageEvidence	`xml:"linkage-evidence,omitempty" json:"linkage_evidence,omitempty" asn1:"optional"`
// }
// type LinkageEvidence struct {
// 	Type int `xml:"type" json:"type"`//Type,IntegerEnum:paired-ends(0),align-genus(1),align-xgenus(2),align-trnscpt(3),within-clone(4),clone-contig(5),map(6),strobe(7),unspecified(8),pcr(9),proximity-ligation(10),other(255)
// }
// type IUPACna string
// type IUPACaa string
// type NCBI2na []byte
// type NCBI4na []byte
// type NCBI8na []byte
// type NCBIpna []byte
// type NCBI8aa []byte
// type NCBIeaa string
// type NCBIpaa []byte
// type NCBIstdaa []byte
type TextannotId struct {
	Name      string `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Accession string `xml:"accession,omitempty" json:"accession,omitempty" asn1:"optional"`
	Release   string `xml:"release,omitempty" json:"release,omitempty" asn1:"optional"`
	Version   int64  `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
}

//AnnotId,ChoiceOption
type AnnotId struct {
	Local   *NCBIGeneral.ObjectId `xml:"local,omitempty" json:"local,omitempty"`
	Ncbi    int64                 `xml:"ncbi,omitempty" json:"ncbi,omitempty"`
	General *NCBIGeneral.Dbtag    `xml:"general,omitempty" json:"general,omitempty"`
	Other   *TextannotId          `xml:"other,omitempty" json:"other,omitempty"`
}

type AnnotDescr []Annotdesc

//Annotdesc,ChoiceOption
type Annotdesc struct {
	Name       string                  `xml:"name,omitempty" json:"name,omitempty"`
	Title      string                  `xml:"title,omitempty" json:"title,omitempty"`
	Comment    string                  `xml:"comment,omitempty" json:"comment,omitempty"`
	Pub        *NCBISeqRef.Pubdesc     `xml:"pub,omitempty" json:"pub,omitempty"`
	User       *NCBIGeneral.UserObject `xml:"user,omitempty" json:"user,omitempty"`
	CreateDate *NCBIGeneral.Date       `xml:"create-date,omitempty" json:"create_date,omitempty"`
	UpdateDate *NCBIGeneral.Date       `xml:"update-date,omitempty" json:"update_date,omitempty"`
	Src        *NCBISeqloc.SeqId       `xml:"src,omitempty" json:"src,omitempty"`
	Align      *AlignDef               `xml:"align,omitempty" json:"align,omitempty"`
	Region     *NCBISeqloc.SeqLoc      `xml:"region,omitempty" json:"region,omitempty"`
}

type AlignDef struct {
	AlignType int                `xml:"align-type" json:"align_type"` //AlignType,IntegerEnum:ref(1),alt(2),blocks(3),other(255)
	Ids       []NCBISeqloc.SeqId `xml:"ids,omitempty" json:"ids,omitempty" asn1:"optional"`
}
type SeqAnnot struct {
	Id   []AnnotId  `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Db   int        `xml:"db,omitempty" json:"db,omitempty" asn1:"optional"` //Db,IntegerEnum:genbank(1),embl(2),ddbj(3),pir(4),sp(5),bbone(6),pdb(7),other(255)
	Name string     `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Desc AnnotDescr `xml:"desc,omitempty" json:"desc,omitempty" asn1:"optional"`
	Data struct {
		Ftable   []NCBISeqfeat.SeqFeat   `xml:"ftable,omitempty" json:"ftable,omitempty"`
		Align    []NCBISeqalign.SeqAlign `xml:"align,omitempty" json:"align,omitempty"`
		Graph    []NCBISeqres.SeqGraph   `xml:"graph,omitempty" json:"graph,omitempty"`
		Ids      []NCBISeqloc.SeqId      `xml:"ids,omitempty" json:"ids,omitempty"`
		Locs     []NCBISeqloc.SeqLoc     `xml:"locs,omitempty" json:"locs,omitempty"`
		SeqTable *NCBISeqTable.SeqTable  `xml:"seq-table,omitempty" json:"seq_table,omitempty"`
	} `xml:"data" json:"data"` //Data,ChoiceOption
}

// package NCBISeqset
// import "ncbiasn/NCBISequence"
// import "ncbiasn/NCBIGeneral"

type BioseqSet struct {
	Id      *NCBIGeneral.ObjectId `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Coll    *NCBIGeneral.Dbtag    `xml:"coll,omitempty" json:"coll,omitempty" asn1:"optional"`
	Level   int64                 `xml:"level,omitempty" json:"level,omitempty" asn1:"optional"`
	Class   string                `xml:"class" json:"class"` //Class,EnumList:not-set(0),nuc-prot(1),segset(2),conset(3),parts(4),gibb(5),gi(6),genbank(7),pir(8),pub-set(9),equiv(10),swissprot(11),pdb-entry(12),mut-set(13),pop-set(14),phy-set(15),eco-set(16),gen-prod-set(17),wgs-set(18),named-annot(19),named-annot-prod(20),read-set(21),paired-end-reads(22),small-genome-set(23),other(255)
	Release string                `xml:"release,omitempty" json:"release,omitempty" asn1:"optional"`
	Date    *NCBIGeneral.Date     `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
	Descr   SeqDescr              `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	SeqSet  []SeqEntry            `xml:"seq-set,omitempty" json:"seq_set,omitempty"`
	Annot   []SeqAnnot            `xml:"annot,omitempty" json:"annot,omitempty" asn1:"optional"`
}

// 适用于Nucleotide/nuccore下载的Full record
//基因系列的入口
type SeqEntry struct {
	Seq *Bioseq    `xml:"seq,omitempty" json:"seq,omitempty"`
	Set *BioseqSet `xml:"set,omitempty" json:"set,omitempty"`
}
