package bank

type UiPaymentGateway struct {
	Reference     string   `json:"reference"`
	MchShortName  string   `json:"mchShortName"`
	ProductName   string   `json:"productName"`
	ProductDesc   string   `json:"productDesc"`
	UserPhone     string   `json:"userPhone"`
	UserRequestIp string   `json:"userRequestIp"`
	Amount        string   `json:"amount"`
	Currency      string   `json:"currency"`
	PayTypes      []string `json:"payTypes"`
	PayMethods    []string `json:"payMethods"`
	CallbackUrl   string   `json:"callbackUrl"`
	ReturnUrl     string   `json:"returnUrl"`
	ExpireAt      string   `json:"expireAt"`
}

type CheckOutResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    CheckOutDataBody `json:"data"`
}

type CheckOutDataBody struct {
	Reference  string `json:"reference"`
	OrderNo    string `json:"orderNo"`
	CashierUrl string `json:"cashierUrl"`
	Amount     int    `json:"amount"`
	Currency   string `json:"currency"`
	Status     string `json:"status"`
}

//for the bank list in a country e.g Ng = banks in nigeria
type Country struct {
	CountryCode string `json:"countryCode"`
}

type BankListResponse struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Data    []BankListDataBody `json:"data"`
}

type BankListDataBody struct {
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	Type                 string `json:"type"`
	Icon                 string `json:"icon"`
	Color                string `json:"color"`
	IsSupportBankAccount bool   `json:"isSupportBankAccount"`
}

//bank transfer
type TransferFund struct {
	Amount    string       `json:"amount"`
	Country   string       `json:"country"`
	Currency  string       `json:"currency"`
	Reason    string       `json:"reason"`
	Receiver  ReceiverBody `json:"receiver"`
	Reference string       `json:"reference"`
}

type ReceiverBody struct {
	BankAccountNumber string `json:"bankAccountNumber"`
	BankCode          string `json:"bankCode"`
	Name              string `json:"name"`
}

type TransferFundResponse struct {
	Reference         string `json:"reference"`
	OrderNo           string `json:"orderNo"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	Fee               string `json:"fee"`
	Status            string `json:"status"`
	FailureReason     string `json:"failureReason"`
	BankCode          string `json:"bankCode"`
	BankAccountNumber string `json:"bankAccountNumber"`
}
