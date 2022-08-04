package NCBIBlastOutput2

type BlastOutput2 struct {
	Report *Report `xml:"report,omitempty" json:"report,omitempty" asn1:"optional"`
	Error  *Err    `xml:"error,omitempty" json:"error,omitempty" asn1:"optional"`
}
type BlastXML2 []BlastOutput2
type Report struct {
	Program      string      `xml:"program" json:"program"`
	Version      string      `xml:"version" json:"version"`
	Reference    string      `xml:"reference" json:"reference"`
	SearchTarget *Target     `xml:"search-target,omitempty" json:"search_target,omitempty"`
	Params       *Parameters `xml:"params,omitempty" json:"params,omitempty"`
	Results      *Results    `xml:"results,omitempty" json:"results,omitempty"`
}
type Err struct {
	Code    int64  `xml:"code" json:"code"`
	Message string `xml:"message,omitempty" json:"message,omitempty" asn1:"optional"`
}

//Target,ChoiceOption
type Target struct {
	Db       string   `xml:"db,omitempty" json:"db,omitempty"`
	Subjects []string `xml:"subjects,omitempty" json:"subjects,omitempty"`
}

//Results,ChoiceOption
type Results struct {
	Iterations []Iteration `xml:"iterations,omitempty" json:"iterations,omitempty"`
	Search     *Search     `xml:"search,omitempty" json:"search,omitempty"`
	Bl2seq     []Search    `xml:"bl2seq,omitempty" json:"bl2seq,omitempty"`
}

type Iteration struct {
	IterNum int64   `xml:"iter-num" json:"iter_num"`
	Search  *Search `xml:"search,omitempty" json:"search,omitempty"`
}
type Search struct {
	QueryId      string      `xml:"query-id,omitempty" json:"query_id,omitempty" asn1:"optional"`
	QueryTitle   string      `xml:"query-title,omitempty" json:"query_title,omitempty" asn1:"optional"`
	QueryLen     int64       `xml:"query-len,omitempty" json:"query_len,omitempty" asn1:"optional"`
	QueryMasking []Range     `xml:"query-masking,omitempty" json:"query_masking,omitempty" asn1:"optional"`
	Hits         []Hit       `xml:"hits,omitempty" json:"hits,omitempty" asn1:"optional"`
	Stat         *Statistics `xml:"stat,omitempty" json:"stat,omitempty" asn1:"optional"`
	Message      string      `xml:"message,omitempty" json:"message,omitempty" asn1:"optional"`
}
type Parameters struct {
	Matrix       string  `xml:"matrix,omitempty" json:"matrix,omitempty" asn1:"optional"`
	Expect       float64 `xml:"expect" json:"expect"`
	Include      float64 `xml:"include,omitempty" json:"include,omitempty" asn1:"optional"`
	ScMatch      int64   `xml:"sc-match,omitempty" json:"sc_match,omitempty" asn1:"optional"`
	ScMismatch   int64   `xml:"sc-mismatch,omitempty" json:"sc_mismatch,omitempty" asn1:"optional"`
	GapOpen      int64   `xml:"gap-open,omitempty" json:"gap_open,omitempty" asn1:"optional"`
	GapExtend    int64   `xml:"gap-extend,omitempty" json:"gap_extend,omitempty" asn1:"optional"`
	Filter       string  `xml:"filter,omitempty" json:"filter,omitempty" asn1:"optional"`
	Pattern      string  `xml:"pattern,omitempty" json:"pattern,omitempty" asn1:"optional"`
	EntrezQuery  string  `xml:"entrez-query,omitempty" json:"entrez_query,omitempty" asn1:"optional"`
	Cbs          int64   `xml:"cbs,omitempty" json:"cbs,omitempty" asn1:"optional"`
	QueryGencode int64   `xml:"query-gencode,omitempty" json:"query_gencode,omitempty" asn1:"optional"`
	DbGencode    int64   `xml:"db-gencode,omitempty" json:"db_gencode,omitempty" asn1:"optional"`
	Bl2seqMode   string  `xml:"bl2seq-mode,omitempty" json:"bl2seq_mode,omitempty" asn1:"optional"`
}
type Range struct {
	From int64 `xml:"from" json:"from"`
	To   int64 `xml:"to" json:"to"`
}
type Statistics struct {
	DbNum    *int64  `xml:"db-num,omitempty" json:"db_num,omitempty" asn1:"optional"`
	DbLen    *int64  `xml:"db-len,omitempty" json:"db_len,omitempty" asn1:"optional"`
	HspLen   int64   `xml:"hsp-len" json:"hsp_len"`
	EffSpace *int64  `xml:"eff-space,omitempty" json:"eff_space,omitempty"`
	Kappa    float64 `xml:"kappa" json:"kappa"`
	Lambda   float64 `xml:"lambda" json:"lambda"`
	Entropy  float64 `xml:"entropy" json:"entropy"`
}
type HitDescr struct {
	Id        string `xml:"id" json:"id"`
	Accession string `xml:"accession,omitempty" json:"accession,omitempty" asn1:"optional"`
	Title     string `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	Taxid     int64  `xml:"taxid,omitempty" json:"taxid,omitempty" asn1:"optional"`
	Sciname   string `xml:"sciname,omitempty" json:"sciname,omitempty" asn1:"optional"`
}
type Hit struct {
	Num         int64      `xml:"num" json:"num"`
	Description []HitDescr `xml:"description,omitempty" json:"description,omitempty"`
	Len         int64      `xml:"len" json:"len"`
	Hsps        []Hsp      `xml:"hsps,omitempty" json:"hsps,omitempty" asn1:"optional"`
}
type Hsp struct {
	Num         int64   `xml:"num" json:"num"`
	BitScore    float64 `xml:"bit-score" json:"bit_score"`
	Score       float64 `xml:"score" json:"score"`
	Evalue      float64 `xml:"evalue" json:"evalue"`
	Identity    int64   `xml:"identity,omitempty" json:"identity,omitempty" asn1:"optional"`
	Positive    int64   `xml:"positive,omitempty" json:"positive,omitempty" asn1:"optional"`
	Density     int64   `xml:"density,omitempty" json:"density,omitempty" asn1:"optional"`
	PatternFrom int64   `xml:"pattern-from,omitempty" json:"pattern_from,omitempty" asn1:"optional"`
	PatternTo   int64   `xml:"pattern-to,omitempty" json:"pattern_to,omitempty" asn1:"optional"`
	QueryFrom   int64   `xml:"query-from" json:"query_from"`
	QueryTo     int64   `xml:"query-to" json:"query_to"`
	QueryStrand string  `xml:"query-strand,omitempty" json:"query_strand,omitempty" asn1:"optional"`
	QueryFrame  int64   `xml:"query-frame,omitempty" json:"query_frame,omitempty" asn1:"optional"`
	HitFrom     int64   `xml:"hit-from" json:"hit_from"`
	HitTo       int64   `xml:"hit-to" json:"hit_to"`
	HitStrand   string  `xml:"hit-strand,omitempty" json:"hit_strand,omitempty" asn1:"optional"`
	HitFrame    int64   `xml:"hit-frame,omitempty" json:"hit_frame,omitempty" asn1:"optional"`
	AlignLen    int64   `xml:"align-len,omitempty" json:"align_len,omitempty" asn1:"optional"`
	Gaps        int64   `xml:"gaps,omitempty" json:"gaps,omitempty" asn1:"optional"`
	Qseq        string  `xml:"qseq" json:"qseq"`
	Hseq        string  `xml:"hseq" json:"hseq"`
	Midline     string  `xml:"midline,omitempty" json:"midline,omitempty" asn1:"optional"`
}
