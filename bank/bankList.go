package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BankList(baseUrl, merchantId, publicKey string) (*BankListResponse, int, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/banks", baseUrl)
	method := "POST"

	payload := Country{}
	payload.CountryCode = "NG"

	jsonReq, jsonErr := json.Marshal(&payload)
	if jsonErr != nil {
		return nil, 0, jsonErr
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("MerchantID", merchantId)
	token := fmt.Sprintf("Bearer %s", publicKey)
	req.Header.Add("Authorization", token)

	resp, respErr := client.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp_body))
	status := resp.StatusCode
	var response BankListResponse
	json.Unmarshal(resp_body, &response)

	return &response, status, nil
}
