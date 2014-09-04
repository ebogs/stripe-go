package stripe

import "encoding/json"

// Account is the resource representing youe Stripe account.
// For more details see https://stripe.com/docs/api/#account.
type Account struct {
	Id            string `json:"id"`
	ChargeEnabled bool   `json:"charge_enabled"`
	Country       string `json:"country"`
	// Currencies is the list of supported currencies.
	Currencies       []string `json:"currencies_supported"`
	DefaultCurrency  string   `json:"default_currency"`
	DetailsSubmitted bool     `json:"details_submitted"`
	TransferEnabled  bool     `json:"transfer_enabled"`
	Name             string   `json:"display_name"`
	Email            string   `json:"email"`
	Statement        string   `json:"statement_description"`
	Timezone         string   `json:"timezone"`
}

// accountClient is the client used to invoke /account APIs.
type accountClient struct {
	api   Api
	token string
}

// Get returns the details of your account.
// For more details see https://stripe.com/docs/api/#retrieve_account.
func (c *accountClient) Get() (*Account, error) {
	account := &Account{}
	err := c.api.Call("GET", "/account", c.token, nil, account)

	return account, err
}

func (a *Account) UnmarshalJSON(data []byte) error {
	type account Account
	var aa account
	err := json.Unmarshal(data, &aa)
	if err == nil {
		*a = Account(aa)
	} else {
		// the id is surrounded by escaped \, so ignore those
		a.Id = string(data[1 : len(data)-1])
	}

	return nil
}
