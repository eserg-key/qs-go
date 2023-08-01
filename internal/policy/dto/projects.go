package dto

type AllProjectOutput struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func NewAllProjectOutput(name string, code string) *AllProjectOutput {
	return &AllProjectOutput{Name: name, Code: code}
}
