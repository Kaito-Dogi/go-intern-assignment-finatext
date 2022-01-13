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
	// URLと処理を紐付ける．
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/address", addressHandler)

	// ポート番号を8080としてAPIサーバを立ち上げる．
	fmt.Println("Server is running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<small>&copy; 2022 Kaito-Dogi</small>")
}

func addressHandler(w http.ResponseWriter, r *http.Request) {
	//クエリパラメータ取得する．
	zn := r.URL.Query().Get("postal_code")

	// URLを作成する．
	url := "http://zip.cgis.biz/xml/zip.php?zn=" + zn

	// XML形式のレスポンスを受け取る．
	xmlResponse := getAddress(url)

	// XML形式のレスポンスをXMLの構造体に変換する．
	xmlAddress := new(types.XmlAddress)
	if err := xml.Unmarshal(xmlResponse, xmlAddress); err != nil {
		log.Fatal(err)
	}

	// 無効な郵便番号だった場合，ステータスコード400を返す．
	if xmlAddress.Result.ResultZipNum == "" {
		http.Error(w, "400 Bad Request", 400)
		return
	}

	// XMLの構造体のフィールドを元に，JSONの構造体を生成する．
	jsonAddress := types.JsonAddress{
		PostalCode:  xmlAddress.Result.ResultZipNum,
		Address:     xmlAddress.ADDRESSValue.Value.State + xmlAddress.ADDRESSValue.Value.City + xmlAddress.ADDRESSValue.Value.Address,
		AddressKana: xmlAddress.ADDRESSValue.Value.StateKana + xmlAddress.ADDRESSValue.Value.CityKana + xmlAddress.ADDRESSValue.Value.AddressKana,
	}

	// JSONの構造体をJSON形式に変換する．
	var jsonResponse bytes.Buffer
	enc := json.NewEncoder(&jsonResponse)
	if err := enc.Encode(&jsonAddress); err != nil {
		log.Fatal(err)
	}

	// JSON形式のResponseを返す．
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResponse.String())
}

// 引数にURLを受け取り，XML形式のレスポンスを返す．
func getAddress(url string) []byte {
	// URLを受け取り，外部APIを叩く．
	xmlResponse, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// ResponseのBodyを読み込む．
	body, err := io.ReadAll(xmlResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	// ResponseのBodyを閉じる．
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Fatal(err)
		}
	}(xmlResponse.Body)

	// Bodyを返す．
	return body
}
