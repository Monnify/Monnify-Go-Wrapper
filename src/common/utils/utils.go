package utils

func GetBaseUrl(isProduction bool) string {
	if isProduction {
		return "https://api.monnify.com"
	}

	return "https://sandbox.monnify.com"
}
