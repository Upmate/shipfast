package billing

type TaxRate struct {
	Country   string  `json:"country"`
	Rate      float64 `json:"rate"`
	Name      string  `json:"name"`
	Inclusive bool    `json:"inclusive"`
}

var TaxRates = map[string]TaxRate{
	"US": {Country: "US", Rate: 0.0, Name: "No federal tax", Inclusive: false},
	"GB": {Country: "GB", Rate: 20.0, Name: "VAT", Inclusive: true},
	"DE": {Country: "DE", Rate: 19.0, Name: "MwSt", Inclusive: true},
	"FR": {Country: "FR", Rate: 20.0, Name: "TVA", Inclusive: true},
	"GR": {Country: "GR", Rate: 24.0, Name: "FPA", Inclusive: true},
}

func CalculateTax(amount int64, country string) int64 {
	rate, ok := TaxRates[country]
	if !ok {
		return 0
	}
	return int64(float64(amount) * rate.Rate / 100)
}
