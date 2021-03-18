package models

import "github.com/ProtonMail/gopenpgp/v2/helper"

type EncryptedMessage struct {
	id   string
	body string
}

type EncryptResult struct {
	message  EncryptedMessage
	password string
}

const PASSWORD string = "hunter2"

func Encrypt(message string) (EncryptResult, error) {
	encrypted, err := helper.EncryptMessageWithPassword([]byte(PASSWORD), message)

	if err != nil {
		return EncryptResult{}, err
	}

	return EncryptResult{
		password: PASSWORD,
		message:  EncryptedMessage{id: "fake-uuid", body: encrypted},
	}, nil
}

func Decrypt(encrypted EncryptedMessage, password string) (string, error) {
	plaintext, err := helper.DecryptMessageWithPassword([]byte(password), encrypted.body)
	if err != nil {
		return "", err
	}
	return plaintext, nil
}
