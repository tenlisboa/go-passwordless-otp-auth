package otp

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Mailer struct {
	mailer *ses.SES
}

func NewMailer() *Mailer {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	if err != nil {
		fmt.Printf("Error on intanciating SES: %s\n", err.Error())
		return nil
	}

	svc := ses.New(sess)

	return &Mailer{
		mailer: svc,
	}
}

type SendEmailInput struct {
	Email   string
	Body    string
	Subject string
}

func (mailer *Mailer) SendEmail(input *SendEmailInput) error {
	emailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(input.Email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(input.Body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(input.Subject),
			},
		},
		Source: aws.String("teydadeydo@gufum.com"),
	}

	_, err := mailer.mailer.SendEmail(emailInput)
	if err != nil {
		return err
	}

	return nil
}
