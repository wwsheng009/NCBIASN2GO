package INSDINSDSeq

type INSDSet []INSDSeq
type INSDSeq struct {
	Locus               string              `xml:"locus,omitempty" json:"locus,omitempty" asn1:"optional"`
	Length              int64               `xml:"length" json:"length"`
	Strandedness        string              `xml:"strandedness,omitempty" json:"strandedness,omitempty" asn1:"optional"`
	Moltype             string              `xml:"moltype" json:"moltype"`
	Topology            string              `xml:"topology,omitempty" json:"topology,omitempty" asn1:"optional"`
	Division            string              `xml:"division,omitempty" json:"division,omitempty" asn1:"optional"`
	UpdateDate          string              `xml:"update-date,omitempty" json:"update_date,omitempty" asn1:"optional"`
	CreateDate          string              `xml:"create-date,omitempty" json:"create_date,omitempty" asn1:"optional"`
	UpdateRelease       string              `xml:"update-release,omitempty" json:"update_release,omitempty" asn1:"optional"`
	CreateRelease       string              `xml:"create-release,omitempty" json:"create_release,omitempty" asn1:"optional"`
	Definition          string              `xml:"definition,omitempty" json:"definition,omitempty" asn1:"optional"`
	PrimaryAccession    string              `xml:"primary-accession,omitempty" json:"primary_accession,omitempty" asn1:"optional"`
	EntryVersion        string              `xml:"entry-version,omitempty" json:"entry_version,omitempty" asn1:"optional"`
	AccessionVersion    string              `xml:"accession-version,omitempty" json:"accession_version,omitempty" asn1:"optional"`
	OtherSeqids         []INSDSeqid         `xml:"other-seqids,omitempty" json:"other_seqids,omitempty" asn1:"optional"`
	SecondaryAccessions []INSDSecondaryAccn `xml:"secondary-accessions,omitempty" json:"secondary_accessions,omitempty" asn1:"optional"`
	Project             string              `xml:"project,omitempty" json:"project,omitempty" asn1:"optional"`
	Keywords            []INSDKeyword       `xml:"keywords,omitempty" json:"keywords,omitempty" asn1:"optional"`
	Segment             string              `xml:"segment,omitempty" json:"segment,omitempty" asn1:"optional"`
	Source              string              `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"`
	Organism            string              `xml:"organism,omitempty" json:"organism,omitempty" asn1:"optional"`
	Taxonomy            string              `xml:"taxonomy,omitempty" json:"taxonomy,omitempty" asn1:"optional"`
	References          []INSDReference     `xml:"references,omitempty" json:"references,omitempty" asn1:"optional"`
	Comment             string              `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	CommentSet          []INSDComment       `xml:"comment-set,omitempty" json:"comment_set,omitempty" asn1:"optional"`
	StrucComments       []INSDStrucComment  `xml:"struc-comments,omitempty" json:"struc_comments,omitempty" asn1:"optional"`
	Primary             string              `xml:"primary,omitempty" json:"primary,omitempty" asn1:"optional"`
	SourceDb            string              `xml:"source-db,omitempty" json:"source_db,omitempty" asn1:"optional"`
	DatabaseReference   string              `xml:"database-reference,omitempty" json:"database_reference,omitempty" asn1:"optional"`
	FeatureTable        []INSDFeature       `xml:"feature-table,omitempty" json:"feature_table,omitempty" asn1:"optional"`
	FeatureSet          []INSDFeatureSet    `xml:"feature-set,omitempty" json:"feature_set,omitempty" asn1:"optional"`
	Sequence            string              `xml:"sequence,omitempty" json:"sequence,omitempty" asn1:"optional"`
	Contig              string              `xml:"contig,omitempty" json:"contig,omitempty" asn1:"optional"`
	AltSeq              []INSDAltSeqData    `xml:"alt-seq,omitempty" json:"alt_seq,omitempty" asn1:"optional"`
	Xrefs               []INSDXref          `xml:"xrefs,omitempty" json:"xrefs,omitempty" asn1:"optional"`
}
type INSDSeqid string
type INSDSecondaryAccn string
type INSDKeyword string
type INSDReference struct {
	Reference  string       `xml:"reference" json:"reference"`
	Position   string       `xml:"position,omitempty" json:"position,omitempty" asn1:"optional"`
	Authors    []INSDAuthor `xml:"authors,omitempty" json:"authors,omitempty" asn1:"optional"`
	Consortium string       `xml:"consortium,omitempty" json:"consortium,omitempty" asn1:"optional"`
	Title      string       `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	Journal    string       `xml:"journal" json:"journal"`
	Xref       []INSDXref   `xml:"xref,omitempty" json:"xref,omitempty" asn1:"optional"`
	Pubmed     int64        `xml:"pubmed,omitempty" json:"pubmed,omitempty" asn1:"optional"`
	Remark     string       `xml:"remark,omitempty" json:"remark,omitempty" asn1:"optional"`
}
type INSDAuthor string
type INSDXref struct {
	Dbname string `xml:"dbname" json:"dbname"`
	Id     string `xml:"id" json:"id"`
}
type INSDComment struct {
	Type       string                 `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"`
	Paragraphs []INSDCommentParagraph `xml:"paragraphs,omitempty" json:"paragraphs,omitempty"`
}
type INSDCommentParagraph string
type INSDStrucComment struct {
	Name  string                 `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Items []INSDStrucCommentItem `xml:"items,omitempty" json:"items,omitempty"`
}
type INSDStrucCommentItem struct {
	Tag   string `xml:"tag,omitempty" json:"tag,omitempty" asn1:"optional"`
	Value string `xml:"value,omitempty" json:"value,omitempty" asn1:"optional"`
	Url   string `xml:"url,omitempty" json:"url,omitempty" asn1:"optional"`
}
type INSDFeatureSet struct {
	AnnotSource string        `xml:"annot-source,omitempty" json:"annot_source,omitempty" asn1:"optional"`
	Features    []INSDFeature `xml:"features,omitempty" json:"features,omitempty"`
}
type INSDFeature struct {
	Key       string          `xml:"key" json:"key"`
	Location  string          `xml:"location" json:"location"`
	Intervals []INSDInterval  `xml:"intervals,omitempty" json:"intervals,omitempty" asn1:"optional"`
	Operator  string          `xml:"operator,omitempty" json:"operator,omitempty" asn1:"optional"`
	Partial5  bool            `xml:"partial5,omitempty" json:"partial5,omitempty" asn1:"optional"`
	Partial3  bool            `xml:"partial3,omitempty" json:"partial3,omitempty" asn1:"optional"`
	Quals     []INSDQualifier `xml:"quals,omitempty" json:"quals,omitempty" asn1:"optional"`
	Xrefs     []INSDXref      `xml:"xrefs,omitempty" json:"xrefs,omitempty" asn1:"optional"`
}
type INSDInterval struct {
	From      int64  `xml:"from,omitempty" json:"from,omitempty" asn1:"optional"`
	To        int64  `xml:"to,omitempty" json:"to,omitempty" asn1:"optional"`
	Point     int64  `xml:"point,omitempty" json:"point,omitempty" asn1:"optional"`
	Iscomp    bool   `xml:"iscomp,omitempty" json:"iscomp,omitempty" asn1:"optional"`
	Interbp   bool   `xml:"interbp,omitempty" json:"interbp,omitempty" asn1:"optional"`
	Accession string `xml:"accession" json:"accession"`
}
type INSDQualifier struct {
	Name  string `xml:"name" json:"name"`
	Value string `xml:"value,omitempty" json:"value,omitempty" asn1:"optional"`
}
type INSDAltSeqData struct {
	Name  string           `xml:"name" json:"name"`
	Items []INSDAltSeqItem `xml:"items,omitempty" json:"items,omitempty" asn1:"optional"`
}
type INSDAltSeqItem struct {
	Interval   *INSDInterval `xml:"interval,omitempty" json:"interval,omitempty" asn1:"optional"`
	Isgap      bool          `xml:"isgap,omitempty" json:"isgap,omitempty" asn1:"optional"`
	GapLength  int64         `xml:"gap-length,omitempty" json:"gap_length,omitempty" asn1:"optional"`
	GapType    string        `xml:"gap-type,omitempty" json:"gap_type,omitempty" asn1:"optional"`
	GapLinkage string        `xml:"gap-linkage,omitempty" json:"gap_linkage,omitempty" asn1:"optional"`
	GapComment string        `xml:"gap-comment,omitempty" json:"gap_comment,omitempty" asn1:"optional"`
	FirstAccn  string        `xml:"first-accn,omitempty" json:"first_accn,omitempty" asn1:"optional"`
	LastAccn   string        `xml:"last-accn,omitempty" json:"last_accn,omitempty" asn1:"optional"`
	Value      string        `xml:"value,omitempty" json:"value,omitempty" asn1:"optional"`
}
