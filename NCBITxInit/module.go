package NCBITxInit

import (
	"ncbiasn/NCBIGene"
	"ncbiasn/NCBIOrganism"
	"ncbiasn/NCBIProtein"
)

type Txinit struct {
	Name             string                `xml:"name" json:"name"`
	Syn              []string              `xml:"syn,omitempty" json:"syn,omitempty" asn1:"optional"`
	Gene             []NCBIGene.GeneRef    `xml:"gene,omitempty" json:"gene,omitempty" asn1:"optional"`
	Protein          []NCBIProtein.ProtRef `xml:"protein,omitempty" json:"protein,omitempty" asn1:"optional"`
	Rna              []string              `xml:"rna,omitempty" json:"rna,omitempty" asn1:"optional"`
	Expression       string                `xml:"expression,omitempty" json:"expression,omitempty" asn1:"optional"`
	Txsystem         string                `xml:"txsystem" json:"txsystem"` //Txsystem,EnumList:unknown(0),pol1(1),pol2(2),pol3(3),bacterial(4),viral(5),rna(6),organelle(7),other(255)
	Txdescr          string                `xml:"txdescr,omitempty" json:"txdescr,omitempty" asn1:"optional"`
	Txorg            *NCBIOrganism.OrgRef  `xml:"txorg,omitempty" json:"txorg,omitempty" asn1:"optional"`
	MappingPrecise   bool                  `xml:"mapping-precise" json:"mapping_precise"`
	LocationAccurate bool                  `xml:"location-accurate" json:"location_accurate"`
	Inittype         string                `xml:"inittype,omitempty" json:"inittype,omitempty" asn1:"optional"` //Inittype,EnumList:unknown(0),single(1),multiple(2),region(3)
	Evidence         []TxEvidence          `xml:"evidence,omitempty" json:"evidence,omitempty" asn1:"optional"`
}
type TxEvidence struct {
	ExpCode          string `xml:"exp-code" json:"exp_code"`                   //ExpCode,EnumList:unknown(0),rna-seq(1),rna-size(2),np-map(3),np-size(4),pe-seq(5),cDNA-seq(6),pe-map(7),pe-size(8),pseudo-seq(9),rev-pe-map(10),other(255)
	ExpressionSystem string `xml:"expression-system" json:"expression_system"` //ExpressionSystem,EnumList:unknown(0),physiological(1),in-vitro(2),oocyte(3),transfection(4),transgenic(5),other(255)
	LowPrecData      bool   `xml:"low-prec-data" json:"low_prec_data"`
	FromHomolog      bool   `xml:"from-homolog" json:"from_homolog"`
}
