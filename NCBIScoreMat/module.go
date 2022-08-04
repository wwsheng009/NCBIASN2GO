package NCBIScoreMat

import (
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBISequence"
)

type BlockProperty struct {
	Type      int    `xml:"type" json:"type"` //Type,IntegerEnum:unassigned(0),threshold(1),minscore(2),maxscore(3),meanscore(4),variance(5),name(10),is-optional(20),other(255)
	Intvalue  int64  `xml:"intvalue,omitempty" json:"intvalue,omitempty" asn1:"optional"`
	Textvalue string `xml:"textvalue,omitempty" json:"textvalue,omitempty" asn1:"optional"`
}
type CoreBlock struct {
	Start    int64           `xml:"start" json:"start"`
	Stop     int64           `xml:"stop" json:"stop"`
	Minstart int64           `xml:"minstart,omitempty" json:"minstart,omitempty" asn1:"optional"`
	Maxstop  int64           `xml:"maxstop,omitempty" json:"maxstop,omitempty" asn1:"optional"`
	Property []BlockProperty `xml:"property,omitempty" json:"property,omitempty" asn1:"optional"`
}
type LoopConstraint struct {
	Minlength int64 `xml:"minlength" json:"minlength" asn1:"default:0"`
	Maxlength int64 `xml:"maxlength" json:"maxlength" asn1:"default:100000"`
}
type CoreDef struct {
	Nblocks         int64            `xml:"nblocks" json:"nblocks"`
	Blocks          []CoreBlock      `xml:"blocks,omitempty" json:"blocks,omitempty"`
	Loops           []LoopConstraint `xml:"loops,omitempty" json:"loops,omitempty"`
	IsDiscontinuous bool             `xml:"isDiscontinuous,omitempty" json:"isDiscontinuous,omitempty" asn1:"optional"`
	Insertions      []int64          `xml:"insertions,omitempty" json:"insertions,omitempty" asn1:"optional"`
}
type SiteAnnot struct {
	StartPosition int64    `xml:"startPosition" json:"startPosition"`
	StopPosition  int64    `xml:"stopPosition" json:"stopPosition"`
	Description   string   `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	Type          int64    `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"`
	Aliases       []string `xml:"aliases,omitempty" json:"aliases,omitempty" asn1:"optional"`
	Motif         string   `xml:"motif,omitempty" json:"motif,omitempty" asn1:"optional"`
	Motifuse      int64    `xml:"motifuse,omitempty" json:"motifuse,omitempty" asn1:"optional"`
}
type SiteAnnotSet []SiteAnnot
type PssmFinalData struct {
	Scores             []int64 `xml:"scores,omitempty" json:"scores,omitempty"`
	Lambda             float64 `xml:"lambda" json:"lambda"`
	Kappa              float64 `xml:"kappa" json:"kappa"`
	H                  float64 `xml:"h" json:"h"`
	ScalingFactor      int64   `xml:"scalingFactor" json:"scalingFactor" asn1:"default:1"`
	LambdaUngapped     float64 `xml:"lambdaUngapped,omitempty" json:"lambdaUngapped,omitempty" asn1:"optional"`
	KappaUngapped      float64 `xml:"kappaUngapped,omitempty" json:"kappaUngapped,omitempty" asn1:"optional"`
	HUngapped          float64 `xml:"hUngapped,omitempty" json:"hUngapped,omitempty" asn1:"optional"`
	WordScoreThreshold float64 `xml:"wordScoreThreshold,omitempty" json:"wordScoreThreshold,omitempty" asn1:"optional"`
}
type PssmIntermediateData struct {
	ResFreqsPerPos         []int64   `xml:"resFreqsPerPos,omitempty" json:"resFreqsPerPos,omitempty" asn1:"optional"`
	WeightedResFreqsPerPos []float64 `xml:"weightedResFreqsPerPos,omitempty" json:"weightedResFreqsPerPos,omitempty" asn1:"optional"`
	FreqRatios             []float64 `xml:"freqRatios,omitempty" json:"freqRatios,omitempty"`
	InformationContent     []float64 `xml:"informationContent,omitempty" json:"informationContent,omitempty" asn1:"optional"`
	GaplessColumnWeights   []float64 `xml:"gaplessColumnWeights,omitempty" json:"gaplessColumnWeights,omitempty" asn1:"optional"`
	Sigma                  []float64 `xml:"sigma,omitempty" json:"sigma,omitempty" asn1:"optional"`
	IntervalSizes          []int64   `xml:"intervalSizes,omitempty" json:"intervalSizes,omitempty" asn1:"optional"`
	NumMatchingSeqs        []int64   `xml:"numMatchingSeqs,omitempty" json:"numMatchingSeqs,omitempty" asn1:"optional"`
	NumIndeptObsr          []float64 `xml:"numIndeptObsr,omitempty" json:"numIndeptObsr,omitempty" asn1:"optional"`
}
type Pssm struct {
	IsProtein        bool                   `xml:"isProtein" json:"isProtein"`
	Identifier       *NCBIGeneral.ObjectId  `xml:"identifier,omitempty" json:"identifier,omitempty" asn1:"optional"`
	NumRows          int64                  `xml:"numRows" json:"numRows"`
	NumColumns       int64                  `xml:"numColumns" json:"numColumns"`
	RowLabels        []string               `xml:"rowLabels,omitempty" json:"rowLabels,omitempty" asn1:"optional"`
	ByRow            bool                   `xml:"byRow" json:"byRow"`
	Query            *NCBISequence.SeqEntry `xml:"query,omitempty" json:"query,omitempty" asn1:"optional"`
	IntermediateData *PssmIntermediateData  `xml:"intermediateData,omitempty" json:"intermediateData,omitempty" asn1:"optional"`
	FinalData        *PssmFinalData         `xml:"finalData,omitempty" json:"finalData,omitempty" asn1:"optional"`
}
type FormatRpsDbParameters struct {
	MatrixName string `xml:"matrixName" json:"matrixName"`
	GapOpen    int64  `xml:"gapOpen,omitempty" json:"gapOpen,omitempty" asn1:"optional"`
	GapExtend  int64  `xml:"gapExtend,omitempty" json:"gapExtend,omitempty" asn1:"optional"`
}
type PssmParameters struct {
	Pseudocount    int64                  `xml:"pseudocount,omitempty" json:"pseudocount,omitempty" asn1:"optional"`
	Rpsdbparams    *FormatRpsDbParameters `xml:"rpsdbparams,omitempty" json:"rpsdbparams,omitempty" asn1:"optional"`
	Constraints    *CoreDef               `xml:"constraints,omitempty" json:"constraints,omitempty" asn1:"optional"`
	BitScoreThresh float64                `xml:"bitScoreThresh,omitempty" json:"bitScoreThresh,omitempty" asn1:"optional"`
	AnnotatedSites SiteAnnotSet           `xml:"annotatedSites,omitempty" json:"annotatedSites,omitempty" asn1:"optional"`
}
type PssmWithParameters struct {
	Pssm   *Pssm           `xml:"pssm,omitempty" json:"pssm,omitempty"`
	Params *PssmParameters `xml:"params,omitempty" json:"params,omitempty" asn1:"optional"`
}
