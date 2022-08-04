package NCBIEntrez2

type Entrez2Dt int64
type Entrez2DbId string
type Entrez2FieldId string
type Entrez2LinkId string
type Entrez2IdList struct {
	Db   Entrez2DbId `xml:"db,omitempty" json:"db,omitempty"`
	Num  int64       `xml:"num" json:"num"`
	Uids []byte      `xml:"uids,omitempty" json:"uids,omitempty" asn1:"optional"`
}
type Entrez2BooleanExp struct {
	Db     Entrez2DbId             `xml:"db,omitempty" json:"db,omitempty"`
	Exp    []Entrez2BooleanElement `xml:"exp,omitempty" json:"exp,omitempty"`
	Limits *Entrez2Limits          `xml:"limits,omitempty" json:"limits,omitempty" asn1:"optional"`
}

//Entrez2BooleanElement,ChoiceOption
type Entrez2BooleanElement struct {
	Str  string              `xml:"str,omitempty" json:"str,omitempty"`
	Op   *Entrez2Operator    `xml:"op,omitempty" json:"op,omitempty"`
	Term *Entrez2BooleanTerm `xml:"term,omitempty" json:"term,omitempty"`
	Ids  *Entrez2IdList      `xml:"ids,omitempty" json:"ids,omitempty"`
	Key  string              `xml:"key,omitempty" json:"key,omitempty"`
}

type Entrez2BooleanTerm struct {
	Field          Entrez2FieldId `xml:"field,omitempty" json:"field,omitempty"`
	Term           string         `xml:"term" json:"term"`
	TermCount      int64          `xml:"term-count,omitempty" json:"term_count,omitempty" asn1:"optional"`
	DoNotExplode   bool           `xml:"do-not-explode" json:"do_not_explode"`
	DoNotTranslate bool           `xml:"do-not-translate" json:"do_not_translate"`
}

//Entrez2Operator,IntegerEnum:and(1),or(2),butnot(3),range(4),left-paren(5),right-paren(6)
type Entrez2Operator int

type Entrez2Request struct {
	Request    *E2Request `xml:"request,omitempty" json:"request,omitempty"`
	Version    int64      `xml:"version" json:"version"`
	Tool       string     `xml:"tool,omitempty" json:"tool,omitempty" asn1:"optional"`
	Cookie     string     `xml:"cookie,omitempty" json:"cookie,omitempty" asn1:"optional"`
	UseHistory bool       `xml:"use-history" json:"use_history"`
}

//E2Request,ChoiceOption
type E2Request struct {
	GetInfo          interface{}         `xml:"-" json:"-"` //GetInfo,NullType
	EvalBoolean      *Entrez2EvalBoolean `xml:"eval-boolean,omitempty" json:"eval_boolean,omitempty"`
	GetDocsum        *Entrez2IdList      `xml:"get-docsum,omitempty" json:"get_docsum,omitempty"`
	GetTermPos       *Entrez2TermQuery   `xml:"get-term-pos,omitempty" json:"get_term_pos,omitempty"`
	GetTermList      *Entrez2TermPos     `xml:"get-term-list,omitempty" json:"get_term_list,omitempty"`
	GetTermHierarchy *Entrez2HierQuery   `xml:"get-term-hierarchy,omitempty" json:"get_term_hierarchy,omitempty"`
	GetLinks         *Entrez2GetLinks    `xml:"get-links,omitempty" json:"get_links,omitempty"`
	GetLinked        *Entrez2GetLinks    `xml:"get-linked,omitempty" json:"get_linked,omitempty"`
	GetLinkCounts    *Entrez2Id          `xml:"get-link-counts,omitempty" json:"get_link_counts,omitempty"`
}

type Entrez2EvalBoolean struct {
	ReturnUIDs  bool               `xml:"return-UIDs" json:"return_UIDs"`
	ReturnParse bool               `xml:"return-parse" json:"return_parse"`
	Query       *Entrez2BooleanExp `xml:"query,omitempty" json:"query,omitempty"`
}
type Entrez2DtFilter struct {
	BeginDate *Entrez2Dt     `xml:"begin-date,omitempty" json:"begin_date,omitempty"`
	EndDate   *Entrez2Dt     `xml:"end-date,omitempty" json:"end_date,omitempty"`
	TypeDate  Entrez2FieldId `xml:"type-date,omitempty" json:"type_date,omitempty"`
}
type Entrez2Limits struct {
	FilterDate *Entrez2DtFilter `xml:"filter-date,omitempty" json:"filter_date,omitempty" asn1:"optional"`
	MaxUIDs    int64            `xml:"max-UIDs,omitempty" json:"max_UIDs,omitempty" asn1:"optional"`
	OffsetUIDs int64            `xml:"offset-UIDs,omitempty" json:"offset_UIDs,omitempty" asn1:"optional"`
}
type Entrez2Id struct {
	Db  Entrez2DbId `xml:"db,omitempty" json:"db,omitempty"`
	Uid int64       `xml:"uid" json:"uid"`
}
type Entrez2TermQuery struct {
	Db    Entrez2DbId    `xml:"db,omitempty" json:"db,omitempty"`
	Field Entrez2FieldId `xml:"field,omitempty" json:"field,omitempty"`
	Term  string         `xml:"term" json:"term"`
}
type Entrez2HierQuery struct {
	Db    Entrez2DbId    `xml:"db,omitempty" json:"db,omitempty"`
	Field Entrez2FieldId `xml:"field,omitempty" json:"field,omitempty"`
	Term  string         `xml:"term,omitempty" json:"term,omitempty" asn1:"optional"`
	Txid  int64          `xml:"txid,omitempty" json:"txid,omitempty" asn1:"optional"`
}
type Entrez2TermPos struct {
	Db            Entrez2DbId    `xml:"db,omitempty" json:"db,omitempty"`
	Field         Entrez2FieldId `xml:"field,omitempty" json:"field,omitempty"`
	FirstTermPos  int64          `xml:"first-term-pos" json:"first_term_pos"`
	NumberOfTerms int64          `xml:"number-of-terms,omitempty" json:"number_of_terms,omitempty" asn1:"optional"`
}
type Entrez2GetLinks struct {
	Uids           *Entrez2IdList `xml:"uids,omitempty" json:"uids,omitempty"`
	Linktype       Entrez2LinkId  `xml:"linktype,omitempty" json:"linktype,omitempty"`
	MaxUIDS        int64          `xml:"max-UIDS,omitempty" json:"max_UIDS,omitempty" asn1:"optional"`
	CountOnly      bool           `xml:"count-only,omitempty" json:"count_only,omitempty" asn1:"optional"`
	ParentsPersist bool           `xml:"parents-persist,omitempty" json:"parents_persist,omitempty" asn1:"optional"`
}
type Entrez2Reply struct {
	Reply  *E2Reply   `xml:"reply,omitempty" json:"reply,omitempty"`
	Dt     *Entrez2Dt `xml:"dt,omitempty" json:"dt,omitempty"`
	Server string     `xml:"server" json:"server"`
	Msg    string     `xml:"msg,omitempty" json:"msg,omitempty" asn1:"optional"`
	Key    string     `xml:"key,omitempty" json:"key,omitempty" asn1:"optional"`
	Cookie string     `xml:"cookie,omitempty" json:"cookie,omitempty" asn1:"optional"`
}

//E2Reply,ChoiceOption
type E2Reply struct {
	Error            string                `xml:"error,omitempty" json:"error,omitempty"`
	GetInfo          *Entrez2Info          `xml:"get-info,omitempty" json:"get_info,omitempty"`
	EvalBoolean      *Entrez2BooleanReply  `xml:"eval-boolean,omitempty" json:"eval_boolean,omitempty"`
	GetDocsum        *Entrez2DocsumList    `xml:"get-docsum,omitempty" json:"get_docsum,omitempty"`
	GetTermPos       int64                 `xml:"get-term-pos,omitempty" json:"get_term_pos,omitempty"`
	GetTermList      *Entrez2TermList      `xml:"get-term-list,omitempty" json:"get_term_list,omitempty"`
	GetTermHierarchy *Entrez2HierNode      `xml:"get-term-hierarchy,omitempty" json:"get_term_hierarchy,omitempty"`
	GetLinks         *Entrez2LinkSet       `xml:"get-links,omitempty" json:"get_links,omitempty"`
	GetLinked        *Entrez2IdList        `xml:"get-linked,omitempty" json:"get_linked,omitempty"`
	GetLinkCounts    *Entrez2LinkCountList `xml:"get-link-counts,omitempty" json:"get_link_counts,omitempty"`
}

type Entrez2Info struct {
	DbCount   int64           `xml:"db-count" json:"db_count"`
	BuildDate *Entrez2Dt      `xml:"build-date,omitempty" json:"build_date,omitempty"`
	DbInfo    []Entrez2DbInfo `xml:"db-info,omitempty" json:"db_info,omitempty"`
}
type Entrez2DbInfo struct {
	DbName           Entrez2DbId              `xml:"db-name,omitempty" json:"db_name,omitempty"`
	DbMenu           string                   `xml:"db-menu" json:"db_menu"`
	DbDescr          string                   `xml:"db-descr" json:"db_descr"`
	DocCount         int64                    `xml:"doc-count" json:"doc_count"`
	FieldCount       int64                    `xml:"field-count" json:"field_count"`
	Fields           []Entrez2FieldInfo       `xml:"fields,omitempty" json:"fields,omitempty"`
	LinkCount        int64                    `xml:"link-count" json:"link_count"`
	Links            []Entrez2LinkInfo        `xml:"links,omitempty" json:"links,omitempty"`
	DocsumFieldCount int64                    `xml:"docsum-field-count" json:"docsum_field_count"`
	DocsumFields     []Entrez2DocsumFieldInfo `xml:"docsum-fields,omitempty" json:"docsum_fields,omitempty"`
}
type Entrez2FieldInfo struct {
	FieldName      Entrez2FieldId `xml:"field-name,omitempty" json:"field_name,omitempty"`
	FieldMenu      string         `xml:"field-menu" json:"field_menu"`
	FieldDescr     string         `xml:"field-descr" json:"field_descr"`
	TermCount      int64          `xml:"term-count" json:"term_count"`
	IsDate         bool           `xml:"is-date,omitempty" json:"is_date,omitempty" asn1:"optional"`
	IsNumerical    bool           `xml:"is-numerical,omitempty" json:"is_numerical,omitempty" asn1:"optional"`
	SingleToken    bool           `xml:"single-token,omitempty" json:"single_token,omitempty" asn1:"optional"`
	HierarchyAvail bool           `xml:"hierarchy-avail,omitempty" json:"hierarchy_avail,omitempty" asn1:"optional"`
	IsRangable     bool           `xml:"is-rangable,omitempty" json:"is_rangable,omitempty" asn1:"optional"`
	IsTruncatable  bool           `xml:"is-truncatable,omitempty" json:"is_truncatable,omitempty" asn1:"optional"`
}
type Entrez2LinkInfo struct {
	LinkName  Entrez2LinkId `xml:"link-name,omitempty" json:"link_name,omitempty"`
	LinkMenu  string        `xml:"link-menu" json:"link_menu"`
	LinkDescr string        `xml:"link-descr" json:"link_descr"`
	DbTo      Entrez2DbId   `xml:"db-to,omitempty" json:"db_to,omitempty"`
	DataSize  int64         `xml:"data-size,omitempty" json:"data_size,omitempty" asn1:"optional"`
}

//Entrez2DocsumFieldType,IntegerEnum:string(1),int(2),float(3),date-pubmed(4)
type Entrez2DocsumFieldType int

type Entrez2DocsumFieldInfo struct {
	FieldName        string                  `xml:"field-name" json:"field_name"`
	FieldDescription string                  `xml:"field-description" json:"field_description"`
	FieldType        *Entrez2DocsumFieldType `xml:"field-type,omitempty" json:"field_type,omitempty"`
}
type Entrez2BooleanReply struct {
	Count int64              `xml:"count" json:"count"`
	Uids  *Entrez2IdList     `xml:"uids,omitempty" json:"uids,omitempty" asn1:"optional"`
	Query *Entrez2BooleanExp `xml:"query,omitempty" json:"query,omitempty" asn1:"optional"`
}
type Entrez2DocsumList struct {
	Count int64           `xml:"count" json:"count"`
	List  []Entrez2Docsum `xml:"list,omitempty" json:"list,omitempty"`
}
type Entrez2Docsum struct {
	Uid        int64               `xml:"uid" json:"uid"`
	DocsumData []Entrez2DocsumData `xml:"docsum-data,omitempty" json:"docsum_data,omitempty"`
}
type Entrez2DocsumData struct {
	FieldName  string `xml:"field-name" json:"field_name"`
	FieldValue string `xml:"field-value" json:"field_value"`
}
type Entrez2TermList struct {
	Pos  int64         `xml:"pos" json:"pos"`
	Num  int64         `xml:"num" json:"num"`
	List []Entrez2Term `xml:"list,omitempty" json:"list,omitempty"`
}
type Entrez2Term struct {
	Term       string `xml:"term" json:"term"`
	Txid       int64  `xml:"txid,omitempty" json:"txid,omitempty" asn1:"optional"`
	Count      int64  `xml:"count" json:"count"`
	IsLeafNode bool   `xml:"is-leaf-node,omitempty" json:"is_leaf_node,omitempty" asn1:"optional"`
}
type Entrez2HierNode struct {
	CannonicalForm string        `xml:"cannonical-form" json:"cannonical_form"`
	LineageCount   int64         `xml:"lineage-count" json:"lineage_count"`
	Lineage        []Entrez2Term `xml:"lineage,omitempty" json:"lineage,omitempty" asn1:"optional"`
	ChildCount     int64         `xml:"child-count" json:"child_count"`
	Children       []Entrez2Term `xml:"children,omitempty" json:"children,omitempty"`
	IsAmbiguous    bool          `xml:"is-ambiguous,omitempty" json:"is_ambiguous,omitempty" asn1:"optional"`
}
type Entrez2LinkSet struct {
	Ids      *Entrez2IdList `xml:"ids,omitempty" json:"ids,omitempty"`
	DataSize int64          `xml:"data-size,omitempty" json:"data_size,omitempty" asn1:"optional"`
	Data     []byte         `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
}
type Entrez2LinkCountList struct {
	LinkTypeCount int64              `xml:"link-type-count" json:"link_type_count"`
	Links         []Entrez2LinkCount `xml:"links,omitempty" json:"links,omitempty"`
}
type Entrez2LinkCount struct {
	LinkType  Entrez2LinkId `xml:"link-type,omitempty" json:"link_type,omitempty"`
	LinkCount int64         `xml:"link-count" json:"link_count"`
}
