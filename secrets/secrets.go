package secrets

import (
	"fmt"

	"github.com/spf13/viper"
)

var AwsAccessKeyId string
var AwsSecretKey string

func init() {
	viper.AutomaticEnv()
	viper.SetConfigName("secrets")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error looking for or reading secrets file: %v\n", err)

	}

	AwsAccessKeyId = viper.GetString("aws_access_key_id")
	if len(AwsAccessKeyId) < 1 {
		panic("AwsAccessKeyId is required")
	}
	AwsSecretKey = viper.GetString("aws_secret_key")
	if len(AwsSecretKey) < 1 {
		panic("AwsSecretKey is required")
	}
}
