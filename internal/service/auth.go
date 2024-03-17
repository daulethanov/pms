package service

import (
	"context"
	"errors"
	"mime/multipart"
	"regexp"
	"time"
	"todo/internal/model"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceInterface interface {
	CreateUser(email string, password string) (string, error)
	LoginUser(email string, password string) (string, error)
	EmailValidate(email string) error
	ViewProfile(id string) (*model.User, error)
	EditImageProfile(userID string, imageFile multipart.File, size int64) error
}

type AuthService struct {
	userCollection *mongo.Collection
	minioClient    *minio.Client
}

func NewAuthService(userCollection *mongo.Collection, minioClient *minio.Client) AuthServiceInterface {
	return &AuthService{
		userCollection: userCollection,
		minioClient:    minioClient,
	}
}

func (a *AuthService) EmailValidate(email string) error {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	if re.MatchString(email) {
		return nil
	} else {
		return errors.New("неверный формат email")
	}
}

func (a *AuthService) CreateUser(email string, password string) (string, error) {
	existingUser := a.userCollection.FindOne(context.TODO(), bson.M{"email": email})
	if existingUser.Err() == nil {
		return "", errors.New("user with this email already exists")
	}

	user := model.User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		CretedAt: time.Now(),
		Active:   true,
	}
	user.BaseUsername(email)

	err := user.PasswordHash(password)
	if err != nil {
		return "", err
	}

	_, err = a.userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	return user.ID.Hex(), nil
}

func (a *AuthService) LoginUser(email string, password string) (string, error) {
	var user model.User
	err := a.userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("пользователь с таким email не найден")
		}
		return "", err
	}

	if err := user.ValidatePassword(user.Password, password); err != nil {
		return "", err
	}
	userID := user.ID.Hex()
	return userID, nil
}

func (a *AuthService) ViewProfile(id string) (*model.User, error) {
	var user model.User
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = a.userCollection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)

	return &user, nil
}

func (a *AuthService) EditImageProfile(userID string, imageFile multipart.File, size int64) error {
	ctx := context.Background()
	profileBucket := "image"
	userID_hex, err  := primitive.ObjectIDFromHex(userID)
	objectName := uuid.New().String()

	_, err = a.minioClient.PutObject(ctx, profileBucket, objectName, imageFile, size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return err
	}
    imageURL, err := a.minioClient.PresignedGetObject(ctx, profileBucket, objectName, time.Second*24*60*60, nil)

    _, err = a.userCollection.UpdateOne(ctx, bson.M{"_id": userID_hex}, bson.M{"$set": bson.M{"image": imageURL.String()}})
    if err != nil {
        return err
    }
	
	return nil
}
