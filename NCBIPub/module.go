package NCBIPub

import (
	"ncbiasn/NCBIBiblio"
	"ncbiasn/NCBIMedline"
)

//Pub,ChoiceOption
type Pub struct {
	Gen     *NCBIBiblio.CitGen        `xml:"gen,omitempty" json:"gen,omitempty"`
	Sub     *NCBIBiblio.CitSub        `xml:"sub,omitempty" json:"sub,omitempty"`
	Medline *NCBIMedline.MedlineEntry `xml:"medline,omitempty" json:"medline,omitempty"`
	Muid    int64                     `xml:"muid,omitempty" json:"muid,omitempty"`
	Article *NCBIBiblio.CitArt        `xml:"article,omitempty" json:"article,omitempty"`
	Journal *NCBIBiblio.CitJour       `xml:"journal,omitempty" json:"journal,omitempty"`
	Book    *NCBIBiblio.CitBook       `xml:"book,omitempty" json:"book,omitempty"`
	Proc    *NCBIBiblio.CitProc       `xml:"proc,omitempty" json:"proc,omitempty"`
	Patent  *NCBIBiblio.CitPat        `xml:"patent,omitempty" json:"patent,omitempty"`
	PatId   *NCBIBiblio.IdPat         `xml:"pat-id,omitempty" json:"pat_id,omitempty"`
	Man     *NCBIBiblio.CitLet        `xml:"man,omitempty" json:"man,omitempty"`
	Equiv   PubEquiv                  `xml:"equiv,omitempty" json:"equiv,omitempty"`
	Pmid    *NCBIBiblio.PubMedId      `xml:"pmid,omitempty" json:"pmid,omitempty"`
}

type PubEquiv []Pub

//PubSet,ChoiceOption
type PubSet struct {
	Pub     []Pub                      `xml:"pub,omitempty" json:"pub,omitempty"`
	Medline []NCBIMedline.MedlineEntry `xml:"medline,omitempty" json:"medline,omitempty"`
	Article []NCBIBiblio.CitArt        `xml:"article,omitempty" json:"article,omitempty"`
	Journal []NCBIBiblio.CitJour       `xml:"journal,omitempty" json:"journal,omitempty"`
	Book    []NCBIBiblio.CitBook       `xml:"book,omitempty" json:"book,omitempty"`
	Proc    []NCBIBiblio.CitProc       `xml:"proc,omitempty" json:"proc,omitempty"`
	Patent  []NCBIBiblio.CitPat        `xml:"patent,omitempty" json:"patent,omitempty"`
}
