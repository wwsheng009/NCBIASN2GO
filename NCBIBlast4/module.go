package NCBIBlast4

import (
	"ncbiasn/NCBIScoreMat"
	"ncbiasn/NCBISeqCommon"
	"ncbiasn/NCBISeqalign"
	"ncbiasn/NCBISeqloc"
	"ncbiasn/NCBISequence"
)

type Blast4Request struct {
	Ident string             `xml:"ident,omitempty" json:"ident,omitempty" asn1:"optional"`
	Body  *Blast4RequestBody `xml:"body,omitempty" json:"body,omitempty"`
}
type Blast4Archive struct {
	Request  *Blast4Request               `xml:"request,omitempty" json:"request,omitempty"`
	Results  *Blast4GetSearchResultsReply `xml:"results,omitempty" json:"results,omitempty"`
	Messages []Blast4Error                `xml:"messages,omitempty" json:"messages,omitempty" asn1:"optional"`
}
type Blast4RequestBody struct {
	FinishParams          *Blast4FinishParamsRequest     `xml:"finish-params,omitempty" json:"finish_params,omitempty"`
	GetDatabases          interface{}                    `xml:"-" json:"-"` //GetDatabases,NullType
	GetMatrices           interface{}                    `xml:"-" json:"-"` //GetMatrices,NullType
	GetParameters         interface{}                    `xml:"-" json:"-"` //GetParameters,NullType
	GetParamsets          interface{}                    `xml:"-" json:"-"` //GetParamsets,NullType
	GetPrograms           interface{}                    `xml:"-" json:"-"` //GetPrograms,NullType
	GetSearchResults      *Blast4GetSearchResultsRequest `xml:"get-search-results,omitempty" json:"get_search_results,omitempty"`
	GetSequences          *Blast4GetSequencesRequest     `xml:"get-sequences,omitempty" json:"get_sequences,omitempty"`
	QueueSearch           *Blast4QueueSearchRequest      `xml:"queue-search,omitempty" json:"queue_search,omitempty"`
	GetRequestInfo        *Blast4GetRequestInfoRequest   `xml:"get-request-info,omitempty" json:"get_request_info,omitempty"`
	GetSequenceParts      *Blast4GetSeqPartsRequest      `xml:"get-sequence-parts,omitempty" json:"get_sequence_parts,omitempty"`
	GetWindowmaskedTaxids interface{}                    `xml:"-" json:"-"` //GetWindowmaskedTaxids,NullType
	GetProtocolInfo       Blast4GetProtocolInfoRequest   `xml:"get-protocol-info,omitempty" json:"get_protocol_info,omitempty"`
	GetSearchInfo         *Blast4GetSearchInfoRequest    `xml:"get-search-info,omitempty" json:"get_search_info,omitempty"`
	GetDatabasesEx        *Blast4GetDatabasesExRequest   `xml:"get-databases-ex,omitempty" json:"get_databases_ex,omitempty"`
}

//Blast4RequestBody,ChoiceOption
type Blast4GetDatabasesExRequest struct {
	Params Blast4Parameters `xml:"params,omitempty" json:"params,omitempty" asn1:"optional"`
}
type Blast4FinishParamsRequest struct {
	Program  string           `xml:"program" json:"program"`
	Service  string           `xml:"service" json:"service"`
	Paramset string           `xml:"paramset,omitempty" json:"paramset,omitempty" asn1:"optional"`
	Params   Blast4Parameters `xml:"params,omitempty" json:"params,omitempty" asn1:"optional"`
}
type Blast4ResultTypes string

//Blast4ResultTypes,EnumList:default(63),alignments(1),phi-alignments(2),masks(4),ka-blocks(8),search-stats(16),pssm(32),simple-results(64)
type Blast4GetSearchResultsRequest struct {
	RequestId   string `xml:"request-id" json:"request_id"`
	ResultTypes int64  `xml:"result-types,omitempty" json:"result_types,omitempty" asn1:"optional"`
}
type Blast4Queries struct {
	Pssm       *NCBIScoreMat.PssmWithParameters `xml:"pssm,omitempty" json:"pssm,omitempty"`
	SeqLocList []NCBISeqloc.SeqLoc              `xml:"seq-loc-list,omitempty" json:"seq_loc_list,omitempty"`
	BioseqSet  *NCBISequence.BioseqSet          `xml:"bioseq-set,omitempty" json:"bioseq_set,omitempty"`
}

//Blast4Queries,ChoiceOption
type Blast4QueueSearchRequest struct {
	Program          string           `xml:"program" json:"program"`
	Service          string           `xml:"service" json:"service"`
	Queries          *Blast4Queries   `xml:"queries,omitempty" json:"queries,omitempty"`
	Subject          *Blast4Subject   `xml:"subject,omitempty" json:"subject,omitempty"`
	Paramset         string           `xml:"paramset,omitempty" json:"paramset,omitempty" asn1:"optional"`
	AlgorithmOptions Blast4Parameters `xml:"algorithm-options,omitempty" json:"algorithm_options,omitempty" asn1:"optional"`
	ProgramOptions   Blast4Parameters `xml:"program-options,omitempty" json:"program_options,omitempty" asn1:"optional"`
	FormatOptions    Blast4Parameters `xml:"format-options,omitempty" json:"format_options,omitempty" asn1:"optional"`
}
type Blast4GetSearchStatusRequest struct {
	RequestId string `xml:"request-id" json:"request_id"`
}
type Blast4GetSearchStatusReply struct {
	Status string `xml:"status" json:"status"`
}
type Blast4GetRequestInfoRequest struct {
	RequestId string `xml:"request-id" json:"request_id"`
}
type Blast4GetRequestInfoReply struct {
	Database         *Blast4Database  `xml:"database,omitempty" json:"database,omitempty"`
	Program          string           `xml:"program" json:"program"`
	Service          string           `xml:"service" json:"service"`
	CreatedBy        string           `xml:"created-by" json:"created_by"`
	Queries          *Blast4Queries   `xml:"queries,omitempty" json:"queries,omitempty"`
	AlgorithmOptions Blast4Parameters `xml:"algorithm-options,omitempty" json:"algorithm_options,omitempty"`
	ProgramOptions   Blast4Parameters `xml:"program-options,omitempty" json:"program_options,omitempty"`
	FormatOptions    Blast4Parameters `xml:"format-options,omitempty" json:"format_options,omitempty" asn1:"optional"`
	Subjects         *Blast4Subject   `xml:"subjects,omitempty" json:"subjects,omitempty" asn1:"optional"`
}
type Blast4GetSearchStrategyRequest struct {
	RequestId string `xml:"request-id" json:"request_id"`
}
type Blast4GetSearchStrategyReply Blast4Request
type Blast4GetSequencesRequest struct {
	Database    *Blast4Database    `xml:"database,omitempty" json:"database,omitempty"`
	SeqIds      []NCBISeqloc.SeqId `xml:"seq-ids,omitempty" json:"seq_ids,omitempty"`
	SkipSeqData bool               `xml:"skip-seq-data" json:"skip_seq_data"`
	TargetOnly  bool               `xml:"target-only,omitempty" json:"target_only,omitempty" asn1:"optional"`
}
type Blast4GetSeqPartsRequest struct {
	Database     *Blast4Database          `xml:"database,omitempty" json:"database,omitempty"`
	SeqLocations []NCBISeqloc.SeqInterval `xml:"seq-locations,omitempty" json:"seq_locations,omitempty"`
}
type Blast4GetProtocolInfoRequest Blast4Parameters
type Blast4GetSearchInfoRequest struct {
	RequestId string           `xml:"request-id" json:"request_id"`
	Info      Blast4Parameters `xml:"info,omitempty" json:"info,omitempty" asn1:"optional"`
}
type Blast4Reply struct {
	Errors []Blast4Error    `xml:"errors,omitempty" json:"errors,omitempty" asn1:"optional"`
	Body   *Blast4ReplyBody `xml:"body,omitempty" json:"body,omitempty"`
}
type Blast4ReplyBody struct {
	FinishParams          Blast4FinishParamsReply          `xml:"finish-params,omitempty" json:"finish_params,omitempty"`
	GetDatabases          Blast4GetDatabasesReply          `xml:"get-databases,omitempty" json:"get_databases,omitempty"`
	GetMatrices           Blast4GetMatricesReply           `xml:"get-matrices,omitempty" json:"get_matrices,omitempty"`
	GetParameters         Blast4GetParametersReply         `xml:"get-parameters,omitempty" json:"get_parameters,omitempty"`
	GetParamsets          Blast4GetParamsetsReply          `xml:"get-paramsets,omitempty" json:"get_paramsets,omitempty"`
	GetPrograms           Blast4GetProgramsReply           `xml:"get-programs,omitempty" json:"get_programs,omitempty"`
	GetSearchResults      *Blast4GetSearchResultsReply     `xml:"get-search-results,omitempty" json:"get_search_results,omitempty"`
	GetSequences          Blast4GetSequencesReply          `xml:"get-sequences,omitempty" json:"get_sequences,omitempty"`
	QueueSearch           *Blast4QueueSearchReply          `xml:"queue-search,omitempty" json:"queue_search,omitempty"`
	GetQueries            *Blast4GetQueriesReply           `xml:"get-queries,omitempty" json:"get_queries,omitempty"`
	GetRequestInfo        *Blast4GetRequestInfoReply       `xml:"get-request-info,omitempty" json:"get_request_info,omitempty"`
	GetSequenceParts      Blast4GetSeqPartsReply           `xml:"get-sequence-parts,omitempty" json:"get_sequence_parts,omitempty"`
	GetWindowmaskedTaxids Blast4GetWindowmaskedTaxidsReply `xml:"get-windowmasked-taxids,omitempty" json:"get_windowmasked_taxids,omitempty"`
	GetProtocolInfo       Blast4GetProtocolInfoReply       `xml:"get-protocol-info,omitempty" json:"get_protocol_info,omitempty"`
	GetSearchInfo         *Blast4GetSearchInfoReply        `xml:"get-search-info,omitempty" json:"get_search_info,omitempty"`
	GetDatabasesEx        Blast4GetDatabasesExReply        `xml:"get-databases-ex,omitempty" json:"get_databases_ex,omitempty"`
}

//Blast4ReplyBody,ChoiceOption
type Blast4FinishParamsReply Blast4Parameters
type Blast4GetWindowmaskedTaxidsReply []int64
type Blast4GetDatabasesReply []Blast4DatabaseInfo
type Blast4GetDatabasesExReply []Blast4DatabaseInfo
type Blast4GetMatricesReply []Blast4MatrixId
type Blast4GetParametersReply []Blast4ParameterInfo
type Blast4GetParamsetsReply []Blast4TaskInfo
type Blast4GetProgramsReply []Blast4ProgramInfo
type Blast4GetSearchResultsReply struct {
	Alignments    NCBISeqalign.SeqAlignSet         `xml:"alignments,omitempty" json:"alignments,omitempty" asn1:"optional"`
	PhiAlignments *Blast4PhiAlignments             `xml:"phi-alignments,omitempty" json:"phi_alignments,omitempty" asn1:"optional"`
	Masks         []Blast4Mask                     `xml:"masks,omitempty" json:"masks,omitempty" asn1:"optional"`
	KaBlocks      []Blast4KaBlock                  `xml:"ka-blocks,omitempty" json:"ka_blocks,omitempty" asn1:"optional"`
	SearchStats   []string                         `xml:"search-stats,omitempty" json:"search_stats,omitempty" asn1:"optional"`
	Pssm          *NCBIScoreMat.PssmWithParameters `xml:"pssm,omitempty" json:"pssm,omitempty" asn1:"optional"`
	SimpleResults *Blast4SimpleResults             `xml:"simple-results,omitempty" json:"simple_results,omitempty" asn1:"optional"`
}
type Blast4GetSequencesReply []NCBISequence.Bioseq
type Blast4SeqPartData struct {
	Id   *NCBISeqloc.SeqId      `xml:"id,omitempty" json:"id,omitempty"`
	Data *NCBISeqCommon.SeqData `xml:"data,omitempty" json:"data,omitempty"`
}
type Blast4GetSeqPartsReply []Blast4SeqPartData
type Blast4QueueSearchReply struct {
	RequestId string `xml:"request-id,omitempty" json:"request_id,omitempty" asn1:"optional"`
}
type Blast4GetQueriesReply struct {
	Queries *Blast4Queries `xml:"queries,omitempty" json:"queries,omitempty"`
}
type Blast4GetProtocolInfoReply Blast4Parameters
type Blast4GetSearchInfoReply struct {
	RequestId string           `xml:"request-id" json:"request_id"`
	Info      Blast4Parameters `xml:"info,omitempty" json:"info,omitempty" asn1:"optional"`
}
type Blast4Error struct {
	Code    int64  `xml:"code" json:"code"`
	Message string `xml:"message,omitempty" json:"message,omitempty" asn1:"optional"`
}
type Blast4ErrorFlags string

//Blast4ErrorFlags,EnumList:warning(1024),error(2048)
type Blast4ErrorCode int

//Blast4ErrorCode,IntegerEnum:conversion-warning(1024),internal-error(2048),not-implemented(2049),not-allowed(2050),bad-request(2051),bad-request-id(2052),search-pending(2053)
type Blast4Cutoff struct {
	EValue   float64 `xml:"e-value,omitempty" json:"e_value,omitempty"`
	RawScore int64   `xml:"raw-score,omitempty" json:"raw_score,omitempty"`
}

//Blast4Cutoff,ChoiceOption
type Blast4Database struct {
	Name string             `xml:"name" json:"name"`
	Type *Blast4ResidueType `xml:"type,omitempty" json:"type,omitempty"`
}
type Blast4Seqtech int

//Blast4Seqtech,IntegerEnum:unknown(0),standard(1),est(2),sts(3),survey(4),genemap(5),physmap(6),derived(7),concept-trans(8),seq-pept(9),both(10),seq-pept-overlap(11),seq-pept-homol(12),concept-trans-a(13),htgs-1(14),htgs-2(15),htgs-3(16),fli-cdna(17),htgs-0(18),htc(19),wgs(20),other(255)
type Blast4DatabaseInfo struct {
	Database     *Blast4Database  `xml:"database,omitempty" json:"database,omitempty"`
	Description  string           `xml:"description" json:"description"`
	LastUpdated  string           `xml:"last-updated" json:"last_updated"`
	TotalLength  *int64           `xml:"total-length,omitempty" json:"total_length,omitempty"`
	NumSequences *int64           `xml:"num-sequences,omitempty" json:"num_sequences,omitempty"`
	Seqtech      *Blast4Seqtech   `xml:"seqtech,omitempty" json:"seqtech,omitempty"`
	Taxid        int64            `xml:"taxid" json:"taxid"`
	Extended     Blast4Parameters `xml:"extended,omitempty" json:"extended,omitempty" asn1:"optional"`
}
type Blast4FrameType string

//Blast4FrameType,EnumList:notset(0),plus1(1),plus2(2),plus3(3),minus1(4),minus2(5),minus3(6)
type Blast4KaBlock struct {
	Lambda float64 `xml:"lambda" json:"lambda"`
	K      float64 `xml:"k" json:"k"`
	H      float64 `xml:"h" json:"h"`
	Gapped bool    `xml:"gapped" json:"gapped"`
}
type Blast4Mask struct {
	Locations []NCBISeqloc.SeqLoc `xml:"locations,omitempty" json:"locations,omitempty"`
	Frame     *Blast4FrameType    `xml:"frame,omitempty" json:"frame,omitempty"`
}
type Blast4MatrixId struct {
	ResidueType *Blast4ResidueType `xml:"residue-type,omitempty" json:"residue_type,omitempty"`
	Name        string             `xml:"name" json:"name"`
}
type Blast4Parameter struct {
	Name  string       `xml:"name" json:"name"`
	Value *Blast4Value `xml:"value,omitempty" json:"value,omitempty"`
}
type Blast4ParameterInfo struct {
	Name string `xml:"name" json:"name"`
	Type string `xml:"type" json:"type"`
}
type Blast4TaskInfo struct {
	Name          string `xml:"name" json:"name"`
	Documentation string `xml:"documentation" json:"documentation"`
}
type Blast4ProgramInfo struct {
	Program  string   `xml:"program" json:"program"`
	Services []string `xml:"services,omitempty" json:"services,omitempty"`
}
type Blast4ResidueType string

//Blast4ResidueType,EnumList:unknown(0),protein(1),nucleotide(2)
type Blast4StrandType string

//Blast4StrandType,EnumList:forward-strand(1),reverse-strand(2),both-strands(3)
type Blast4Subject struct {
	Database   string                `xml:"database,omitempty" json:"database,omitempty"`
	Sequences  []NCBISequence.Bioseq `xml:"sequences,omitempty" json:"sequences,omitempty"`
	SeqLocList []NCBISeqloc.SeqLoc   `xml:"seq-loc-list,omitempty" json:"seq_loc_list,omitempty"`
}

//Blast4Subject,ChoiceOption
type Blast4Parameters []Blast4Parameter
type Blast4PhiAlignments struct {
	NumAlignments int64               `xml:"num-alignments" json:"num_alignments"`
	SeqLocs       []NCBISeqloc.SeqLoc `xml:"seq-locs,omitempty" json:"seq_locs,omitempty"`
}
type Blast4Value struct {
	BigInteger     *int64                            `xml:"big-integer,omitempty" json:"big_integer,omitempty"`
	Bioseq         *NCBISequence.Bioseq              `xml:"bioseq,omitempty" json:"bioseq,omitempty"`
	Boolean        bool                              `xml:"boolean,omitempty" json:"boolean,omitempty"`
	Cutoff         *Blast4Cutoff                     `xml:"cutoff,omitempty" json:"cutoff,omitempty"`
	Integer        int64                             `xml:"integer,omitempty" json:"integer,omitempty"`
	Matrix         *NCBIScoreMat.PssmWithParameters  `xml:"matrix,omitempty" json:"matrix,omitempty"`
	Real           float64                           `xml:"real,omitempty" json:"real,omitempty"`
	SeqAlign       *NCBISeqalign.SeqAlign            `xml:"seq-align,omitempty" json:"seq_align,omitempty"`
	SeqId          *NCBISeqloc.SeqId                 `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
	SeqLoc         *NCBISeqloc.SeqLoc                `xml:"seq-loc,omitempty" json:"seq_loc,omitempty"`
	StrandType     *Blast4StrandType                 `xml:"strand-type,omitempty" json:"strand_type,omitempty"`
	String         string                            `xml:"string,omitempty" json:"string,omitempty"`
	BigIntegerList []int64                           `xml:"big-integer-list,omitempty" json:"big_integer_list,omitempty"`
	BioseqList     []NCBISequence.Bioseq             `xml:"bioseq-list,omitempty" json:"bioseq_list,omitempty"`
	BooleanList    []bool                            `xml:"boolean-list,omitempty" json:"boolean_list,omitempty"`
	CutoffList     []Blast4Cutoff                    `xml:"cutoff-list,omitempty" json:"cutoff_list,omitempty"`
	IntegerList    []int64                           `xml:"integer-list,omitempty" json:"integer_list,omitempty"`
	MatrixList     []NCBIScoreMat.PssmWithParameters `xml:"matrix-list,omitempty" json:"matrix_list,omitempty"`
	RealList       []float64                         `xml:"real-list,omitempty" json:"real_list,omitempty"`
	SeqAlignList   []NCBISeqalign.SeqAlign           `xml:"seq-align-list,omitempty" json:"seq_align_list,omitempty"`
	SeqIdList      []NCBISeqloc.SeqId                `xml:"seq-id-list,omitempty" json:"seq_id_list,omitempty"`
	SeqLocList     []NCBISeqloc.SeqLoc               `xml:"seq-loc-list,omitempty" json:"seq_loc_list,omitempty"`
	StrandTypeList []Blast4StrandType                `xml:"strand-type-list,omitempty" json:"strand_type_list,omitempty"`
	StringList     []string                          `xml:"string-list,omitempty" json:"string_list,omitempty"`
	BioseqSet      *NCBISequence.BioseqSet           `xml:"bioseq-set,omitempty" json:"bioseq_set,omitempty"`
	SeqAlignSet    NCBISeqalign.SeqAlignSet          `xml:"seq-align-set,omitempty" json:"seq_align_set,omitempty"`
	QueryMask      *Blast4Mask                       `xml:"query-mask,omitempty" json:"query_mask,omitempty"`
}

//Blast4Value,ChoiceOption
type Blast4SimpleResults struct {
	AllAlignments []Blast4AlignmentsForQuery `xml:"all-alignments,omitempty" json:"all_alignments,omitempty"`
}
type Blast4AlignmentsForQuery struct {
	QueryId    string                  `xml:"query-id" json:"query_id"`
	Alignments []Blast4SimpleAlignment `xml:"alignments,omitempty" json:"alignments,omitempty"`
}
type Blast4SimpleAlignment struct {
	SubjectId        string       `xml:"subject-id" json:"subject_id"`
	EValue           float64      `xml:"e-value" json:"e_value"`
	BitScore         float64      `xml:"bit-score" json:"bit_score"`
	NumIdentities    int64        `xml:"num-identities,omitempty" json:"num_identities,omitempty" asn1:"optional"`
	NumIndels        int64        `xml:"num-indels,omitempty" json:"num_indels,omitempty" asn1:"optional"`
	FullQueryRange   *Blast4Range `xml:"full-query-range,omitempty" json:"full_query_range,omitempty"`
	FullSubjectRange *Blast4Range `xml:"full-subject-range,omitempty" json:"full_subject_range,omitempty"`
}
type Blast4Range struct {
	Start  int64 `xml:"start,omitempty" json:"start,omitempty" asn1:"optional"`
	End    int64 `xml:"end,omitempty" json:"end,omitempty" asn1:"optional"`
	Strand int64 `xml:"strand,omitempty" json:"strand,omitempty" asn1:"optional"`
}
