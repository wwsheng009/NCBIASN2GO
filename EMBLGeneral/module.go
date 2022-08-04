package EMBLGeneral

import "ncbiasn/NCBIGeneral"

//EMBLDbname,ChoiceOption
type EMBLDbname struct {
	Code string `xml:"code,omitempty" json:"code,omitempty"` //Code,EnumList:embl(0),genbank(1),ddbj(2),geninfo(3),medline(4),swissprot(5),pir(6),pdb(7),epd(8),ecd(9),tfd(10),flybase(11),prosite(12),enzyme(13),mim(14),ecoseq(15),hiv(16),other(255)
	Name string `xml:"name,omitempty" json:"name,omitempty"`
}

type EMBLXref struct {
	Dbname *EMBLDbname            `xml:"dbname,omitempty" json:"dbname,omitempty"`
	Id     []NCBIGeneral.ObjectId `xml:"id,omitempty" json:"id,omitempty"`
}
type EMBLBlock struct {
	Class        string            `xml:"class" json:"class"`                                 //Class,EnumList:not-set(0),standard(1),unannotated(2),other(255)
	Div          string            `xml:"div,omitempty" json:"div,omitempty" asn1:"optional"` //Div,EnumList:fun(0),inv(1),mam(2),org(3),phg(4),pln(5),pri(6),pro(7),rod(8),syn(9),una(10),vrl(11),vrt(12),pat(13),est(14),sts(15),other(255)
	CreationDate *NCBIGeneral.Date `xml:"creation-date,omitempty" json:"creation_date,omitempty"`
	UpdateDate   *NCBIGeneral.Date `xml:"update-date,omitempty" json:"update_date,omitempty"`
	ExtraAcc     []string          `xml:"extra-acc,omitempty" json:"extra_acc,omitempty" asn1:"optional"`
	Keywords     []string          `xml:"keywords,omitempty" json:"keywords,omitempty" asn1:"optional"`
	Xref         []EMBLXref        `xml:"xref,omitempty" json:"xref,omitempty" asn1:"optional"`
}
