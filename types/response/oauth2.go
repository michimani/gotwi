package response

type Token struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}
