package NCBIObjPrt

type PrintTemplate struct {
	Name      TemplateName `xml:"name,omitempty" json:"name,omitempty"`
	Labelfrom string       `xml:"labelfrom,omitempty" json:"labelfrom,omitempty" asn1:"optional"`
	Format    *PrintFormat `xml:"format,omitempty" json:"format,omitempty"`
}
type TemplateName string
type PrintTemplateSet []PrintTemplate
type PrintFormat struct {
	Asn1   string     `xml:"asn1" json:"asn1"`
	Label  string     `xml:"label,omitempty" json:"label,omitempty" asn1:"optional"`
	Prefix string     `xml:"prefix,omitempty" json:"prefix,omitempty" asn1:"optional"`
	Suffix string     `xml:"suffix,omitempty" json:"suffix,omitempty" asn1:"optional"`
	Form   *PrintForm `xml:"form,omitempty" json:"form,omitempty"`
}

//PrintForm,ChoiceOption
type PrintForm struct {
	Block       *PrintFormBlock   `xml:"block,omitempty" json:"block,omitempty"`
	Boolean     *PrintFormBoolean `xml:"boolean,omitempty" json:"boolean,omitempty"`
	Enum        *PrintFormEnum    `xml:"enum,omitempty" json:"enum,omitempty"`
	Text        *PrintFormText    `xml:"text,omitempty" json:"text,omitempty"`
	UseTemplate TemplateName      `xml:"use-template,omitempty" json:"use_template,omitempty"`
	User        *UserFormat       `xml:"user,omitempty" json:"user,omitempty"`
	Null        interface{}       `xml:"-" json:"-"` //Null,NullType
}

type UserFormat struct {
	Printfunc   string `xml:"printfunc" json:"printfunc"`
	Defaultfunc string `xml:"defaultfunc,omitempty" json:"defaultfunc,omitempty" asn1:"optional"`
}
type PrintFormBlock struct {
	Separator  string        `xml:"separator,omitempty" json:"separator,omitempty" asn1:"optional"`
	Components []PrintFormat `xml:"components,omitempty" json:"components,omitempty"`
}
type PrintFormBoolean struct {
	True  string `xml:"true,omitempty" json:"true,omitempty" asn1:"optional"`
	False string `xml:"false,omitempty" json:"false,omitempty" asn1:"optional"`
}
type PrintFormEnum struct {
	Values []string `xml:"values,omitempty" json:"values,omitempty" asn1:"optional"`
}
type PrintFormText struct {
	Textfunc string `xml:"textfunc,omitempty" json:"textfunc,omitempty" asn1:"optional"`
}
