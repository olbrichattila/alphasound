package contracts

import "alphagen/internal/dto"

type AlphaGenerator interface {
	Generate(fileName string) error
}

type Config interface {
	Get() (dto.InputParams, error)
}
