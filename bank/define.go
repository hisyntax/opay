package bank

type PayWithBankCard struct {
	Reference         string `json:"reference"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	Country           string `json:"country"`
	PayType           string `json:"payType"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	CustomerEmail     string `json:"customerEmail"`
	CardNumber        string `json:"cardNumber"`
	CardDateMonth     string `json:"cardDateMonth"`
	CardDateYear      string `json:"cardDateYear"`
	CardCVC           string `json:"cardCVC"`
	Return3dsUrl      string `json:"return3dsUrl"`
	BankAccountNumber string `json:"bankAccountNumber"`
	BankCode          string `json:"bankCode"`
	Reason            string `json:"reason"`
	CallbackUrl       string `json:"callbackUrl"`
	ExpireAt          string `json:"expireAt"`
	BillingZip        string `json:"billingZip"`
	BillingCity       string `json:"billingCity"`
	BillingAddress    string `json:"billingAddress"`
	BillingState      string `json:"billingState"`
	BillingCountry    string `json:"billingCountry"`
}

type UiPaymentGateway struct {
	Country     string   `json:"country"`
	Reference   string   `json:"reference"`
	Amount      Amount   `json:"amount"`
	ReturnUrl   string   `json:"returnUrl"`
	CallbackUrl string   `json:"callbackUrl"`
	CancelUrl   string   `json:"cancelUrl"`
	ExpireAt    int      `json:"expireAt"`
	UserInfo    UserInfo `json:"userInfo"`
	Product     Product  `json:"product"`
	PayMethod   string   `json:"payMethod"`
}

type Amount struct {
	Total    int    `json:"total"`
	Currency string `json:"currency"`
}
type UserInfo struct {
	UserEmail  string `json:"userEmail"`
	UserId     string `json:"userId"`
	UserMobile string `json:"userMobile"`
	UserName   string `json:"userName"`
}

type Product struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}
