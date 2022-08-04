package Docsum34

type Assay struct {
	Attlist struct {
		Handle     string `xml:"handle,omitempty" json:"handle,omitempty" asn1:"optional"`
		Batch      string `xml:"batch,omitempty" json:"batch,omitempty" asn1:"optional"`
		BatchId    int64  `xml:"batchId,omitempty" json:"batchId,omitempty" asn1:"optional"`
		BatchType  string `xml:"batchType,omitempty" json:"batchType,omitempty" asn1:"optional"` //BatchType,EnumList:snpassay(1),validation(2),doublehit(3)
		MolType    string `xml:"molType,omitempty" json:"molType,omitempty" asn1:"optional"`     //MolType,EnumList:genomic(1),cDNA(2),mito(3),chloro(4)
		SampleSize int64  `xml:"sampleSize,omitempty" json:"sampleSize,omitempty" asn1:"optional"`
		Population string `xml:"population,omitempty" json:"population,omitempty" asn1:"optional"`
		LinkoutUrl string `xml:"linkoutUrl,omitempty" json:"linkoutUrl,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	Method struct {
		EMethod struct {
			Attlist struct {
				Name string `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
				Id   string `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
			} `xml:"attlist" json:"attlist"`
			Exception string `xml:"exception" json:"exception"`
		} `xml:"eMethod,omitempty" json:"eMethod,omitempty" asn1:"optional"`
	} `xml:"method" json:"method"`
	Taxonomy struct {
		Attlist struct {
			Id       int64  `xml:"id" json:"id"`
			Organism string `xml:"organism,omitempty" json:"organism,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		Taxonomy interface{} `xml:"-" json:"-"` //Taxonomy,NullType
	} `xml:"taxonomy" json:"taxonomy"`
	Strains  []string `xml:"strains,omitempty" json:"strains,omitempty" asn1:"optional"`
	Comment  string   `xml:"comment,omitempty" json:"comment,omitempty" asn1:"optional"`
	Citation []string `xml:"citation,omitempty" json:"citation,omitempty" asn1:"optional"`
}
type Assembly struct {
	Attlist struct {
		DbSnpBuild     int64  `xml:"dbSnpBuild" json:"dbSnpBuild"`
		GenomeBuild    string `xml:"genomeBuild" json:"genomeBuild"`
		GroupLabel     string `xml:"groupLabel,omitempty" json:"groupLabel,omitempty" asn1:"optional"`
		AssemblySource string `xml:"assemblySource,omitempty" json:"assemblySource,omitempty" asn1:"optional"`
		Current        bool   `xml:"current,omitempty" json:"current,omitempty" asn1:"optional"`
		Reference      bool   `xml:"reference,omitempty" json:"reference,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	Component []Component `xml:"component,omitempty" json:"component,omitempty" asn1:"optional"`
	SnpStat   struct {
		Attlist struct {
			MapWeight           string `xml:"mapWeight" json:"mapWeight"` //MapWeight,EnumList:unmapped(1),unique-in-contig(2),two-hits-in-contig(3),less-10-hits(4),multiple-hits(5)
			ChromCount          int64  `xml:"chromCount,omitempty" json:"chromCount,omitempty" asn1:"optional"`
			PlacedContigCount   int64  `xml:"placedContigCount,omitempty" json:"placedContigCount,omitempty" asn1:"optional"`
			UnplacedContigCount int64  `xml:"unplacedContigCount,omitempty" json:"unplacedContigCount,omitempty" asn1:"optional"`
			SeqlocCount         int64  `xml:"seqlocCount,omitempty" json:"seqlocCount,omitempty" asn1:"optional"`
			HapCount            int64  `xml:"hapCount,omitempty" json:"hapCount,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		SnpStat interface{} `xml:"-" json:"-"` //SnpStat,NullType
	} `xml:"snpStat" json:"snpStat"`
}
type BaseURL struct {
	Attlist struct {
		UrlId        int64  `xml:"urlId,omitempty" json:"urlId,omitempty" asn1:"optional"`
		ResourceName string `xml:"resourceName,omitempty" json:"resourceName,omitempty" asn1:"optional"`
		ResourceId   string `xml:"resourceId,omitempty" json:"resourceId,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	BaseURL string `xml:"baseURL" json:"baseURL"`
}
type Component struct {
	Attlist struct {
		ComponentType string `xml:"componentType,omitempty" json:"componentType,omitempty" asn1:"optional"` //ComponentType,EnumList:contig(1),mrna(2)
		CtgId         int64  `xml:"ctgId,omitempty" json:"ctgId,omitempty" asn1:"optional"`
		Accession     string `xml:"accession,omitempty" json:"accession,omitempty" asn1:"optional"`
		Name          string `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
		Chromosome    string `xml:"chromosome,omitempty" json:"chromosome,omitempty" asn1:"optional"`
		Start         int64  `xml:"start,omitempty" json:"start,omitempty" asn1:"optional"`
		End           int64  `xml:"end,omitempty" json:"end,omitempty" asn1:"optional"`
		Orientation   string `xml:"orientation,omitempty" json:"orientation,omitempty" asn1:"optional"` //Orientation,EnumList:fwd(1),rev(2),unknown(3)
		Gi            string `xml:"gi,omitempty" json:"gi,omitempty" asn1:"optional"`
		GroupTerm     string `xml:"groupTerm,omitempty" json:"groupTerm,omitempty" asn1:"optional"`
		ContigLabel   string `xml:"contigLabel,omitempty" json:"contigLabel,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	MapLoc []MapLoc `xml:"mapLoc,omitempty" json:"mapLoc,omitempty"`
}
type ExchangeSet struct {
	Attlist struct {
		SetType     string `xml:"setType,omitempty" json:"setType,omitempty" asn1:"optional"`
		SetDepth    string `xml:"setDepth,omitempty" json:"setDepth,omitempty" asn1:"optional"`
		SpecVersion string `xml:"specVersion,omitempty" json:"specVersion,omitempty" asn1:"optional"`
		DbSnpBuild  int64  `xml:"dbSnpBuild,omitempty" json:"dbSnpBuild,omitempty" asn1:"optional"`
		Generated   string `xml:"generated,omitempty" json:"generated,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	SourceDatabase struct {
		Attlist struct {
			TaxId        int64  `xml:"taxId" json:"taxId"`
			Organism     string `xml:"organism" json:"organism"`
			DbSnpOrgAbbr string `xml:"dbSnpOrgAbbr,omitempty" json:"dbSnpOrgAbbr,omitempty" asn1:"optional"`
			GpipeOrgAbbr string `xml:"gpipeOrgAbbr,omitempty" json:"gpipeOrgAbbr,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		SourceDatabase interface{} `xml:"-" json:"-"` //SourceDatabase,NullType
	} `xml:"sourceDatabase,omitempty" json:"sourceDatabase,omitempty" asn1:"optional"`
	Rs    []Rs   `xml:"rs,omitempty" json:"rs,omitempty" asn1:"optional"`
	Assay *Assay `xml:"assay,omitempty" json:"assay,omitempty" asn1:"optional"`
	Query struct {
		Attlist struct {
			Date   string `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
			String string `xml:"string,omitempty" json:"string,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		Query interface{} `xml:"-" json:"-"` //Query,NullType
	} `xml:"query,omitempty" json:"query,omitempty" asn1:"optional"`
	Summary struct {
		Attlist struct {
			NumRsIds       int64 `xml:"numRsIds,omitempty" json:"numRsIds,omitempty" asn1:"optional"`
			TotalSeqLength int64 `xml:"totalSeqLength,omitempty" json:"totalSeqLength,omitempty" asn1:"optional"`
			NumContigHits  int64 `xml:"numContigHits,omitempty" json:"numContigHits,omitempty" asn1:"optional"`
			NumGeneHits    int64 `xml:"numGeneHits,omitempty" json:"numGeneHits,omitempty" asn1:"optional"`
			NumGiHits      int64 `xml:"numGiHits,omitempty" json:"numGiHits,omitempty" asn1:"optional"`
			Num3dStructs   int64 `xml:"num3dStructs,omitempty" json:"num3dStructs,omitempty" asn1:"optional"`
			NumAlleleFreqs int64 `xml:"numAlleleFreqs,omitempty" json:"numAlleleFreqs,omitempty" asn1:"optional"`
			NumStsHits     int64 `xml:"numStsHits,omitempty" json:"numStsHits,omitempty" asn1:"optional"`
			NumUnigeneCids int64 `xml:"numUnigeneCids,omitempty" json:"numUnigeneCids,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		Summary interface{} `xml:"-" json:"-"` //Summary,NullType
	} `xml:"summary,omitempty" json:"summary,omitempty" asn1:"optional"`
	BaseURL []BaseURL `xml:"baseURL,omitempty" json:"baseURL,omitempty" asn1:"optional"`
}
type FxnSet struct {
	Attlist struct {
		GeneId       int64  `xml:"geneId,omitempty" json:"geneId,omitempty" asn1:"optional"`
		Symbol       string `xml:"symbol,omitempty" json:"symbol,omitempty" asn1:"optional"`
		MrnaAcc      string `xml:"mrnaAcc,omitempty" json:"mrnaAcc,omitempty" asn1:"optional"`
		MrnaVer      int64  `xml:"mrnaVer,omitempty" json:"mrnaVer,omitempty" asn1:"optional"`
		ProtAcc      string `xml:"protAcc,omitempty" json:"protAcc,omitempty" asn1:"optional"`
		ProtVer      int64  `xml:"protVer,omitempty" json:"protVer,omitempty" asn1:"optional"`
		FxnClass     string `xml:"fxnClass,omitempty" json:"fxnClass,omitempty" asn1:"optional"` //FxnClass,EnumList:locus-region(1),coding-unknown(2),synonymous-codon(3),non-synonymous-codon(4),mrna-utr(5),intron-variant(6),splice-region-variant(7),reference(8),coding-exception(9),coding-sequence-variant(10),nc-transcript-variant(11),downstream-variant-500B(12),upstream-variant-2KB(13),nonsense(14),missense(15),frameshift-variant(16),utr-variant-3-prime(17),utr-variant-5-prime(18),splice-acceptor-variant(19),splice-donor-variant(20),cds-indel(21),stop-gained(22),stop-lost(23),complex-change-in-transcript(24),incomplete-terminal-codon-variant(25),nmd-transcript-variant(26),mature-miRNA-variant(27),upstream-variant-5KB(28),downstream-variant-5KB(29),intergenic(30)
		ReadingFrame int64  `xml:"readingFrame,omitempty" json:"readingFrame,omitempty" asn1:"optional"`
		Allele       string `xml:"allele,omitempty" json:"allele,omitempty" asn1:"optional"`
		Residue      string `xml:"residue,omitempty" json:"residue,omitempty" asn1:"optional"`
		AaPosition   int64  `xml:"aaPosition,omitempty" json:"aaPosition,omitempty" asn1:"optional"`
		MrnaPosition int64  `xml:"mrnaPosition,omitempty" json:"mrnaPosition,omitempty" asn1:"optional"`
		SoTerm       string `xml:"soTerm,omitempty" json:"soTerm,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	FxnSet interface{} `xml:"-" json:"-"` //FxnSet,NullType
}
type MapLoc struct {
	Attlist struct {
		AsnFrom                int64   `xml:"asnFrom" json:"asnFrom"`
		AsnTo                  int64   `xml:"asnTo" json:"asnTo"`
		LocType                string  `xml:"locType" json:"locType"` //LocType,EnumList:insertion(1),exact(2),deletion(3),range-ins(4),range-exact(5),range-del(6)
		AlnQuality             float64 `xml:"alnQuality,omitempty" json:"alnQuality,omitempty" asn1:"optional"`
		Orient                 string  `xml:"orient,omitempty" json:"orient,omitempty" asn1:"optional"` //Orient,EnumList:forward(1),reverse(2)
		PhysMapInt             int64   `xml:"physMapInt,omitempty" json:"physMapInt,omitempty" asn1:"optional"`
		LeftFlankNeighborPos   int64   `xml:"leftFlankNeighborPos,omitempty" json:"leftFlankNeighborPos,omitempty" asn1:"optional"`
		RightFlankNeighborPos  int64   `xml:"rightFlankNeighborPos,omitempty" json:"rightFlankNeighborPos,omitempty" asn1:"optional"`
		LeftContigNeighborPos  int64   `xml:"leftContigNeighborPos,omitempty" json:"leftContigNeighborPos,omitempty" asn1:"optional"`
		RightContigNeighborPos int64   `xml:"rightContigNeighborPos,omitempty" json:"rightContigNeighborPos,omitempty" asn1:"optional"`
		NumberOfMismatches     int64   `xml:"numberOfMismatches,omitempty" json:"numberOfMismatches,omitempty" asn1:"optional"`
		NumberOfDeletions      int64   `xml:"numberOfDeletions,omitempty" json:"numberOfDeletions,omitempty" asn1:"optional"`
		NumberOfInsertions     int64   `xml:"numberOfInsertions,omitempty" json:"numberOfInsertions,omitempty" asn1:"optional"`
		RefAllele              string  `xml:"refAllele,omitempty" json:"refAllele,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	FxnSet []FxnSet `xml:"fxnSet,omitempty" json:"fxnSet,omitempty" asn1:"optional"`
}
type PrimarySequence struct {
	Attlist struct {
		DbSnpBuild int64  `xml:"dbSnpBuild" json:"dbSnpBuild"`
		Gi         int64  `xml:"gi" json:"gi"`
		Source     string `xml:"source,omitempty" json:"source,omitempty" asn1:"optional"` //Source,EnumList:submitter(1),blastmb(2),xm(3),remap(4),hgvs(5)
		Accession  string `xml:"accession,omitempty" json:"accession,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	MapLoc []MapLoc `xml:"mapLoc,omitempty" json:"mapLoc,omitempty"`
}
type Rs struct {
	Attlist struct {
		RsId         int64  `xml:"rsId" json:"rsId"`
		SnpClass     string `xml:"snpClass" json:"snpClass"` //SnpClass,EnumList:snp(1),in-del(2),heterozygous(3),microsatellite(4),named-locus(5),no-variation(6),mixed(7),multinucleotide-polymorphism(8)
		SnpType      string `xml:"snpType" json:"snpType"`   //SnpType,EnumList:notwithdrawn(1),artifact(2),gene-duplication(3),duplicate-submission(4),notspecified(5),ambiguous-location(6),low-map-quality(7)
		MolType      string `xml:"molType" json:"molType"`   //MolType,EnumList:genomic(1),cDNA(2),mito(3),chloro(4),unknown(5)
		ValidProbMin int64  `xml:"validProbMin,omitempty" json:"validProbMin,omitempty" asn1:"optional"`
		ValidProbMax int64  `xml:"validProbMax,omitempty" json:"validProbMax,omitempty" asn1:"optional"`
		Genotype     bool   `xml:"genotype,omitempty" json:"genotype,omitempty" asn1:"optional"`
		BitField     string `xml:"bitField,omitempty" json:"bitField,omitempty" asn1:"optional"`
		TaxId        int64  `xml:"taxId,omitempty" json:"taxId,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	Het struct {
		Attlist struct {
			Type     string  `xml:"type" json:"type"` //Type,EnumList:est(1),obs(2)
			Value    float64 `xml:"value" json:"value"`
			StdError float64 `xml:"stdError,omitempty" json:"stdError,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		Het interface{} `xml:"-" json:"-"` //Het,NullType
	} `xml:"het,omitempty" json:"het,omitempty" asn1:"optional"`
	Validation struct {
		Attlist struct {
			ByCluster     bool `xml:"byCluster,omitempty" json:"byCluster,omitempty" asn1:"optional"`
			ByFrequency   bool `xml:"byFrequency,omitempty" json:"byFrequency,omitempty" asn1:"optional"`
			ByOtherPop    bool `xml:"byOtherPop,omitempty" json:"byOtherPop,omitempty" asn1:"optional"`
			By2Hit2Allele bool `xml:"by2Hit2Allele,omitempty" json:"by2Hit2Allele,omitempty" asn1:"optional"`
			ByHapMap      bool `xml:"byHapMap,omitempty" json:"byHapMap,omitempty" asn1:"optional"`
			By1000G       bool `xml:"by1000G,omitempty" json:"by1000G,omitempty" asn1:"optional"`
			Suspect       bool `xml:"suspect,omitempty" json:"suspect,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		OtherPopBatchId      []int64  `xml:"otherPopBatchId,omitempty" json:"otherPopBatchId,omitempty" asn1:"optional"`
		TwoHit2AlleleBatchId []int64  `xml:"twoHit2AlleleBatchId,omitempty" json:"twoHit2AlleleBatchId,omitempty" asn1:"optional"`
		FrequencyClass       []int64  `xml:"frequencyClass,omitempty" json:"frequencyClass,omitempty" asn1:"optional"`
		HapMapPhase          []int64  `xml:"hapMapPhase,omitempty" json:"hapMapPhase,omitempty" asn1:"optional"`
		TGPPhase             []int64  `xml:"tGPPhase,omitempty" json:"tGPPhase,omitempty" asn1:"optional"`
		SuspectEvidence      []string `xml:"suspectEvidence,omitempty" json:"suspectEvidence,omitempty" asn1:"optional"`
	} `xml:"validation" json:"validation"`
	Create struct {
		Attlist struct {
			Build int64  `xml:"build,omitempty" json:"build,omitempty" asn1:"optional"`
			Date  string `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		Create interface{} `xml:"-" json:"-"` //Create,NullType
	} `xml:"create" json:"create"`
	Update struct {
		Attlist struct {
			Build int64  `xml:"build,omitempty" json:"build,omitempty" asn1:"optional"`
			Date  string `xml:"date,omitempty" json:"date,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		Update interface{} `xml:"-" json:"-"` //Update,NullType
	} `xml:"update,omitempty" json:"update,omitempty" asn1:"optional"`
	Sequence struct {
		Attlist struct {
			ExemplarSs      int64  `xml:"exemplarSs" json:"exemplarSs"`
			AncestralAllele string `xml:"ancestralAllele,omitempty" json:"ancestralAllele,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		Seq5     string `xml:"seq5,omitempty" json:"seq5,omitempty" asn1:"optional"`
		Observed string `xml:"observed" json:"observed"`
		Seq3     string `xml:"seq3,omitempty" json:"seq3,omitempty" asn1:"optional"`
	} `xml:"sequence" json:"sequence"`
	Ss              []Ss              `xml:"ss,omitempty" json:"ss,omitempty"`
	Assembly        []Assembly        `xml:"assembly,omitempty" json:"assembly,omitempty" asn1:"optional"`
	PrimarySequence []PrimarySequence `xml:"primarySequence,omitempty" json:"primarySequence,omitempty" asn1:"optional"`
	RsStruct        []RsStruct        `xml:"rsStruct,omitempty" json:"rsStruct,omitempty" asn1:"optional"`
	RsLinkout       []RsLinkout       `xml:"rsLinkout,omitempty" json:"rsLinkout,omitempty" asn1:"optional"`
	MergeHistory    []struct {
		Attlist struct {
			RsId       int64 `xml:"rsId" json:"rsId"`
			BuildId    int64 `xml:"buildId,omitempty" json:"buildId,omitempty" asn1:"optional"`
			OrientFlip bool  `xml:"orientFlip,omitempty" json:"orientFlip,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		MergeHistory interface{} `xml:"-" json:"-"` //MergeHistory,NullType
	} `xml:"mergeHistory,omitempty" json:"mergeHistory,omitempty" asn1:"optional"`
	Hgvs         []string `xml:"hgvs,omitempty" json:"hgvs,omitempty" asn1:"optional"`
	AlleleOrigin []struct {
		Attlist struct {
			Allele string `xml:"allele,omitempty" json:"allele,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		AlleleOrigin int64 `xml:"alleleOrigin" json:"alleleOrigin"`
	} `xml:"alleleOrigin,omitempty" json:"alleleOrigin,omitempty" asn1:"optional"`
	Phenotype []struct {
		ClinicalSignificance []string `xml:"clinicalSignificance,omitempty" json:"clinicalSignificance,omitempty" asn1:"optional"`
	} `xml:"phenotype,omitempty" json:"phenotype,omitempty" asn1:"optional"`
	BioSource []struct {
		Genome []string `xml:"genome,omitempty" json:"genome,omitempty" asn1:"optional"`
		Origin []string `xml:"origin,omitempty" json:"origin,omitempty" asn1:"optional"`
	} `xml:"bioSource,omitempty" json:"bioSource,omitempty" asn1:"optional"`
	Frequency []struct {
		Attlist struct {
			Freq       float64 `xml:"freq,omitempty" json:"freq,omitempty" asn1:"optional"`
			Allele     string  `xml:"allele,omitempty" json:"allele,omitempty" asn1:"optional"`
			PopId      int64   `xml:"popId,omitempty" json:"popId,omitempty" asn1:"optional"`
			SampleSize int64   `xml:"sampleSize,omitempty" json:"sampleSize,omitempty" asn1:"optional"`
		} `xml:"attlist" json:"attlist"`
		Frequency interface{} `xml:"-" json:"-"` //Frequency,NullType
	} `xml:"frequency,omitempty" json:"frequency,omitempty" asn1:"optional"`
}
type RsLinkout struct {
	Attlist struct {
		ResourceId string `xml:"resourceId" json:"resourceId"`
		LinkValue  string `xml:"linkValue" json:"linkValue"`
	} `xml:"attlist" json:"attlist"`
	RsLinkout interface{} `xml:"-" json:"-"` //RsLinkout,NullType
}
type RsStruct struct {
	Attlist struct {
		ProtAcc       string `xml:"protAcc,omitempty" json:"protAcc,omitempty" asn1:"optional"`
		ProtGi        int64  `xml:"protGi,omitempty" json:"protGi,omitempty" asn1:"optional"`
		ProtLoc       int64  `xml:"protLoc,omitempty" json:"protLoc,omitempty" asn1:"optional"`
		ProtResidue   string `xml:"protResidue,omitempty" json:"protResidue,omitempty" asn1:"optional"`
		RsResidue     string `xml:"rsResidue,omitempty" json:"rsResidue,omitempty" asn1:"optional"`
		StructGi      int64  `xml:"structGi,omitempty" json:"structGi,omitempty" asn1:"optional"`
		StructLoc     int64  `xml:"structLoc,omitempty" json:"structLoc,omitempty" asn1:"optional"`
		StructResidue string `xml:"structResidue,omitempty" json:"structResidue,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	RsStruct interface{} `xml:"-" json:"-"` //RsStruct,NullType
}
type Ss struct {
	Attlist struct {
		SsId                 int64  `xml:"ssId" json:"ssId"`
		Handle               string `xml:"handle" json:"handle"`
		BatchId              int64  `xml:"batchId" json:"batchId"`
		LocSnpId             string `xml:"locSnpId,omitempty" json:"locSnpId,omitempty" asn1:"optional"`
		SubSnpClass          string `xml:"subSnpClass,omitempty" json:"subSnpClass,omitempty" asn1:"optional"` //SubSnpClass,EnumList:snp(1),in-del(2),heterozygous(3),microsatellite(4),named-locus(5),no-variation(6),mixed(7),multinucleotide-polymorphism(8)
		Orient               string `xml:"orient,omitempty" json:"orient,omitempty" asn1:"optional"`           //Orient,EnumList:forward(1),reverse(2)
		Strand               string `xml:"strand,omitempty" json:"strand,omitempty" asn1:"optional"`           //Strand,EnumList:top(1),bottom(2)
		MolType              string `xml:"molType,omitempty" json:"molType,omitempty" asn1:"optional"`         //MolType,EnumList:genomic(1),cDNA(2),mito(3),chloro(4),unknown(5)
		BuildId              int64  `xml:"buildId,omitempty" json:"buildId,omitempty" asn1:"optional"`
		MethodClass          string `xml:"methodClass,omitempty" json:"methodClass,omitempty" asn1:"optional"` //MethodClass,EnumList:dHPLC(1),hybridize(2),computed(3),sSCP(4),other(5),unknown(6),rFLP(7),sequence(8)
		Validated            string `xml:"validated,omitempty" json:"validated,omitempty" asn1:"optional"`     //Validated,EnumList:by-submitter(1),by-frequency(2),by-cluster(3)
		LinkoutUrl           string `xml:"linkoutUrl,omitempty" json:"linkoutUrl,omitempty" asn1:"optional"`
		SsAlias              string `xml:"ssAlias,omitempty" json:"ssAlias,omitempty" asn1:"optional"`
		AlleleOrigin         int64  `xml:"alleleOrigin,omitempty" json:"alleleOrigin,omitempty" asn1:"optional"`
		ClinicalSignificance string `xml:"clinicalSignificance,omitempty" json:"clinicalSignificance,omitempty" asn1:"optional"`
	} `xml:"attlist" json:"attlist"`
	Sequence struct {
		Seq5     string `xml:"seq5,omitempty" json:"seq5,omitempty" asn1:"optional"`
		Observed string `xml:"observed" json:"observed"`
		Seq3     string `xml:"seq3,omitempty" json:"seq3,omitempty" asn1:"optional"`
	} `xml:"sequence" json:"sequence"`
}
