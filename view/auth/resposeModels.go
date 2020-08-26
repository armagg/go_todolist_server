package auth


type signUpResponse struct{
	Token string `json:"token"`
	Message string `json:"msg"`
}

type signInResponse struct{
	Token string
	message string
}