package otp

import (
	"math/rand"
)

type OtpService struct {
	repo interface{}
}

func NewOtpService(repo interface{}) *OtpService {
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

func (otps *OtpService) Execute() string {
	return otps.generateOtp()
}
