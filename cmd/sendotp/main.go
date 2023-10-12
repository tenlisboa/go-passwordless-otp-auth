package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tenlisboa/passwordless/internal/pkg/otp"
)

type Body struct {
	Email string `json:"email"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO: Extrair email do body da requisicao - OK
	var body Body
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 422,
			Body:       "Requisição inválida",
		}, err
	}

	// TODO: gerar codigo OTP - OK
	// TODO: Salvar codigo em um cache - OK
	// TODO: Enviar um email para o usuário - OK
	repo := otp.NewOtpRepository()
	mailer := otp.NewMailer()
	svc := otp.NewOtpService(repo, mailer)
	err = svc.SendOtp(body.Email)
	if err != nil {
		fmt.Printf("Error in otp service execution: %s", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal error",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Otp enviado com sucesso",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
