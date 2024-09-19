package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	VND = "VND"
)

var supportedCurrencies = []string{USD, EUR, CAD, VND}

func IsSupportedCurrency(currency string) bool {
	for _, c := range supportedCurrencies {
		if currency == c {
			return true
		}
	}
	return false
}
