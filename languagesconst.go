package countries

// TypeLanguageCode for Typer interface
const TypeLanguageCode string = "countries.LanguageCode"

// TypeLanguage for Typer interface
const TypeLanguage string = "countries.Language"

// UnknownLanguageCode is the unknown language code.
const UnknownLanguageCode string = "Unknown"

// ISO 639-1 language code name, 2 codes present e.g. LanguageRU == LanguageRussian = "ru".
const (
	LanguageUnknown          LanguageCode = "Unknown"
	LanguageAfar             LanguageCode = "aa"
	LanguageAbkhazian        LanguageCode = "ab"
	LanguageAvestan          LanguageCode = "ae"
	LanguageAfrikaans        LanguageCode = "af"
	LanguageAkan             LanguageCode = "ak"
	LanguageAmharic          LanguageCode = "am"
	LanguageAragonese        LanguageCode = "an"
	LanguageArabic           LanguageCode = "ar"
	LanguageAssamese         LanguageCode = "as"
	LanguageAvaric           LanguageCode = "av"
	LanguageAymara           LanguageCode = "ay"
	LanguageAzerbaijani      LanguageCode = "az"
	LanguageBashkir          LanguageCode = "ba"
	LanguageBelarusian       LanguageCode = "be"
	LanguageBulgarian        LanguageCode = "bg"
	LanguageBihariLanguages  LanguageCode = "bh"
	LanguageBislama          LanguageCode = "bi"
	LanguageBambara          LanguageCode = "bm"
	LanguageBengali          LanguageCode = "bn"
	LanguageTibetan          LanguageCode = "bo"
	LanguageBreton           LanguageCode = "br"
	LanguageBosnian          LanguageCode = "bs"
	LanguageCatalan          LanguageCode = "ca"
	LanguageChechen          LanguageCode = "ce"
	LanguageChamorro         LanguageCode = "ch"
	LanguageCorsican         LanguageCode = "co"
	LanguageCree             LanguageCode = "cr"
	LanguageCzech            LanguageCode = "cs"
	LanguageChurchSlavic     LanguageCode = "cu"
	LanguageChuvash          LanguageCode = "cv"
	LanguageWelsh            LanguageCode = "cy"
	LanguageDanish           LanguageCode = "da"
	LanguageGerman           LanguageCode = "de"
	LanguageDivehi           LanguageCode = "dv"
	LanguageDzongkha         LanguageCode = "dz"
	LanguageEwe              LanguageCode = "ee"
	LanguageGreek            LanguageCode = "el"
	LanguageEnglish          LanguageCode = "en"
	LanguageEsperanto        LanguageCode = "eo"
	LanguageSpanish          LanguageCode = "es"
	LanguageEstonian         LanguageCode = "et"
	LanguageBasque           LanguageCode = "eu"
	LanguagePersian          LanguageCode = "fa"
	LanguageFulah            LanguageCode = "ff"
	LanguageFinnish          LanguageCode = "fi"
	LanguageFijian           LanguageCode = "fj"
	LanguageFaroese          LanguageCode = "fo"
	LanguageFrench           LanguageCode = "fr"
	LanguageWesternFrisian   LanguageCode = "fy"
	LanguageIrish            LanguageCode = "ga"
	LanguageScottishGaelic   LanguageCode = "gd"
	LanguageGalician         LanguageCode = "gl"
	LanguageGuarani          LanguageCode = "gn"
	LanguageGujarati         LanguageCode = "gu"
	LanguageManx             LanguageCode = "gv"
	LanguageHausa            LanguageCode = "ha"
	LanguageHebrew           LanguageCode = "he"
	LanguageHindi            LanguageCode = "hi"
	LanguageHiriMotu         LanguageCode = "ho"
	LanguageCroatian         LanguageCode = "hr"
	LanguageHaitianCreole    LanguageCode = "ht"
	LanguageHungarian        LanguageCode = "hu"
	LanguageArmenian         LanguageCode = "hy"
	LanguageHerero           LanguageCode = "hz"
	LanguageInterlingua      LanguageCode = "ia"
	LanguageIndonesian       LanguageCode = "id"
	LanguageInterlingue      LanguageCode = "ie"
	LanguageIgbo             LanguageCode = "ig"
	LanguageSichuanYi        LanguageCode = "ii"
	LanguageInupiaq          LanguageCode = "ik"
	LanguageIdo              LanguageCode = "io"
	LanguageIcelandic        LanguageCode = "is"
	LanguageItalian          LanguageCode = "it"
	LanguageInuktitut        LanguageCode = "iu"
	LanguageJapanese         LanguageCode = "ja"
	LanguageJavanese         LanguageCode = "jv"
	LanguageGeorgian         LanguageCode = "ka"
	LanguageKongo            LanguageCode = "kg"
	LanguageKikuyu           LanguageCode = "ki"
	LanguageKuanyama         LanguageCode = "kj"
	LanguageKazakh           LanguageCode = "kk"
	LanguageKalaallisut      LanguageCode = "kl"
	LanguageCentralKhmer     LanguageCode = "km"
	LanguageKannada          LanguageCode = "kn"
	LanguageKorean           LanguageCode = "ko"
	LanguageKanuri           LanguageCode = "kr"
	LanguageKashmiri         LanguageCode = "ks"
	LanguageKurdish          LanguageCode = "ku"
	LanguageKomi             LanguageCode = "kv"
	LanguageCornish          LanguageCode = "kw"
	LanguageKirghiz          LanguageCode = "ky"
	LanguageLatin            LanguageCode = "la"
	LanguageLuxembourgish    LanguageCode = "lb"
	LanguageGanda            LanguageCode = "lg"
	LanguageLimburgish       LanguageCode = "li"
	LanguageLingala          LanguageCode = "ln"
	LanguageLao              LanguageCode = "lo"
	LanguageLithuanian       LanguageCode = "lt"
	LanguageLubaKatanga      LanguageCode = "lu"
	LanguageLatvian          LanguageCode = "lv"
	LanguageMalagasy         LanguageCode = "mg"
	LanguageMarshallese      LanguageCode = "mh"
	LanguageMaori            LanguageCode = "mi"
	LanguageMacedonian       LanguageCode = "mk"
	LanguageMalayalam        LanguageCode = "ml"
	LanguageMongolian        LanguageCode = "mn"
	LanguageMarathi          LanguageCode = "mr"
	LanguageMalay            LanguageCode = "ms"
	LanguageMaltese          LanguageCode = "mt"
	LanguageBurmese          LanguageCode = "my"
	LanguageNauru            LanguageCode = "na"
	LanguageNorwegianBokmal  LanguageCode = "nb"
	LanguageNorthNdebele     LanguageCode = "nd"
	LanguageNepali           LanguageCode = "ne"
	LanguageNdonga           LanguageCode = "ng"
	LanguageDutch            LanguageCode = "nl"
	LanguageNorwegianNynorsk LanguageCode = "nn"
	LanguageNorwegian        LanguageCode = "no"
	LanguageSouthNdebele     LanguageCode = "ns"
	LanguageNavajo           LanguageCode = "nv"
	LanguageChichewa         LanguageCode = "ny"
	LanguageOccitan          LanguageCode = "oc"
	LanguageOjibwe           LanguageCode = "oj"
	LanguageOromo            LanguageCode = "om"
	LanguageOriya            LanguageCode = "or"
	LanguageOssetian         LanguageCode = "os"
	LanguagePanjabi          LanguageCode = "pa"
	LanguagePali             LanguageCode = "pi"
	LanguagePolish           LanguageCode = "pl"
	LanguagePushto           LanguageCode = "ps"
	LanguagePortuguese       LanguageCode = "pt"
	LanguageQuechua          LanguageCode = "qu"
	LanguageRomansh          LanguageCode = "rm"
	LanguageRundi            LanguageCode = "rn"
	LanguageRomanian         LanguageCode = "ro"
	LanguageRussian          LanguageCode = "ru"
	LanguageKinyarwanda      LanguageCode = "rw"
	LanguageSanskrit         LanguageCode = "sa"
	LanguageSardinian        LanguageCode = "sc"
	LanguageSindhi           LanguageCode = "sd"
	LanguageNorthernSami     LanguageCode = "se"
	LanguageSango            LanguageCode = "sg"
	LanguageSinhalese        LanguageCode = "si"
	LanguageSlovak           LanguageCode = "sk"
	LanguageSlovenian        LanguageCode = "sl"
	LanguageSamoan           LanguageCode = "sm"
	LanguageShona            LanguageCode = "sn"
	LanguageSomali           LanguageCode = "so"
	LanguageAlbanian         LanguageCode = "sq"
	LanguageSerbian          LanguageCode = "sr"
	LanguageSwati            LanguageCode = "ss"
	LanguageSouthernSotho    LanguageCode = "st"
	LanguageSundanese        LanguageCode = "su"
	LanguageSwedish          LanguageCode = "sv"
	LanguageSwahili          LanguageCode = "sw"
	LanguageTamil            LanguageCode = "ta"
	LanguageTelugu           LanguageCode = "te"
	LanguageTajik            LanguageCode = "tg"
	LanguageThai             LanguageCode = "th"
	LanguageTigrinya         LanguageCode = "ti"
	LanguageTurkmen          LanguageCode = "tk"
	LanguageTagalog          LanguageCode = "tl"
	LanguageTswana           LanguageCode = "tn"
	LanguageTonga            LanguageCode = "to"
	LanguageTurkish          LanguageCode = "tr"
	LanguageTsonga           LanguageCode = "ts"
	LanguageTatar            LanguageCode = "tt"
	LanguageTwi              LanguageCode = "tw"
	LanguageTahitian         LanguageCode = "ty"
	LanguageUighur           LanguageCode = "ug"
	LanguageUkrainian        LanguageCode = "uk"
	LanguageUrdu             LanguageCode = "ur"
	LanguageUzbek            LanguageCode = "uz"
	LanguageVenda            LanguageCode = "ve"
	LanguageVietnamese       LanguageCode = "vi"
	LanguageVolapuk          LanguageCode = "vo"
	LanguageWalloon          LanguageCode = "wa"
	LanguageWolof            LanguageCode = "wo"
	LanguageXhosa            LanguageCode = "xh"
	LanguageYiddish          LanguageCode = "yi"
	LanguageYoruba           LanguageCode = "yo"
	LanguageZhuang           LanguageCode = "za"
	LanguageChinese          LanguageCode = "zh"
	LanguageZulu             LanguageCode = "zu"
)

// ISO 639-1 language code name, 2 codes present e.g. LanguageRU == LanguageRussian = "ru".
const (
	LanguageAA LanguageCode = "aa"
	LanguageAB LanguageCode = "ab"
	LanguageAE LanguageCode = "ae"
	LanguageAF LanguageCode = "af"
	LanguageAK LanguageCode = "ak"
	LanguageAM LanguageCode = "am"
	LanguageAN LanguageCode = "an"
	LanguageAR LanguageCode = "ar"
	LanguageAS LanguageCode = "as"
	LanguageAV LanguageCode = "av"
	LanguageAY LanguageCode = "ay"
	LanguageAZ LanguageCode = "az"
	LanguageBA LanguageCode = "ba"
	LanguageBE LanguageCode = "be"
	LanguageBG LanguageCode = "bg"
	LanguageBH LanguageCode = "bh"
	LanguageBI LanguageCode = "bi"
	LanguageBM LanguageCode = "bm"
	LanguageBN LanguageCode = "bn"
	LanguageBO LanguageCode = "bo"
	LanguageBR LanguageCode = "br"
	LanguageBS LanguageCode = "bs"
	LanguageCA LanguageCode = "ca"
	LanguageCE LanguageCode = "ce"
	LanguageCH LanguageCode = "ch"
	LanguageCO LanguageCode = "co"
	LanguageCR LanguageCode = "cr"
	LanguageCS LanguageCode = "cs"
	LanguageCU LanguageCode = "cu"
	LanguageCV LanguageCode = "cv"
	LanguageCY LanguageCode = "cy"
	LanguageDA LanguageCode = "da"
	LanguageDE LanguageCode = "de"
	LanguageDV LanguageCode = "dv"
	LanguageDZ LanguageCode = "dz"
	LanguageEE LanguageCode = "ee"
	LanguageEL LanguageCode = "el"
	LanguageEN LanguageCode = "en"
	LanguageEO LanguageCode = "eo"
	LanguageES LanguageCode = "es"
	LanguageET LanguageCode = "et"
	LanguageEU LanguageCode = "eu"
	LanguageFA LanguageCode = "fa"
	LanguageFF LanguageCode = "ff"
	LanguageFI LanguageCode = "fi"
	LanguageFJ LanguageCode = "fj"
	LanguageFO LanguageCode = "fo"
	LanguageFR LanguageCode = "fr"
	LanguageFY LanguageCode = "fy"
	LanguageGA LanguageCode = "ga"
	LanguageGD LanguageCode = "gd"
	LanguageGL LanguageCode = "gl"
	LanguageGN LanguageCode = "gn"
	LanguageGU LanguageCode = "gu"
	LanguageGV LanguageCode = "gv"
	LanguageHA LanguageCode = "ha"
	LanguageHE LanguageCode = "he"
	LanguageHI LanguageCode = "hi"
	LanguageHO LanguageCode = "ho"
	LanguageHR LanguageCode = "hr"
	LanguageHT LanguageCode = "ht"
	LanguageHU LanguageCode = "hu"
	LanguageHY LanguageCode = "hy"
	LanguageHZ LanguageCode = "hz"
	LanguageIA LanguageCode = "ia"
	LanguageID LanguageCode = "id"
	LanguageIE LanguageCode = "ie"
	LanguageIG LanguageCode = "ig"
	LanguageII LanguageCode = "ii"
	LanguageIK LanguageCode = "ik"
	LanguageIO LanguageCode = "io"
	LanguageIS LanguageCode = "is"
	LanguageIT LanguageCode = "it"
	LanguageIU LanguageCode = "iu"
	LanguageJA LanguageCode = "ja"
	LanguageJV LanguageCode = "jv"
	LanguageKA LanguageCode = "ka"
	LanguageKG LanguageCode = "kg"
	LanguageKI LanguageCode = "ki"
	LanguageKJ LanguageCode = "kj"
	LanguageKK LanguageCode = "kk"
	LanguageKL LanguageCode = "kl"
	LanguageKM LanguageCode = "km"
	LanguageKN LanguageCode = "kn"
	LanguageKO LanguageCode = "ko"
	LanguageKR LanguageCode = "kr"
	LanguageKS LanguageCode = "ks"
	LanguageKU LanguageCode = "ku"
	LanguageKV LanguageCode = "kv"
	LanguageKW LanguageCode = "kw"
	LanguageKY LanguageCode = "ky"
	LanguageLA LanguageCode = "la"
	LanguageLB LanguageCode = "lb"
	LanguageLG LanguageCode = "lg"
	LanguageLI LanguageCode = "li"
	LanguageLN LanguageCode = "ln"
	LanguageLO LanguageCode = "lo"
	LanguageLT LanguageCode = "lt"
	LanguageLU LanguageCode = "lu"
	LanguageLV LanguageCode = "lv"
	LanguageMG LanguageCode = "mg"
	LanguageMH LanguageCode = "mh"
	LanguageMI LanguageCode = "mi"
	LanguageMK LanguageCode = "mk"
	LanguageML LanguageCode = "ml"
	LanguageMN LanguageCode = "mn"
	LanguageMR LanguageCode = "mr"
	LanguageMS LanguageCode = "ms"
	LanguageMT LanguageCode = "mt"
	LanguageMY LanguageCode = "my"
	LanguageNA LanguageCode = "na"
	LanguageNB LanguageCode = "nb"
	LanguageND LanguageCode = "nd"
	LanguageNE LanguageCode = "ne"
	LanguageNG LanguageCode = "ng"
	LanguageNL LanguageCode = "nl"
	LanguageNN LanguageCode = "nn"
	LanguageNO LanguageCode = "no"
	LanguageNS LanguageCode = "ns"
	LanguageNV LanguageCode = "nv"
	LanguageNY LanguageCode = "ny"
	LanguageOC LanguageCode = "oc"
	LanguageOJ LanguageCode = "oj"
	LanguageOM LanguageCode = "om"
	LanguageOR LanguageCode = "or"
	LanguageOS LanguageCode = "os"
	LanguagePA LanguageCode = "pa"
	LanguagePI LanguageCode = "pi"
	LanguagePL LanguageCode = "pl"
	LanguagePS LanguageCode = "ps"
	LanguagePT LanguageCode = "pt"
	LanguageQU LanguageCode = "qu"
	LanguageRM LanguageCode = "rm"
	LanguageRN LanguageCode = "rn"
	LanguageRO LanguageCode = "ro"
	LanguageRU LanguageCode = "ru"
	LanguageRW LanguageCode = "rw"
	LanguageSA LanguageCode = "sa"
	LanguageSC LanguageCode = "sc"
	LanguageSD LanguageCode = "sd"
	LanguageSE LanguageCode = "se"
	LanguageSG LanguageCode = "sg"
	LanguageSI LanguageCode = "si"
	LanguageSK LanguageCode = "sk"
	LanguageSL LanguageCode = "sl"
	LanguageSM LanguageCode = "sm"
	LanguageSN LanguageCode = "sn"
	LanguageSO LanguageCode = "so"
	LanguageSQ LanguageCode = "sq"
	LanguageSR LanguageCode = "sr"
	LanguageSS LanguageCode = "ss"
	LanguageST LanguageCode = "st"
	LanguageSU LanguageCode = "su"
	LanguageSV LanguageCode = "sv"
	LanguageSW LanguageCode = "sw"
	LanguageTA LanguageCode = "ta"
	LanguageTE LanguageCode = "te"
	LanguageTG LanguageCode = "tg"
	LanguageTH LanguageCode = "th"
	LanguageTI LanguageCode = "ti"
	LanguageTK LanguageCode = "tk"
	LanguageTL LanguageCode = "tl"
	LanguageTN LanguageCode = "tn"
	LanguageTO LanguageCode = "to"
	LanguageTR LanguageCode = "tr"
	LanguageTS LanguageCode = "ts"
	LanguageTT LanguageCode = "tt"
	LanguageTW LanguageCode = "tw"
	LanguageTY LanguageCode = "ty"
	LanguageUG LanguageCode = "ug"
	LanguageUK LanguageCode = "uk"
	LanguageUR LanguageCode = "ur"
	LanguageUZ LanguageCode = "uz"
	LanguageVE LanguageCode = "ve"
	LanguageVI LanguageCode = "vi"
	LanguageVO LanguageCode = "vo"
	LanguageWA LanguageCode = "wa"
	LanguageWO LanguageCode = "wo"
	LanguageXH LanguageCode = "xh"
	LanguageYI LanguageCode = "yi"
	LanguageYO LanguageCode = "yo"
	LanguageZA LanguageCode = "za"
	LanguageZH LanguageCode = "zh"
	LanguageZU LanguageCode = "zu"
)
