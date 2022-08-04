package NCBIMim

type MimEntries []MimEntry
type MimSet struct {
	ReleaseDate *MimDate   `xml:"releaseDate,omitempty" json:"releaseDate,omitempty"`
	MimEntries  []MimEntry `xml:"mimEntries,omitempty" json:"mimEntries,omitempty"`
}
type MimEntry struct {
	MimNumber            string              `xml:"mimNumber" json:"mimNumber"`
	MimType              int                 `xml:"mimType" json:"mimType"` //MimType,IntegerEnum:none(0),star(1),caret(2),pound(3),plus(4),perc(5)
	Title                string              `xml:"title" json:"title"`
	Copyright            string              `xml:"copyright,omitempty" json:"copyright,omitempty" asn1:"optional"`
	Symbol               string              `xml:"symbol,omitempty" json:"symbol,omitempty" asn1:"optional"`
	Locus                string              `xml:"locus,omitempty" json:"locus,omitempty" asn1:"optional"`
	Synonyms             []string            `xml:"synonyms,omitempty" json:"synonyms,omitempty" asn1:"optional"`
	Aliases              []string            `xml:"aliases,omitempty" json:"aliases,omitempty" asn1:"optional"`
	Included             []string            `xml:"included,omitempty" json:"included,omitempty" asn1:"optional"`
	SeeAlso              []MimCit            `xml:"seeAlso,omitempty" json:"seeAlso,omitempty" asn1:"optional"`
	Text                 []MimText           `xml:"text,omitempty" json:"text,omitempty" asn1:"optional"`
	Textfields           []MimText           `xml:"textfields,omitempty" json:"textfields,omitempty" asn1:"optional"`
	HasSummary           bool                `xml:"hasSummary,omitempty" json:"hasSummary,omitempty" asn1:"optional"`
	Summary              []MimText           `xml:"summary,omitempty" json:"summary,omitempty" asn1:"optional"`
	SummaryAttribution   []MimEditItem       `xml:"summaryAttribution,omitempty" json:"summaryAttribution,omitempty" asn1:"optional"`
	SummaryEditHistory   []MimEditItem       `xml:"summaryEditHistory,omitempty" json:"summaryEditHistory,omitempty" asn1:"optional"`
	SummaryCreationDate  *MimEditItem        `xml:"summaryCreationDate,omitempty" json:"summaryCreationDate,omitempty" asn1:"optional"`
	AllelicVariants      []MimAllelicVariant `xml:"allelicVariants,omitempty" json:"allelicVariants,omitempty" asn1:"optional"`
	HasSynopsis          bool                `xml:"hasSynopsis,omitempty" json:"hasSynopsis,omitempty" asn1:"optional"`
	ClinicalSynopsis     []MimIndexTerm      `xml:"clinicalSynopsis,omitempty" json:"clinicalSynopsis,omitempty" asn1:"optional"`
	SynopsisAttribution  []MimEditItem       `xml:"synopsisAttribution,omitempty" json:"synopsisAttribution,omitempty" asn1:"optional"`
	SynopsisEditHistory  []MimEditItem       `xml:"synopsisEditHistory,omitempty" json:"synopsisEditHistory,omitempty" asn1:"optional"`
	SynopsisCreationDate *MimEditItem        `xml:"synopsisCreationDate,omitempty" json:"synopsisCreationDate,omitempty" asn1:"optional"`
	EditHistory          []MimEditItem       `xml:"editHistory,omitempty" json:"editHistory,omitempty" asn1:"optional"`
	CreationDate         *MimEditItem        `xml:"creationDate,omitempty" json:"creationDate,omitempty" asn1:"optional"`
	References           []MimReference      `xml:"references,omitempty" json:"references,omitempty" asn1:"optional"`
	Attribution          []MimEditItem       `xml:"attribution,omitempty" json:"attribution,omitempty" asn1:"optional"`
	NumGeneMaps          int64               `xml:"numGeneMaps" json:"numGeneMaps"`
	MedlineLinks         *MimLink            `xml:"medlineLinks,omitempty" json:"medlineLinks,omitempty" asn1:"optional"`
	ProteinLinks         *MimLink            `xml:"proteinLinks,omitempty" json:"proteinLinks,omitempty" asn1:"optional"`
	NucleotideLinks      *MimLink            `xml:"nucleotideLinks,omitempty" json:"nucleotideLinks,omitempty" asn1:"optional"`
	StructureLinks       *MimLink            `xml:"structureLinks,omitempty" json:"structureLinks,omitempty" asn1:"optional"`
	GenomeLinks          *MimLink            `xml:"genomeLinks,omitempty" json:"genomeLinks,omitempty" asn1:"optional"`
}
type MimText struct {
	Label     string   `xml:"label" json:"label"`
	Text      string   `xml:"text" json:"text"`
	Neighbors *MimLink `xml:"neighbors,omitempty" json:"neighbors,omitempty" asn1:"optional"`
}
type MimAllelicVariant struct {
	Number      string    `xml:"number" json:"number"`
	Name        string    `xml:"name" json:"name"`
	Aliases     []string  `xml:"aliases,omitempty" json:"aliases,omitempty" asn1:"optional"`
	Mutation    []MimText `xml:"mutation,omitempty" json:"mutation,omitempty" asn1:"optional"`
	Description []MimText `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	SnpLinks    *MimLink  `xml:"snpLinks,omitempty" json:"snpLinks,omitempty" asn1:"optional"`
}
type MimLink struct {
	Num         int64  `xml:"num" json:"num"`
	Uids        string `xml:"uids" json:"uids"`
	NumRelevant int64  `xml:"numRelevant,omitempty" json:"numRelevant,omitempty" asn1:"optional"`
}
type MimAuthor struct {
	Name  string `xml:"name" json:"name"`
	Index int64  `xml:"index" json:"index"`
}
type MimCit struct {
	Number int64  `xml:"number" json:"number"`
	Author string `xml:"author" json:"author"`
	Others string `xml:"others" json:"others"`
	Year   int64  `xml:"year" json:"year"`
}
type MimReference struct {
	Number        int64       `xml:"number" json:"number"`
	OrigNumber    int64       `xml:"origNumber,omitempty" json:"origNumber,omitempty" asn1:"optional"`
	Type          string      `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"` //Type,EnumList:not-set(0),citation(1),book(2),personal-communication(3),book-citation(4)
	Authors       []MimAuthor `xml:"authors,omitempty" json:"authors,omitempty"`
	PrimaryAuthor string      `xml:"primaryAuthor" json:"primaryAuthor"`
	OtherAuthors  string      `xml:"otherAuthors" json:"otherAuthors"`
	CitationTitle string      `xml:"citationTitle" json:"citationTitle"`
	CitationType  int64       `xml:"citationType,omitempty" json:"citationType,omitempty" asn1:"optional"`
	BookTitle     string      `xml:"bookTitle,omitempty" json:"bookTitle,omitempty" asn1:"optional"`
	Editors       []MimAuthor `xml:"editors,omitempty" json:"editors,omitempty" asn1:"optional"`
	Volume        string      `xml:"volume,omitempty" json:"volume,omitempty" asn1:"optional"`
	Edition       string      `xml:"edition,omitempty" json:"edition,omitempty" asn1:"optional"`
	Journal       string      `xml:"journal,omitempty" json:"journal,omitempty" asn1:"optional"`
	Series        string      `xml:"series,omitempty" json:"series,omitempty" asn1:"optional"`
	Publisher     string      `xml:"publisher,omitempty" json:"publisher,omitempty" asn1:"optional"`
	Place         string      `xml:"place,omitempty" json:"place,omitempty" asn1:"optional"`
	CommNote      string      `xml:"commNote,omitempty" json:"commNote,omitempty" asn1:"optional"`
	PubDate       *MimDate    `xml:"pubDate,omitempty" json:"pubDate,omitempty"`
	Pages         []MimPage   `xml:"pages,omitempty" json:"pages,omitempty" asn1:"optional"`
	MiscInfo      string      `xml:"miscInfo,omitempty" json:"miscInfo,omitempty" asn1:"optional"`
	PubmedUID     int64       `xml:"pubmedUID,omitempty" json:"pubmedUID,omitempty" asn1:"optional"`
	Ambiguous     bool        `xml:"ambiguous" json:"ambiguous"`
	NoLink        bool        `xml:"noLink,omitempty" json:"noLink,omitempty" asn1:"optional"`
}
type MimIndexTerm struct {
	Key   string   `xml:"key" json:"key"`
	Terms []string `xml:"terms,omitempty" json:"terms,omitempty"`
}
type MimEditItem struct {
	Author  string   `xml:"author" json:"author"`
	ModDate *MimDate `xml:"modDate,omitempty" json:"modDate,omitempty"`
}
type MimDate struct {
	Year  int64 `xml:"year" json:"year"`
	Month int64 `xml:"month,omitempty" json:"month,omitempty" asn1:"optional"`
	Day   int64 `xml:"day,omitempty" json:"day,omitempty" asn1:"optional"`
}
type MimPage struct {
	From string `xml:"from" json:"from"`
	To   string `xml:"to,omitempty" json:"to,omitempty" asn1:"optional"`
}
