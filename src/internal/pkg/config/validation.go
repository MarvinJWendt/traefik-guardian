package config

import (
	"fmt"
	"os"
	"strings"
)

type EnvVariable struct {
	Name           string
	Required       bool
	DefaultValue   string
	Value          string
	PossibleValues []string
	Validator      func(v EnvVariable) bool
}

func (v *EnvVariable) String() string {
	return v.Value
}

func (v *EnvVariable) ToBool() bool {
	return strings.ToLower(v.Value) == "true"
}

func (v *EnvVariable) Validate() error {
	v.Value = os.Getenv(v.Name)

	if v.Value == "" {
		v.Value = v.DefaultValue
	}

	if v.Required && v.Value == "" {
		return NewValidationError(v.Name, "required but not set", v.PossibleValues)
	}

	if !v.Validator(*v) {
		return NewValidationError(v.Name, v.Value, v.PossibleValues)
	}

	return nil
}

var (
	ValidateNotEmptyString = func(v EnvVariable) bool {
		return v.Value != ""
	}
	ValidateAny = func(v EnvVariable) bool {
		return true
	}
	ValidateStrictPossibleValues = func(v EnvVariable) bool {
		for _, possibleValue := range v.PossibleValues {
			if v.Value == possibleValue {
				return true
			}
		}

		return false
	}
	ValidateCaseInsensitivePossibleValues = func(v EnvVariable) bool {
		for _, possibleValue := range v.PossibleValues {
			if strings.EqualFold(v.Value, possibleValue) {
				return true
			}
		}

		return false
	}
)

type ValidationError struct {
	KeyName        string
	AcceptedValues []string
	ProvidedValue  string
}

func NewValidationError(keyName, providedValue string, acceptedValues []string) *ValidationError {
	return &ValidationError{
		KeyName:        keyName,
		AcceptedValues: acceptedValues,
		ProvidedValue:  providedValue,
	}
}

func (e ValidationError) Error() string {
	return e.String()
}

func (e ValidationError) String() string {
	return fmt.Sprintf(
		"invalid value for key '%s': '%s', accepted values: %v",
		e.KeyName,
		e.ProvidedValue,
		e.AcceptedValues,
	)
}
