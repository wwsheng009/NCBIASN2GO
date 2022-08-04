package NCBIBlastDL

import "ncbiasn/NCBISeqloc"

type BlastDefLineSet []BlastDefLine
type BlastDefLine struct {
	Title       string             `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
	Seqid       []NCBISeqloc.SeqId `xml:"seqid,omitempty" json:"seqid,omitempty"`
	Taxid       int64              `xml:"taxid,omitempty" json:"taxid,omitempty" asn1:"optional"`
	Memberships []int64            `xml:"memberships,omitempty" json:"memberships,omitempty" asn1:"optional"`
	Links       []int64            `xml:"links,omitempty" json:"links,omitempty" asn1:"optional"`
	OtherInfo   []int64            `xml:"other-info,omitempty" json:"other_info,omitempty" asn1:"optional"`
}
type BlastFilterProgram int

//BlastFilterProgram,IntegerEnum:not-set(0),dust(10),seg(20),windowmasker(30),repeat(40),other(100),max(255)
type BlastMaskList struct {
	Masks []NCBISeqloc.SeqLoc `xml:"masks,omitempty" json:"masks,omitempty"`
	More  bool                `xml:"more" json:"more"`
}
type BlastDbMaskInfo struct {
	AlgoId      int64               `xml:"algo-id" json:"algo_id"`
	AlgoProgram *BlastFilterProgram `xml:"algo-program,omitempty" json:"algo_program,omitempty"`
	AlgoOptions string              `xml:"algo-options" json:"algo_options"`
	Masks       *BlastMaskList      `xml:"masks,omitempty" json:"masks,omitempty"`
}
type BlastDbMetadata struct {
	Version              string   `xml:"version" json:"version"`
	Dbname               string   `xml:"dbname" json:"dbname"`
	Dbtype               string   `xml:"dbtype" json:"dbtype"`
	DbVersion            int64    `xml:"db-version" json:"db_version" asn1:"default:5"`
	Description          string   `xml:"description,omitempty" json:"description,omitempty" asn1:"optional"`
	NumberOfLetters      int64    `xml:"number-of-letters" json:"number_of_letters"`
	NumberOfSequences    int64    `xml:"number-of-sequences" json:"number_of_sequences"`
	LastUpdated          string   `xml:"last-updated" json:"last_updated"`
	NumberOfVolumes      int64    `xml:"number-of-volumes" json:"number_of_volumes"`
	NumberOfTaxids       int64    `xml:"number-of-taxids,omitempty" json:"number_of_taxids,omitempty" asn1:"optional"`
	BytesTotal           int64    `xml:"bytes-total" json:"bytes_total"`
	BytesToCache         int64    `xml:"bytes-to-cache" json:"bytes_to_cache"`
	BytesTotalCompressed int64    `xml:"bytes-total-compressed,omitempty" json:"bytes_total_compressed,omitempty" asn1:"optional"`
	Files                []string `xml:"files,omitempty" json:"files,omitempty" asn1:"optional"`
	CompressedFiles      []string `xml:"compressed-files,omitempty" json:"compressed_files,omitempty" asn1:"optional"`
}
