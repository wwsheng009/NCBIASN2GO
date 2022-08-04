package NCBIID2Access

import (
	"ncbiasn/NCBISeqSplit"
	"ncbiasn/NCBISeqloc"
)

type ID2RequestPacket []ID2Request
type ID2Request struct {
	SerialNumber int64     `xml:"serial-number,omitempty" json:"serial_number,omitempty" asn1:"optional"`
	Params       ID2Params `xml:"params,omitempty" json:"params,omitempty" asn1:"optional"`
	Request      struct {
		Init        interface{}            `xml:"-" json:"-"` //Init,NullType
		GetPackages *ID2RequestGetPackages `xml:"get-packages,omitempty" json:"get_packages,omitempty"`
		GetSeqId    *ID2RequestGetSeqId    `xml:"get-seq-id,omitempty" json:"get_seq_id,omitempty"`
		GetBlobId   *ID2RequestGetBlobId   `xml:"get-blob-id,omitempty" json:"get_blob_id,omitempty"`
		GetBlobInfo *ID2RequestGetBlobInfo `xml:"get-blob-info,omitempty" json:"get_blob_info,omitempty"`
		RegetBlob   *ID2RequestReGetBlob   `xml:"reget-blob,omitempty" json:"reget_blob,omitempty"`
		GetChunks   *ID2SRequestGetChunks  `xml:"get-chunks,omitempty" json:"get_chunks,omitempty"`
	} `xml:"request" json:"request"` //Request,ChoiceOption
}
type ID2RequestGetPackages struct {
	Names      []string    `xml:"names,omitempty" json:"names,omitempty" asn1:"optional"`
	NoContents interface{} `xml:"-" json:"-" asn1:"optional"` //NoContents,NullType
}
type ID2RequestGetSeqId struct {
	SeqId     *ID2SeqId `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
	SeqIdType int       `xml:"seq-id-type" json:"seq_id_type"` //SeqIdType,IntegerEnum:any(0),gi(1),text(2),general(4),all(127),label(128),taxid(256),hash(512),seq-length(1024),seq-mol(2048)
}
type ID2SeqId struct {
	String string            `xml:"string,omitempty" json:"string,omitempty"`
	SeqId  *NCBISeqloc.SeqId `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
}

//ID2SeqId,ChoiceOption
type ID2RequestGetBlobId struct {
	SeqId    *ID2RequestGetSeqId `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
	Sources  []string            `xml:"sources,omitempty" json:"sources,omitempty" asn1:"optional"`
	External interface{}         `xml:"-" json:"-" asn1:"optional"` //External,NullType
}
type ID2RequestGetBlobInfo struct {
	BlobId struct {
		BlobId  *ID2BlobId `xml:"blob-id,omitempty" json:"blob_id,omitempty"`
		Resolve struct {
			Request      *ID2RequestGetBlobId `xml:"request,omitempty" json:"request,omitempty"`
			ExcludeBlobs []ID2BlobId          `xml:"exclude-blobs,omitempty" json:"exclude_blobs,omitempty" asn1:"optional"`
		} `xml:"resolve,omitempty" json:"resolve,omitempty"`
	} `xml:"blob-id" json:"blob_id"` //BlobId,ChoiceOption
	GetSeqIds interface{}        `xml:"-" json:"-" asn1:"optional"` //GetSeqIds,NullType
	GetData   *ID2GetBlobDetails `xml:"get-data,omitempty" json:"get_data,omitempty" asn1:"optional"`
}
type ID2RequestReGetBlob struct {
	BlobId       *ID2BlobId `xml:"blob-id,omitempty" json:"blob_id,omitempty"`
	SplitVersion int64      `xml:"split-version" json:"split_version"`
	Offset       int64      `xml:"offset" json:"offset"`
}
type ID2SRequestGetChunks struct {
	BlobId       *ID2BlobId                 `xml:"blob-id,omitempty" json:"blob_id,omitempty"`
	Chunks       []NCBISeqSplit.ID2SChunkId `xml:"chunks,omitempty" json:"chunks,omitempty"`
	SplitVersion int64                      `xml:"split-version,omitempty" json:"split_version,omitempty" asn1:"optional"`
}
type ID2GetBlobDetails struct {
	Location      *NCBISeqloc.SeqLoc `xml:"location,omitempty" json:"location,omitempty" asn1:"optional"`
	SeqClassLevel int64              `xml:"seq-class-level" json:"seq_class_level" asn1:"default:1"`
	DescrLevel    int64              `xml:"descr-level" json:"descr_level" asn1:"default:1"`
	DescrTypeMask int64              `xml:"descr-type-mask" json:"descr_type_mask" asn1:"default:0"`
	AnnotTypeMask int64              `xml:"annot-type-mask" json:"annot_type_mask" asn1:"default:0"`
	FeatTypeMask  int64              `xml:"feat-type-mask" json:"feat_type_mask" asn1:"default:0"`
	SequenceLevel string             `xml:"sequence-level" json:"sequence_level"` //SequenceLevel,EnumList:none(0),seq-map(1),seq-data(2)
}
type ID2Reply struct {
	SerialNumber int64       `xml:"serial-number,omitempty" json:"serial_number,omitempty" asn1:"optional"`
	Params       ID2Params   `xml:"params,omitempty" json:"params,omitempty" asn1:"optional"`
	Error        []ID2Error  `xml:"error,omitempty" json:"error,omitempty" asn1:"optional"`
	EndOfReply   interface{} `xml:"-" json:"-" asn1:"optional"` //EndOfReply,NullType
	Reply        struct {
		Init          interface{}            `xml:"-" json:"-"` //Init,NullType
		Empty         interface{}            `xml:"-" json:"-"` //Empty,NullType
		GetPackage    *ID2ReplyGetPackage    `xml:"get-package,omitempty" json:"get_package,omitempty"`
		GetSeqId      *ID2ReplyGetSeqId      `xml:"get-seq-id,omitempty" json:"get_seq_id,omitempty"`
		GetBlobId     *ID2ReplyGetBlobId     `xml:"get-blob-id,omitempty" json:"get_blob_id,omitempty"`
		GetBlobSeqIds *ID2ReplyGetBlobSeqIds `xml:"get-blob-seq-ids,omitempty" json:"get_blob_seq_ids,omitempty"`
		GetBlob       *ID2ReplyGetBlob       `xml:"get-blob,omitempty" json:"get_blob,omitempty"`
		RegetBlob     *ID2ReplyReGetBlob     `xml:"reget-blob,omitempty" json:"reget_blob,omitempty"`
		GetSplitInfo  *ID2SReplyGetSplitInfo `xml:"get-split-info,omitempty" json:"get_split_info,omitempty"`
		GetChunk      *ID2SReplyGetChunk     `xml:"get-chunk,omitempty" json:"get_chunk,omitempty"`
	} `xml:"reply" json:"reply"` //Reply,ChoiceOption
	Discard string `xml:"discard,omitempty" json:"discard,omitempty" asn1:"optional"` //Discard,EnumList:reply(0),last-octet-string(1),nothing(2)
}
type ID2Error struct {
	Severity   string `xml:"severity" json:"severity"` //Severity,EnumList:warning(1),failed-command(2),failed-connection(3),failed-server(4),no-data(5),restricted-data(6),unsupported-command(7),invalid-arguments(8)
	RetryDelay int64  `xml:"retry-delay,omitempty" json:"retry_delay,omitempty" asn1:"optional"`
	Message    string `xml:"message,omitempty" json:"message,omitempty" asn1:"optional"`
}
type ID2ReplyGetPackage struct {
	Name   string    `xml:"name" json:"name"`
	Params ID2Params `xml:"params,omitempty" json:"params,omitempty" asn1:"optional"`
}
type ID2ReplyGetSeqId struct {
	Request    *ID2RequestGetSeqId `xml:"request,omitempty" json:"request,omitempty"`
	SeqId      []NCBISeqloc.SeqId  `xml:"seq-id,omitempty" json:"seq_id,omitempty" asn1:"optional"`
	EndOfReply interface{}         `xml:"-" json:"-" asn1:"optional"` //EndOfReply,NullType
}
type ID2BlobState string

//ID2BlobState,EnumList:live(0),suppressed-temp(1),suppressed(2),dead(3),protected(4),withdrawn(5)
type ID2ReplyGetBlobId struct {
	SeqId        *NCBISeqloc.SeqId               `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
	BlobId       *ID2BlobId                      `xml:"blob-id,omitempty" json:"blob_id,omitempty" asn1:"optional"`
	SplitVersion int64                           `xml:"split-version" json:"split_version" asn1:"default:0"`
	AnnotInfo    []NCBISeqSplit.ID2SSeqAnnotInfo `xml:"annot-info,omitempty" json:"annot_info,omitempty" asn1:"optional"`
	EndOfReply   interface{}                     `xml:"-" json:"-" asn1:"optional"` //EndOfReply,NullType
	BlobState    int64                           `xml:"blob-state,omitempty" json:"blob_state,omitempty" asn1:"optional"`
}
type ID2ReplyGetBlobSeqIds struct {
	BlobId *ID2BlobId    `xml:"blob-id,omitempty" json:"blob_id,omitempty"`
	Ids    *ID2ReplyData `xml:"ids,omitempty" json:"ids,omitempty" asn1:"optional"`
}
type ID2ReplyGetBlob struct {
	BlobId       *ID2BlobId    `xml:"blob-id,omitempty" json:"blob_id,omitempty"`
	SplitVersion int64         `xml:"split-version" json:"split_version" asn1:"default:0"`
	Data         *ID2ReplyData `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
	BlobState    int64         `xml:"blob-state,omitempty" json:"blob_state,omitempty" asn1:"optional"`
}
type ID2SReplyGetSplitInfo struct {
	BlobId       *ID2BlobId    `xml:"blob-id,omitempty" json:"blob_id,omitempty"`
	SplitVersion int64         `xml:"split-version" json:"split_version"`
	Data         *ID2ReplyData `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
	BlobState    int64         `xml:"blob-state,omitempty" json:"blob_state,omitempty" asn1:"optional"`
}
type ID2ReplyReGetBlob struct {
	BlobId       *ID2BlobId    `xml:"blob-id,omitempty" json:"blob_id,omitempty"`
	SplitVersion int64         `xml:"split-version" json:"split_version"`
	Offset       int64         `xml:"offset" json:"offset"`
	Data         *ID2ReplyData `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
}
type ID2SReplyGetChunk struct {
	BlobId  *ID2BlobId                `xml:"blob-id,omitempty" json:"blob_id,omitempty"`
	ChunkId *NCBISeqSplit.ID2SChunkId `xml:"chunk-id,omitempty" json:"chunk_id,omitempty"`
	Data    *ID2ReplyData             `xml:"data,omitempty" json:"data,omitempty" asn1:"optional"`
}
type ID2ReplyData struct {
	DataType        int    `xml:"data-type" json:"data_type"`               //DataType,IntegerEnum:seq-entry(0),seq-annot(1),id2s-split-info(2),id2s-chunk(3)
	DataFormat      int    `xml:"data-format" json:"data_format"`           //DataFormat,IntegerEnum:asn-binary(0),asn-text(1),xml(2)
	DataCompression int    `xml:"data-compression" json:"data_compression"` //DataCompression,IntegerEnum:none(0),gzip(1),nlmzip(2),bzip2(3)
	Data            []byte `xml:"data,omitempty" json:"data,omitempty"`
}
type ID2BlobSeqIds []ID2BlobSeqId
type ID2BlobSeqId struct {
	SeqId    *NCBISeqloc.SeqId `xml:"seq-id,omitempty" json:"seq_id,omitempty"`
	Replaced interface{}       `xml:"-" json:"-" asn1:"optional"` //Replaced,NullType
}
type ID2BlobId struct {
	Sat     int64 `xml:"sat" json:"sat"`
	SubSat  int   `xml:"sub-sat" json:"sub_sat"` //SubSat,IntegerEnum:main(0),snp(1),snp-graph(4),cdd(8),mgc(16),hprd(32),sts(64),trna(128),exon(512)
	SatKey  int64 `xml:"sat-key" json:"sat_key"`
	Version int64 `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
}
type ID2Params []ID2Param
type ID2Param struct {
	Name  string   `xml:"name" json:"name"`
	Value []string `xml:"value,omitempty" json:"value,omitempty" asn1:"optional"`
	Type  string   `xml:"type" json:"type"` //Type,EnumList:set-value(1),get-value(2),force-value(3),use-package(4)
}
