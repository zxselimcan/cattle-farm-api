package utils

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/google/uuid"
)

func SendVerificationMail(email string) error {

	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	from := os.Getenv("SMTP_MAIL")
	password := os.Getenv("SMTP_PASSWORD")

	subject := "Verification"
	auth := smtp.PlainAuth("", from, password, host)

	verificationKey := fmt.Sprintf(
		"%x%x",
		sha256.Sum256([]byte(email + "-" + uuid.New().String() + uuid.New().String())[:]),
		sha256.Sum256([]byte(uuid.New().String() + uuid.New().String() + "-" + email)[:]),
	)

	body := fmt.Sprintf(`
	<html>
		<a href="http://dev.easyserv.net/email-verification?token=%v" target="_blank">CLICK</a>
	</html>
	`, verificationKey)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	c, err := smtp.Dial(host + ":" + port)
	if err != nil {
		return err
	}

	c.StartTLS(tlsconfig)

	if err = c.Auth(auth); err != nil {
		fmt.Println(err)
		return err
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	if err = c.Rcpt(email); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(
		fmt.Sprintf("MIME-Version: %v\r\n", "1.0") +
			fmt.Sprintf("Content-type: %v\r\n", "text/html; charset=UTF-8") +
			fmt.Sprintf("From: %v\r\n", from) +
			fmt.Sprintf("To: %v\r\n", email) +
			fmt.Sprintf("Subject: %v\r\n", subject) +
			fmt.Sprintf("%v\r\n", body),
	))

	if err != nil {
		return err
	}

	statusCMD := Redis.Set(context.Background(),
		verificationKey,
		fmt.Sprintf("%v", email),
		time.Duration(time.Hour*24),
	)
	if statusCMD.Err() != nil {
		return statusCMD.Err()
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil

}
