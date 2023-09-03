package otp

import (
	"math/rand"
)

type OtpService struct {
	repo OtpRepository
}

func NewOtpService(repo OtpRepository) *OtpService {
	return &OtpService{
		repo: repo,
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

	return nil
}
