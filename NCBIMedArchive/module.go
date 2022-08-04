package NCBIMedArchive

import (
	"ncbiasn/NCBIBiblio"
	"ncbiasn/NCBIMedlars"
	"ncbiasn/NCBIMedline"
	"ncbiasn/NCBIPub"
	"ncbiasn/NCBIPubMed"
)

//MlaRequest,ChoiceOption
type MlaRequest struct {
	Init         interface{}            `xml:"-" json:"-" asn1:"explicit,tag:0"`
	Getmle       int64                  `xml:"getmle,omitempty" json:"getmle,omitempty" asn1:"explicit,tag:1"`
	Getpub       int64                  `xml:"getpub,omitempty" json:"getpub,omitempty" asn1:"explicit,tag:2"`
	Gettitle     *TitleMsg              `xml:"gettitle,omitempty" json:"gettitle,omitempty" asn1:"explicit,tag:3"`
	Citmatch     *NCBIPub.Pub           `xml:"citmatch,omitempty" json:"citmatch,omitempty" asn1:"explicit,tag:4"`
	Fini         interface{}            `xml:"-" json:"-" asn1:"explicit,tag:5"`
	Getmriuids   int64                  `xml:"getmriuids,omitempty" json:"getmriuids,omitempty" asn1:"explicit,tag:6"`
	Getaccuids   *NCBIMedline.MedlineSi `xml:"getaccuids,omitempty" json:"getaccuids,omitempty" asn1:"explicit,tag:7"`
	Uidtopmid    int64                  `xml:"uidtopmid,omitempty" json:"uidtopmid,omitempty" asn1:"explicit,tag:8"`
	Pmidtouid    *NCBIBiblio.PubMedId   `xml:"pmidtouid,omitempty" json:"pmidtouid,omitempty" asn1:"explicit,tag:9"`
	Getmlepmid   *NCBIBiblio.PubMedId   `xml:"getmlepmid,omitempty" json:"getmlepmid,omitempty" asn1:"explicit,tag:10"`
	Getpubpmid   *NCBIBiblio.PubMedId   `xml:"getpubpmid,omitempty" json:"getpubpmid,omitempty" asn1:"explicit,tag:11"`
	Citmatchpmid *NCBIPub.Pub           `xml:"citmatchpmid,omitempty" json:"citmatchpmid,omitempty" asn1:"explicit,tag:12"`
	Getmripmids  int64                  `xml:"getmripmids,omitempty" json:"getmripmids,omitempty" asn1:"explicit,tag:13"`
	Getaccpmids  *NCBIMedline.MedlineSi `xml:"getaccpmids,omitempty" json:"getaccpmids,omitempty" asn1:"explicit,tag:14"`
	Citlstpmids  *NCBIPub.Pub           `xml:"citlstpmids,omitempty" json:"citlstpmids,omitempty" asn1:"explicit,tag:15"`
	Getmleuid    int64                  `xml:"getmleuid,omitempty" json:"getmleuid,omitempty" asn1:"explicit,tag:16"`
	Getmlrpmid   *NCBIBiblio.PubMedId   `xml:"getmlrpmid,omitempty" json:"getmlrpmid,omitempty" asn1:"explicit,tag:17"`
	Getmlruid    int64                  `xml:"getmlruid,omitempty" json:"getmlruid,omitempty" asn1:"explicit,tag:18"`
}

//TitleType,EnumList:not-set(0),name(1),tsub(2),trans(3),jta(4),iso-jta(5),ml-jta(6),coden(7),issn(8),abr(9),isbn(10),all(255)
type TitleType string

type TitleMsg struct {
	Type  *TitleType       `xml:"type,omitempty" json:"type,omitempty"`
	Title NCBIBiblio.Title `xml:"title,omitempty" json:"title,omitempty"`
}
type TitleMsgList struct {
	Num    int64      `xml:"num" json:"num"`
	Titles []TitleMsg `xml:"titles,omitempty" json:"titles,omitempty"`
}

//ErrorVal,EnumList:not-found(0),operational-error(1),cannot-connect-jrsrv(2),cannot-connect-pmdb(3),journal-not-found(4),citation-not-found(5),citation-ambiguous(6),citation-too-many(7),cannot-connect-searchbackend-jrsrv(8),cannot-connect-searchbackend-pmdb(9),cannot-connect-docsumbackend(10)
type ErrorVal string

//MlaBack,ChoiceOption
type MlaBack struct {
	Init     interface{}               `xml:"-" json:"-" asn1:"explicit,tag:0"`
	Error    *ErrorVal                 `xml:"error,omitempty" json:"error,omitempty" asn1:"explicit,tag:1"`
	Getmle   *NCBIMedline.MedlineEntry `xml:"getmle,omitempty" json:"getmle,omitempty" asn1:"explicit,tag:2"`
	Getpub   *NCBIPub.Pub              `xml:"getpub,omitempty" json:"getpub,omitempty" asn1:"explicit,tag:3"`
	Gettitle *TitleMsgList             `xml:"gettitle,omitempty" json:"gettitle,omitempty" asn1:"explicit,tag:4"`
	Citmatch int64                     `xml:"citmatch,omitempty" json:"citmatch,omitempty" asn1:"explicit,tag:5"`
	Fini     interface{}               `xml:"-" json:"-" asn1:"explicit,tag:6"`
	Getuids  []int64                   `xml:"getuids,omitempty" json:"getuids,omitempty" asn1:"explicit,tag:7"`
	Getpmids []int64                   `xml:"getpmids,omitempty" json:"getpmids,omitempty" asn1:"explicit,tag:8"`
	Outuid   int64                     `xml:"outuid,omitempty" json:"outuid,omitempty" asn1:"explicit,tag:9"`
	Outpmid  *NCBIBiblio.PubMedId      `xml:"outpmid,omitempty" json:"outpmid,omitempty" asn1:"explicit,tag:10"`
	Getpme   *NCBIPubMed.PubmedEntry   `xml:"getpme,omitempty" json:"getpme,omitempty" asn1:"explicit,tag:11"`
	Getmlr   *NCBIMedlars.MedlarsEntry `xml:"getmlr,omitempty" json:"getmlr,omitempty" asn1:"explicit,tag:12"`
}
