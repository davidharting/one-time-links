package models

import (
	"fmt"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type EncryptedMessage struct {
	Id     string `dynamo:"partition_key"`
	Body   string `dynamo:"body"`
	Ignore string `dynamo:"sort_key"`
}

type EncryptResult struct {
	Message  EncryptedMessage
	Password string
}

func (result EncryptResult) link() string {
	return fmt.Sprintf("/message/?id=%v&password=%v", result.Message.Id, result.Password)
}

const PASSWORD string = "hunter2"

func Encrypt(message string) (EncryptResult, error) {
	encrypted, err := helper.EncryptMessageWithPassword([]byte(PASSWORD), message)

	if err != nil {
		return EncryptResult{}, err
	}

	id, err := gonanoid.New()
	if err != nil {
		return EncryptResult{}, err
	}

	return EncryptResult{
		Password: PASSWORD,
		Message:  EncryptedMessage{Id: id, Body: encrypted, Ignore: "unused"},
	}, nil
}

func Decrypt(encrypted EncryptedMessage, password string) (string, error) {
	plaintext, err := helper.DecryptMessageWithPassword([]byte(password), encrypted.Body)
	if err != nil {
		return "", err
	}
	return plaintext, nil
}
