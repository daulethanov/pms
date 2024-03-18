package mail

import (
	"crypto/tls"
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type smtpConfig struct {
	smtp     string
	port     int
	email    string
	password string
	username string
}

func SendMessageEditPassword(to string, code int) error {
	config := smtpConfig{
		smtp:     viper.GetString("smtp.smtp"),
		port:     viper.GetInt("smtp.port"),
		email:    viper.GetString("smtp.email"),
		password: viper.GetString("smtp.password"),
		username: viper.GetString("smtp.username"),
	}

	m := gomail.NewMessage()
	from := fmt.Sprintf("%s <%s>", config.username, config.email)
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Код для изменения пароля")
	body := fmt.Sprintf("Код для изменения пароля: %d", code)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(config.smtp, config.port, config.email, config.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Используйте только в тестовых целях!

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}