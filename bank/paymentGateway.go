package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InitializePaymentGateway(baseUrl, merchantId, publicKey, reference, mchShortName, productName, productDesc, userPhone, userRequestIp, amount, callbackUrl, returnUrl, expireAt string) (*CheckOutResponse, int, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/cashier/initialize", baseUrl)
	method := "POST"

	payload := UiPaymentGateway{}
	payload.Reference = reference
	payload.MchShortName = mchShortName
	payload.ProductName = productName
	payload.ProductDesc = productDesc
	payload.UserPhone = userPhone
	payload.UserRequestIp = userRequestIp
	payload.Amount = amount
	payload.Currency = "NGN"
	// payload.PayTypes = "BalancePayment"
	// payload.PayMethods = "bankCard"
	payload.PayTypes = []string{"BalancePayment", "BonusPayment", "OWealth"}
	fmt.Println(payload.PayTypes)
	payload.PayMethods = []string{"account", "qrcode", "bankCard", "bankAccount", "bankTransfer", "bankUSSD"}
	payload.CallbackUrl = callbackUrl
	payload.ReturnUrl = returnUrl
	payload.ExpireAt = expireAt

	jsonReq, jsonErr := json.Marshal(&payload)
	if jsonErr != nil {
		fmt.Println(jsonErr)
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
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp_body))

	status := resp.StatusCode
	var response CheckOutResponse
	json.Unmarshal(resp_body, &response)

	return &response, status, nil
}
