package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tenlisboa/passwordless/internal/pkg/otp"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO: Extrair email do body da requisicao
	// TODO: gerar codigo OTP
	// TODO: Salvar codigo em um cache
	// TODO: Enviar um email para o usuário

	svc := otp.NewOtpService(1)
	code := svc.Execute()

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       code,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
