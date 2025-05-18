package gofair

// Accounts API Operations
const (
	getAccountFunds     = "getAccountFunds/"
	getAccountStatement = "getAccountStatement/"
)

// Account object
type Account struct {
	Client *Client
}

func (a *Account) GetAccountFunds() (AccountFundsResponse, error) {
	// create url
	url := createURL(Endpoints.Account, getAccountFunds)

	// build request
	params := struct {
		Wallet string
	}{
		Wallet: "UK",
	}

	var response AccountFundsResponse

	// make request
	err := a.Client.request(url, params, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (a *Account) getAccountStatement(filter AccountStatementFilter) (AccountStatementReport, error) {
	// create url
	url := createURL(Endpoints.Account, getAccountStatement)

	// build request
	params := struct {
		Locale        string
		FromRecord    int
		RecordCount   int
		ItemDateRange TimeRange
		Wallet        string
	}{
		filter.Locale,
		filter.FromRecord,
		filter.RecordCount,
		filter.ItemDateRange,
		"UK",
	}

	var response AccountStatementReport

	// make request
	err := a.Client.request(url, params, &response)
	if err != nil {
		return response, err
	}
	return response, err
}
