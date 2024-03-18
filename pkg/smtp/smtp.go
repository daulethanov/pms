package smtp

// import (
// 	"crypto/tls"
// 	"github.com/spf13/viper"
// 	"gopkg.in/gomail.v2"
// )

// type smtpConfig struct {
// 	smtp     string
// 	port     int
// 	email    string
// 	password string
// }

// func NewSMTPClient() {
// 	config := smtpConfig{
// 		smtp:     viper.GetString("smtp.server"),
// 		port:     viper.GetInt("smtp.port"),
// 		email:    viper.GetString("stmp.email"),
// 		password: viper.GetString("smtp.password"),
// 	}
// 	d := gomail.NewDialer(config.smtp, config.port, config.email, config.password)
// 	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
// }
