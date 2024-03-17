package minio

import (
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"log"
)

type minioConfig struct {
	endpoint  string
	accessKey string
	secretKey string
}

func NewMinioClient() (*minio.Client, error) {

	config := minioConfig{
		endpoint:  viper.GetString("minio.endpoint"),
		accessKey: viper.GetString("minio.access-key"),
		secretKey: viper.GetString("minio.secret-key"),
	}
	minioClient, err := minio.New(config.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.accessKey, config.secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient, nil
}
