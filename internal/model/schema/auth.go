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

type EditPasswordSchema struct {
	Email    string `json:"email"`
}

type EditPasswordCodeSchema struct {
	Code    int `json:"code"`
}


type EditPasswordConfirmSchema struct {
	Password string `json:"password"`
}

