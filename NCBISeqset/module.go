package NCBISeqset

// moved to NCBISequence
// import "ncbiasn/NCBISequence"
// import "ncbiasn/NCBIGeneral"

// type BioseqSet struct {
// 	Id	*NCBIGeneral.ObjectId	`xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
// 	Coll	*NCBIGeneral.Dbtag	`xml:"coll,omitempty" json:"coll,omitempty" asn1:"optional"`
// 	Level	int64			`xml:"level,omitempty" json:"level,omitempty" asn1:"optional"`
// 	Class	string			`xml:"class" json:"class"`//Class,EnumList:not-set(0),nuc-prot(1),segset(2),conset(3),parts(4),gibb(5),gi(6),genbank(7),pir(8),pub-set(9),equiv(10),swissprot(11),pdb-entry(12),mut-set(13),pop-set(14),phy-set(15),eco-set(16),gen-prod-set(17),wgs-set(18),named-annot(19),named-annot-prod(20),read-set(21),paired-end-reads(22),small-genome-set(23),other(255)
// 	Release	string			`xml:"release,omitempty" json:"release,omitempty" asn1:"optional"`
// 	Date	*NCBIGeneral.Date	`xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
// 	Descr	NCBISequence.SeqDescr	`xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
// 	SeqSet	[]SeqEntry		`xml:"seq-set,omitempty" json:"seq_set,omitempty"`
// 	Annot	[]NCBISequence.SeqAnnot	`xml:"annot,omitempty" json:"annot,omitempty" asn1:"optional"`
// }
// type SeqEntry struct {
// 	Seq	*NCBISequence.Bioseq	`xml:"seq,omitempty" json:"seq,omitempty"`
// 	Set	*BioseqSet		`xml:"set,omitempty" json:"set,omitempty"`
// }
//SeqEntry,ChoiceOption
