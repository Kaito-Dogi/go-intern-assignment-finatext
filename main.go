package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AddressXml struct {
	Result struct {
		ResultZipNum string `xml:"result_zip_num,attr"`
	} `xml:"result"`
	AddressValue struct {
		Value struct {
			State       string `xml:"state,attr"`
			City        string `xml:"city,attr"`
			Address     string `xml:"address,attr"`
			StateKana   string `xml:"state_kana,attr"`
			CityKana    string `xml:"city_kana,attr"`
			AddressKana string `xml:"address_kana,attr"`
		} `xml:"value"`
	} `xml:"ADDRESS_value"`
}

type AddressJson struct {
	PostalCode  string `json:"postal_code"`
	Address     string `json:"address"`
	AddressKana string `json:"address_kana"`
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/address", addressHandler)
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "<h1>I'm Kaito-Dogi!</h1>")
}

func addressHandler(writer http.ResponseWriter, request *http.Request) {
	//クエリパラメータ取得する．
	zn := request.URL.Query().Get("postal_code")

	// 外部APIを叩く．
	xmlResponse, err := http.Get("http://zip.cgis.biz/xml/zip.php?zn=" + zn)
	if err != nil {
		log.Fatal("Get Http Error: ", err)
	}

	body, err := io.ReadAll(xmlResponse.Body)
	if err != nil {
		log.Fatal("IO Read Error: ", err)
	}

	defer xmlResponse.Body.Close()

	addressXml := new(AddressXml)
	if err := xml.Unmarshal([]byte(body), addressXml); err != nil {
		log.Fatal("XML Unmarshal Error: ", err)
	}

	addressJson := AddressJson{
		PostalCode:  addressXml.Result.ResultZipNum,
		Address:     addressXml.AddressValue.Value.State + addressXml.AddressValue.Value.City + addressXml.AddressValue.Value.Address,
		AddressKana: addressXml.AddressValue.Value.StateKana + addressXml.AddressValue.Value.CityKana + addressXml.AddressValue.Value.AddressKana,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&addressJson); err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(writer, buf.String())

}
