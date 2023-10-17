package otp

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

func (otps *OtpService) SendOtp(email string) error {
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

func (otps *OtpService) VerifyOtp(email, code string) error {
	otp, err := otps.repo.GetByEmail(email)
	if err != nil {
		return err
	}

	// TODO: Verifica se o OTP enviado é o mesmo - OK
	if otp.Code != code {
		return errors.New("código otp inválido")
	}

	// TODO: Se o OTP já foi verificado - OK
	if otp.VerifiedAt != 0 {
		return errors.New("código otp já foi verificado")
	}
	// TODO: Se o OTP não está expirado - OK
	now := time.Now().Unix()
	if otp.ExpiresIn-now < 0 {
		return errors.New("código expirado")
	}
	// TODO: Verifica OTP - OK
	otp.VerifiedAt = now
	err = otps.repo.Save(*otp)
	if err != nil {
		return err
	}

	return nil
}
