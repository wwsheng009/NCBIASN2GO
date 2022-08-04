package NCBIBlastOutput

type BlastOutput struct {
	Program    string      `xml:"program" json:"program"`
	Version    string      `xml:"version" json:"version"`
	Reference  string      `xml:"reference" json:"reference"`
	Db         string      `xml:"db" json:"db"`
	QueryID    string      `xml:"query-ID" json:"query_ID"`
	QueryDef   string      `xml:"query-def" json:"query_def"`
	QueryLen   int64       `xml:"query-len" json:"query_len"`
	QuerySeq   string      `xml:"query-seq,omitempty" json:"query_seq,omitempty" asn1:"optional"`
	Param      *Parameters `xml:"param,omitempty" json:"param,omitempty"`
	Iterations []Iteration `xml:"iterations,omitempty" json:"iterations,omitempty"`
	Mbstat     *Statistics `xml:"mbstat,omitempty" json:"mbstat,omitempty" asn1:"optional"`
}
type Iteration struct {
	IterNum  int64       `xml:"iter-num" json:"iter_num"`
	QueryID  string      `xml:"query-ID,omitempty" json:"query_ID,omitempty" asn1:"optional"`
	QueryDef string      `xml:"query-def,omitempty" json:"query_def,omitempty" asn1:"optional"`
	QueryLen int64       `xml:"query-len,omitempty" json:"query_len,omitempty" asn1:"optional"`
	Hits     []Hit       `xml:"hits,omitempty" json:"hits,omitempty" asn1:"optional"`
	Stat     *Statistics `xml:"stat,omitempty" json:"stat,omitempty" asn1:"optional"`
	Message  string      `xml:"message,omitempty" json:"message,omitempty" asn1:"optional"`
}
type Parameters struct {
	Matrix      string  `xml:"matrix,omitempty" json:"matrix,omitempty" asn1:"optional"`
	Expect      float64 `xml:"expect" json:"expect"`
	Include     float64 `xml:"include,omitempty" json:"include,omitempty" asn1:"optional"`
	ScMatch     int64   `xml:"sc-match,omitempty" json:"sc_match,omitempty" asn1:"optional"`
	ScMismatch  int64   `xml:"sc-mismatch,omitempty" json:"sc_mismatch,omitempty" asn1:"optional"`
	GapOpen     int64   `xml:"gap-open" json:"gap_open"`
	GapExtend   int64   `xml:"gap-extend" json:"gap_extend"`
	Filter      string  `xml:"filter,omitempty" json:"filter,omitempty" asn1:"optional"`
	Pattern     string  `xml:"pattern,omitempty" json:"pattern,omitempty" asn1:"optional"`
	EntrezQuery string  `xml:"entrez-query,omitempty" json:"entrez_query,omitempty" asn1:"optional"`
}
type Statistics struct {
	DbNum    int64   `xml:"db-num" json:"db_num"`
	DbLen    *int64  `xml:"db-len,omitempty" json:"db_len,omitempty"`
	HspLen   int64   `xml:"hsp-len" json:"hsp_len"`
	EffSpace float64 `xml:"eff-space" json:"eff_space"`
	Kappa    float64 `xml:"kappa" json:"kappa"`
	Lambda   float64 `xml:"lambda" json:"lambda"`
	Entropy  float64 `xml:"entropy" json:"entropy"`
}
type Hit struct {
	Num       int64  `xml:"num" json:"num"`
	Id        string `xml:"id" json:"id"`
	Def       string `xml:"def" json:"def"`
	Accession string `xml:"accession" json:"accession"`
	Len       int64  `xml:"len" json:"len"`
	Hsps      []Hsp  `xml:"hsps,omitempty" json:"hsps,omitempty" asn1:"optional"`
}
type Hsp struct {
	Num         int64   `xml:"num" json:"num"`
	BitScore    float64 `xml:"bit-score" json:"bit_score"`
	Score       float64 `xml:"score" json:"score"`
	Evalue      float64 `xml:"evalue" json:"evalue"`
	QueryFrom   int64   `xml:"query-from" json:"query_from"`
	QueryTo     int64   `xml:"query-to" json:"query_to"`
	HitFrom     int64   `xml:"hit-from" json:"hit_from"`
	HitTo       int64   `xml:"hit-to" json:"hit_to"`
	PatternFrom int64   `xml:"pattern-from,omitempty" json:"pattern_from,omitempty" asn1:"optional"`
	PatternTo   int64   `xml:"pattern-to,omitempty" json:"pattern_to,omitempty" asn1:"optional"`
	QueryFrame  int64   `xml:"query-frame,omitempty" json:"query_frame,omitempty" asn1:"optional"`
	HitFrame    int64   `xml:"hit-frame,omitempty" json:"hit_frame,omitempty" asn1:"optional"`
	Identity    int64   `xml:"identity,omitempty" json:"identity,omitempty" asn1:"optional"`
	Positive    int64   `xml:"positive,omitempty" json:"positive,omitempty" asn1:"optional"`
	Gaps        int64   `xml:"gaps,omitempty" json:"gaps,omitempty" asn1:"optional"`
	AlignLen    int64   `xml:"align-len,omitempty" json:"align_len,omitempty" asn1:"optional"`
	Density     int64   `xml:"density,omitempty" json:"density,omitempty" asn1:"optional"`
	Qseq        string  `xml:"qseq" json:"qseq"`
	Hseq        string  `xml:"hseq" json:"hseq"`
	Midline     string  `xml:"midline,omitempty" json:"midline,omitempty" asn1:"optional"`
}
