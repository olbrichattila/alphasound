package main

import (
	"fmt"
	"os"

	"alphagen/internal/repositories/config/yaml"
	"alphagen/internal/repositories/generator"
)

func main() {
	configFileName, wavFileName, err := getParameters()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	config, err := yaml.New(configFileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	params, err := config.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gen := generator.New(params)

	gen.Generate(wavFileName)
}

func getParameters() (string, string, error) {
	if len(os.Args) != 3 {
		return "", "", fmt.Errorf("usage alphagen <config> <output>\nNote: without extension")
	}

	return os.Args[1] + ".yaml", os.Args[2] + ".wav", nil
}
