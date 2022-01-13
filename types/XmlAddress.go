package types

type XmlAddress struct {
	Result struct {
		ResultZipNum string `xml:"result_zip_num,attr"`
	} `xml:"result"`
	ADDRESSValue struct {
		Value struct {
			StateKana   string `xml:"state_kana,attr"`
			CityKana    string `xml:"city_kana,attr"`
			AddressKana string `xml:"address_kana,attr"`
			State       string `xml:"state,attr"`
			City        string `xml:"city,attr"`
			Address     string `xml:"address,attr"`
		} `xml:"value"`
	} `xml:"ADDRESS_value"`
}
