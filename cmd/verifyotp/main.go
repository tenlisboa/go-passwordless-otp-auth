package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tenlisboa/passwordless/internal/pkg/otp"
)

type Body struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body Body
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 422,
			Body:       "Requisição inválida",
		}, nil
	}

	// TODO: Instancia o repo - OK
	repo := otp.NewOtpRepository()
	svc := otp.NewOtpService(repo, nil)
	// TODO: Chama o serviço - OK
	err = svc.VerifyOtp(body.Email, body.Code)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Código validado com sucesso",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
