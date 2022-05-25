package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InitializePaymentGateway(baseUrl, userEmail, userID, userName, merchantId, publicKey, ref, productName, productDesc, userPhone, callBackUrl, returnUrl, cancelUrl string, amount, expireAt int) (*CheckOutResponse, int, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/api/v1/international/cashier/create", baseUrl)
	method := "POST"

	payload := UiPaymentGateway{}

	// request payload
	payload.Country = "NG" // by default country is set
	payload.Reference = ref
	payload.Amount.Total = amount
	payload.Amount.Currency = "NGN"
	payload.ReturnUrl = returnUrl
	payload.CallbackUrl = callBackUrl
	payload.CancelUrl = cancelUrl
	payload.ExpireAt = expireAt // 1 = 1 minute while 10 = 10 minutes
	payload.UserInfo.UserEmail = userEmail
	payload.UserInfo.UserId = userID //this is an optional field as it can be an empty string like so ""
	payload.UserInfo.UserMobile = userPhone
	payload.UserInfo.UserName = userName
	payload.Product.Name = productName
	payload.Product.Description = productDesc
	payload.PayMethod = "" //this is an optional field as it would return all the payment options to the user if not set
	///////////////////////////////////////

	jsonReq, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewReader(jsonReq))
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("MerchantId", merchantId)
	token := fmt.Sprintf("Bearer %s", publicKey)
	req.Header.Add("Authorization", token)

	resp, respErr := client.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	status := resp.StatusCode
	fmt.Println(string(resp_body))

	var response CheckOutResponse
	json.Unmarshal(resp_body, &response)
	fmt.Printf("%+v\n", response)

	return &response, status, nil
}

func GetTransactionStatus(baseUrl, merchantId, signature, ref string) (*TransactionStatusResponse, int, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/api/v1/international/cashier/status", baseUrl)
	method := "POST"

	payload := TransactionStatusRequest{}
	payload.Country = "NG"
	payload.Reference = ref
	jsonReq, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewReader(jsonReq))
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

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	status := resp.StatusCode
	fmt.Println(string(resp_body))

	var response TransactionStatusResponse
	json.Unmarshal(resp_body, &response)

	return &response, status, nil
}
