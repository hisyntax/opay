package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BankList(baseUrl, merchantId, publicKey string) {
	client := http.Client{}
	url := fmt.Sprintf("%s/banks", baseUrl)
	method := "POST"

	payload := Country{}
	payload.CountryCode = "NG"

	jsonReq, jsonErr := json.Marshal(&payload)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if reqErr != nil {
		fmt.Println(reqErr)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("MerchantId", merchantId)
	token := fmt.Sprintf("Bearer %s", publicKey)
	req.Header.Add("Authorization", token)

	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Println(respErr)
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp_body))

}
