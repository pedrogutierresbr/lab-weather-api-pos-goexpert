package usecase

import (
	"errors"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"
)

type CEPInputDTO struct {
	CEP string
}

type CEPOutputDTO struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

type GetCEPUsecase struct {
	CEPRepository entity.CEPRepositoryInterface
}

func NewCepUseCase(cepRepository entity.CEPRepositoryInterface) *GetCEPUsecase {
	return &GetCEPUsecase{
		CEPRepository: cepRepository,
	}
}

func (c *GetCEPUsecase) Execute(input CEPInputDTO) (CEPOutputDTO, error) {
	cep := entity.CEP{
		CEP: input.CEP,
	}

	isValidCEP := c.CEPRepository.IsValidCEP(cep.CEP)
	if !isValidCEP {
		return CEPOutputDTO{}, errors.New("CEP inválido")
	}

	response, err := c.CEPRepository.GetCEP(cep.CEP)
	if err != nil {
		return CEPOutputDTO{}, err
	}

	cepOutput, err := c.CEPRepository.ConvertResponse(response)
	if err != nil {
		return CEPOutputDTO{}, err
	}

	dto := CEPOutputDTO{
		CEP:         cepOutput.CEP,
		Logradouro:  cepOutput.Logradouro,
		Complemento: cepOutput.Complemento,
		Bairro:      cepOutput.Bairro,
		Localidade:  cepOutput.Localidade,
		UF:          cepOutput.UF,
		IBGE:        cepOutput.IBGE,
		GIA:         cepOutput.GIA,
		DDD:         cepOutput.DDD,
		SIAFI:       cepOutput.SIAFI,
	}
	return dto, nil
}
