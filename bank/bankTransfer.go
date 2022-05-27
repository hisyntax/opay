package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FundDisburstMent(baseUrl, merchantId, signature, reference, amount, name, bankCode, bankAccountNumber, reason string) (*TransferFundResponse, int, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/transfer/toBank", baseUrl)
	method := "POST"

	payload := TransferFund{}
	payload.Reference = reference
	payload.Amount = amount
	payload.Currency = "NGN"
	payload.Country = "NG"
	payload.Receiver.Name = name
	payload.Receiver.BankCode = bankCode
	payload.Receiver.BankAccountNumber = bankAccountNumber
	payload.Reason = reason

	jsonReq, jsonErr := json.Marshal(&payload)
	if jsonErr != nil {
		return nil, 0, jsonErr
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("MerchantId", merchantId)
	token := fmt.Sprintf("Bearer %s", signature)
	req.Header.Add("Authorization", token)

	resp, respErr := client.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	resp_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp_body))

	var response TransferFundResponse
	status := resp.StatusCode
	json.Unmarshal(resp_body, &response)

	return &response, status, nil
}
