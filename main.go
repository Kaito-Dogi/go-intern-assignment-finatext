package main

import (
	"TaskFinatext/types"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

	// responseのbodyを読み込む．
	body, err := io.ReadAll(xmlResponse.Body)
	if err != nil {
		log.Fatal("IO Read Error: ", err)
	}

	// responseのリソースを解放する（調べる）．
	defer xmlResponse.Body.Close()

	// xmlを構造体にする．
	addressXml := new(types.AddressXMl)
	if err := xml.Unmarshal([]byte(body), addressXml); err != nil {
		log.Fatal("XML Unmarshal Error: ", err)
	}

	// jsonの構造体を生成する．
	addressJson := types.AddressJSON{
		PostalCode:  addressXml.Result.ResultZipNum,
		Address:     addressXml.ADDRESSValue.Value.State + addressXml.ADDRESSValue.Value.City + addressXml.ADDRESSValue.Value.Address,
		AddressKana: addressXml.ADDRESSValue.Value.StateKana + addressXml.ADDRESSValue.Value.CityKana + addressXml.ADDRESSValue.Value.AddressKana,
	}

	// 構造体をjsonに変換する．
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&addressJson); err != nil {
		log.Fatal(err)
	}

	// responseを出力する．
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(writer, buf.String())

}

func getAddress(url string) {

}
