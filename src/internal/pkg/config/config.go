package config

import "github.com/sirupsen/logrus"

var (
	Debug = EnvVariable{
		Name:           "DEBUG",
		Required:       false,
		DefaultValue:   "false",
		PossibleValues: []string{"true", "false"},
		Validator:      ValidateCaseInsensitivePossibleValues,
	}

	AuthDomain = EnvVariable{
		Name:           "AUTH_DOMAIN",
		Required:       true,
		DefaultValue:   "",
		PossibleValues: []string{"*"},
		Validator:      ValidateNotEmptyString,
	}

	LoginPageTitle = EnvVariable{
		Name:           "LOGIN_PAGE_TITLE",
		Required:       false,
		DefaultValue:   "Traefik Auth Provider | Login",
		PossibleValues: []string{"*"},
		Validator:      ValidateAny,
	}
)

func Initialize() error {
	var envVariables = []*EnvVariable{&Debug, &AuthDomain, &LoginPageTitle}

	for _, variable := range envVariables {
		err := variable.Validate()
		if err != nil {
			return err
		}

		logrus.Info("Config: ", variable.Name, " = ", variable.Value)
	}

	return nil
}
