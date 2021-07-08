package main

import (
	"os"
	"plugin/exceltopb"

	"plugin/conf"

	"gopkg.in/yaml.v2"
)

func main() {
	exceltopb.ReadExcelSheet("../excel/test.xlsx")
}

func mustLoadConfig() *conf.Config {
	path := "setting.yaml"
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var data []byte
	_, err = f.Read(data)
	if err != nil {
		panic(err)
	}
	var config *conf.Config
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	return config
}
