package models

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func createTable() dynamo.Table {
	session, _ := session.NewSession()
	db := dynamo.New(session, &aws.Config{Region: aws.String("ca-central-1")})
	table := db.Table("one_time_links")
	return table
}

func GetMessage(id string) (EncryptedMessage, error) {
	t := createTable()
	var m EncryptedMessage
	err := t.Get("partition_key", id).Range("sort_key", dynamo.Equal, "unused").One(&m)

	// TODO: Decrypt
	if err != nil {
		return EncryptedMessage{}, err
	}
	return m, nil
}

func init() {
	if len(os.Getenv("AWS_ACCESS_KEY_ID")) < 1 {
		panic("AWS_ACCESS_KEY is required")
	}
	if len(os.Getenv("AWS_SECRET_ACCESS_KEY")) < 1 {
		panic("AWS_SECRET_ACCESS_KEY is required")
	}
}
