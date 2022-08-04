package MMDBCommon

import (
	"ncbiasn/NCBIGeneral"
)

type MmdbId int64
type BiostrucId struct {
	MmdbId        MmdbId                `xml:"mmdb-id,omitempty" json:"mmdb_id,omitempty"`
	OtherDatabase *NCBIGeneral.Dbtag    `xml:"other-database,omitempty" json:"other_database,omitempty"`
	LocalId       *NCBIGeneral.ObjectId `xml:"local-id,omitempty" json:"local_id,omitempty"`
}
type ModelId int64
type ModelCoordinateSetId int64
