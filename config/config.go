package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	InitialBalanceAmount float32 `json:"initialBalanceAmount"`
	MinumumBalanceAmount float32 `json:"minumumBalanceAmount"`
}

var c = &Config{}

func init() {
	file, err := os.Open(".config/" + env + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(read, c)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return c
}
