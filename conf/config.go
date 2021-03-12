package conf

import (
	"encoding/json"
	"os"
)

type FunCSSConf struct {
	DB struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"Name"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"db"`
}

func LoadFunCSSConf(file string) (*FunCSSConf, error) {
	f1, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	funCSSConf := &FunCSSConf{}
	err = json.NewDecoder(f1).Decode(funCSSConf)
	if err != nil {
		return nil, err
	}
	return funCSSConf, nil
}
