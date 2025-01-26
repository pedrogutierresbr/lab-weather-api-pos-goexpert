package cepvalidator

import "regexp"

func IsValidCEP(cep string) bool {
	return regexp.MustCompile("^[0-9]{8}$").MatchString(cep)
}
