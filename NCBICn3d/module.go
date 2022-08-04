package NCBICn3d

import (
	"ncbiasn/MMDB"
	"ncbiasn/MMDBChemicalGraph"
)

//Cn3dBackboneType,EnumList:off(1),trace(2),partial(3),complete(4)
type Cn3dBackboneType string

//Cn3dDrawingStyle,EnumList:wire(1),tubes(2),ball-and-stick(3),space-fill(4),wire-worm(5),tube-worm(6),with-arrows(7),without-arrows(8)
type Cn3dDrawingStyle string

//Cn3dColorScheme,EnumList:element(1),object(2),molecule(3),domain(4),residue(20),secondary-structure(5),user-select(6),aligned(7),identity(8),variety(9),weighted-variety(10),information-content(11),fit(12),block-fit(17),block-z-fit(18),block-row-fit(19),temperature(13),hydrophobicity(14),charge(15),rainbow(16)
type Cn3dColorScheme string

type Cn3dColor struct {
	ScaleFactor int64 `xml:"scale-factor" json:"scale_factor" asn1:"default:255"`
	Red         int64 `xml:"red" json:"red"`
	Green       int64 `xml:"green" json:"green"`
	Blue        int64 `xml:"blue" json:"blue"`
	Alpha       int64 `xml:"alpha" json:"alpha" asn1:"default:255"`
}
type Cn3dBackboneStyle struct {
	Type        *Cn3dBackboneType `xml:"type,omitempty" json:"type,omitempty"`
	Style       *Cn3dDrawingStyle `xml:"style,omitempty" json:"style,omitempty"`
	ColorScheme *Cn3dColorScheme  `xml:"color-scheme,omitempty" json:"color_scheme,omitempty"`
	UserColor   *Cn3dColor        `xml:"user-color,omitempty" json:"user_color,omitempty"`
}
type Cn3dGeneralStyle struct {
	IsOn        bool              `xml:"is-on" json:"is_on"`
	Style       *Cn3dDrawingStyle `xml:"style,omitempty" json:"style,omitempty"`
	ColorScheme *Cn3dColorScheme  `xml:"color-scheme,omitempty" json:"color_scheme,omitempty"`
	UserColor   *Cn3dColor        `xml:"user-color,omitempty" json:"user_color,omitempty"`
}
type Cn3dBackboneLabelStyle struct {
	Spacing int64  `xml:"spacing" json:"spacing"`
	Type    string `xml:"type" json:"type"`     //Type,EnumList:one-letter(1),three-letter(2)
	Number  string `xml:"number" json:"number"` //Number,EnumList:none(0),sequential(1),pdb(2)
	Termini bool   `xml:"termini" json:"termini"`
	White   bool   `xml:"white" json:"white"`
}
type Cn3dStyleSettings struct {
	Name                  string                  `xml:"name,omitempty" json:"name,omitempty" asn1:"optional"`
	ProteinBackbone       *Cn3dBackboneStyle      `xml:"protein-backbone,omitempty" json:"protein_backbone,omitempty"`
	NucleotideBackbone    *Cn3dBackboneStyle      `xml:"nucleotide-backbone,omitempty" json:"nucleotide_backbone,omitempty"`
	ProteinSidechains     *Cn3dGeneralStyle       `xml:"protein-sidechains,omitempty" json:"protein_sidechains,omitempty"`
	NucleotideSidechains  *Cn3dGeneralStyle       `xml:"nucleotide-sidechains,omitempty" json:"nucleotide_sidechains,omitempty"`
	Heterogens            *Cn3dGeneralStyle       `xml:"heterogens,omitempty" json:"heterogens,omitempty"`
	Solvents              *Cn3dGeneralStyle       `xml:"solvents,omitempty" json:"solvents,omitempty"`
	Connections           *Cn3dGeneralStyle       `xml:"connections,omitempty" json:"connections,omitempty"`
	HelixObjects          *Cn3dGeneralStyle       `xml:"helix-objects,omitempty" json:"helix_objects,omitempty"`
	StrandObjects         *Cn3dGeneralStyle       `xml:"strand-objects,omitempty" json:"strand_objects,omitempty"`
	VirtualDisulfidesOn   bool                    `xml:"virtual-disulfides-on" json:"virtual_disulfides_on"`
	VirtualDisulfideColor *Cn3dColor              `xml:"virtual-disulfide-color,omitempty" json:"virtual_disulfide_color,omitempty"`
	HydrogensOn           bool                    `xml:"hydrogens-on" json:"hydrogens_on"`
	BackgroundColor       *Cn3dColor              `xml:"background-color,omitempty" json:"background_color,omitempty"`
	ScaleFactor           int64                   `xml:"scale-factor" json:"scale_factor"`
	SpaceFillProportion   int64                   `xml:"space-fill-proportion" json:"space_fill_proportion"`
	BallRadius            int64                   `xml:"ball-radius" json:"ball_radius"`
	StickRadius           int64                   `xml:"stick-radius" json:"stick_radius"`
	TubeRadius            int64                   `xml:"tube-radius" json:"tube_radius"`
	TubeWormRadius        int64                   `xml:"tube-worm-radius" json:"tube_worm_radius"`
	HelixRadius           int64                   `xml:"helix-radius" json:"helix_radius"`
	StrandWidth           int64                   `xml:"strand-width" json:"strand_width"`
	StrandThickness       int64                   `xml:"strand-thickness" json:"strand_thickness"`
	ProteinLabels         *Cn3dBackboneLabelStyle `xml:"protein-labels,omitempty" json:"protein_labels,omitempty" asn1:"optional"`
	NucleotideLabels      *Cn3dBackboneLabelStyle `xml:"nucleotide-labels,omitempty" json:"nucleotide_labels,omitempty" asn1:"optional"`
	IonLabels             bool                    `xml:"ion-labels,omitempty" json:"ion_labels,omitempty" asn1:"optional"`
}
type Cn3dStyleSettingsSet []Cn3dStyleSettings
type Cn3dStyleTableId int64
type Cn3dStyleTableItem struct {
	Id    *Cn3dStyleTableId  `xml:"id,omitempty" json:"id,omitempty"`
	Style *Cn3dStyleSettings `xml:"style,omitempty" json:"style,omitempty"`
}
type Cn3dStyleDictionary struct {
	GlobalStyle *Cn3dStyleSettings   `xml:"global-style,omitempty" json:"global_style,omitempty"`
	StyleTable  []Cn3dStyleTableItem `xml:"style-table,omitempty" json:"style_table,omitempty" asn1:"optional"`
}
type Cn3dResidueRange struct {
	From *MMDBChemicalGraph.ResidueId `xml:"from,omitempty" json:"from,omitempty"`
	To   *MMDBChemicalGraph.ResidueId `xml:"to,omitempty" json:"to,omitempty"`
}
type Cn3dMoleculeLocation struct {
	MoleculeId *MMDBChemicalGraph.MoleculeId `xml:"molecule-id,omitempty" json:"molecule_id,omitempty"`
	Residues   []Cn3dResidueRange            `xml:"residues,omitempty" json:"residues,omitempty" asn1:"optional"`
}
type Cn3dObjectLocation struct {
	StructureId *MMDB.BiostrucId       `xml:"structure-id,omitempty" json:"structure_id,omitempty"`
	Residues    []Cn3dMoleculeLocation `xml:"residues,omitempty" json:"residues,omitempty"`
}
type Cn3dUserAnnotation struct {
	Name        string               `xml:"name" json:"name"`
	Description string               `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	StyleId     *Cn3dStyleTableId    `xml:"style-id,omitempty" json:"style_id,omitempty"`
	Residues    []Cn3dObjectLocation `xml:"residues,omitempty" json:"residues,omitempty"`
	IsOn        bool                 `xml:"is-on" json:"is_on"`
}
type Cn3dGLMatrix struct {
	M0  float64 `xml:"m0" json:"m0"`
	M1  float64 `xml:"m1" json:"m1"`
	M2  float64 `xml:"m2" json:"m2"`
	M3  float64 `xml:"m3" json:"m3"`
	M4  float64 `xml:"m4" json:"m4"`
	M5  float64 `xml:"m5" json:"m5"`
	M6  float64 `xml:"m6" json:"m6"`
	M7  float64 `xml:"m7" json:"m7"`
	M8  float64 `xml:"m8" json:"m8"`
	M9  float64 `xml:"m9" json:"m9"`
	M10 float64 `xml:"m10" json:"m10"`
	M11 float64 `xml:"m11" json:"m11"`
	M12 float64 `xml:"m12" json:"m12"`
	M13 float64 `xml:"m13" json:"m13"`
	M14 float64 `xml:"m14" json:"m14"`
	M15 float64 `xml:"m15" json:"m15"`
}
type Cn3dVector struct {
	X float64 `xml:"x" json:"x"`
	Y float64 `xml:"y" json:"y"`
	Z float64 `xml:"z" json:"z"`
}
type Cn3dViewSettings struct {
	CameraDistance float64       `xml:"camera-distance" json:"camera_distance"`
	CameraAngleRad float64       `xml:"camera-angle-rad" json:"camera_angle_rad"`
	CameraLookAtX  float64       `xml:"camera-look-at-X" json:"camera_look_at_X"`
	CameraLookAtY  float64       `xml:"camera-look-at-Y" json:"camera_look_at_Y"`
	CameraClipNear float64       `xml:"camera-clip-near" json:"camera_clip_near"`
	CameraClipFar  float64       `xml:"camera-clip-far" json:"camera_clip_far"`
	Matrix         *Cn3dGLMatrix `xml:"matrix,omitempty" json:"matrix,omitempty"`
	RotationCenter *Cn3dVector   `xml:"rotation-center,omitempty" json:"rotation_center,omitempty"`
}
type Cn3dUserAnnotations struct {
	Annotations []Cn3dUserAnnotation `xml:"annotations,omitempty" json:"annotations,omitempty" asn1:"optional"`
	View        *Cn3dViewSettings    `xml:"view,omitempty" json:"view,omitempty" asn1:"optional"`
}
