package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCEPUsecase(t *testing.T) {
	mockCEPRepo := new(MockCEPRepository)

	mockCEPRepo.On("IsValidCEP", "38050600").Return(true)
	mockCEPRepo.On("IsValidCEP", "00000000").Return(false)

	usecase := NewValidateCEPUsecase(mockCEPRepo)

	input := ValidateCEPInputDTO{CEP: "38050600"}
	output := usecase.Execute(input)
	assert.True(t, output.IsValid, "O CEP deveria ser válido")

	input = ValidateCEPInputDTO{CEP: "00000000"}
	output = usecase.Execute(input)
	assert.False(t, output.IsValid, "O CEP deveria ser inválido")

	mockCEPRepo.AssertExpectations(t)
}
