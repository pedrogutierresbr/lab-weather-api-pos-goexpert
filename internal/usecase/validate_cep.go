package usecase

import "github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"

type ValidateCEPInputDTO struct {
	CEP string
}

type ValidateCEPOutputDTO struct {
	IsValid bool
}

type ValidateCEPUsecase struct {
	CEPRepository entity.CEPRepositoryInterface
}

func NewValidateCEPUsecase(cepRepository entity.CEPRepositoryInterface) *ValidateCEPUsecase {
	return &ValidateCEPUsecase{
		CEPRepository: cepRepository,
	}
}

func (c *ValidateCEPUsecase) Execute(input ValidateCEPInputDTO) ValidateCEPOutputDTO {
	isValid := c.CEPRepository.IsValidCEP(input.CEP)

	return ValidateCEPOutputDTO{
		IsValid: isValid,
	}
}
