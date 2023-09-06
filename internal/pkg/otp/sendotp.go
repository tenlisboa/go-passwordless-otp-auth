package otp

import (
	"fmt"
	"math/rand"
)

type OtpService struct {
	repo   OtpRepository
	mailer *Mailer
}

func NewOtpService(repo OtpRepository, mailer *Mailer) *OtpService {
	return &OtpService{
		repo:   repo,
		mailer: mailer,
	}
}

func (otps *OtpService) generateOtp() string {
	bytes := "abcdefghijklmnopqrstuvxz0123456789"
	size := 10
	code := make([]byte, size)
	for i := range code {
		code[i] = bytes[rand.Intn(len(bytes))]
	}

	return string(code)
}

func (otps *OtpService) Execute(email string) error {
	code := otps.generateOtp()

	otp := &Otp{
		Email: email,
		Code:  code,
	}

	err := otps.repo.Save(*otp)
	if err != nil {
		return err
	}

	err = otps.mailer.SendEmail(&SendEmailInput{
		Subject: "Seu código de autentição.",
		Body:    fmt.Sprintf("Não compartilhe este código com ninguém: %s\n", code),
		Email:   email,
	})

	if err != nil {
		return err
	}

	return nil
}
