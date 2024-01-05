package apiapp

type ApiAppConfiguration struct {
	token string
}

func NewApiService(token string) *ApiAppConfiguration {
	return &ApiAppConfiguration{
		token: token,
	}
}