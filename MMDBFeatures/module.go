package MMDBFeatures

import (
	"ncbiasn/MMDBChemicalGraph"
	"ncbiasn/MMDBCommon"
	"ncbiasn/NCBIGeneral"
	"ncbiasn/NCBIPub"
)

type BiostrucFeatureSet struct {
	Id       *BiostrucFeatureSetId     `xml:"id,omitempty" json:"id,omitempty"`
	Descr    []BiostrucFeatureSetDescr `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	Features []BiostrucFeature         `xml:"features,omitempty" json:"features,omitempty"`
}
type BiostrucFeatureSetId int64

//BiostrucFeatureSetDescr,ChoiceOption
type BiostrucFeatureSetDescr struct {
	Name         string       `xml:"name,omitempty" json:"name,omitempty"`
	PdbComment   string       `xml:"pdb-comment,omitempty" json:"pdb_comment,omitempty"`
	OtherComment string       `xml:"other-comment,omitempty" json:"other_comment,omitempty"`
	Attribution  *NCBIPub.Pub `xml:"attribution,omitempty" json:"attribution,omitempty"`
}

type BiostrucFeature struct {
	Id   *BiostrucFeatureId `xml:"id,omitempty" json:"id,omitempty" asn1:"optional"`
	Name string             `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	Type int                `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"` //Type,IntegerEnum:helix(1),strand(2),sheet(3),turn(4),site(5),footnote(6),comment(7),interaction(8),subgraph(100),region(101),core(102),supercore(103),color(150),render(151),label(152),transform(153),camera(154),script(155),alignment(200),similarity(201),multalign(202),indirect(203),cn3dstate(254),other(255)
	//Property,ChoiceOption
	Property struct {
		Color     *ColorProp              `xml:"color,omitempty" json:"color,omitempty"`
		Render    *RenderProp             `xml:"render,omitempty" json:"render,omitempty"`
		Transform *Transform              `xml:"transform,omitempty" json:"transform,omitempty"`
		Camera    *Camera                 `xml:"camera,omitempty" json:"camera,omitempty"`
		Script    BiostrucScript          `xml:"script,omitempty" json:"script,omitempty"`
		User      *NCBIGeneral.UserObject `xml:"user,omitempty" json:"user,omitempty"`
	} `xml:"property,omitempty" json:"property,omitempty" asn1:"optional"`
	//Location,ChoiceOption
	Location struct {
		Subgraph    *ChemGraphPntrs       `xml:"subgraph,omitempty" json:"subgraph,omitempty"`
		Region      *RegionPntrs          `xml:"region,omitempty" json:"region,omitempty"`
		Alignment   *ChemGraphAlignment   `xml:"alignment,omitempty" json:"alignment,omitempty"`
		Interaction *ChemGraphInteraction `xml:"interaction,omitempty" json:"interaction,omitempty"`
		Similarity  *RegionSimilarity     `xml:"similarity,omitempty" json:"similarity,omitempty"`
		Indirect    *OtherFeature         `xml:"indirect,omitempty" json:"indirect,omitempty"`
	} `xml:"location,omitempty" json:"location,omitempty" asn1:"optional"`
}
type OtherFeature struct {
	BiostrucId *MMDBCommon.BiostrucId `xml:"biostruc-id,omitempty" json:"biostruc_id,omitempty"`
	Set        *BiostrucFeatureSetId  `xml:"set,omitempty" json:"set,omitempty"`
	Feature    *BiostrucFeatureId     `xml:"feature,omitempty" json:"feature,omitempty"`
}
type BiostrucFeatureId int64
type BiostrucMoleculePntr struct {
	BiostrucId *MMDBCommon.BiostrucId        `xml:"biostruc-id,omitempty" json:"biostruc_id,omitempty"`
	MoleculeId *MMDBChemicalGraph.MoleculeId `xml:"molecule-id,omitempty" json:"molecule_id,omitempty"`
}

//ChemGraphPntrs,ChoiceOption
type ChemGraphPntrs struct {
	Atoms     *AtomPntrs     `xml:"atoms,omitempty" json:"atoms,omitempty"`
	Residues  *ResiduePntrs  `xml:"residues,omitempty" json:"residues,omitempty"`
	Molecules *MoleculePntrs `xml:"molecules,omitempty" json:"molecules,omitempty"`
}

type AtomPntrs struct {
	NumberOfPtrs int64                          `xml:"number-of-ptrs" json:"number_of_ptrs"`
	MoleculeIds  []MMDBChemicalGraph.MoleculeId `xml:"molecule-ids,omitempty" json:"molecule_ids,omitempty"`
	ResidueIds   []MMDBChemicalGraph.ResidueId  `xml:"residue-ids,omitempty" json:"residue_ids,omitempty"`
	AtomIds      []MMDBChemicalGraph.AtomId     `xml:"atom-ids,omitempty" json:"atom_ids,omitempty"`
}

//ResiduePntrs,ChoiceOption
type ResiduePntrs struct {
	Explicit *ResidueExplicitPntrs `xml:"explicit,omitempty" json:"explicit,omitempty"`
	Interval []ResidueIntervalPntr `xml:"interval,omitempty" json:"interval,omitempty"`
}

type ResidueExplicitPntrs struct {
	NumberOfPtrs int64                          `xml:"number-of-ptrs" json:"number_of_ptrs"`
	MoleculeIds  []MMDBChemicalGraph.MoleculeId `xml:"molecule-ids,omitempty" json:"molecule_ids,omitempty"`
	ResidueIds   []MMDBChemicalGraph.ResidueId  `xml:"residue-ids,omitempty" json:"residue_ids,omitempty"`
}
type ResidueIntervalPntr struct {
	MoleculeId *MMDBChemicalGraph.MoleculeId `xml:"molecule-id,omitempty" json:"molecule_id,omitempty"`
	From       *MMDBChemicalGraph.ResidueId  `xml:"from,omitempty" json:"from,omitempty"`
	To         *MMDBChemicalGraph.ResidueId  `xml:"to,omitempty" json:"to,omitempty"`
}
type MoleculePntrs struct {
	NumberOfPtrs int64                          `xml:"number-of-ptrs" json:"number_of_ptrs"`
	MoleculeIds  []MMDBChemicalGraph.MoleculeId `xml:"molecule-ids,omitempty" json:"molecule_ids,omitempty"`
}
type RegionPntrs struct {
	ModelId *MMDBCommon.ModelId `xml:"model-id,omitempty" json:"model_id,omitempty"`
	Region  struct {
		Site     []RegionCoordinates `xml:"site,omitempty" json:"site,omitempty"`
		Boundary []RegionBoundary    `xml:"boundary,omitempty" json:"boundary,omitempty"`
	} `xml:"region" json:"region"` //Region,ChoiceOption
}
type RegionCoordinates struct {
	ModelCoordSetId   *MMDBCommon.ModelCoordinateSetId `xml:"model-coord-set-id,omitempty" json:"model_coord_set_id,omitempty"`
	NumberOfCoords    int64                            `xml:"number-of-coords,omitempty" json:"number_of_coords,omitempty" asn1:"optional"`
	CoordinateIndices []int64                          `xml:"coordinate-indices,omitempty" json:"coordinate_indices,omitempty" asn1:"optional"`
}
type RegionBoundary struct {
	Sphere   *Sphere   `xml:"sphere,omitempty" json:"sphere,omitempty"`
	Cone     *Cone     `xml:"cone,omitempty" json:"cone,omitempty"`
	Cylinder *Cylinder `xml:"cylinder,omitempty" json:"cylinder,omitempty"`
	Brick    *Brick    `xml:"brick,omitempty" json:"brick,omitempty"`
}

//RegionBoundary,ChoiceOption
type ChemGraphAlignment struct {
	Dimension   int64                   `xml:"dimension" json:"dimension" asn1:"default:2"`
	BiostrucIds []MMDBCommon.BiostrucId `xml:"biostruc-ids,omitempty" json:"biostruc_ids,omitempty"`
	Alignment   []ChemGraphPntrs        `xml:"alignment,omitempty" json:"alignment,omitempty"`
	Domain      []ChemGraphPntrs        `xml:"domain,omitempty" json:"domain,omitempty" asn1:"optional"`
	Transform   []Transform             `xml:"transform,omitempty" json:"transform,omitempty" asn1:"optional"`
	Aligndata   []AlignStats            `xml:"aligndata,omitempty" json:"aligndata,omitempty" asn1:"optional"`
}
type ChemGraphInteraction struct {
	Type              int                    `xml:"type,omitempty" json:"type,omitempty" asn1:"optional"` //Type,IntegerEnum:protein-protein(1),protein-dna(2),protein-rna(3),protein-chemical(4),dna-dna(5),dna-rna(6),dna-chemical(7),rna-rna(8),rna-chemical(9),other(255)
	DistanceThreshold *RealValue             `xml:"distance-threshold,omitempty" json:"distance_threshold,omitempty" asn1:"optional"`
	Interactors       []BiostrucMoleculePntr `xml:"interactors,omitempty" json:"interactors,omitempty"`
	ResidueContacts   []ChemGraphPntrs       `xml:"residue-contacts,omitempty" json:"residue_contacts,omitempty" asn1:"optional"`
	AtomContacts      []ChemGraphPntrs       `xml:"atom-contacts,omitempty" json:"atom_contacts,omitempty" asn1:"optional"`
	AtomDistance      []RealValue            `xml:"atom-distance,omitempty" json:"atom_distance,omitempty" asn1:"optional"`
}
type AlignStats struct {
	Descr       string `xml:"descr,omitempty" json:"descr,omitempty" asn1:"optional"`
	ScaleFactor int64  `xml:"scale-factor,omitempty" json:"scale_factor,omitempty" asn1:"optional"`
	VastScore   int64  `xml:"vast-score,omitempty" json:"vast_score,omitempty" asn1:"optional"`
	VastMlogp   int64  `xml:"vast-mlogp,omitempty" json:"vast_mlogp,omitempty" asn1:"optional"`
	AlignRes    int64  `xml:"align-res,omitempty" json:"align_res,omitempty" asn1:"optional"`
	Rmsd        int64  `xml:"rmsd,omitempty" json:"rmsd,omitempty" asn1:"optional"`
	BlastScore  int64  `xml:"blast-score,omitempty" json:"blast_score,omitempty" asn1:"optional"`
	BlastMlogp  int64  `xml:"blast-mlogp,omitempty" json:"blast_mlogp,omitempty" asn1:"optional"`
	OtherScore  int64  `xml:"other-score,omitempty" json:"other_score,omitempty" asn1:"optional"`
}
type RegionSimilarity struct {
	Dimension   int64                   `xml:"dimension" json:"dimension" asn1:"default:2"`
	BiostrucIds []MMDBCommon.BiostrucId `xml:"biostruc-ids,omitempty" json:"biostruc_ids,omitempty"`
	Similarity  []RegionPntrs           `xml:"similarity,omitempty" json:"similarity,omitempty"`
	Transform   []Transform             `xml:"transform,omitempty" json:"transform,omitempty"`
}
type Sphere struct {
	Center *ModelSpacePoint `xml:"center,omitempty" json:"center,omitempty"`
	Radius *RealValue       `xml:"radius,omitempty" json:"radius,omitempty"`
}
type Cone struct {
	AxisTop      *ModelSpacePoint `xml:"axis-top,omitempty" json:"axis_top,omitempty"`
	AxisBottom   *ModelSpacePoint `xml:"axis-bottom,omitempty" json:"axis_bottom,omitempty"`
	RadiusBottom *RealValue       `xml:"radius-bottom,omitempty" json:"radius_bottom,omitempty"`
}
type Cylinder struct {
	AxisTop    *ModelSpacePoint `xml:"axis-top,omitempty" json:"axis_top,omitempty"`
	AxisBottom *ModelSpacePoint `xml:"axis-bottom,omitempty" json:"axis_bottom,omitempty"`
	Radius     *RealValue       `xml:"radius,omitempty" json:"radius,omitempty"`
}
type Brick struct {
	Corner000 *ModelSpacePoint `xml:"corner-000,omitempty" json:"corner_000,omitempty"`
	Corner001 *ModelSpacePoint `xml:"corner-001,omitempty" json:"corner_001,omitempty"`
	Corner010 *ModelSpacePoint `xml:"corner-010,omitempty" json:"corner_010,omitempty"`
	Corner011 *ModelSpacePoint `xml:"corner-011,omitempty" json:"corner_011,omitempty"`
	Corner100 *ModelSpacePoint `xml:"corner-100,omitempty" json:"corner_100,omitempty"`
	Corner101 *ModelSpacePoint `xml:"corner-101,omitempty" json:"corner_101,omitempty"`
	Corner110 *ModelSpacePoint `xml:"corner-110,omitempty" json:"corner_110,omitempty"`
	Corner111 *ModelSpacePoint `xml:"corner-111,omitempty" json:"corner_111,omitempty"`
}
type ModelSpacePoint struct {
	ScaleFactor int64 `xml:"scale-factor" json:"scale_factor"`
	X           int64 `xml:"x" json:"x"`
	Y           int64 `xml:"y" json:"y"`
	Z           int64 `xml:"z" json:"z"`
}
type RealValue struct {
	ScaleFactor        int64 `xml:"scale-factor" json:"scale_factor"`
	ScaledIntegerValue int64 `xml:"scaled-integer-value" json:"scaled_integer_value"`
}
type Transform struct {
	Id    int64  `xml:"id" json:"id"`
	Moves []Move `xml:"moves,omitempty" json:"moves,omitempty"`
}
type Move struct {
	Rotate    *RotMatrix   `xml:"rotate,omitempty" json:"rotate,omitempty"`
	Translate *TransMatrix `xml:"translate,omitempty" json:"translate,omitempty"`
}

//Move,ChoiceOption
type RotMatrix struct {
	ScaleFactor int64 `xml:"scale-factor" json:"scale_factor"`
	Rot11       int64 `xml:"rot-11" json:"rot_11"`
	Rot12       int64 `xml:"rot-12" json:"rot_12"`
	Rot13       int64 `xml:"rot-13" json:"rot_13"`
	Rot21       int64 `xml:"rot-21" json:"rot_21"`
	Rot22       int64 `xml:"rot-22" json:"rot_22"`
	Rot23       int64 `xml:"rot-23" json:"rot_23"`
	Rot31       int64 `xml:"rot-31" json:"rot_31"`
	Rot32       int64 `xml:"rot-32" json:"rot_32"`
	Rot33       int64 `xml:"rot-33" json:"rot_33"`
}
type TransMatrix struct {
	ScaleFactor int64 `xml:"scale-factor" json:"scale_factor"`
	Tran1       int64 `xml:"tran-1" json:"tran_1"`
	Tran2       int64 `xml:"tran-2" json:"tran_2"`
	Tran3       int64 `xml:"tran-3" json:"tran_3"`
}
type Camera struct {
	X         int64     `xml:"x" json:"x"`
	Y         int64     `xml:"y" json:"y"`
	Distance  int64     `xml:"distance" json:"distance"`
	Angle     int64     `xml:"angle" json:"angle"`
	Scale     int64     `xml:"scale" json:"scale"`
	Modelview *GLMatrix `xml:"modelview,omitempty" json:"modelview,omitempty"`
}
type GLMatrix struct {
	Scale int64 `xml:"scale" json:"scale"`
	M11   int64 `xml:"m11" json:"m11"`
	M12   int64 `xml:"m12" json:"m12"`
	M13   int64 `xml:"m13" json:"m13"`
	M14   int64 `xml:"m14" json:"m14"`
	M21   int64 `xml:"m21" json:"m21"`
	M22   int64 `xml:"m22" json:"m22"`
	M23   int64 `xml:"m23" json:"m23"`
	M24   int64 `xml:"m24" json:"m24"`
	M31   int64 `xml:"m31" json:"m31"`
	M32   int64 `xml:"m32" json:"m32"`
	M33   int64 `xml:"m33" json:"m33"`
	M34   int64 `xml:"m34" json:"m34"`
	M41   int64 `xml:"m41" json:"m41"`
	M42   int64 `xml:"m42" json:"m42"`
	M43   int64 `xml:"m43" json:"m43"`
	M44   int64 `xml:"m44" json:"m44"`
}
type ColorProp struct {
	R    int64  `xml:"r,omitempty" json:"r,omitempty" asn1:"optional"`
	G    int64  `xml:"g,omitempty" json:"g,omitempty" asn1:"optional"`
	B    int64  `xml:"b,omitempty" json:"b,omitempty" asn1:"optional"`
	Name string `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
}
type RenderProp int

//RenderProp,IntegerEnum:default(0),wire(1),space(2),stick(3),ballNStick(4),thickWire(5),hide(9),name(10),number(11),pdbNumber(12),objWireFrame(150),objPolygons(151),colorsetCPK(225),colorsetbyChain(226),colorsetbyTemp(227),colorsetbyRes(228),colorsetbyLen(229),colorsetbySStru(230),colorsetbyHydro(231),colorsetbyObject(246),colorsetbyDomain(247),other(255)
type BiostrucScript []BiostrucScriptStep
type BiostrucScriptStep struct {
	StepId     *StepId        `xml:"step-id,omitempty" json:"step_id,omitempty"`
	StepName   string         `xml:"step-name,omitempty" json:"step_name,omitempty" asn1:"optional"`
	FeatureDo  []OtherFeature `xml:"feature-do,omitempty" json:"feature_do,omitempty" asn1:"optional"`
	CameraMove *Transform     `xml:"camera-move,omitempty" json:"camera_move,omitempty" asn1:"optional"`
	Pause      int64          `xml:"pause" json:"pause" asn1:"default:10"`
	Waitevent  bool           `xml:"waitevent" json:"waitevent"`
	Extra      int64          `xml:"extra" json:"extra"`
	Jump       *StepId        `xml:"jump,omitempty" json:"jump,omitempty" asn1:"optional"`
}
type StepId int64
