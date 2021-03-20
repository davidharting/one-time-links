package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptIsSymmetric(t *testing.T) {
	original := "encryption is fun, but not too easy"

	result, err := Encrypt(original)
	assert.Nil(t, err, "Encryption should not result in an error")

	plaintext, err := Decrypt(result.Message, result.Password)
	assert.Nil(t, err, "Decryption should not result in an error")

	assert.Equal(t, original, plaintext, "Decrypted message should match the original message")
}
