package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"
	cepvalidator "github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/pkg/cep_validator"
)

type CEPRepository struct{}

func NewCEPRepository() *CEPRepository {
	return &CEPRepository{}
}

func (r *CEPRepository) IsValid(cep string) bool {
	return cepvalidator.IsValidCEP(cep)
}

func (r *CEPRepository) GetCEP(cep string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https:viacep.com.br/ws/%s/json/", cep), nil)
	if err != nil {
		log.Printf("Falha ao criar requisição HTTP: %v", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Falha ao realizar requisição HTTP: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	ctx_err := ctx.Err()
	if ctx_err != nil {
		log.Printf("Timeout ao realizar requisição HTTP: %v", ctx_err)
		return nil, err
	}

	response, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Falha ao ler resposta HTTP: %v", err)
		return nil, err
	}

	return response, nil
}

func (r *CEPRepository) ConvertResponse(cep_data []byte) (*entity.CEP, error) {
	var cep entity.CEP
	err := json.Unmarshal(cep_data, &cep)
	if err != nil {
		log.Printf("Falha ao converter resposta HTTP para CEP: %v", err)
		return nil, err
	}

	return &cep, nil
}
