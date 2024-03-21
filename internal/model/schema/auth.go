package schema

import "errors"

type CreateUserSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *CreateUserSchema) Validate() error {
    if e.Email == "" {
        return errors.New("поле Email не заполнено")
    }
    if e.Password == "" {
        return errors.New("поле Password не заполнено")
    }
    return nil
}


type LoginUserSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *LoginUserSchema) Validate() error {
    if e.Email == "" {
        return errors.New("поле Email не заполнено")
    }
    if e.Password == "" {
        return errors.New("поле Password не заполнено")
    }
    return nil
}



type NewTokenSchema struct{
	RefreshToken string `json:"refresh_token"`
}

func(e *NewTokenSchema) Validate() error{
	if e.RefreshToken== "" {
		return errors.New("поле Token не заполнено")
    }
	return nil

}


type EditPasswordSchema struct {
	Email    string `json:"email"`
}

func(e *EditPasswordSchema) Validate() error{
	if e.Email == "" {
		return errors.New("поле Email не заполнено")
    }
	return nil

}



type EditPasswordCodeSchema struct {
	Code    int `json:"code"`
}


func(e *EditPasswordCodeSchema) Validate() error{
	if e.Code == 0 {
		return errors.New("поле Code не заполнено")
    }
	return nil
}


type EditPasswordConfirmSchema struct {
	Password string `json:"password"`
}


func(e *EditPasswordConfirmSchema) Validate() error{
	if e.Password == "" {
		return errors.New("поле Password не заполнено")
    }
	return nil
}
