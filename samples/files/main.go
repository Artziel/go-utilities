package main

import (
	"fmt"

	Util "github.com/artziel/go-utilities"
)

type YamlFile struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Status   string `yaml:"status"`
}

func main() {

	file := YamlFile{}

	if err := Util.ReadYAML("./sample.yaml", file); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", file)
}
