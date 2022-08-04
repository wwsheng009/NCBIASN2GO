package NCBIFeatDef

type FeatDef struct {
	Typelabel    string       `xml:"typelabel" json:"typelabel"`
	Menulabel    string       `xml:"menulabel" json:"menulabel"`
	FeatdefKey   int64        `xml:"featdef-key" json:"featdef_key"`
	SeqfeatKey   int64        `xml:"seqfeat-key" json:"seqfeat_key"`
	Entrygroup   int64        `xml:"entrygroup" json:"entrygroup"`
	Displaygroup int64        `xml:"displaygroup" json:"displaygroup"`
	Molgroup     *FeatMolType `xml:"molgroup,omitempty" json:"molgroup,omitempty"`
}

//FeatMolType,EnumList:aa(1),na(2),both(3)
type FeatMolType string

type FeatDefSet []FeatDef
type FeatDispGroup struct {
	Groupkey  int64  `xml:"groupkey" json:"groupkey"`
	Groupname string `xml:"groupname" json:"groupname"`
}
type FeatDispGroupSet []FeatDispGroup
type FeatDefGroupSet struct {
	Groups FeatDispGroupSet `xml:"groups,omitempty" json:"groups,omitempty"`
	Defs   FeatDefSet       `xml:"defs,omitempty" json:"defs,omitempty"`
}
