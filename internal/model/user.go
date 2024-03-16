package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	
)

type User struct{
	ID primitive.ObjectID `bson:"_id"`
	Username string `bson:"username"`
	Email string `bson:"email"`
	Password string `bson:"password"`
	CretedAt time.Time `bson:"created_at"`
	Active bool `bson:"active"`
}

type Profile struct{
	ID primitive.ObjectID `bson:"_id"`
	UserID User `bson:"user_id"`
	InviteLink string `bson:"invite_link"`
}


func (u *User)  PasswordHash(password string)  error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return  err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) ValidatePassword(hash string, password string) error {
    if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
        return err
    }
    return nil
}

func (u *User) BaseUsername(email string) {
	var username string
	for _, i := range(email) {
		if i == '@'{
			break
		}
		username += string(i)
	}
	u.Username = username
}


