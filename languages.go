package countries

import (
	"encoding/json"
	"fmt"
)

type LanguageCode string // string for database/sql/driver.Valuer compatibility

type Language struct {
	Name string
	Code LanguageCode
}

func (LanguageCode) Type() string {
	return TypeLanguageCode
}

func (lc LanguageCode) String() string { //nolint:gocyclo
	switch lc {
	case LanguageAA:
		return "Afar"
	case LanguageAB:
		return "Abkhazian"
	case LanguageAE:
		return "Avestan"
	case LanguageAF:
		return "Afrikaans"
	case LanguageAK:
		return "Akan"
	case LanguageAM:
		return "Amharic"
	case LanguageAN:
		return "Aragonese"
	case LanguageAR:
		return "Arabic"
	case LanguageAS:
		return "Assamese"
	case LanguageAV:
		return "Avaric"
	case LanguageAY:
		return "Aymara"
	case LanguageAZ:
		return "Azerbaijani"
	case LanguageBA:
		return "Bashkir"
	case LanguageBE:
		return "Belarusian"
	case LanguageBG:
		return "Bulgarian"
	case LanguageBH:
		return "Bihari"
	case LanguageBI:
		return "Bislama"
	case LanguageBM:
		return "Bambara"
	case LanguageBN:
		return "Bengali"
	case LanguageBO:
		return "Tibetan"
	case LanguageBR:
		return "Breton"
	case LanguageBS:
		return "Bosnian"
	case LanguageCA:
		return "Catalan"
	case LanguageCE:
		return "Chechen"
	case LanguageCH:
		return "Chamorro"
	case LanguageCO:
		return "Corsican"
	case LanguageCR:
		return "Cree"
	case LanguageCS:
		return "Czech"
	case LanguageCU:
		return "Old Church Slavonic"
	case LanguageCV:
		return "Chuvash"
	case LanguageCY:
		return "Welsh"
	case LanguageDA:
		return "Danish"
	case LanguageDE:
		return "German"
	case LanguageDV:
		return "Divehi"
	case LanguageDZ:
		return "Dzongkha"
	case LanguageEE:
		return "Ewe"
	case LanguageEL:
		return "Greek"
	case LanguageEN:
		return "English"
	case LanguageEO:
		return "Esperanto"
	case LanguageES:
		return "Spanish"
	case LanguageET:
		return "Estonian"
	case LanguageEU:
		return "Basque"
	case LanguageFA:
		return "Persian"
	case LanguageFF:
		return "Fulah"
	case LanguageFI:
		return "Finnish"
	case LanguageFJ:
		return "Fijian"
	case LanguageFO:
		return "Faroese"
	case LanguageFR:
		return "French"
	case LanguageFY:
		return "Western Frisian"
	case LanguageGA:
		return "Irish"
	case LanguageGD:
		return "Scottish Gaelic"
	case LanguageGL:
		return "Galician"
	case LanguageGN:
		return "Guarani"
	case LanguageGU:
		return "Gujarati"
	case LanguageGV:
		return "Manx"
	case LanguageHA:
		return "Hausa"
	case LanguageHE:
		return "Hebrew"
	case LanguageHI:
		return "Hindi"
	case LanguageHO:
		return "Hiri Motu"
	case LanguageHR:
		return "Croatian"
	case LanguageHT:
		return "Haitian"
	case LanguageHU:
		return "Hungarian"
	case LanguageHY:
		return "Armenian"
	case LanguageHZ:
		return "Herero"
	case LanguageIA:
		return "Interlingua"
	case LanguageID:
		return "Indonesian"
	case LanguageIE:
		return "Interlingue"
	case LanguageIG:
		return "Igbo"
	case LanguageII:
		return "Sichuan Yi"
	case LanguageIK:
		return "Inupiaq"
	case LanguageIO:
		return "Ido"
	case LanguageIS:
		return "Icelandic"
	case LanguageIT:
		return "Italian"
	case LanguageIU:
		return "Inuktitut"
	case LanguageJA:
		return "Japanese"
	case LanguageJV:
		return "Javanese"
	case LanguageKA:
		return "Georgian"
	case LanguageKG:
		return "Kongo"
	case LanguageKI:
		return "Kikuyu"
	case LanguageKJ:
		return "Kwanyama"
	case LanguageKK:
		return "Kazakh"
	case LanguageKL:
		return "Kalaallisut"
	case LanguageKM:
		return "Khmer"
	case LanguageKN:
		return "Kannada"
	case LanguageKO:
		return "Korean"
	case LanguageKR:
		return "Kanuri"
	case LanguageKS:
		return "Kashmiri"
	case LanguageKU:
		return "Kurdish"
	case LanguageKV:
		return "Komi"
	case LanguageKW:
		return "Cornish"
	case LanguageKY:
		return "Kyrgyz"
	case LanguageLA:
		return "Latin"
	case LanguageLB:
		return "Luxembourgish"
	case LanguageLG:
		return "Ganda"
	case LanguageLI:
		return "Limburgish"
	case LanguageLN:
		return "Lingala"
	case LanguageLO:
		return "Lao"
	case LanguageLT:
		return "Lithuanian"
	case LanguageLU:
		return "Luba-Katanga"
	case LanguageLV:
		return "Latvian"
	case LanguageMG:
		return "Malagasy"
	case LanguageMH:
		return "Marshallese"
	case LanguageMI:
		return "Maori"
	case LanguageMK:
		return "Macedonian"
	case LanguageML:
		return "Malayalam"
	case LanguageMN:
		return "Mongolian"
	case LanguageMR:
		return "Marathi"
	case LanguageMS:
		return "Malay"
	case LanguageMT:
		return "Maltese"
	case LanguageMY:
		return "Burmese"
	case LanguageNA:
		return "Nauru"
	case LanguageNB:
		return "Norwegian Bokmål"
	case LanguageND:
		return "North Ndebele"
	case LanguageNE:
		return "Nepali"
	case LanguageNG:
		return "Ndonga"
	case LanguageNL:
		return "Dutch"
	case LanguageNN:
		return "Norwegian Nynorsk"
	case LanguageNO:
		return "Norwegian"
	case LanguageNR:
		return "South Ndebele"
	case LanguageNS:
		return "Northern Sotho"
	case LanguageNV:
		return "Navajo"
	case LanguageNY:
		return "Chichewa"
	case LanguageOC:
		return "Occitan"
	case LanguageOJ:
		return "Ojibwa"
	case LanguageOM:
		return "Oromo"
	case LanguageOR:
		return "Oriya"
	case LanguageOS:
		return "Ossetian"
	case LanguagePA:
		return "Panjabi"
	case LanguagePI:
		return "Pali"
	case LanguagePL:
		return "Polish"
	case LanguagePS:
		return "Pashto"
	case LanguagePT:
		return "Portuguese"
	case LanguageQU:
		return "Quechua"
	case LanguageRM:
		return "Romansh"
	case LanguageRN:
		return "Rundi"
	case LanguageRO:
		return "Romanian"
	case LanguageRU:
		return "Russian"
	case LanguageRW:
		return "Kinyarwanda"
	case LanguageSA:
		return "Sanskrit"
	case LanguageSC:
		return "Sardinian"
	case LanguageSD:
		return "Sindhi"
	case LanguageSE:
		return "Northern Sami"
	case LanguageSG:
		return "Sango"
	case LanguageSI:
		return "Sinhala"
	case LanguageSK:
		return "Slovak"
	case LanguageSL:
		return "Slovene"
	case LanguageSM:
		return "Samoan"
	case LanguageSN:
		return "Shona"
	case LanguageSO:
		return "Somali"
	case LanguageSQ:
		return "Albanian"
	case LanguageSR:
		return "Serbian"
	case LanguageSS:
		return "Swati"
	case LanguageST:
		return "Southern Sotho"
	case LanguageSU:
		return "Sundanese"
	case LanguageSV:
		return "Swedish"
	case LanguageSW:
		return "Swahili"
	case LanguageTA:
		return "Tamil"
	case LanguageTE:
		return "Telugu"
	case LanguageTG:
		return "Tajik"
	case LanguageTH:
		return "Thai"
	case LanguageTI:
		return "Tigrinya"
	case LanguageTK:
		return "Turkmen"
	case LanguageTL:
		return "Tagalog"
	case LanguageTN:
		return "Tswana"
	case LanguageTO:
		return "Tonga"
	case LanguageTR:
		return "Turkish"
	case LanguageTS:
		return "Tsonga"
	case LanguageTT:
		return "Tatar"
	case LanguageTVL:
		return "Tuvaluan"
	case LanguageTW:
		return "Twi"
	case LanguageTY:
		return "Tahitian"
	case LanguageUG:
		return "Uighur"
	case LanguageUK:
		return "Ukrainian"
	case LanguageUR:
		return "Urdu"
	case LanguageUZ:
		return "Uzbek"
	case LanguageVE:
		return "Venda"
	case LanguageVI:
		return "Vietnamese"
	case LanguageVO:
		return "Volapük"
	case LanguageWA:
		return "Walloon"
	case LanguageWO:
		return "Wolof"
	case LanguageXH:
		return "Xhosa"
	case LanguageYI:
		return "Yiddish"
	case LanguageYO:
		return "Yoruba"
	case LanguageZA:
		return "Zhuang"
	case LanguageZH:
		return "Chinese"
	case LanguageZU:
		return "Zulu"
	default:
		return UnknownLanguageCode
	}
}

func (lc LanguageCode) Alpha2() string {
	if len(string(lc)) == 2 {
		return string(lc)
	}

	return "unknown"
}

func (lc LanguageCode) Alpha3() string {
	switch lc {
	case LanguageAA:
		return "aar"
	case LanguageAB:
		return "abk"
	case LanguageAE:
		return "ave"
	case LanguageAF:
		return "afr"
	case LanguageAK:
		return "aka"
	case LanguageAM:
		return "amh"
	case LanguageAN:
		return "arg"
	case LanguageAR:
		return "ara"
	case LanguageAS:
		return "asm"
	case LanguageAV:
		return "ava"
	case LanguageAY:
		return "aym"
	case LanguageAZ:
		return "aze"
	case LanguageBA:
		return "bak"
	case LanguageBE:
		return "bel"
	case LanguageBG:
		return "bul"
	case LanguageBH:
		return "bih"
	case LanguageBI:
		return "bis"
	case LanguageBM:
		return "bam"
	case LanguageBN:
		return "ben"
	case LanguageBO:
		return "bod"
	case LanguageBR:
		return "bre"
	case LanguageBS:
		return "bos"
	case LanguageCA:
		return "cat"
	case LanguageCE:
		return "che"
	case LanguageCH:
		return "cha"
	case LanguageCO:
		return "cos"
	case LanguageCR:
		return "cre"
	case LanguageCS:
		return "ces"
	case LanguageCU:
		return "chu"
	case LanguageCV:
		return "chv"
	case LanguageCY:
		return "cym"
	case LanguageDA:
		return "dan"
	case LanguageDE:
		return "deu"
	case LanguageDV:
		return "div"
	case LanguageDZ:
		return "dzo"
	case LanguageEE:
		return "ewe"
	case LanguageEL:
		return "ell"
	case LanguageEN:
		return "eng"
	case LanguageEO:
		return "epo"
	case LanguageES:
		return "spa"
	case LanguageET:
		return "est"
	case LanguageEU:
		return "eus"
	case LanguageFA:
		return "fas"
	case LanguageFF:
		return "ful"
	case LanguageFI:
		return "fin"
	case LanguageFJ:
		return "fij"
	case LanguageFO:
		return "fao"
	case LanguageFR:
		return "fra"
	case LanguageFY:
		return "fry"
	case LanguageGA:
		return "gle"
	case LanguageGD:
		return "gla"
	case LanguageGL:
		return "glg"
	case LanguageGN:
		return "grn"
	case LanguageGU:
		return "guj"
	case LanguageGV:
		return "glv"
	case LanguageHA:
		return "hau"
	case LanguageHE:
		return "heb"
	case LanguageHI:
		return "hin"
	case LanguageHO:
		return "hmo"
	case LanguageHR:
		return "hrv"
	case LanguageHT:
		return "hat"
	case LanguageHU:
		return "hun"
	case LanguageHY:
		return "hye"
	case LanguageHZ:
		return "her"
	case LanguageIA:
		return "ina"
	case LanguageID:
		return "ind"
	case LanguageIE:
		return "ile"
	case LanguageIG:
		return "ibo"
	case LanguageII:
		return "iii"
	case LanguageIK:
		return "ipk"
	case LanguageIO:
		return "ido"
	case LanguageIS:
		return "isl"
	case LanguageIT:
		return "ita"
	case LanguageIU:
		return "iku"
	case LanguageJA:
		return "jpn"
	case LanguageJV:
		return "jav"
	case LanguageKA:
		return "kat"
	case LanguageKG:
		return "kon"
	case LanguageKI:
		return "kik"
	case LanguageKJ:
		return "kua"
	case LanguageKK:
		return "kaz"
	case LanguageKL:
		return "kal"
	case LanguageKM:
		return "khm"
	case LanguageKN:
		return "kan"
	case LanguageKO:
		return "kor"
	case LanguageKR:
		return "kau"
	case LanguageKS:
		return "kas"
	case LanguageKU:
		return "kur"
	case LanguageKV:
		return "kom"
	case LanguageKW:
		return "cor"
	case LanguageKY:
		return "kir"
	case LanguageLA:
		return "lat"
	case LanguageLB:
		return "ltz"
	case LanguageLG:
		return "lug"
	case LanguageLI:
		return "lim"
	case LanguageLN:
		return "lin"
	case LanguageLO:
		return "lao"
	case LanguageLT:
		return "lit"
	case LanguageLU:
		return "lub"
	case LanguageLV:
		return "lav"
	case LanguageMG:
		return "mlg"
	case LanguageMH:
		return "mah"
	case LanguageMI:
		return "mri"
	case LanguageMK:
		return "mkd"
	case LanguageML:
		return "mal"
	case LanguageMN:
		return "mon"
	case LanguageMR:
		return "mar"
	case LanguageMS:
		return "msa"
	case LanguageMT:
		return "mlt"
	case LanguageMY:
		return "mya"
	case LanguageNA:
		return "nau"
	case LanguageNB:
		return "nob"
	case LanguageND:
		return "nde"
	case LanguageNE:
		return "nep"
	case LanguageNG:
		return "ndo"
	case LanguageNL:
		return "nld"
	case LanguageNN:
		return "nno"
	case LanguageNO:
		return "nor"
	case LanguageNR:
		return "nbl"
	case LanguageNS:
		return "nso"
	case LanguageNV:
		return "nav"
	case LanguageNY:
		return "nya"
	case LanguageOC:
		return "oci"
	case LanguageOJ:
		return "oji"
	case LanguageOM:
		return "orm"
	case LanguageOR:
		return "ori"
	case LanguageOS:
		return "oss"
	case LanguagePA:
		return "pan"
	case LanguagePI:
		return "pli"
	case LanguagePL:
		return "pol"
	case LanguagePS:
		return "pus"
	case LanguagePT:
		return "por"
	case LanguageQU:
		return "que"
	case LanguageRM:
		return "roh"
	case LanguageRN:
		return "run"
	case LanguageRO:
		return "ron"
	case LanguageRU:
		return "rus"
	case LanguageRW:
		return "kin"
	case LanguageSA:
		return "san"
	case LanguageSC:
		return "srd"
	case LanguageSD:
		return "snd"
	case LanguageSE:
		return "sme"
	case LanguageSG:
		return "sag"
	case LanguageSI:
		return "sin"
	case LanguageSK:
		return "slk"
	case LanguageSL:
		return "slv"
	case LanguageSM:
		return "smo"
	case LanguageSN:
		return "sna"
	case LanguageSO:
		return "som"
	case LanguageSQ:
		return "sqi"
	case LanguageSR:
		return "srp"
	case LanguageSS:
		return "ssw"
	case LanguageST:
		return "sot"
	case LanguageSU:
		return "sun"
	case LanguageSV:
		return "swe"
	case LanguageSW:
		return "swa"
	case LanguageTA:
		return "tam"
	case LanguageTE:
		return "tel"
	case LanguageTG:
		return "tgk"
	case LanguageTH:
		return "tha"
	case LanguageTI:
		return "tir"
	case LanguageTK:
		return "tuk"
	case LanguageTL:
		return "tgl"
	case LanguageTN:
		return "tsn"
	case LanguageTO:
		return "ton"
	case LanguageTR:
		return "tur"
	case LanguageTS:
		return "tso"
	case LanguageTT:
		return "tat"
	case LanguageTW:
		return "twi"
	case LanguageTVL:
		return "tvl"
	case LanguageTY:
		return "tah"
	case LanguageUG:
		return "uig"
	case LanguageUK:
		return "ukr"
	case LanguageUR:
		return "urd"
	case LanguageUZ:
		return "uzb"
	case LanguageVE:
		return "ven"
	case LanguageVI:
		return "vie"
	case LanguageVO:
		return "vol"
	case LanguageWA:
		return "wln"
	case LanguageWO:
		return "wol"
	case LanguageXH:
		return "xho"
	case LanguageYI:
		return "yid"
	case LanguageYO:
		return "yor"
	case LanguageZA:
		return "zha"
	case LanguageZH:
		return "zho"
	case LanguageZU:
		return "zul"
	}
	return "unknown"
}

// StringRus returns the language name in Russian.
// StringRus возвращает название языка на русском языке.
func (lc LanguageCode) StringRus() string { //nolint:gocyclo
	switch lc {
	case LanguageAA:
		return "Афар"
	case LanguageAB:
		return "Абхазский"
	case LanguageAE:
		return "Авестийский"
	case LanguageAF:
		return "Африкаанс"
	case LanguageAK:
		return "Акан"
	case LanguageAM:
		return "Амхарский"
	case LanguageAN:
		return "Арагонский"
	case LanguageAR:
		return "Арабский"
	case LanguageAS:
		return "Ассамский"
	case LanguageAV:
		return "Аварский"
	case LanguageAY:
		return "Аймара"
	case LanguageAZ:
		return "Азербайджанский"
	case LanguageBA:
		return "Башкирский"
	case LanguageBE:
		return "Белорусский"
	case LanguageBG:
		return "Болгарский"
	case LanguageBH:
		return "Бихари"
	case LanguageBI:
		return "Бислама"
	case LanguageBM:
		return "Бамбара"
	case LanguageBN:
		return "Бенгальский"
	case LanguageBO:
		return "Тибетский"
	case LanguageBR:
		return "Бретонский"
	case LanguageBS:
		return "Боснийский"
	case LanguageCA:
		return "Каталанский"
	case LanguageCE:
		return "Чеченский"
	case LanguageCH:
		return "Чаморро"
	case LanguageCO:
		return "Корсиканский"
	case LanguageCR:
		return "Кри"
	case LanguageCS:
		return "Чешский"
	case LanguageCU:
		return "Церковнославянский"
	case LanguageCV:
		return "Чувашский"
	case LanguageCY:
		return "Валлийский"
	case LanguageDA:
		return "Датский"
	case LanguageDE:
		return "Немецкий"
	case LanguageDV:
		return "Дивехи"
	case LanguageDZ:
		return "Дзонг-кэ"
	case LanguageEE:
		return "Эве"
	case LanguageEL:
		return "Греческий"
	case LanguageEN:
		return "Английский"
	case LanguageEO:
		return "Эсперанто"
	case LanguageES:
		return "Испанский"
	case LanguageET:
		return "Эстонский"
	case LanguageEU:
		return "Баскский"
	case LanguageFA:
		return "Персидский"
	case LanguageFF:
		return "Фула"
	case LanguageFI:
		return "Финский"
	case LanguageFJ:
		return "Фиджи"
	case LanguageFO:
		return "Фарерский"
	case LanguageFR:
		return "Французский"
	case LanguageFY:
		return "Западнофризский"
	case LanguageGA:
		return "Ирландский"
	case LanguageGD:
		return "Шотландский гэльский"
	case LanguageGL:
		return "Галисийский"
	case LanguageGN:
		return "Гуарани"
	case LanguageGU:
		return "Гуджарати"
	case LanguageGV:
		return "Мэнский"
	case LanguageHA:
		return "Хауса"
	case LanguageHE:
		return "Иврит"
	case LanguageHI:
		return "Хинди"
	case LanguageHO:
		return "Хири-моту"
	case LanguageHR:
		return "Хорватский"
	case LanguageHT:
		return "Гаитянский"
	case LanguageHU:
		return "Венгерский"
	case LanguageHY:
		return "Армянский"
	case LanguageHZ:
		return "Гереро"
	case LanguageIA:
		return "Интерлингва"
	case LanguageID:
		return "Индонезийский"
	case LanguageIE:
		return "Интерлингве"
	case LanguageIG:
		return "Игбо"
	case LanguageII:
		return "Сычуань И"
	case LanguageIK:
		return "Инупиак"
	case LanguageIO:
		return "Идо"
	case LanguageIS:
		return "Исландский"
	case LanguageIT:
		return "Итальянский"
	case LanguageIU:
		return "Инуктитут"
	case LanguageJA:
		return "Японский"
	case LanguageJV:
		return "Яванский"
	case LanguageKA:
		return "Грузинский"
	case LanguageKG:
		return "Конго"
	case LanguageKI:
		return "Кикуйю"
	case LanguageKJ:
		return "Кваньяма"
	case LanguageKK:
		return "Казахский"
	case LanguageKL:
		return "Гренландский"
	case LanguageKM:
		return "Кхмерский"
	case LanguageKN:
		return "Каннада"
	case LanguageKO:
		return "Корейский"
	case LanguageKR:
		return "Канури"
	case LanguageKS:
		return "Кашмири"
	case LanguageKU:
		return "Курдский"
	case LanguageKV:
		return "Коми"
	case LanguageKW:
		return "Корнский"
	case LanguageKY:
		return "Киргизский"
	case LanguageLA:
		return "Латинский"
	case LanguageLB:
		return "Люксембургский"
	case LanguageLG:
		return "Ганда"
	case LanguageLI:
		return "Лимбургский"
	case LanguageLN:
		return "Лингала"
	case LanguageLO:
		return "Лаосский"
	case LanguageLT:
		return "Литовский"
	case LanguageLU:
		return "Луба-катанга"
	case LanguageLV:
		return "Латышский"
	case LanguageMG:
		return "Малагасийский"
	case LanguageMH:
		return "Маршалльский"
	case LanguageMI:
		return "Маори"
	case LanguageMK:
		return "Македонский"
	case LanguageML:
		return "Малаялам"
	case LanguageMN:
		return "Монгольский"
	case LanguageMR:
		return "Маратхи"
	case LanguageMS:
		return "Малайский"
	case LanguageMT:
		return "Мальтийский"
	case LanguageMY:
		return "Бирманский"
	case LanguageNA:
		return "Науру"
	case LanguageNB:
		return "Норвежский букмол"
	case LanguageND:
		return "Северный ндебеле"
	case LanguageNE:
		return "Непальский"
	case LanguageNG:
		return "Ндонга"
	case LanguageNL:
		return "Голландский"
	case LanguageNN:
		return "Норвежский нюнорск"
	case LanguageNO:
		return "Норвежский"
	case LanguageNR:
		return "Южный ндебеле"
	case LanguageNS:
		return "Северный сото"
	case LanguageNV:
		return "Навахо"
	case LanguageNY:
		return "Чичева"
	case LanguageOC:
		return "Окситанский"
	case LanguageOJ:
		return "Оджибве"
	case LanguageOM:
		return "Оромо"
	case LanguageOR:
		return "Ория"
	case LanguageOS:
		return "Осетинский"
	case LanguagePA:
		return "Панджаби"
	case LanguagePI:
		return "Пали"
	case LanguagePL:
		return "Польский"
	case LanguagePS:
		return "Пушту"
	case LanguagePT:
		return "Португальский"
	case LanguageQU:
		return "Кечуа"
	case LanguageRM:
		return "Ретороманский"
	case LanguageRN:
		return "Рунди"
	case LanguageRO:
		return "Румынский"
	case LanguageRU:
		return "Русский"
	case LanguageRW:
		return "Киньяруанда"
	case LanguageSA:
		return "Санскрит"
	case LanguageSC:
		return "Сардинский"
	case LanguageSD:
		return "Синдхи"
	case LanguageSE:
		return "Северносаамский"
	case LanguageSG:
		return "Санго"
	case LanguageSI:
		return "Сингальский"
	case LanguageSK:
		return "Словацкий"
	case LanguageSL:
		return "Словенский"
	case LanguageSM:
		return "Самоанский"
	case LanguageSN:
		return "Шона"
	case LanguageSO:
		return "Сомалийский"
	case LanguageSQ:
		return "Албанский"
	case LanguageSR:
		return "Сербский"
	case LanguageSS:
		return "Свази"
	case LanguageST:
		return "Сесото"
	case LanguageSU:
		return "Сунданский"
	case LanguageSV:
		return "Шведский"
	case LanguageSW:
		return "Суахили"
	case LanguageTA:
		return "Тамильский"
	case LanguageTE:
		return "Телугу"
	case LanguageTG:
		return "Таджикский"
	case LanguageTH:
		return "Тайский"
	case LanguageTI:
		return "Тигринья"
	case LanguageTK:
		return "Туркменский"
	case LanguageTL:
		return "Тагальский"
	case LanguageTN:
		return "Тсвана"
	case LanguageTO:
		return "Тонганский"
	case LanguageTR:
		return "Турецкий"
	case LanguageTS:
		return "Тсонга"
	case LanguageTT:
		return "Татарский"
	case LanguageTW:
		return "Тви"
	case LanguageTVL:
		return "тувалуанский"
	case LanguageTY:
		return "Таитянский"
	case LanguageUG:
		return "Уйгурский"
	case LanguageUK:
		return "Украинский"
	case LanguageUR:
		return "Урду"
	case LanguageUZ:
		return "Узбекский"
	case LanguageVE:
		return "Венда"
	case LanguageVI:
		return "Вьетнамский"
	case LanguageVO:
		return "Волапюк"
	case LanguageWA:
		return "Валлонский"
	case LanguageWO:
		return "Волоф"
	case LanguageXH:
		return "Коса"
	case LanguageYI:
		return "Идиш"
	case LanguageYO:
		return "Йоруба"
	case LanguageZA:
		return "Чжуанский"
	case LanguageZH:
		return "Китайский"
	case LanguageZU:
		return "Зулу"
	default:
		return "Неизвестный"
	}
}

// StringLocal returns the language name in the local language.
func (lc LanguageCode) StringLocal() string {
	switch lc {
	case LanguageAA:
		return "Qafar af"
	case LanguageAB:
		return "Аҧсуа"
	case LanguageAE:
		return "Avesta"
	case LanguageAF:
		return "Afrikaans"
	case LanguageAK:
		return "Ákán"
	case LanguageAM:
		return "አማርኛ"
	case LanguageAN:
		return "Aragonés"
	case LanguageAR:
		return "العربية"
	case LanguageAS:
		return "অসমীয়া"
	case LanguageAV:
		return "Авар мацӀ"
	case LanguageAY:
		return "Aymara"
	case LanguageAZ:
		return "Azərbaycan dili"
	case LanguageBA:
		return "башҡорт теле"
	case LanguageBE:
		return "беларуская"
	case LanguageBG:
		return "български"
	case LanguageBH:
		return "भोजपुरी"
	case LanguageBI:
		return "Bislama"
	case LanguageBM:
		return "ߓߡߊߣߊ߲"
	case LanguageBN:
		return "বাংলা"
	case LanguageBO:
		return "བོད་ཡིག"
	case LanguageBR:
		return "Brezhoneg"
	case LanguageBS:
		return "Босански"
	case LanguageCA:
		return "català"
	case LanguageCE:
		return "нохчийн мотт"
	case LanguageCH:
		return "Chamoru"
	case LanguageCO:
		return "corsu"
	case LanguageCR:
		return "ᓀᐦᐃᔭᐍᐏᐣ"
	case LanguageCS:
		return "čeština"
	case LanguageCU:
		return "ѩзыкъ словѣньскъ"
	case LanguageCV:
		return "чӑваш чӗлхи"
	case LanguageCY:
		return "Cymraeg"
	case LanguageDA:
		return "Dansk"
	case LanguageDE:
		return "Deutsch"
	case LanguageDV:
		return "ދިވެހި"
	case LanguageDZ:
		return "རྫོང་ཁ"
	case LanguageEE:
		return "Eʋegbe"
	case LanguageEL:
		return "Ελληνικά"
	case LanguageEN:
		return "English"
	case LanguageEO:
		return "Esperanto"
	case LanguageES:
		return "Español"
	case LanguageET:
		return "Eesti"
	case LanguageEU:
		return "Euskara"
	case LanguageFA:
		return "فارسی"
	case LanguageFF:
		return "Fulfulde"
	case LanguageFI:
		return "Suomi"
	case LanguageFJ:
		return "Vosa Vakaviti"
	case LanguageFO:
		return "Føroyskt"
	case LanguageFR:
		return "Français"
	case LanguageFY:
		return "Frysk"
	case LanguageGA:
		return "Gaeilge"
	case LanguageGD:
		return "Gàidhlig"
	case LanguageGL:
		return "Galego"
	case LanguageGN:
		return "Avañe'ẽ"
	case LanguageGU:
		return "ગુજરાતી"
	case LanguageGV:
		return "Gaelg"
	case LanguageHA:
		return "Hausa"
	case LanguageHE:
		return "עברית"
	case LanguageHI:
		return "हिन्दी"
	case LanguageHO:
		return "Hiri Motu"
	case LanguageHR:
		return "Hrvatski"
	case LanguageHT:
		return "Kreyòl ayisyen"
	case LanguageHU:
		return "Magyar"
	case LanguageHY:
		return "Հայերեն"
	case LanguageHZ:
		return "Otjiherero"
	case LanguageIA:
		return "Interlingua"
	case LanguageID:
		return "Bahasa Indonesia"
	case LanguageIE:
		return "Interlingue"
	case LanguageIG:
		return "Asụsụ Igbo"
	case LanguageII:
		return "ꆇꉙ"
	case LanguageIK:
		return "Iñupiatun"
	case LanguageIO:
		return "Ido"
	case LanguageIS:
		return "Íslenska"
	case LanguageIT:
		return "Italiano"
	case LanguageIU:
		return "ᐃᓄᒃᑎᑐᑦ"
	case LanguageJA:
		return "日本語"
	case LanguageJV:
		return "ꦧꦱꦗꦮ"
	case LanguageKA:
		return "ქართული"
	case LanguageKG:
		return "Kikongo"
	case LanguageKI:
		return "Gĩkũyũ"
	case LanguageKJ:
		return "Kuanyama"
	case LanguageKK:
		return "қазақ тілі"
	case LanguageKL:
		return "Kalaallisut"
	case LanguageKM:
		return "ខ្មែរ"
	case LanguageKN:
		return "ಕನ್ನಡ"
	case LanguageKO:
		return "한국어"
	case LanguageKR:
		return "Kanuri"
	case LanguageKS:
		return "कश्मीरी"
	case LanguageKU:
		return "Kurdî"
	case LanguageKV:
		return "коми кыв"
	case LanguageKW:
		return "Kernewek"
	case LanguageKY:
		return "Кыргызча"
	case LanguageLA:
		return "Latine"
	case LanguageLB:
		return "Lëtzebuergesch"
	case LanguageLG:
		return "Luganda"
	case LanguageLI:
		return "Limburgs"
	case LanguageLN:
		return "Lingála"
	case LanguageLO:
		return "ພາສາລາວ"
	case LanguageLT:
		return "Lietuvių kalba"
	case LanguageLU:
		return "Tshiluba"
	case LanguageLV:
		return "Latviešu valoda"
	case LanguageMG:
		return "Fiteny malagasy"
	case LanguageMH:
		return "Kajin M̧ajeļ"
	case LanguageMI:
		return "Te reo Māori"
	case LanguageMK:
		return "македонски јазик"
	case LanguageML:
		return "മലയാളം"
	case LanguageMN:
		return "монгол"
	case LanguageMR:
		return "मराठी"
	case LanguageMS:
		return "Bahasa Melayu"
	case LanguageMT:
		return "Malti"
	case LanguageMY:
		return "ဗမာစာ"
	case LanguageNA:
		return "Ekakairũ Naoero"
	case LanguageNB:
		return "Norsk bokmål"
	case LanguageND:
		return "Ndebele"
	case LanguageNE:
		return "नेपाली"
	case LanguageNG:
		return "Owambo"
	case LanguageNL:
		return "Nederlands"
	case LanguageNN:
		return "Norsk nynorsk"
	case LanguageNO:
		return "Norsk"
	case LanguageNR:
		return "isiNdebele"
	case LanguageNS:
		return "Sesotho sa Leboa"
	case LanguageNV:
		return "Diné bizaad"
	case LanguageNY:
		return "chichewa"
	case LanguageOC:
		return "Occitan"
	case LanguageOJ:
		return "ᐊᓂᔑᓈᐯᒧᐎᓐ"
	case LanguageOM:
		return "Afaan Oromoo"
	case LanguageOR:
		return "ଓଡ଼ିଆ"
	case LanguageOS:
		return "ирон æвзаг"
	case LanguagePA:
		return "ਪੰਜਾਬੀ"
	case LanguagePI:
		return "पाऴि"
	case LanguagePL:
		return "Polski"
	case LanguagePS:
		return "پښتو"
	case LanguagePT:
		return "Português"
	case LanguageQU:
		return "Runa Simi"
	case LanguageRM:
		return "Rumantsch grischun"
	case LanguageRN:
		return "Ikirundi"
	case LanguageRO:
		return "Română"
	case LanguageRU:
		return "Русский"
	case LanguageRW:
		return "Ikinyarwanda"
	case LanguageSA:
		return "संस्कृतम्"
	case LanguageSC:
		return "Sardu"
	case LanguageSD:
		return "सिन्धी"
	case LanguageSE:
		return "Davvisámegiella"
	case LanguageSG:
		return "Yângâ tî sängö"
	case LanguageSI:
		return "සිංහල"
	case LanguageSK:
		return "Slovenčina"
	case LanguageSL:
		return "Slovenščina"
	case LanguageSM:
		return "Gagana fa'a Samoa"
	case LanguageSN:
		return "chiShona"
	case LanguageSO:
		return "Soomaaliga"
	case LanguageSQ:
		return "Shqip"
	case LanguageSR:
		return "Српски"
	case LanguageSS:
		return "SiSwati"
	case LanguageST:
		return "Sesotho"
	case LanguageSU:
		return "Basa Sunda"
	case LanguageSV:
		return "Svenska"
	case LanguageSW:
		return "Kiswahili"
	case LanguageTA:
		return "தமிழ்"
	case LanguageTE:
		return "తెలుగు"
	case LanguageTG:
		return "тоҷикӣ"
	case LanguageTH:
		return "ไทย"
	case LanguageTI:
		return "ትግርኛ"
	case LanguageTK:
		return "Türkmen"
	case LanguageTL:
		return "Wikang Tagalog"
	case LanguageTN:
		return "Setswana"
	case LanguageTO:
		return "lea faka-Tonga"
	case LanguageTR:
		return "Türkçe"
	case LanguageTS:
		return "Xitsonga"
	case LanguageTT:
		return "татар теле"
	case LanguageTW:
		return "Twi"
	case LanguageTVL:
		return "Te Ggana Tuuvalu"
	case LanguageTY:
		return "Reo Tahiti"
	case LanguageUG:
		return "ئۇيغۇرچە"
	case LanguageUK:
		return "українська"
	case LanguageUR:
		return "اردو"
	case LanguageUZ:
		return "Oʻzbek"
	case LanguageVE:
		return "Tshivenḓa"
	case LanguageVI:
		return "Tiếng Việt"
	case LanguageVO:
		return "Volapük"
	case LanguageWA:
		return "walon"
	case LanguageWO:
		return "Wollof"
	case LanguageXH:
		return "isiXhosa"
	case LanguageYI:
		return "ייִדיש"
	case LanguageYO:
		return "Yorùbá"
	case LanguageZA:
		return "Saɯ cueŋƅ"
	case LanguageZH:
		return "中文"
	case LanguageZU:
		return "isiZulu"
	}
	return "unknown"
}

func (lc LanguageCode) IsValid() bool {
	return lc.String() != UnknownLanguageCode
}

func TotalLanguages() int {
	return len(AllLanguages())
}

func (lc LanguageCode) Info() *Language {
	return &Language{
		Name: lc.String(),
		Code: lc,
	}
}

func (*Language) Type() string {
	return TypeLanguage
}

func (l Language) Value() (Value, error) {
	return json.Marshal(l)
}

func (l *Language) Scan(src interface{}) error {
	if l == nil {
		return fmt.Errorf("countries::Scan: Language scan err: language == nil")
	}
	switch src := src.(type) {
	case *Language:
		*l = *src
	case Language:
		*l = src
	}
	return nil
}

func AllLanguages() []LanguageCode {
	return []LanguageCode{
		LanguageAA, LanguageAB, LanguageAE, LanguageAF, LanguageAK, LanguageAM, LanguageAN, LanguageAR, LanguageAS, LanguageAV, LanguageAY, LanguageAZ,
		LanguageBA, LanguageBE, LanguageBG, LanguageBH, LanguageBI, LanguageBM, LanguageBN, LanguageBO, LanguageBR, LanguageBS, LanguageCA, LanguageCE,
		LanguageCH, LanguageCO, LanguageCR, LanguageCS, LanguageCU, LanguageCV, LanguageCY, LanguageDA, LanguageDE, LanguageDV, LanguageDZ, LanguageEE,
		LanguageEL, LanguageEN, LanguageEO, LanguageES, LanguageET, LanguageEU, LanguageFA, LanguageFF, LanguageFI, LanguageFJ, LanguageFO, LanguageFR,
		LanguageFY, LanguageGA, LanguageGD, LanguageGL, LanguageGN, LanguageGU, LanguageGV, LanguageHA, LanguageHE, LanguageHI, LanguageHO, LanguageHR,
		LanguageHT, LanguageHU, LanguageHY, LanguageHZ, LanguageIA, LanguageID, LanguageIE, LanguageIG, LanguageII, LanguageIK, LanguageIO, LanguageIS,
		LanguageIT, LanguageIU, LanguageJA, LanguageJV, LanguageKA, LanguageKG, LanguageKI, LanguageKJ, LanguageKK, LanguageKL, LanguageKM, LanguageKN,
		LanguageKO, LanguageKR, LanguageKS, LanguageKU, LanguageKV, LanguageKW, LanguageKY, LanguageLA, LanguageLB, LanguageLG, LanguageLI, LanguageLN,
		LanguageLO, LanguageLT, LanguageLU, LanguageLV, LanguageMG, LanguageMH, LanguageMI, LanguageMK, LanguageML, LanguageMN, LanguageMR, LanguageMS,
		LanguageMT, LanguageMY, LanguageNA, LanguageNB, LanguageND, LanguageNE, LanguageNG, LanguageNL, LanguageNN, LanguageNO, LanguageNR, LanguageNS,
		LanguageNV, LanguageNY, LanguageOC, LanguageOJ, LanguageOM, LanguageOR, LanguageOS, LanguagePA, LanguagePI, LanguagePL, LanguagePS, LanguagePT,
		LanguageQU, LanguageRM, LanguageRN, LanguageRO, LanguageRU, LanguageRW, LanguageSA, LanguageSC, LanguageSD, LanguageSE, LanguageSG, LanguageSI,
		LanguageSK, LanguageSL, LanguageSM, LanguageSN, LanguageSO, LanguageSQ, LanguageSR, LanguageSS, LanguageST, LanguageSU, LanguageSV, LanguageSW,
		LanguageTA, LanguageTE, LanguageTG, LanguageTH, LanguageTI, LanguageTK, LanguageTL, LanguageTN, LanguageTO, LanguageTR, LanguageTS, LanguageTT,
		LanguageTW, LanguageTVL, LanguageTY, LanguageUG, LanguageUK, LanguageUR, LanguageUZ, LanguageVE, LanguageVI, LanguageVO, LanguageWA, LanguageWO,
		LanguageXH, LanguageYI, LanguageYO, LanguageZA, LanguageZH, LanguageZU,
	}
}

func AllLanguagesInfo() []*Language {
	all := AllLanguages()
	languages := make([]*Language, 0, len(all))
	for _, v := range all {
		languages = append(languages, v.Info())
	}
	return languages
}

func LanguageCodeByName(name string) LanguageCode { //nolint:gocyclo
	switch textPrepare(name) {
	case "AA", "AFAR":
		return LanguageAA
	case "AB", "ABKHAZIAN":
		return LanguageAB
	case "AE", "AVESTAN":
		return LanguageAE
	case "AF", "AFRIKAANS":
		return LanguageAF
	case "AK", "AKAN":
		return LanguageAK
	case "AM", "AMHARIC":
		return LanguageAM
	case "AN", "ARAGONESE":
		return LanguageAN
	case "AR", "ARABIC":
		return LanguageAR
	case "AS", "ASSAMESE":
		return LanguageAS
	case "AV", "AVARIC":
		return LanguageAV
	case "AY", "AYMARA":
		return LanguageAY
	case "AZ", "AZERBAIJANI":
		return LanguageAZ
	case "BA", "BASHKIR":
		return LanguageBA
	case "BE", "BELARUSIAN":
		return LanguageBE
	case "BG", "BULGARIAN":
		return LanguageBG
	case "BH", "BIHARI":
		return LanguageBH
	case "BI", "BISLAMA":
		return LanguageBI
	case "BM", "BAMBARA":
		return LanguageBM
	case "BN", "BENGALI":
		return LanguageBN
	case "BO", "TIBETAN":
		return LanguageBO
	case "BR", "BRETON":
		return LanguageBR
	case "BS", "BOSNIAN":
		return LanguageBS
	case "CA", "CATALAN":
		return LanguageCA
	case "CE", "CHECHEN":
		return LanguageCE
	case "CH", "CHAMORRO":
		return LanguageCH
	case "CO", "CORSICAN":
		return LanguageCO
	case "CR", "CREE":
		return LanguageCR
	case "CS", "CZECH":
		return LanguageCS
	case "CU", "OLDCHURCHSLAVONIC", "CHURCHSLAVONIC", "CHURCHSLAVIC":
		return LanguageCU
	case "CV", "CHUVASH":
		return LanguageCV
	case "CY", "WELSH":
		return LanguageCY
	case "DA", "DANISH":
		return LanguageDA
	case "DE", "GERMAN":
		return LanguageDE
	case "DV", "DIVEHI":
		return LanguageDV
	case "DZ", "DZONGKHA":
		return LanguageDZ
	case "EE", "EWE":
		return LanguageEE
	case "EL", "GREEK":
		return LanguageEL
	case "EN", "ENGLISH":
		return LanguageEN
	case "EO", "ESPERANTO":
		return LanguageEO
	case "ES", "SPANISH":
		return LanguageES
	case "ET", "ESTONIAN":
		return LanguageET
	case "EU", "BASQUE":
		return LanguageEU
	case "FA", "PERSIAN":
		return LanguageFA
	case "FF", "FULAH":
		return LanguageFF
	case "FI", "FINNISH":
		return LanguageFI
	case "FJ", "FIJIAN":
		return LanguageFJ
	case "FO", "FAROESE":
		return LanguageFO
	case "FR", "FRENCH":
		return LanguageFR
	case "FY", "WESTERNFRISIAN":
		return LanguageFY
	case "GA", "IRISH":
		return LanguageGA
	case "GD", "SCOTTISHGAELIC":
		return LanguageGD
	case "GL", "GALICIAN":
		return LanguageGL
	case "GN", "GUARANI":
		return LanguageGN
	case "GU", "GUJARATI":
		return LanguageGU
	case "GV", "MANX":
		return LanguageGV
	case "HA", "HAUSA":
		return LanguageHA
	case "HE", "HEBREW":
		return LanguageHE
	case "HI", "HINDI":
		return LanguageHI
	case "HO", "HIRIMOTU":
		return LanguageHO
	case "HR", "CROATIAN":
		return LanguageHR
	case "HT", "HAITIAN":
		return LanguageHT
	case "HU", "HUNGARIAN":
		return LanguageHU
	case "HY", "ARMENIAN":
		return LanguageHY
	case "HZ", "HERERO":
		return LanguageHZ
	case "IA", "INTERLINGUA":
		return LanguageIA
	case "ID", "INDONESIAN":
		return LanguageID
	case "IE", "INTERLINGUE":
		return LanguageIE
	case "IG", "IGBO":
		return LanguageIG
	case "II", "SICHUANYI":
		return LanguageII
	case "IK", "INUPIAQ":
		return LanguageIK
	case "IO", "IDO":
		return LanguageIO
	case "IS", "ICELANDIC":
		return LanguageIS
	case "IT", "ITALIAN":
		return LanguageIT
	case "IU", "INUKTITUT":
		return LanguageIU
	case "JA", "JAPANESE":
		return LanguageJA
	case "JV", "JAVANESE":
		return LanguageJV
	case "KA", "GEORGIAN":
		return LanguageKA
	case "KG", "KONGO":
		return LanguageKG
	case "KI", "KIKUYU":
		return LanguageKI
	case "KJ", "KWANYAMA":
		return LanguageKJ
	case "KK", "KAZAKH":
		return LanguageKK
	case "KL", "KALAALLISUT":
		return LanguageKL
	case "KM", "KHMER":
		return LanguageKM
	case "KN", "KANNADA":
		return LanguageKN
	case "KO", "KOREAN":
		return LanguageKO
	case "KR", "KANURI":
		return LanguageKR
	case "KS", "KASHMIRI":
		return LanguageKS
	case "KU", "KURDISH":
		return LanguageKU
	case "KV", "KOMI":
		return LanguageKV
	case "KW", "CORNISH":
		return LanguageKW
	case "KY", "KYRGYZ":
		return LanguageKY
	case "LA", "LATIN":
		return LanguageLA
	case "LB", "LUXEMBOURGISH":
		return LanguageLB
	case "LG", "GANDA":
		return LanguageLG
	case "LI", "LIMBURGISH":
		return LanguageLI
	case "LN", "LINGALA":
		return LanguageLN
	case "LO", "LAO":
		return LanguageLO
	case "LT", "LITHUANIAN":
		return LanguageLT
	case "LU", "LUBAKATANGA":
		return LanguageLU
	case "LV", "LATVIAN":
		return LanguageLV
	case "MG", "MALAGASY":
		return LanguageMG
	case "MH", "MARSHALLESE":
		return LanguageMH
	case "MI", "MAORI":
		return LanguageMI
	case "MK", "MACEDONIAN":
		return LanguageMK
	case "ML", "MALAYALAM":
		return LanguageML
	case "MN", "MONGOLIAN":
		return LanguageMN
	case "MR", "MARATHI":
		return LanguageMR
	case "MS", "MALAY":
		return LanguageMS
	case "MT", "MALTESE":
		return LanguageMT
	case "MY", "BURMESE":
		return LanguageMY
	case "NA", "NAURU":
		return LanguageNA
	case "NB", "NORWEGIANBOKMAL":
		return LanguageNB
	case "ND", "NORTHNDEBELE":
		return LanguageND
	case "NE", "NEPALI":
		return LanguageNE
	case "NG", "NDONGA":
		return LanguageNG
	case "NL", "DUTCH":
		return LanguageNL
	case "NN", "NORWEGIANNYNORSK":
		return LanguageNN
	case "NO", "NORWEGIAN":
		return LanguageNO
	case "NR", "SOUTHNDEBELE":
		return LanguageNR
	case "NS", "NORTHERNSOTHO":
		return LanguageNS
	case "NV", "NAVAJO":
		return LanguageNV
	case "NY", "CHICHEWA":
		return LanguageNY
	case "OC", "OCCITAN":
		return LanguageOC
	case "OJ", "OJIBWA":
		return LanguageOJ
	case "OM", "OROMO":
		return LanguageOM
	case "OR", "ORIYA":
		return LanguageOR
	case "OS", "OSSETIAN":
		return LanguageOS
	case "PA", "PANJABI":
		return LanguagePA
	case "PI", "PALI":
		return LanguagePI
	case "PL", "POLISH":
		return LanguagePL
	case "PS", "PASHTO":
		return LanguagePS
	case "PT", "PORTUGUESE":
		return LanguagePT
	case "QU", "QUECHUA":
		return LanguageQU
	case "RM", "ROMANSH":
		return LanguageRM
	case "RN", "RUNDI":
		return LanguageRN
	case "RO", "ROMANIAN":
		return LanguageRO
	case "RU", "RUSSIAN":
		return LanguageRU
	case "RW", "KINYARWANDA":
		return LanguageRW
	case "SA", "SANSKRIT":
		return LanguageSA
	case "SC", "SARDINIAN":
		return LanguageSC
	case "SD", "SINDHI":
		return LanguageSD
	case "SE", "NORTHERNSAMI":
		return LanguageSE
	case "SG", "SANGO":
		return LanguageSG
	case "SI", "SINHALA":
		return LanguageSI
	case "SK", "SLOVAK":
		return LanguageSK
	case "SL", "SLOVENE":
		return LanguageSL
	case "SM", "SAMOAN":
		return LanguageSM
	case "SN", "SHONA":
		return LanguageSN
	case "SO", "SOMALI":
		return LanguageSO
	case "SQ", "ALBANIAN":
		return LanguageSQ
	case "SR", "SERBIAN":
		return LanguageSR
	case "SS", "SWATI":
		return LanguageSS
	case "ST", "SOUTHERNSOTHO":
		return LanguageST
	case "SU", "SUNDANESE":
		return LanguageSU
	case "SV", "SWEDISH":
		return LanguageSV
	case "SW", "SWAHILI":
		return LanguageSW
	case "TA", "TAMIL":
		return LanguageTA
	case "TE", "TELUGU":
		return LanguageTE
	case "TG", "TAJIK":
		return LanguageTG
	case "TH", "THAI":
		return LanguageTH
	case "TI", "TIGRINYA":
		return LanguageTI
	case "TK", "TURKMEN":
		return LanguageTK
	case "TL", "TAGALOG":
		return LanguageTL
	case "TN", "TSWANA":
		return LanguageTN
	case "TO", "TONGAN":
		return LanguageTO
	case "TR", "TURKISH":
		return LanguageTR
	case "TS", "TSONGA":
		return LanguageTS
	case "TT", "TATAR":
		return LanguageTT
	case "TW", "TWI":
		return LanguageTW
	case "TVL", "TUVALUAN":
		return LanguageTVL
	case "TY", "TAHITIAN":
		return LanguageTY
	case "UG", "UIGHUR":
		return LanguageUG
	case "UK", "UKRAINIAN":
		return LanguageUK
	case "UR", "URDU":
		return LanguageUR
	case "UZ", "UZBEK":
		return LanguageUZ
	case "VE", "VENDA":
		return LanguageVE
	case "VI", "VIETNAMESE":
		return LanguageVI
	case "VO", "VOLAPUK":
		return LanguageVO
	case "WA", "WALLOON":
		return LanguageWA
	case "WO", "WOLOF":
		return LanguageWO
	case "XH", "XHOSA":
		return LanguageXH
	case "YI", "YIDDISH":
		return LanguageYI
	case "YO", "YORUBA":
		return LanguageYO
	case "ZA", "ZHUANG":
		return LanguageZA
	case "ZH", "CHINESE":
		return LanguageZH
	case "ZU", "ZULU":
		return LanguageZU
	default:
		return LanguageUnknown
	}
}
