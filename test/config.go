package test

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host      string
	AccessKey string
	SecretKey string
	AccountId int64
	SubUid    int64
}

var config *Config

func init() {

	filePtr, err := os.Open("config.json")
	if err != nil {
		return
	}
	defer filePtr.Close()

	config = new(Config)
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(config)
}
