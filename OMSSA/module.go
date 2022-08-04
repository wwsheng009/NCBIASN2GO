package OMSSA

import "ncbiasn/NCBISequence"

type NameValue struct {
	Name  string `xml:"name" json:"name"`
	Value string `xml:"value" json:"value"`
}
type MSSpectrum struct {
	Number      int64       `xml:"number" json:"number"`
	Charge      []int64     `xml:"charge,omitempty" json:"charge,omitempty"`
	Precursormz int64       `xml:"precursormz" json:"precursormz"`
	Mz          []int64     `xml:"mz,omitempty" json:"mz,omitempty"`
	Abundance   []int64     `xml:"abundance,omitempty" json:"abundance,omitempty"`
	Iscale      float64     `xml:"iscale" json:"iscale"`
	Ids         []string    `xml:"ids,omitempty" json:"ids,omitempty" asn1:"optional"`
	Namevalue   []NameValue `xml:"namevalue,omitempty" json:"namevalue,omitempty" asn1:"optional"`
}
type MSSpectrumset []MSSpectrum

//MSEnzymes,IntegerEnum:trypsin(0),argc(1),cnbr(2),chymotrypsin(3),formicacid(4),lysc(5),lysc-p(6),pepsin-a(7),tryp-cnbr(8),tryp-chymo(9),trypsin-p(10),whole-protein(11),aspn(12),gluc(13),aspngluc(14),top-down(15),semi-tryptic(16),no-enzyme(17),chymotrypsin-p(18),aspn-de(19),gluc-de(20),lysn(21),thermolysin-p(22),semi-chymotrypsin(23),semi-gluc(24),max(25),none(255)
type MSEnzymes int

//MSMod,IntegerEnum:methylk(0),oxym(1),carboxymethylc(2),carbamidomethylc(3),deamidationkq(4),propionamidec(5),phosphorylations(6),phosphorylationt(7),phosphorylationy(8),ntermmcleave(9),ntermacetyl(10),ntermmethyl(11),ntermtrimethyl(12),methythiold(13),methylq(14),trimethylk(15),methyld(16),methyle(17),ctermpepmethyl(18),trideuteromethyld(19),trideuteromethyle(20),ctermpeptrideuteromethyl(21),nformylmet(22),twoamino3oxobutanoicacid(23),acetylk(24),ctermamide(25),bmethylthiold(26),carbamidomethylk(27),carbamidometylh(28),carbamidomethyld(29),carbamidomethyle(30),carbamylk(31),ntermcarbamyl(32),citrullinationr(33),cysteicacidc(34),diiodinationy(35),dimethylk(36),dimethylr(37),ntermpepdimethyl(38),dihydroxyf(39),thioacetylk(40),ntermpeptioacetyl(41),farnesylationc(42),formylk(43),ntermpepformyl(44),formylkynureninw(45),phef(46),gammacarboxyld(47),gammacarboxyle(48),geranylgeranylc(49),ntermpepglucuronylg(50),glutathionec(51),glyglyk(52),guanidinationk(53),his2asnh(54),his2asph(55),ctermpephsem(56),ctermpephselactm(57),hydroxykynureninw(58),hydroxylationd(59),hydroxylationk(60),hydroxylationn(61),hydroxylationp(62),hydroxylationf(63),hydroxylationy(64),iodinationy(65),kynureninw(66),lipoylk(67),ctermpepmeester(68),meesterd(69),meestere(70),meesters(71),meestery(72),methylc(73),methylh(74),methyln(75),ntermpepmethyl(76),methylr(77),ntermpepmyristoyeylationg(78),ntermpepmyristoyl4hg(79),ntermpepmyristoylationg(80),myristoylationk(81),ntermformyl(82),nemc(83),nipcam(84),nitrow(85),nitroy(86),ctermpepo18(87),ctermpepdio18(88),oxyh(89),oxyw(90),ppantetheines(91),palmitoylationc(92),palmitoylationk(93),palmitoylations(94),palmitoylationt(95),phospholosss(96),phospholosst(97),phospholossy(98),phosphoneutrallossc(99),phosphoneutrallossd(100),phosphoneutrallossh(101),propionylk(102),ntermpeppropionyl(103),propionylheavyk(104),ntermpeppropionylheavy(105),pyridylk(106),ntermpeppyridyl(107),ntermpeppyrocmc(108),ntermpeppyroe(109),ntermpeppyroq(110),pyroglutamicp(111),spyridylethylc(112),semetm(113),sulfationy(114),suphonem(115),triiodinationy(116),trimethylationr(117),ntermpeptripalmitatec(118),usermod1(119),usermod2(120),usermod3(121),usermod4(122),usermod5(123),usermod6(124),usermod7(125),usermod8(126),usermod9(127),usermod10(128),icatlight(129),icatheavy(130),camthiopropanoylk(131),phosphoneutrallosss(132),phosphoneutrallosst(133),phosphoetdlosss(134),phosphoetdlosst(135),arg-13c6(136),arg-13c6-15n4(137),lys-13c6(138),oxy18(139),beta-elim-s(140),beta-elim-t(141),usermod11(142),usermod12(143),usermod13(144),usermod14(145),usermod15(146),usermod16(147),usermod17(148),usermod18(149),usermod19(150),usermod20(151),usermod21(152),usermod22(153),usermod23(154),usermod24(155),usermod25(156),usermod26(157),usermod27(158),usermod28(159),usermod29(160),usermod30(161),sulfinicacid(162),arg2orn(163),dehydro(164),carboxykynurenin(165),sumoylation(166),iTRAQ114nterm(167),iTRAQ114K(168),iTRAQ114Y(169),iTRAQ115nterm(170),iTRAQ115K(171),iTRAQ115Y(172),iTRAQ116nterm(173),iTRAQ116K(174),iTRAQ116Y(175),iTRAQ117nterm(176),iTRAQ117K(177),iTRAQ117Y(178),mmts(179),lys-2H4(180),lys-13C615N2(181),hexNAcN(182),dHexHexNAcN(183),hexNAcS(184),hexNAcT(185),mod186(186),mod187(187),mod188(188),mod189(189),mod190(190),mod191(191),mod192(192),mod193(193),mod194(194),mod195(195),mod196(196),mod197(197),mod198(198),mod199(199),mod200(200),mod201(201),mod202(202),mod203(203),mod204(204),mod205(205),mod206(206),mod207(207),mod208(208),mod209(209),mod210(210),mod211(211),mod212(212),mod213(213),mod214(214),mod215(215),mod216(216),mod217(217),mod218(218),mod219(219),mod220(220),mod221(221),mod222(222),mod223(223),mod224(224),mod225(225),mod226(226),mod227(227),mod228(228),mod229(229),mod230(230),max(231),unknown(9999),none(10000)
type MSMod int

type MSModType int

//MSModType,IntegerEnum:modaa(0),modn(1),modnaa(2),modc(3),modcaa(4),modnp(5),modnpaa(6),modcp(7),modcpaa(8),modmax(9)
type MSMassSet struct {
	Monomass    float64 `xml:"monomass" json:"monomass"`
	Averagemass float64 `xml:"averagemass" json:"averagemass"`
	N15mass     float64 `xml:"n15mass" json:"n15mass"`
}
type MSModSpec struct {
	Mod         *MSMod     `xml:"mod,omitempty" json:"mod,omitempty"`
	Type        *MSModType `xml:"type,omitempty" json:"type,omitempty"`
	Name        string     `xml:"name" json:"name"`
	Monomass    float64    `xml:"monomass" json:"monomass"`
	Averagemass float64    `xml:"averagemass" json:"averagemass"`
	N15mass     float64    `xml:"n15mass" json:"n15mass"`
	Residues    []string   `xml:"residues,omitempty" json:"residues,omitempty" asn1:"optional"`
	Neutralloss *MSMassSet `xml:"neutralloss,omitempty" json:"neutralloss,omitempty" asn1:"optional"`
	Unimod      int64      `xml:"unimod,omitempty" json:"unimod,omitempty" asn1:"optional"`
	PsiMs       string     `xml:"psi-ms,omitempty" json:"psi_ms,omitempty" asn1:"optional"`
}
type MSModSpecSet []MSModSpec
type MSCalcPlusOne int

//MSCalcPlusOne,IntegerEnum:dontcalc(0),calc(1)
type MSCalcCharge int

//MSCalcCharge,IntegerEnum:calculate(0),usefile(1),userange(2)
type MSChargeHandle struct {
	Calcplusone      *MSCalcPlusOne `xml:"calcplusone,omitempty" json:"calcplusone,omitempty" asn1:"default:1"`
	Calccharge       *MSCalcCharge  `xml:"calccharge,omitempty" json:"calccharge,omitempty" asn1:"default:2"`
	Mincharge        int64          `xml:"mincharge" json:"mincharge" asn1:"default:2"`
	Maxcharge        int64          `xml:"maxcharge" json:"maxcharge" asn1:"default:3"`
	Considermult     int64          `xml:"considermult" json:"considermult" asn1:"default:3"`
	Plusone          float64        `xml:"plusone" json:"plusone"`
	Maxproductcharge int64          `xml:"maxproductcharge,omitempty" json:"maxproductcharge,omitempty" asn1:"optional"`
	Prodlesspre      bool           `xml:"prodlesspre,omitempty" json:"prodlesspre,omitempty" asn1:"optional"`
	Negative         int64          `xml:"negative" json:"negative" asn1:"default:1"`
}
type MSSearchType int

//MSSearchType,IntegerEnum:monoisotopic(0),average(1),monon15(2),exact(3),multiisotope(4),max(5)
type MSZdependence int

//MSZdependence,IntegerEnum:independent(0),linearwithz(1),max(2)
type MSIterativeSettings struct {
	Researchthresh float64 `xml:"researchthresh" json:"researchthresh"`
	Subsetthresh   float64 `xml:"subsetthresh" json:"subsetthresh"`
	Replacethresh  float64 `xml:"replacethresh" json:"replacethresh"`
}
type MSLibrarySettings struct {
	Libnames          []string `xml:"libnames,omitempty" json:"libnames,omitempty"`
	Presearch         bool     `xml:"presearch" json:"presearch"`
	Useomssascore     bool     `xml:"useomssascore" json:"useomssascore"`
	Usereplicatescore bool     `xml:"usereplicatescore" json:"usereplicatescore"`
	Qtofscore         bool     `xml:"qtofscore" json:"qtofscore"`
}
type MSSearchSettings struct {
	Precursorsearchtype *MSSearchType        `xml:"precursorsearchtype,omitempty" json:"precursorsearchtype,omitempty"`
	Productsearchtype   *MSSearchType        `xml:"productsearchtype,omitempty" json:"productsearchtype,omitempty"`
	Ionstosearch        []MSIonType          `xml:"ionstosearch,omitempty" json:"ionstosearch,omitempty"`
	Peptol              float64              `xml:"peptol" json:"peptol"`
	Msmstol             float64              `xml:"msmstol" json:"msmstol"`
	Zdep                *MSZdependence       `xml:"zdep,omitempty" json:"zdep,omitempty"`
	Cutoff              float64              `xml:"cutoff" json:"cutoff"`
	Cutlo               float64              `xml:"cutlo" json:"cutlo"`
	Cuthi               float64              `xml:"cuthi" json:"cuthi"`
	Cutinc              float64              `xml:"cutinc" json:"cutinc"`
	Singlewin           int64                `xml:"singlewin" json:"singlewin"`
	Doublewin           int64                `xml:"doublewin" json:"doublewin"`
	Singlenum           int64                `xml:"singlenum" json:"singlenum"`
	Doublenum           int64                `xml:"doublenum" json:"doublenum"`
	Fixed               []MSMod              `xml:"fixed,omitempty" json:"fixed,omitempty"`
	Variable            []MSMod              `xml:"variable,omitempty" json:"variable,omitempty"`
	Enzyme              *MSEnzymes           `xml:"enzyme,omitempty" json:"enzyme,omitempty"`
	Missedcleave        int64                `xml:"missedcleave" json:"missedcleave"`
	Hitlistlen          int64                `xml:"hitlistlen" json:"hitlistlen" asn1:"default:25"`
	Db                  string               `xml:"db" json:"db"`
	Tophitnum           int64                `xml:"tophitnum" json:"tophitnum"`
	Minhit              int64                `xml:"minhit" json:"minhit" asn1:"default:2"`
	Minspectra          int64                `xml:"minspectra" json:"minspectra" asn1:"default:4"`
	Scale               int64                `xml:"scale" json:"scale" asn1:"default:100"`
	Maxmods             int64                `xml:"maxmods" json:"maxmods" asn1:"default:64"`
	Taxids              []int64              `xml:"taxids,omitempty" json:"taxids,omitempty" asn1:"optional"`
	Chargehandling      *MSChargeHandle      `xml:"chargehandling,omitempty" json:"chargehandling,omitempty" asn1:"optional"`
	Usermods            MSModSpecSet         `xml:"usermods,omitempty" json:"usermods,omitempty" asn1:"optional"`
	Pseudocount         int64                `xml:"pseudocount" json:"pseudocount" asn1:"default:1"`
	Searchb1            int64                `xml:"searchb1" json:"searchb1" asn1:"default:0"`
	Searchctermproduct  int64                `xml:"searchctermproduct" json:"searchctermproduct" asn1:"default:0"`
	Maxproductions      int64                `xml:"maxproductions" json:"maxproductions" asn1:"default:0"`
	Minnoenzyme         int64                `xml:"minnoenzyme" json:"minnoenzyme" asn1:"default:4"`
	Maxnoenzyme         int64                `xml:"maxnoenzyme" json:"maxnoenzyme" asn1:"default:0"`
	Exactmass           float64              `xml:"exactmass,omitempty" json:"exactmass,omitempty" asn1:"optional"`
	Settingid           int64                `xml:"settingid,omitempty" json:"settingid,omitempty" asn1:"optional"`
	Iterativesettings   *MSIterativeSettings `xml:"iterativesettings,omitempty" json:"iterativesettings,omitempty" asn1:"optional"`
	Precursorcull       int64                `xml:"precursorcull,omitempty" json:"precursorcull,omitempty" asn1:"optional"`
	Infiles             []MSInFile           `xml:"infiles,omitempty" json:"infiles,omitempty" asn1:"optional"`
	Outfiles            []MSOutFile          `xml:"outfiles,omitempty" json:"outfiles,omitempty" asn1:"optional"`
	Nocorrelationscore  int64                `xml:"nocorrelationscore,omitempty" json:"nocorrelationscore,omitempty" asn1:"optional"`
	Probfollowingion    float64              `xml:"probfollowingion,omitempty" json:"probfollowingion,omitempty" asn1:"optional"`
	Nmethionine         bool                 `xml:"nmethionine,omitempty" json:"nmethionine,omitempty" asn1:"optional"`
	Automassadjust      float64              `xml:"automassadjust,omitempty" json:"automassadjust,omitempty" asn1:"optional"`
	Lomasscutoff        float64              `xml:"lomasscutoff,omitempty" json:"lomasscutoff,omitempty" asn1:"optional"`
	Libsearchsettings   *MSLibrarySettings   `xml:"libsearchsettings,omitempty" json:"libsearchsettings,omitempty" asn1:"optional"`
	Noprolineions       []MSIonType          `xml:"noprolineions,omitempty" json:"noprolineions,omitempty" asn1:"optional"`
	Reversesearch       bool                 `xml:"reversesearch,omitempty" json:"reversesearch,omitempty" asn1:"optional"`
	Othersettings       []NameValue          `xml:"othersettings,omitempty" json:"othersettings,omitempty" asn1:"optional"`
	Numisotopes         int64                `xml:"numisotopes,omitempty" json:"numisotopes,omitempty" asn1:"optional"`
	Pepppm              bool                 `xml:"pepppm,omitempty" json:"pepppm,omitempty" asn1:"optional"`
	Msmsppm             bool                 `xml:"msmsppm,omitempty" json:"msmsppm,omitempty" asn1:"optional"`
	Reportedhitcount    int64                `xml:"reportedhitcount,omitempty" json:"reportedhitcount,omitempty" asn1:"optional"`
}
type MSSerialDataFormat int

//MSSerialDataFormat,IntegerEnum:none(0),asntext(1),asnbinary(2),xml(3),csv(4),pepxml(5),xmlbz2(6)
type MSOutFile struct {
	Outfile        string              `xml:"outfile" json:"outfile"`
	Outfiletype    *MSSerialDataFormat `xml:"outfiletype,omitempty" json:"outfiletype,omitempty"`
	Includerequest bool                `xml:"includerequest" json:"includerequest"`
}
type MSSpectrumFileType int

//MSSpectrumFileType,IntegerEnum:dta(0),dtablank(1),dtaxml(2),asc(3),pkl(4),pks(5),sciex(6),mgf(7),unknown(8),oms(9),omx(10),xml(11),omxbz2(12)
type MSInFile struct {
	Infile     string              `xml:"infile" json:"infile"`
	Infiletype *MSSpectrumFileType `xml:"infiletype,omitempty" json:"infiletype,omitempty"`
}
type MSSearchSettingsSet []MSSearchSettings
type MSRequest struct {
	Spectra      MSSpectrumset       `xml:"spectra,omitempty" json:"spectra,omitempty"`
	Settings     *MSSearchSettings   `xml:"settings,omitempty" json:"settings,omitempty"`
	Rid          string              `xml:"rid,omitempty" json:"rid,omitempty" asn1:"optional"`
	Moresettings MSSearchSettingsSet `xml:"moresettings,omitempty" json:"moresettings,omitempty" asn1:"optional"`
	Modset       MSModSpecSet        `xml:"modset,omitempty" json:"modset,omitempty" asn1:"optional"`
}
type MSIonType int

//MSIonType,IntegerEnum:a(0),b(1),c(2),x(3),y(4),z(5),parent(6),internal(7),immonium(8),unknown(9),adot(10),x-CO2(11),adot-CO2(12),max(13)
type MSIonNeutralLoss int

//MSIonNeutralLoss,IntegerEnum:water(0),ammonia(1)
type MSIonIsotopicType int

//MSIonIsotopicType,IntegerEnum:monoisotopic(0),c13(1),c13two(2),c13three(3),c13four(4)
type MSImmonium struct {
	Parent  string `xml:"parent" json:"parent"`
	Product string `xml:"product,omitempty" json:"product,omitempty" asn1:"optional"`
}
type MSIon struct {
	Neutralloss *MSIonNeutralLoss  `xml:"neutralloss,omitempty" json:"neutralloss,omitempty" asn1:"optional"`
	Isotope     *MSIonIsotopicType `xml:"isotope,omitempty" json:"isotope,omitempty" asn1:"optional"`
	Internal    string             `xml:"internal,omitempty" json:"internal,omitempty" asn1:"optional"`
	Immonium    *MSImmonium        `xml:"immonium,omitempty" json:"immonium,omitempty" asn1:"optional"`
}
type MSIonAnnot struct {
	Suspect        bool    `xml:"suspect,omitempty" json:"suspect,omitempty" asn1:"optional"`
	Massdiff       float64 `xml:"massdiff,omitempty" json:"massdiff,omitempty" asn1:"optional"`
	Missingisotope bool    `xml:"missingisotope,omitempty" json:"missingisotope,omitempty" asn1:"optional"`
}
type MSMZHit struct {
	Ion        *MSIonType  `xml:"ion,omitempty" json:"ion,omitempty"`
	Charge     int64       `xml:"charge" json:"charge"`
	Number     int64       `xml:"number" json:"number"`
	Mz         int64       `xml:"mz" json:"mz"`
	Index      int64       `xml:"index,omitempty" json:"index,omitempty" asn1:"optional"`
	Moreion    *MSIon      `xml:"moreion,omitempty" json:"moreion,omitempty" asn1:"optional"`
	Annotation *MSIonAnnot `xml:"annotation,omitempty" json:"annotation,omitempty" asn1:"optional"`
}
type MSPepHit struct {
	Start      int64  `xml:"start" json:"start"`
	Stop       int64  `xml:"stop" json:"stop"`
	Gi         int64  `xml:"gi,omitempty" json:"gi,omitempty" asn1:"optional"`
	Accession  string `xml:"accession,omitempty" json:"accession,omitempty" asn1:"optional"`
	Defline    string `xml:"defline,omitempty" json:"defline,omitempty" asn1:"optional"`
	Protlength int64  `xml:"protlength,omitempty" json:"protlength,omitempty" asn1:"optional"`
	Oid        int64  `xml:"oid,omitempty" json:"oid,omitempty" asn1:"optional"`
	Reversed   bool   `xml:"reversed,omitempty" json:"reversed,omitempty" asn1:"optional"`
	Pepstart   string `xml:"pepstart,omitempty" json:"pepstart,omitempty" asn1:"optional"`
	Pepstop    string `xml:"pepstop,omitempty" json:"pepstop,omitempty" asn1:"optional"`
}
type MSModHit struct {
	Site    int64  `xml:"site" json:"site"`
	Modtype *MSMod `xml:"modtype,omitempty" json:"modtype,omitempty"`
}
type MSScoreSet struct {
	Name  string  `xml:"name" json:"name"`
	Value float64 `xml:"value" json:"value"`
}
type MSHits struct {
	Evalue       float64      `xml:"evalue" json:"evalue"`
	Pvalue       float64      `xml:"pvalue" json:"pvalue"`
	Charge       int64        `xml:"charge" json:"charge"`
	Pephits      []MSPepHit   `xml:"pephits,omitempty" json:"pephits,omitempty"`
	Mzhits       []MSMZHit    `xml:"mzhits,omitempty" json:"mzhits,omitempty" asn1:"optional"`
	Pepstring    string       `xml:"pepstring,omitempty" json:"pepstring,omitempty" asn1:"optional"`
	Mass         int64        `xml:"mass,omitempty" json:"mass,omitempty" asn1:"optional"`
	Mods         []MSModHit   `xml:"mods,omitempty" json:"mods,omitempty" asn1:"optional"`
	Pepstart     string       `xml:"pepstart,omitempty" json:"pepstart,omitempty" asn1:"optional"`
	Pepstop      string       `xml:"pepstop,omitempty" json:"pepstop,omitempty" asn1:"optional"`
	Protlength   int64        `xml:"protlength,omitempty" json:"protlength,omitempty" asn1:"optional"`
	Theomass     int64        `xml:"theomass,omitempty" json:"theomass,omitempty" asn1:"optional"`
	Oid          int64        `xml:"oid,omitempty" json:"oid,omitempty" asn1:"optional"`
	Scores       []MSScoreSet `xml:"scores,omitempty" json:"scores,omitempty" asn1:"optional"`
	Libaccession string       `xml:"libaccession,omitempty" json:"libaccession,omitempty" asn1:"optional"`
}
type MSHitError int

//MSHitError,IntegerEnum:none(0),generalerr(1),unable2read(2),notenuffpeaks(3)
type MSUserAnnot int

//MSUserAnnot,IntegerEnum:none(0),delete(1),flag(2)
type MSHitSet struct {
	Number         int64        `xml:"number" json:"number"`
	Error          *MSHitError  `xml:"error,omitempty" json:"error,omitempty" asn1:"optional"`
	Hits           []MSHits     `xml:"hits,omitempty" json:"hits,omitempty" asn1:"optional"`
	Ids            []string     `xml:"ids,omitempty" json:"ids,omitempty" asn1:"optional"`
	Namevalue      []NameValue  `xml:"namevalue,omitempty" json:"namevalue,omitempty" asn1:"optional"`
	Settingid      int64        `xml:"settingid,omitempty" json:"settingid,omitempty" asn1:"optional"`
	Userannotation *MSUserAnnot `xml:"userannotation,omitempty" json:"userannotation,omitempty" asn1:"optional"`
}
type MSResponseError int

//MSResponseError,IntegerEnum:none(0),generalerr(1),noblastdb(2),noinput(3)
type MSBioseq struct {
	Oid int64                `xml:"oid" json:"oid"`
	Seq *NCBISequence.Bioseq `xml:"seq,omitempty" json:"seq,omitempty"`
}
type MSBioseqSet []MSBioseq
type MSResponse struct {
	Hitsets   []MSHitSet       `xml:"hitsets,omitempty" json:"hitsets,omitempty"`
	Scale     int64            `xml:"scale" json:"scale" asn1:"default:100"`
	Rid       string           `xml:"rid,omitempty" json:"rid,omitempty" asn1:"optional"`
	Error     *MSResponseError `xml:"error,omitempty" json:"error,omitempty" asn1:"optional"`
	Version   string           `xml:"version,omitempty" json:"version,omitempty" asn1:"optional"`
	Email     string           `xml:"email,omitempty" json:"email,omitempty" asn1:"optional"`
	Dbversion int64            `xml:"dbversion,omitempty" json:"dbversion,omitempty" asn1:"optional"`
	Bioseqs   MSBioseqSet      `xml:"bioseqs,omitempty" json:"bioseqs,omitempty" asn1:"optional"`
}
type MSSearch struct {
	Request  []MSRequest  `xml:"request,omitempty" json:"request,omitempty" asn1:"optional"`
	Response []MSResponse `xml:"response,omitempty" json:"response,omitempty" asn1:"optional"`
}
