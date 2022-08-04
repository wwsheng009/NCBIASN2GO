package NCBIBiblio

import "ncbiasn/NCBIGeneral"

type ArticleId struct {
	Pubmed  *PubMedId          `xml:"pubmed,omitempty" json:"pubmed,omitempty"`
	Medline *MedlineUID        `xml:"medline,omitempty" json:"medline,omitempty"`
	Doi     DOI                `xml:"doi,omitempty" json:"doi,omitempty"`
	Pii     PII                `xml:"pii,omitempty" json:"pii,omitempty"`
	Pmcid   *PmcID             `xml:"pmcid,omitempty" json:"pmcid,omitempty"`
	Pmcpid  PmcPid             `xml:"pmcpid,omitempty" json:"pmcpid,omitempty"`
	Pmpid   PmPid              `xml:"pmpid,omitempty" json:"pmpid,omitempty"`
	Other   *NCBIGeneral.Dbtag `xml:"other,omitempty" json:"other,omitempty"`
}

//ArticleId,ChoiceOption
type PubMedId int64
type MedlineUID int64
type DOI string
type PII string
type PmcID int64
type PmcPid string
type PmPid string
type ArticleIdSet []ArticleId
type PubStatus int //PubStatus,IntegerEnum:received(1),accepted(2),epublish(3),ppublish(4),revised(5),pmc(6),pmcr(7),pubmed(8),pubmedr(9),aheadofprint(10),premedline(11),medline(12),other(255)
type PubStatusDate struct {
	Pubstatus *PubStatus        `xml:"pubstatus,omitempty" json:"pubstatus,omitempty"`
	Date      *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty"`
}
type PubStatusDateSet []PubStatusDate
type CitArt struct {
	Title   Title     `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	Authors *AuthList `xml:"authors,omitempty" json:"authors,omitempty" asn1:"optional"`
	From    struct {
		Journal *CitJour `xml:"journal,omitempty" json:"journal,omitempty"`
		Book    *CitBook `xml:"book,omitempty" json:"book,omitempty"`
		Proc    *CitProc `xml:"proc,omitempty" json:"proc,omitempty"`
	} `xml:"from" json:"from"` //From,ChoiceOption
	Ids ArticleIdSet `xml:"ids,omitempty" json:"ids,omitempty" asn1:"optional"`
}
type CitJour struct {
	Title Title    `xml:"title,omitempty" json:"title,omitempty"`
	Imp   *Imprint `xml:"imp,omitempty" json:"imp,omitempty"`
}
type CitBook struct {
	Title   Title     `xml:"title,omitempty" json:"title,omitempty"`
	Coll    Title     `xml:"coll,omitempty" json:"coll,omitempty" asn1:"optional"`
	Authors *AuthList `xml:"authors,omitempty" json:"authors,omitempty"`
	Imp     *Imprint  `xml:"imp,omitempty" json:"imp,omitempty"`
}
type CitProc struct {
	Book *CitBook `xml:"book,omitempty" json:"book,omitempty"`
	Meet *Meeting `xml:"meet,omitempty" json:"meet,omitempty"`
}
type CitPat struct {
	Title      string            `xml:"title" json:"title"`
	Authors    *AuthList         `xml:"authors,omitempty" json:"authors,omitempty"`
	Country    string            `xml:"country" json:"country"`
	DocType    string            `xml:"doc-type" json:"doc_type"`
	Number     string            `xml:"number,omitempty" json:"number,omitempty" asn1:"optional"`
	DateIssue  *NCBIGeneral.Date `xml:"date-issue,omitempty" json:"date_issue,omitempty" asn1:"optional"`
	Class      []string          `xml:"class,omitempty" json:"class,omitempty" asn1:"optional"`
	AppNumber  string            `xml:"app-number,omitempty" json:"app_number,omitempty" asn1:"optional"`
	AppDate    *NCBIGeneral.Date `xml:"app-date,omitempty" json:"app_date,omitempty" asn1:"optional"`
	Applicants *AuthList         `xml:"applicants,omitempty" json:"applicants,omitempty" asn1:"optional"`
	Assignees  *AuthList         `xml:"assignees,omitempty" json:"assignees,omitempty" asn1:"optional"`
	Priority   []PatentPriority  `xml:"priority,omitempty" json:"priority,omitempty" asn1:"optional"`
	Abstract   string            `xml:"abstract,omitempty" json:"abstract,omitempty" asn1:"optional"`
}
type PatentPriority struct {
	Country string            `xml:"country" json:"country"`
	Number  string            `xml:"number" json:"number"`
	Date    *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty"`
}
type IdPat struct {
	Country string `xml:"country" json:"country"`
	//Id,ChoiceOption
	Id struct {
		Number    string `xml:"number,omitempty" json:"number,omitempty"`
		AppNumber string `xml:"app-number,omitempty" json:"app_number,omitempty"`
	} `xml:"id" json:"id"`
	DocType string `xml:"doc-type,omitempty" json:"doc_type,omitempty" asn1:"optional"`
}
type CitLet struct {
	Cit   *CitBook `xml:"cit,omitempty" json:"cit,omitempty"`
	ManId string   `xml:"man-id,omitempty" json:"man_id,omitempty" asn1:"optional"`
	Type  string   `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"` //Type,EnumList:manuscript(1),letter(2),thesis(3)
}
type CitSub struct {
	Authors *AuthList         `xml:"authors,omitempty" json:"authors,omitempty"`
	Imp     *Imprint          `xml:"imp,omitempty" json:"imp,omitempty" asn1:"optional"`
	Medium  string            `xml:"medium,omitempty" json:"medium,omitempty" asn1:"optional"` //Medium,EnumList:paper(1),tape(2),floppy(3),email(4),other(255)
	Date    *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
	Descr   string            `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
}
type CitGen struct {
	Cit          string            `xml:"cit,omitempty" json:"cit,omitempty" asn1:"optional"`
	Authors      *AuthList         `xml:"authors,omitempty" json:"authors,omitempty" asn1:"optional"`
	Muid         int64             `xml:"muid,omitempty" json:"muid,omitempty" asn1:"optional"`
	Journal      Title             `xml:"journal,omitempty" json:"journal,omitempty" asn1:"optional"`
	Volume       string            `xml:"volume,omitempty" json:"volume,omitempty" asn1:"optional"`
	Issue        string            `xml:"issue,omitempty" json:"issue,omitempty" asn1:"optional"`
	Pages        string            `xml:"pages,omitempty" json:"pages,omitempty" asn1:"optional"`
	Date         *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
	SerialNumber int64             `xml:"serial-number,omitempty" json:"serial_number,omitempty" asn1:"optional"`
	Title        string            `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	Pmid         *PubMedId         `xml:"pmid,omitempty" json:"pmid,omitempty" asn1:"optional"`
}
type AuthList struct {
	Names struct {
		Std []Author `xml:"std,omitempty" json:"std,omitempty"`
		Ml  []string `xml:"ml,omitempty" json:"ml,omitempty"`
		Str []string `xml:"str,omitempty" json:"str,omitempty"`
	} `xml:"names" json:"names"` //Names,ChoiceOption
	Affil *Affil `xml:"affil,omitempty" json:"affil,omitempty" asn1:"optional"`
}
type Author struct {
	Name   *NCBIGeneral.PersonId `xml:"name,omitempty" json:"name,omitempty"`
	Level  string                `xml:"level,omitempty" json:"level,omitempty" asn1:"optional"` //Level,EnumList:primary(1),secondary(2)
	Role   string                `xml:"role,omitempty" json:"role,omitempty" asn1:"optional"`   //Role,EnumList:compiler(1),editor(2),patent-assignee(3),translator(4)
	Affil  *Affil                `xml:"affil,omitempty" json:"affil,omitempty" asn1:"optional"`
	IsCorr bool                  `xml:"is-corr,omitempty" json:"is_corr,omitempty" asn1:"optional"`
}

//Affil,ChoiceOption
type Affil struct {
	Str string `xml:"str,omitempty" json:"str,omitempty"`
	Std struct {
		Affil      string `xml:"affil,omitempty" json:"affil,omitempty" asn1:"optional"`
		Div        string `xml:"div,omitempty" json:"div,omitempty" asn1:"optional"`
		City       string `xml:"city,omitempty" json:"city,omitempty" asn1:"optional"`
		Sub        string `xml:"sub,omitempty" json:"sub,omitempty" asn1:"optional"`
		Country    string `xml:"country,omitempty" json:"country,omitempty" asn1:"optional"`
		Street     string `xml:"street,omitempty" json:"street,omitempty" asn1:"optional"`
		Email      string `xml:"email,omitempty" json:"email,omitempty" asn1:"optional"`
		Fax        string `xml:"fax,omitempty" json:"fax,omitempty" asn1:"optional"`
		Phone      string `xml:"phone,omitempty" json:"phone,omitempty" asn1:"optional"`
		PostalCode string `xml:"postal-code,omitempty" json:"postal_code,omitempty" asn1:"optional"`
	} `xml:"std,omitempty" json:"std,omitempty"`
}

type Title []struct {
	Name   string `xml:"name,omitempty" json:"name,omitempty"`
	Tsub   string `xml:"tsub,omitempty" json:"tsub,omitempty"`
	Trans  string `xml:"trans,omitempty" json:"trans,omitempty"`
	Jta    string `xml:"jta,omitempty" json:"jta,omitempty"`
	IsoJta string `xml:"iso-jta,omitempty" json:"iso_jta,omitempty"`
	MlJta  string `xml:"ml-jta,omitempty" json:"ml_jta,omitempty"`
	Coden  string `xml:"coden,omitempty" json:"coden,omitempty"`
	Issn   string `xml:"issn,omitempty" json:"issn,omitempty"`
	Abr    string `xml:"abr,omitempty" json:"abr,omitempty"`
	Isbn   string `xml:"isbn,omitempty" json:"isbn,omitempty"`
}
type Imprint struct {
	Date      *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty"`
	Volume    string            `xml:"volume,omitempty" json:"volume,omitempty" asn1:"optional"`
	Issue     string            `xml:"issue,omitempty" json:"issue,omitempty" asn1:"optional"`
	Pages     string            `xml:"pages,omitempty" json:"pages,omitempty" asn1:"optional"`
	Section   string            `xml:"section,omitempty" json:"section,omitempty" asn1:"optional"`
	Pub       *Affil            `xml:"pub,omitempty" json:"pub,omitempty" asn1:"optional"`
	Cprt      *NCBIGeneral.Date `xml:"cprt,omitempty" json:"cprt,omitempty" asn1:"optional"`
	PartSup   string            `xml:"part-sup,omitempty" json:"part_sup,omitempty" asn1:"optional"`
	Language  string            `xml:"language" json:"language" asn1:"default:ENG"`
	Prepub    string            `xml:"prepub,omitempty" json:"prepub,omitempty" asn1:"optional"` //Prepub,EnumList:submitted(1),in-press(2),other(255)
	PartSupi  string            `xml:"part-supi,omitempty" json:"part_supi,omitempty" asn1:"optional"`
	Retract   *CitRetract       `xml:"retract,omitempty" json:"retract,omitempty" asn1:"optional"`
	Pubstatus *PubStatus        `xml:"pubstatus,omitempty" json:"pubstatus,omitempty" asn1:"optional"`
	History   PubStatusDateSet  `xml:"history,omitempty" json:"history,omitempty" asn1:"optional"`
}
type CitRetract struct {
	Type string `xml:"type" json:"type"` //Type,EnumList:retracted(1),notice(2),in-error(3),erratum(4)
	Exp  string `xml:"exp,omitempty" json:"exp,omitempty" asn1:"optional"`
}
type Meeting struct {
	Number string            `xml:"number" json:"number"`
	Date   *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty"`
	Place  *Affil            `xml:"place,omitempty" json:"place,omitempty" asn1:"optional"`
}
