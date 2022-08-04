package NCBIRsite

import "ncbiasn/NCBIGeneral"

//RsiteRef,ChoiceOption
type RsiteRef struct {
	Str string             `xml:"str,omitempty" json:"str,omitempty"`
	Db  *NCBIGeneral.Dbtag `xml:"db,omitempty" json:"db,omitempty"`
}
