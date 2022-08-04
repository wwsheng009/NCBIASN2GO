package MMDB

import (
	"ncbiasn/MMDBChemicalGraph"
	"ncbiasn/MMDBFeatures"
	"ncbiasn/MMDBStructuralModel"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIPub"
)

type Biostruc struct {
	Id            []BiostrucId                        `xml:"id,omitempty" json:"id,omitempty"`
	Descr         []BiostrucDescr                     `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	ChemicalGraph *MMDBChemicalGraph.BiostrucGraph    `xml:"chemical-graph,omitempty" json:"chemical_graph,omitempty"`
	Features      []MMDBFeatures.BiostrucFeatureSet   `xml:"features,omitempty" json:"features,omitempty" asn1:"optional"`
	Model         []MMDBStructuralModel.BiostrucModel `xml:"model,omitempty" json:"model,omitempty" asn1:"optional"`
}

//BiostrucId,ChoiceOption
type BiostrucId struct {
	MmdbId        *MmdbId               `xml:"mmdb-id,omitempty" json:"mmdb_id,omitempty"`
	OtherDatabase *NCBIGeneral.Dbtag    `xml:"other-database,omitempty" json:"other_database,omitempty"`
	LocalId       *NCBIGeneral.ObjectId `xml:"local-id,omitempty" json:"local_id,omitempty"`
}
type MmdbId int64
type BiostrucDescr struct {
	Name         string           `xml:"name,omitempty" json:"name,omitempty"`
	PdbComment   string           `xml:"pdb-comment,omitempty" json:"pdb_comment,omitempty"`
	OtherComment string           `xml:"other-comment,omitempty" json:"other_comment,omitempty"`
	History      *BiostrucHistory `xml:"history,omitempty" json:"history,omitempty"`
	Attribution  *NCBIPub.Pub     `xml:"attribution,omitempty" json:"attribution,omitempty"`
}

//BiostrucDescr,ChoiceOption
type BiostrucHistory struct {
	Replaces   *BiostrucReplace `xml:"replaces,omitempty" json:"replaces,omitempty" asn1:"optional"`
	ReplacedBy *BiostrucReplace `xml:"replaced-by,omitempty" json:"replaced_by,omitempty" asn1:"optional"`
	DataSource *BiostrucSource  `xml:"data-source,omitempty" json:"data_source,omitempty" asn1:"optional"`
}
type BiostrucReplace struct {
	Id   *BiostrucId       `xml:"id,omitempty" json:"id,omitempty"`
	Date *NCBIGeneral.Date `xml:"date,omitempty" json:"date,omitempty"`
}
type BiostrucSource struct {
	NameOfDatabase string `xml:"name-of-database" json:"name_of_database"`
	//VersionOfDatabase,ChoiceOption
	VersionOfDatabase struct {
		ReleaseDate *NCBIGeneral.Date `xml:"release-date,omitempty" json:"release_date,omitempty"`
		ReleaseCode string            `xml:"release-code,omitempty" json:"release_code,omitempty"`
	} `xml:"version-of-database,omitempty" json:"version_of_database,omitempty" asn1:"optional"`
	DatabaseEntryId      *BiostrucId       `xml:"database-entry-id,omitempty" json:"database_entry_id,omitempty"`
	DatabaseEntryDate    *NCBIGeneral.Date `xml:"database-entry-date,omitempty" json:"database_entry_date,omitempty"`
	DatabaseEntryHistory []string          `xml:"database-entry-history,omitempty" json:"database_entry_history,omitempty" asn1:"optional"`
}
type BiostrucSet struct {
	Id        []BiostrucId    `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Descr     []BiostrucDescr `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Biostrucs []Biostruc      `xml:"biostrucs,omitempty" json:"biostrucs,omitempty"`
}
type BiostrucAnnotSet struct {
	Id       []BiostrucId                      `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Descr    []BiostrucDescr                   `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Features []MMDBFeatures.BiostrucFeatureSet `xml:"features,omitempty" json:"features,omitempty"`
}
type BiostrucResidueGraphSet struct {
	Id            []BiostrucId                     `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Descr         []MMDBChemicalGraph.BiomolDescr  `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	ResidueGraphs []MMDBChemicalGraph.ResidueGraph `xml:"residue-graphs,omitempty" json:"residue_graphs,omitempty"`
}
