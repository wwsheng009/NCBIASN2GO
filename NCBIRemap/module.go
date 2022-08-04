package NCBIRemap

import "ncbiasn/NCBISeqloc"

type RemapDt int64
type RemapDbId string
type RemapRequest struct {
	Request *RMRequest `xml:"request,omitempty" json:"request,omitempty"`
	Version int64      `xml:"version" json:"version"`
	Tool    string     `xml:"tool,omitempty" json:"tool,omitempty" asn1:"optional"`
}

//RMRequest,ChoiceOption
type RMRequest struct {
	Remap          *RemapQuery `xml:"remap,omitempty" json:"remap,omitempty"`
	MapsToBuilds   string      `xml:"maps-to-builds,omitempty" json:"maps_to_builds,omitempty"`
	MapsFromBuilds string      `xml:"maps-from-builds,omitempty" json:"maps_from_builds,omitempty"`
	AllBuilds      interface{} `xml:"-" json:"-"` //AllBuilds,NullType
}

type RemapQuery struct {
	FromBuild string              `xml:"from-build" json:"from_build"`
	ToBuild   string              `xml:"to-build" json:"to_build"`
	Locs      []NCBISeqloc.SeqLoc `xml:"locs,omitempty" json:"locs,omitempty"`
}
type RemapReply struct {
	Reply  *RMReply `xml:"reply,omitempty" json:"reply,omitempty"`
	Dt     *RemapDt `xml:"dt,omitempty" json:"dt,omitempty"`
	Server string   `xml:"server" json:"server"`
	Msg    string   `xml:"msg,omitempty" json:"msg,omitempty" asn1:"optional"`
}

//RMReply,ChoiceOption
type RMReply struct {
	Error          string      `xml:"error,omitempty" json:"error,omitempty"`
	Remap          RemapResult `xml:"remap,omitempty" json:"remap,omitempty"`
	MapsToBuilds   []string    `xml:"maps-to-builds,omitempty" json:"maps_to_builds,omitempty"`
	MapsFromBuilds []string    `xml:"maps-from-builds,omitempty" json:"maps_from_builds,omitempty"`
	AllBuilds      []string    `xml:"all-builds,omitempty" json:"all_builds,omitempty"`
}

type RemapResult []NCBISeqloc.SeqLoc
