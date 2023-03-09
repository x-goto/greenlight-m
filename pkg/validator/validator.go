package validator

import "regexp"

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) AddError(key, value string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = value
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Check(check bool, key, value string) {
	if !check {
		v.AddError(key, value)
	}
}

func IsValueInList[T comparable](value T, list ...T) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

func MatchesRX(str string, rx *regexp.Regexp) bool {
	return rx.MatchString(str)
}

func IsSetUnique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, v := range values {
		uniqueValues[v] = true
	}

	return len(uniqueValues) == len(values)
}
