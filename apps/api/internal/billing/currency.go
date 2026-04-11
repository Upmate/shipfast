package billing

import "fmt"

type Currency struct {
	Code     string `json:"code"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

var Currencies = map[string]Currency{
	"USD": {Code: "USD", Symbol: "$", Decimals: 2},
	"EUR": {Code: "EUR", Symbol: "€", Decimals: 2},
	"GBP": {Code: "GBP", Symbol: "£", Decimals: 2},
}

func FormatAmount(amount int64, currencyCode string) string {
	cur, ok := Currencies[currencyCode]
	if !ok {
		cur = Currencies["USD"]
	}
	return fmt.Sprintf("%s%.2f", cur.Symbol, float64(amount)/100)
}
