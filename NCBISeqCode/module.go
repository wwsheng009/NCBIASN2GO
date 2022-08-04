package NCBISeqCode

type SeqCodeType string

//SeqCodeType,EnumList:iupacna(1),iupacaa(2),ncbi2na(3),ncbi4na(4),ncbi8na(5),ncbipna(6),ncbi8aa(7),ncbieaa(8),ncbipaa(9),iupacaa3(10),ncbistdaa(11)
type SeqMapTable struct {
	From    *SeqCodeType `xml:"from,omitempty" json:"from,omitempty"`
	To      *SeqCodeType `xml:"to,omitempty" json:"to,omitempty"`
	Num     int64        `xml:"num" json:"num"`
	StartAt int64        `xml:"start-at" json:"start_at" asn1:"default:0"`
	Table   []int64      `xml:"table,omitempty" json:"table,omitempty"`
}
type SeqCodeTable struct {
	Code      *SeqCodeType `xml:"code,omitempty" json:"code,omitempty"`
	Num       int64        `xml:"num" json:"num"`
	OneLetter bool         `xml:"one-letter" json:"one_letter"`
	StartAt   int64        `xml:"start-at" json:"start_at" asn1:"default:0"`
	Table     []struct {
		Symbol string `xml:"symbol" json:"symbol"`
		Name   string `xml:"name" json:"name"`
	} `xml:"table,omitempty" json:"table,omitempty"`
	Comps []int64 `xml:"comps,omitempty" json:"comps,omitempty" asn1:"optional"`
}
type SeqCodeSet struct {
	Codes []SeqCodeTable `xml:"codes,omitempty" json:"codes,omitempty" asn1:"optional"`
	Maps  []SeqMapTable  `xml:"maps,omitempty" json:"maps,omitempty" asn1:"optional"`
}
