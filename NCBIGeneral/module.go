package NCBIGeneral

//Date,ChoiceOption
type Date struct {
	Str string   `xml:"str,omitempty" json:"str,omitempty"`
	Std *DateStd `xml:"std,omitempty" json:"std,omitempty"`
}
type DateStd struct {
	Year   int64  `xml:"year" json:"year"`
	Month  int64  `xml:"month,omitempty" json:"month,omitempty" asn1:"optional"`
	Day    int64  `xml:"day,omitempty" json:"day,omitempty" asn1:"optional"`
	Season string `xml:"season,omitempty" json:"season,omitempty" asn1:"optional"`
	Hour   int64  `xml:"hour,omitempty" json:"hour,omitempty" asn1:"optional"`
	Minute int64  `xml:"minute,omitempty" json:"minute,omitempty" asn1:"optional"`
	Second int64  `xml:"second,omitempty" json:"second,omitempty" asn1:"optional"`
}
type Dbtag struct {
	Db  string    `xml:"db" json:"db"`
	Tag *ObjectId `xml:"tag,omitempty" json:"tag,omitempty"`
}

//ObjectId,ChoiceOption
type ObjectId struct {
	Id  int64  `xml:"id,omitempty" json:"id,omitempty"`
	Str string `xml:"str,omitempty" json:"str,omitempty"`
}

//PersonId,ChoiceOption
type PersonId struct {
	Dbtag      *Dbtag   `xml:"dbtag,omitempty" json:"dbtag,omitempty"`
	Name       *NameStd `xml:"name,omitempty" json:"name,omitempty"`
	Ml         string   `xml:"ml,omitempty" json:"ml,omitempty"`
	Str        string   `xml:"str,omitempty" json:"str,omitempty"`
	Consortium string   `xml:"consortium,omitempty" json:"consortium,omitempty"`
}

type NameStd struct {
	Last     string `xml:"last" json:"last"`
	First    string `xml:"first,omitempty" json:"first,omitempty" asn1:"optional"`
	Middle   string `xml:"middle,omitempty" json:"middle,omitempty" asn1:"optional"`
	Full     string `xml:"full,omitempty" json:"full,omitempty" asn1:"optional"`
	Initials string `xml:"initials,omitempty" json:"initials,omitempty" asn1:"optional"`
	Suffix   string `xml:"suffix,omitempty" json:"suffix,omitempty" asn1:"optional"`
	Title    string `xml:"title,omitempty" json:"title,omitempty" asn1:"optional"`
}

//IntFuzz,ChoiceOption
type IntFuzz struct {
	PM    int64 `xml:"p-m,omitempty" json:"p_m,omitempty"`
	Range struct {
		Max int64 `xml:"max" json:"max"`
		Min int64 `xml:"min" json:"min"`
	} `xml:"range,omitempty" json:"range,omitempty"`
	Pct int64 `xml:"pct,omitempty" json:"pct,omitempty"`
	//Lim,EnumList:unk(0),gt(1),lt(2),tr(3),tl(4),circle(5),other(255)
	Lim string  `xml:"lim,omitempty" json:"lim,omitempty"`
	Alt []int64 `xml:"alt,omitempty" json:"alt,omitempty"`
}

type UserObject struct {
	Class string      `xml:"class,omitempty" json:"class,omitempty" asn1:"optional"`
	Type  *ObjectId   `xml:"type,omitempty" json:"type,omitempty"`
	Data  []UserField `xml:"data,omitempty" json:"data,omitempty"`
}
type UserField struct {
	Label *ObjectId `xml:"label,omitempty" json:"label,omitempty"`
	Num   int64     `xml:"num,omitempty" json:"num,omitempty" asn1:"optional"`
	//Data,ChoiceOption
	Data struct {
		Str     string       `xml:"str,omitempty" json:"str,omitempty"`
		Int     int64        `xml:"int,omitempty" json:"int,omitempty"`
		Real    float64      `xml:"real,omitempty" json:"real,omitempty"`
		Bool    bool         `xml:"bool,omitempty" json:"bool,omitempty"`
		Os      []byte       `xml:"os,omitempty" json:"os,omitempty"`
		Object  *UserObject  `xml:"object,omitempty" json:"object,omitempty"`
		Strs    []string     `xml:"strs,omitempty" json:"strs,omitempty"`
		Ints    []int64      `xml:"ints,omitempty" json:"ints,omitempty"`
		Reals   []float64    `xml:"reals,omitempty" json:"reals,omitempty"`
		Oss     []byte       `xml:"oss,omitempty" json:"oss,omitempty"`
		Fields  []UserField  `xml:"fields,omitempty" json:"fields,omitempty"`
		Objects []UserObject `xml:"objects,omitempty" json:"objects,omitempty"`
	} `xml:"data" json:"data"`
}
