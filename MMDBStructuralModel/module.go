package MMDBStructuralModel

import (
	"ncbiasn/MMDBCommon"
	"ncbiasn/MMDBFeatures"
	"ncbiasn/NCBIPub"
)

type BiostrucModel struct {
	Id               *ModelId             `xml:"id,omitempty" json:"id,omitempty"`
	Type             *ModelType           `xml:"type,omitempty" json:"type,omitempty"`
	Descr            []ModelDescr         `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	ModelSpace       *ModelSpace          `xml:"model-space,omitempty" json:"model_space,omitempty" asn1:"optional"`
	ModelCoordinates []ModelCoordinateSet `xml:"model-coordinates,omitempty" json:"model_coordinates,omitempty" asn1:"optional"`
}
type ModelId int64

//ModelType,IntegerEnum:ncbi-vector(1),ncbi-backbone(2),ncbi-all-atom(3),pdb-model(4),other(255)
type ModelType int

//ModelDescr,ChoiceOption
type ModelDescr struct {
	Name         string       `xml:"name,omitempty" json:"name,omitempty"`
	PdbReso      string       `xml:"pdb-reso,omitempty" json:"pdb_reso,omitempty"`
	PdbMethod    string       `xml:"pdb-method,omitempty" json:"pdb_method,omitempty"`
	PdbComment   string       `xml:"pdb-comment,omitempty" json:"pdb_comment,omitempty"`
	OtherComment string       `xml:"other-comment,omitempty" json:"other_comment,omitempty"`
	Attribution  *NCBIPub.Pub `xml:"attribution,omitempty" json:"attribution,omitempty"`
}

type ModelSpace struct {
	CoordinateUnits      string          `xml:"coordinate-units" json:"coordinate_units"`                                                 //CoordinateUnits,EnumList:angstroms(1),nanometers(2),other(3),unknown(255)
	ThermalFactorUnits   string          `xml:"thermal-factor-units,omitempty" json:"thermal_factor_units,omitempty" asn1:"optional"`     //ThermalFactorUnits,EnumList:b(1),u(2),other(3),unknown(255)
	OccupancyFactorUnits string          `xml:"occupancy-factor-units,omitempty" json:"occupancy_factor_units,omitempty" asn1:"optional"` //OccupancyFactorUnits,EnumList:fractional(1),electrons(2),other(3),unknown(255)
	DensityUnits         string          `xml:"density-units,omitempty" json:"density_units,omitempty" asn1:"optional"`                   //DensityUnits,EnumList:electrons-per-unit-volume(1),arbitrary-scale(2),other(3),unknown(255)
	ReferenceFrame       *ReferenceFrame `xml:"reference-frame,omitempty" json:"reference_frame,omitempty" asn1:"optional"`
}
type ReferenceFrame struct {
	BiostrucId          *MMDBCommon.BiostrucId  `xml:"biostruc-id,omitempty" json:"biostruc_id,omitempty"`
	RotationTranslation *MMDBFeatures.Transform `xml:"rotation-translation,omitempty" json:"rotation_translation,omitempty" asn1:"optional"`
}
type ModelCoordinateSet struct {
	Id          *ModelCoordinateSetId `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Descr       []ModelDescr          `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Coordinates struct {
		Literal   *Coordinates                     `xml:"literal,omitempty" json:"literal,omitempty"`
		Reference *MMDBFeatures.ChemGraphAlignment `xml:"reference,omitempty" json:"reference,omitempty"`
	} `xml:"coordinates" json:"coordinates"` //Coordinates,ChoiceOption
}
type ModelCoordinateSetId int64
type Coordinates struct {
	Atomic  *AtomicCoordinates  `xml:"atomic,omitempty" json:"atomic,omitempty"`
	Surface *SurfaceCoordinates `xml:"surface,omitempty" json:"surface,omitempty"`
	Density *DensityCoordinates `xml:"density,omitempty" json:"density,omitempty"`
}

//Coordinates,ChoiceOption
type AtomicCoordinates struct {
	NumberOfPoints     int64                     `xml:"number-of-points" json:"number_of_points"`
	Atoms              *MMDBFeatures.AtomPntrs   `xml:"atoms,omitempty" json:"atoms,omitempty"`
	Sites              *ModelSpacePoints         `xml:"sites,omitempty" json:"sites,omitempty"`
	TemperatureFactors *AtomicTemperatureFactors `xml:"temperature-factors,omitempty" json:"temperature_factors,omitempty" asn1:"optional"`
	Occupancies        *AtomicOccupancies        `xml:"occupancies,omitempty" json:"occupancies,omitempty" asn1:"optional"`
	AlternateConfIds   AlternateConformationIds  `xml:"alternate-conf-ids,omitempty" json:"alternate_conf_ids,omitempty" asn1:"optional"`
	ConfEnsembles      []ConformationEnsemble    `xml:"conf-ensembles,omitempty" json:"conf_ensembles,omitempty" asn1:"optional"`
}
type ModelSpacePoints struct {
	ScaleFactor int64   `xml:"scale-factor" json:"scale_factor"`
	X           []int64 `xml:"x,omitempty" json:"x,omitempty"`
	Y           []int64 `xml:"y,omitempty" json:"y,omitempty"`
	Z           []int64 `xml:"z,omitempty" json:"z,omitempty"`
}
type AtomicTemperatureFactors struct {
	Isotropic   *IsotropicTemperatureFactors   `xml:"isotropic,omitempty" json:"isotropic,omitempty"`
	Anisotropic *AnisotropicTemperatureFactors `xml:"anisotropic,omitempty" json:"anisotropic,omitempty"`
}

//AtomicTemperatureFactors,ChoiceOption
type IsotropicTemperatureFactors struct {
	ScaleFactor int64   `xml:"scale-factor" json:"scale_factor"`
	B           []int64 `xml:"b,omitempty" json:"b,omitempty"`
}
type AnisotropicTemperatureFactors struct {
	ScaleFactor int64   `xml:"scale-factor" json:"scale_factor"`
	B11         []int64 `xml:"b-11,omitempty" json:"b_11,omitempty"`
	B12         []int64 `xml:"b-12,omitempty" json:"b_12,omitempty"`
	B13         []int64 `xml:"b-13,omitempty" json:"b_13,omitempty"`
	B22         []int64 `xml:"b-22,omitempty" json:"b_22,omitempty"`
	B23         []int64 `xml:"b-23,omitempty" json:"b_23,omitempty"`
	B33         []int64 `xml:"b-33,omitempty" json:"b_33,omitempty"`
}
type AtomicOccupancies struct {
	ScaleFactor int64   `xml:"scale-factor" json:"scale_factor"`
	O           []int64 `xml:"o,omitempty" json:"o,omitempty"`
}
type AlternateConformationIds []AlternateConformationId
type AlternateConformationId string
type ConformationEnsemble struct {
	Name       string                    `xml:"name" json:"name"`
	AltConfIds []AlternateConformationId `xml:"alt-conf-ids,omitempty" json:"alt_conf_ids,omitempty"`
}
type SurfaceCoordinates struct {
	Contents *MMDBFeatures.ChemGraphPntrs `xml:"contents,omitempty" json:"contents,omitempty"`
	Surface  struct {
		Sphere    *MMDBFeatures.Sphere   `xml:"sphere,omitempty" json:"sphere,omitempty"`
		Cone      *MMDBFeatures.Cone     `xml:"cone,omitempty" json:"cone,omitempty"`
		Cylinder  *MMDBFeatures.Cylinder `xml:"cylinder,omitempty" json:"cylinder,omitempty"`
		Brick     *MMDBFeatures.Brick    `xml:"brick,omitempty" json:"brick,omitempty"`
		Tmesh     *TMesh                 `xml:"tmesh,omitempty" json:"tmesh,omitempty"`
		Triangles *Triangles             `xml:"triangles,omitempty" json:"triangles,omitempty"`
	} `xml:"surface" json:"surface"` //Surface,ChoiceOption
}
type TMesh struct {
	NumberOfPoints int64   `xml:"number-of-points" json:"number_of_points"`
	ScaleFactor    int64   `xml:"scale-factor" json:"scale_factor"`
	Swap           []bool  `xml:"swap,omitempty" json:"swap,omitempty"`
	X              []int64 `xml:"x,omitempty" json:"x,omitempty"`
	Y              []int64 `xml:"y,omitempty" json:"y,omitempty"`
	Z              []int64 `xml:"z,omitempty" json:"z,omitempty"`
}
type Triangles struct {
	NumberOfPoints    int64   `xml:"number-of-points" json:"number_of_points"`
	ScaleFactor       int64   `xml:"scale-factor" json:"scale_factor"`
	X                 []int64 `xml:"x,omitempty" json:"x,omitempty"`
	Y                 []int64 `xml:"y,omitempty" json:"y,omitempty"`
	Z                 []int64 `xml:"z,omitempty" json:"z,omitempty"`
	NumberOfTriangles int64   `xml:"number-of-triangles" json:"number_of_triangles"`
	V1                []int64 `xml:"v1,omitempty" json:"v1,omitempty"`
	V2                []int64 `xml:"v2,omitempty" json:"v2,omitempty"`
	V3                []int64 `xml:"v3,omitempty" json:"v3,omitempty"`
}
type DensityCoordinates struct {
	Contents       *MMDBFeatures.ChemGraphPntrs `xml:"contents,omitempty" json:"contents,omitempty"`
	GridCorners    *MMDBFeatures.Brick          `xml:"grid-corners,omitempty" json:"grid_corners,omitempty"`
	GridStepsX     int64                        `xml:"grid-steps-x" json:"grid_steps_x"`
	GridStepsY     int64                        `xml:"grid-steps-y" json:"grid_steps_y"`
	GridStepsZ     int64                        `xml:"grid-steps-z" json:"grid_steps_z"`
	FastestVarying string                       `xml:"fastest-varying" json:"fastest_varying"` //FastestVarying,EnumList:x(1),y(2),z(3)
	SlowestVarying string                       `xml:"slowest-varying" json:"slowest_varying"` //SlowestVarying,EnumList:x(1),y(2),z(3)
	ScaleFactor    int64                        `xml:"scale-factor" json:"scale_factor"`
	Density        []int64                      `xml:"density,omitempty" json:"density,omitempty"`
}
