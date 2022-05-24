package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func BankCardTransaction(baseUrl, userEmail, userID, userName, merchantId, publicKey, ref, productName, productDesc, userPhone, callBackUrl, returnUrl, cancelUrl string, amount, expireAt int) {
	client := http.Client{}
	url := fmt.Sprintf("%s/", baseUrl)
	method := "POST"

	payload := UiPaymentGateway{}

	///////////////////////////////////////
	payload.Country = "NG"
	payload.Reference = ref
	payload.Amount.Total = amount
	payload.Amount.Currency = "NGN"
	payload.ReturnUrl = returnUrl
	payload.CallbackUrl = callBackUrl
	payload.CancelUrl = cancelUrl
	payload.ExpireAt = expireAt
	payload.UserInfo.UserEmail = userEmail
	payload.UserInfo.UserId = userID
	payload.UserInfo.UserMobile = userPhone
	payload.UserInfo.UserName = userName
	payload.Product.Name = productName
	payload.Product.Description = productDesc
	payload.PayMethod = "BankCard"

	jsonReq, _ := json.Marshal(payload)

	req, reqErr := http.NewRequest(method, url, bytes.NewReader(jsonReq))
	if reqErr != nil {
		fmt.Println(reqErr)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("MerchantId", merchantId)
	fmt.Println(merchantId)
	token := fmt.Sprintf("Bearer %s", publicKey)
	fmt.Println(token)
	req.Header.Add("Authorization", token)

	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Println(respErr)
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println("response body...")
	fmt.Println(resp.Body)
	fmt.Println("output...")
	fmt.Println(string(resp_body))
}

func BankCard(baseUrl, merchantId, publicKey, ref, mchShortName, productName, productDesc, userPhone, userRequestIp, amount, callBackUrl, returnUrl, expireAt string) {
	payload := UiPaymentGateway{}
	//add all the fields in the payload
	//
	// payload.Reference = ref
	// payload.MchShortName = mchShortName
	// payload.ProductName = productName
	// payload.ProductDesc = productDesc
	// payload.UserPhone = userPhone
	// payload.UserRequestIp = userRequestIp
	// payload.Amount = amount
	// payload.Currency = "NGN"
	// payload.CallbackUrl = callBackUrl
	// payload.ReturnUrl = returnUrl
	// payload.ExpireAt = expireAt
	// payload.PayTypes = []string{"BalancePayment", "BonusPayment", "OWealth"}
	// payload.PayMethods = []string{"account", "qrcode", "bankCard", "bankAccount", "bankTransfer", "bankUSSD"}
	fmt.Println(payload)

	jsonReq, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/api/v3/cashier/initialize", baseUrl)
	contentType := "application/json"
	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	resp.Header.Add("MerchantId", merchantId)
	fmt.Println(merchantId)
	token := fmt.Sprintf("Bearer %s", publicKey)
	fmt.Println(token)
	resp.Header.Add("Authorization", token)

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	fmt.Println()

	var response UiPaymentGateway
	json.Unmarshal(bodyBytes, &response)
	fmt.Printf("%+v\n", response)
}
