package demo

import (
	"encoding/xml"
	"strings"
)

type Reading struct {
	RType         string `xml:"r_type,attr"`
	ReadingString string `xml:",innerxml"`
}

type Meaning struct {
	MLang         string `xml:"m_lang,attr"`
	MeaningString string `xml:",innerxml"`
}

type DicRef struct {
	DrType    string `xml:"dr_type,attr"`
	RefNumber string `xml:",innerxml"`
}

type Kanji struct {
	Literal     string    `xml:"literal"`
	Grade       int       `xml:"misc>grade"`
	StrokeCount int       `xml:"misc>stroke_count"`
	Freq        int       `xml:"misc>freq"`
	JLPT        int       `xml:"misc>jlpt"`
	DicRefs     []DicRef  `xml:"dic_number>dic_ref"`
	Readings    []Reading `xml:"reading_meaning>rmgroup>reading"`
	Meanings    []Meaning `xml:"reading_meaning>rmgroup>meaning"`
	Nanori      []string  `xml:"nanori"`
}

var theXML string = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE kanjidic2 [
<!ELEMENT variant (#PCDATA)>
<!ATTLIST variant var_type CDATA #REQUIRED>
	<!-- 
	oneill - Japanese Names (O'Neill) - numeric
	-->

]>
<header>

<file_version>4</file_version>
<database_version>2011-536</database_version>
<date_of_creation>2012-06-19</date_of_creation>
</header>
<character>
<literal>本</literal>
<misc>
<grade>1</grade>
<stroke_count>5</stroke_count>
<variant var_type="jis208">52-81</variant>
<freq>10</freq>
<jlpt>4</jlpt>
</misc>
<dic_number>
<dic_ref dr_type="nelson_c">96</dic_ref>
</dic_number>
<query_code>
<q_code qc_type="skip">4-5-3</q_code>
</query_code>
<reading_meaning>
<rmgroup>
<reading r_type="ja_on">ホン</reading>
<reading r_type="ja_kun">もと</reading>
<meaning>book</meaning>
<meaning m_lang="fr">livre</meaning>
</rmgroup>
<nanori>まと</nanori>
</reading_meaning>
</character>
</kanjidic2>
`

func ParseKanjiDic2() (kanjiList []Kanji) {
	decoder := xml.NewDecoder(strings.NewReader(theXML))
	for {
		token, _ := decoder.Token()
		//fmt.Println(token)
		if token == nil {
			break
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			if startElement.Name.Local == "character" {
				var kanji Kanji
				decoder.DecodeElement(&kanji, &startElement)
				kanjiList = append(kanjiList, kanji)
			}
		}
	}
	return
}
