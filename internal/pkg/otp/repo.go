package otp

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type otpRepository struct {
	tableName string
	Api       dynamodbiface.DynamoDBAPI
}

type Otp struct {
	Email      string
	Code       string
	VerifiedAt int64
	ExpiresIn  int64
}

type OtpRepository interface {
	Save(entity Otp) error
	GetByEmail(email string) (*Otp, error)
}

func NewOtpRepository() OtpRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	api := dynamodb.New(sess)

	return &otpRepository{
		tableName: "Otp",
		Api:       api,
	}
}

func (repo *otpRepository) GetByEmail(email string) (*Otp, error) {
	pk, _ := dynamodbattribute.MarshalMap(map[string]string{
		"Email": email,
	})
	input := &dynamodb.GetItemInput{
		Key:       pk,
		TableName: &repo.tableName,
	}

	res, err := repo.Api.GetItem(input)
	if err != nil {
		return nil, err
	}

	var otp Otp
	err = dynamodbattribute.UnmarshalMap(res.Item, &otp)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func (repo *otpRepository) Save(entity Otp) error {
	ttl := time.Now().Add(2 * time.Minute).Unix()

	if entity.ExpiresIn > 0 {
		ttl = entity.ExpiresIn
	}

	item, err := dynamodbattribute.MarshalMap(&Otp{
		Email:      entity.Email,
		Code:       entity.Code,
		ExpiresIn:  ttl,
		VerifiedAt: entity.VerifiedAt,
	})
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: &repo.tableName,
		Item:      item,
	}

	_, err = repo.Api.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
