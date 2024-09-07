package internal

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type shopifyConfLi struct {
	Name        string `yaml:"name"`
	AccessToken string `yaml:"accessToken"`
}

type conf struct {
	Shopify []shopifyConfLi `yaml:"shopify"`
}

func ReadConf(path string) (*conf, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	t := &conf{}
	err = yaml.Unmarshal(byteValue, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
