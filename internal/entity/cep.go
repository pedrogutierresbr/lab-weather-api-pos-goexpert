package entity

type CEP struct {
	CEP         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	UF          string
	Estado      string
	IBGE        string
	GIA         string
	DDD         string
	SIAFI       string
}

func NewCEP(cep, logradouro, complemento, bairro, localidade, uf, estado, ibge, gia, ddd, siafi string) *CEP {
	newCep := &CEP{
		CEP:         cep,
		Logradouro:  logradouro,
		Complemento: complemento,
		Bairro:      bairro,
		Localidade:  localidade,
		UF:          uf,
		Estado:      estado,
		IBGE:        ibge,
		GIA:         gia,
		DDD:         ddd,
		SIAFI:       siafi,
	}

	return newCep
}
