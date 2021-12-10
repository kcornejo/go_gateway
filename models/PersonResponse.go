package models

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Xsi     string   `xml:"xsi,attr"`
	S       string   `xml:"s,attr"`
	Link    struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		ID   string `xml:"id,attr"`
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Body struct {
		Text              string `xml:",chardata"`
		GetByNameResponse struct {
			Text            string `xml:",chardata"`
			Xmlns           string `xml:"xmlns,attr"`
			GetByNameResult struct {
				Text   string `xml:",chardata"`
				Schema struct {
					Text                 string `xml:",chardata"`
					Xmlns                string `xml:"xmlns,attr"`
					S                    string `xml:"s,attr"`
					Msdata               string `xml:"msdata,attr"`
					ID                   string `xml:"id,attr"`
					AttributeFormDefault string `xml:"attributeFormDefault,attr"`
					ElementFormDefault   string `xml:"elementFormDefault,attr"`
				} `xml:"schema"`
				Diffgram struct {
					Text       string `xml:",chardata"`
					Msdata     string `xml:"msdata,attr"`
					Diffgr     string `xml:"diffgr,attr"`
					ListByName struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
						SQL   []struct {
							Text     string `xml:",chardata"`
							AttrID   string `xml:"id,attr"`
							RowOrder string `xml:"rowOrder,attr"`
							ID       string `xml:"ID"`
							Name     string `xml:"Name"`
							DOB      string `xml:"DOB"`
							SSN      string `xml:"SSN"`
						} `xml:"SQL"`
					} `xml:"ListByName"`
				} `xml:"diffgram"`
			} `xml:"GetByNameResult"`
		} `xml:"GetByNameResponse"`
	} `xml:"Body"`
}
