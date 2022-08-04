package NCBISubmit

import (
	"ncbiasn/NCBIBiblio"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBISeqloc"
	"ncbiasn/NCBISequence"
)

type SeqSubmit struct {
	Sub  *SubmitBlock `xml:"sub,omitempty" json:"sub,omitempty"`
	Data struct {
		Entrys []NCBISequence.SeqEntry `xml:"entrys,omitempty" json:"entrys,omitempty"`
		Annots []NCBISequence.SeqAnnot `xml:"annots,omitempty" json:"annots,omitempty"`
		Delete []NCBISeqloc.SeqId      `xml:"delete,omitempty" json:"delete,omitempty"`
	} `xml:"data" json:"data"` //Data,ChoiceOption
}
type SubmitBlock struct {
	Contact *ContactInfo       `xml:"contact,omitempty" json:"contact,omitempty"`
	Cit     *NCBIBiblio.CitSub `xml:"cit,omitempty" json:"cit,omitempty"`
	Hup     bool               `xml:"hup" json:"hup"`
	Reldate *NCBIGeneral.Date  `xml:"reldate,omitempty" json:"reldate,omitempty" asn1:"optional"`
	Subtype int                `xml:"subtype,omitempty" json:"subtype,omitempty" asn1:"optional"` //Subtype,IntegerEnum:new(1),update(2),revision(3),other(255)
	Tool    string             `xml:"tool,omitempty" json:"tool,omitempty" asn1:"optional"`
	UserTag string             `xml:"user-tag,omitempty" json:"user_tag,omitempty" asn1:"optional"`
	Comment string             `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
}
type ContactInfo struct {
	Name          string                `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Address       []string              `xml:"address,omitempty" json:"address,omitempty" asn1:"optional"`
	Phone         string                `xml:"phone,omitempty" json:"phone,omitempty" asn1:"optional"`
	Fax           string                `xml:"fax,omitempty" json:"fax,omitempty" asn1:"optional"`
	Email         string                `xml:"email,omitempty" json:"email,omitempty" asn1:"optional"`
	Telex         string                `xml:"telex,omitempty" json:"telex,omitempty" asn1:"optional"`
	OwnerId       *NCBIGeneral.ObjectId `xml:"owner-id,omitempty" json:"owner_id,omitempty" asn1:"optional"`
	Password      []byte                `xml:"password,omitempty" json:"password,omitempty" asn1:"optional"`
	LastName      string                `xml:"last-name,omitempty" json:"last_name,omitempty" asn1:"optional"`
	FirstName     string                `xml:"first-name,omitempty" json:"first_name,omitempty" asn1:"optional"`
	MiddleInitial string                `xml:"middle-initial,omitempty" json:"middle_initial,omitempty" asn1:"optional"`
	Contact       *NCBIBiblio.Author    `xml:"contact,omitempty" json:"contact,omitempty" asn1:"optional"`
}
