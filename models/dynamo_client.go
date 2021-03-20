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

func GetMessage(id string, password string) (plaintext string, err error) {
	t := createTable()
	var encrypted EncryptedMessage
	err = t.Get("partition_key", id).Range("sort_key", dynamo.Equal, "unused").One(&encrypted)
	if err != nil {
		return "", err
	}
	plaintext, err = Decrypt(encrypted, password)
	if err != nil {
		return "", err
	}
	return plaintext, nil
}

func EncryptAndSave(message string) (link string, err error) {
	result, err := Encrypt(message)
	if err != nil {
		return "", err
	}
	err = save(result.Message)
	if err != nil {
		return "", err
	}
	return result.link(), nil
}

func save(m EncryptedMessage) error {
	t := createTable()
	err := t.Put(m).Run()
	return err
}

func init() {
	if len(os.Getenv("AWS_ACCESS_KEY_ID")) < 1 {
		panic("AWS_ACCESS_KEY is required")
	}
	if len(os.Getenv("AWS_SECRET_ACCESS_KEY")) < 1 {
		panic("AWS_SECRET_ACCESS_KEY is required")
	}
}
