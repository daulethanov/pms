package schema


type CreateUserSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewTokenSchema struct{
	RefreshToken string `json:"refresh_token"`
}
