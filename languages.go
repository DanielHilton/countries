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
		LanguageMT, LanguageMY, LanguageNA, LanguageNB, LanguageND, LanguageNE, LanguageNG, LanguageNL, LanguageNN, LanguageNO, LanguageNS, LanguageNV,
		LanguageNY, LanguageOC, LanguageOJ, LanguageOM, LanguageOR, LanguageOS, LanguagePA, LanguagePI, LanguagePL, LanguagePS, LanguagePT, LanguageQU,
		LanguageRM, LanguageRN, LanguageRO, LanguageRU, LanguageRW, LanguageSA, LanguageSC, LanguageSD, LanguageSE, LanguageSG, LanguageSI, LanguageSK,
		LanguageSL, LanguageSM, LanguageSN, LanguageSO, LanguageSQ, LanguageSR, LanguageSS, LanguageST, LanguageSU, LanguageSV, LanguageSW, LanguageTA,
		LanguageTE, LanguageTG, LanguageTH, LanguageTI, LanguageTK, LanguageTL, LanguageTN, LanguageTO, LanguageTR, LanguageTS, LanguageTT, LanguageTW,
		LanguageTY, LanguageUG, LanguageUK, LanguageUR, LanguageUZ, LanguageVE, LanguageVI, LanguageVO, LanguageWA, LanguageWO, LanguageXH, LanguageYI,
		LanguageYO, LanguageZA, LanguageZH, LanguageZU,
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
