package monnify

type MonnifyOptions struct {
	ApiKey       string
	SecretKey    string
	IsProduction bool
}

func New(options *MonnifyOptions) string {
	baseUrl := GetBaseUrl(options.IsProduction)
	cache := NewCache()
	_ = NewToken(cache, baseUrl, options.ApiKey+":"+options.SecretKey)

	return "Hello"
}
