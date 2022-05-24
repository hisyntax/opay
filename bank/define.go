package bank

// body := []byte("{\n  \"amount\": 10,\n  \"reference\": \"reference12934\",\n  \"narration\": \"911 Transaction\",\n  \"bankCode\": \"058\",\n  \"accountNumber\": \"0111946768\",\n  \"currency\": \"NGN\",\n  \"walletId\": \"4794983C91374AD6B3ECD76F2BEA296D\"\n}")

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

// {
// 	"reference": "test_20191123132233",
// 	"mchShortName": "Jerry's shop",
// 	"productName": "Apple AirPods Pro",
// 	"productDesc": "The best wireless earphone in history",
// 	"userPhone": "+2349876543210",
// 	"userRequestIp": "123.123.123.123",
// 	"amount": "100",
// 	"currency": "NGN",
// 	"payTypes": ["BalancePayment", "BonusPayment", "OWealth"],
// 	"payMethods": ["account", "qrcode", "bankCard", "bankAccount", "bankTransfer", "bankUSSD"],
// 	"callbackUrl": "https://you.domain.com/callbackUrl",
// 	"returnUrl": "https://you.domain.com/returnUrl",
// 	"expireAt": "10"
//   }
