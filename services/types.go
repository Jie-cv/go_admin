package services

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResonse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
