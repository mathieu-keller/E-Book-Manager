package epub

type Ncx struct {
	Points []NavPoint `xml:"navMap>navPoint" json:"points"`
}

type NavPoint struct {
	Text    string     `xml:"navLabel>text" json:"text"`
	Content Content    `xml:"content" json:"content"`
	Points  []NavPoint `xml:"navPoint" json:"points"`
}

type Content struct {
	Src string `xml:"src,attr" json:"src"`
}
