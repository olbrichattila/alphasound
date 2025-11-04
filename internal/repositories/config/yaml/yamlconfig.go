package yaml

import (
	"fmt"
	"os"

	"alphagen/internal/contracts"
	"alphagen/internal/dto"

	yamlReader "gopkg.in/yaml.v3"
)

func New(fileName string) (contracts.Config, error) {
	if fileName == "" {
		return nil, fmt.Errorf("yaml config, file name not set")
	}

	return &yconf{
		fileName: fileName,
	}, nil
}

type yconf struct {
	fileName string
}

// Get implements contracts.Config.
func (y *yconf) Get() (dto.InputParams, error) {
	data, err := os.ReadFile(y.fileName)
	if err != nil {
		return dto.InputParams{}, fmt.Errorf("error reading YAML file: %w", err)
	}

	var params dto.InputParams
	if err := yamlReader.Unmarshal(data, &params); err != nil {
		return dto.InputParams{}, fmt.Errorf("error parsing YAML: %w", err)
	}

	return params, nil
}
