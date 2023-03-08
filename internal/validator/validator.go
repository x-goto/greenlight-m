package validator

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, err string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = err
	}
}

func (v *Validator) Check(ok bool, key, err string) {
	if !ok {
		v.AddError(key, err)
	}
}
