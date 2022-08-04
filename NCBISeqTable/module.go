package NCBISeqTable

import "ncbiasn/NCBISeqloc"

type SeqTableColumnInfo struct {
	Title     string `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	FieldId   int    `xml:"field-id,omitempty" json:"field_id,omitempty" asn1:"optional"` //FieldId,IntegerEnum:location(0),location-id(1),location-gi(2),location-from(3),location-to(4),location-strand(5),location-fuzz-from-lim(6),location-fuzz-to-lim(7),product(10),product-id(11),product-gi(12),product-from(13),product-to(14),product-strand(15),product-fuzz-from-lim(16),product-fuzz-to-lim(17),id-local(20),xref-id-local(21),partial(22),comment(23),title(24),ext(25),qual(26),dbxref(27),data-imp-key(30),data-region(31),data-cdregion-frame(32),ext-type(40),qual-qual(41),qual-val(42),dbxref-db(43),dbxref-tag(44)
	FieldName string `xml:"field-name,omitempty" json:"field_name,omitempty" asn1:"optional"`
}
type CommonStringTable struct {
	Strings []string `xml:"strings,omitempty" json:"strings,omitempty"`
	Indexes []int64  `xml:"indexes,omitempty" json:"indexes,omitempty"`
}
type CommonBytesTable struct {
	Bytes   []byte  `xml:"bytes,omitempty" json:"bytes,omitempty"`
	Indexes []int64 `xml:"indexes,omitempty" json:"indexes,omitempty"`
}

//SeqTableMultiData,ChoiceOption
type SeqTableMultiData struct {
	Int          []int64                  `xml:"int,omitempty" json:"int,omitempty"`
	Real         []float64                `xml:"real,omitempty" json:"real,omitempty"`
	String       []string                 `xml:"string,omitempty" json:"string,omitempty"`
	Bytes        []byte                   `xml:"bytes,omitempty" json:"bytes,omitempty"`
	CommonString *CommonStringTable       `xml:"common-string,omitempty" json:"common_string,omitempty"`
	CommonBytes  *CommonBytesTable        `xml:"common-bytes,omitempty" json:"common_bytes,omitempty"`
	Bit          []byte                   `xml:"bit,omitempty" json:"bit,omitempty"`
	Loc          []NCBISeqloc.SeqLoc      `xml:"loc,omitempty" json:"loc,omitempty"`
	Id           []NCBISeqloc.SeqId       `xml:"id,omitempty" json:"id,omitempty"`
	Interval     []NCBISeqloc.SeqInterval `xml:"interval,omitempty" json:"interval,omitempty"`
}

//SeqTableSingleData,ChoiceOption
type SeqTableSingleData struct {
	Int      int64                   `xml:"int,omitempty" json:"int,omitempty"`
	Real     float64                 `xml:"real,omitempty" json:"real,omitempty"`
	String   string                  `xml:"string,omitempty" json:"string,omitempty" asn1:"utf8"`
	Bytes    []byte                  `xml:"bytes,omitempty" json:"bytes,omitempty"`
	Bit      bool                    `xml:"bit,omitempty" json:"bit,omitempty"`
	Loc      *NCBISeqloc.SeqLoc      `xml:"loc,omitempty" json:"loc,omitempty"`
	Id       *NCBISeqloc.SeqId       `xml:"id,omitempty" json:"id,omitempty"`
	Interval *NCBISeqloc.SeqInterval `xml:"interval,omitempty" json:"interval,omitempty"`
}

//SeqTableSparseIndex,ChoiceOption
type SeqTableSparseIndex struct {
	Indexes []int64 `xml:"indexes,omitempty" json:"indexes,omitempty"`
	BitSet  []byte  `xml:"bit-set,omitempty" json:"bit_set,omitempty"`
}

type SeqTableColumn struct {
	Header      *SeqTableColumnInfo  `xml:"header,omitempty" json:"header,omitempty"`
	Data        *SeqTableMultiData   `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
	Sparse      *SeqTableSparseIndex `xml:"sparse,omitempty" json:"sparse,omitempty" asn1:"optional"`
	Default     *SeqTableSingleData  `xml:"default,omitempty" json:"default,omitempty" asn1:"optional"`
	SparseOther *SeqTableSingleData  `xml:"sparse-other,omitempty" json:"sparse_other,omitempty" asn1:"optional"`
}
type SeqTable struct {
	FeatType    int64            `xml:"feat-type" json:"feat_type"`
	FeatSubtype int64            `xml:"feat-subtype,omitempty" json:"feat_subtype,omitempty" asn1:"optional"`
	NumRows     int64            `xml:"num-rows" json:"num_rows"`
	Columns     []SeqTableColumn `xml:"columns,omitempty" json:"columns,omitempty"`
}
