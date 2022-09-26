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
		DefaultValue:   "Traefik Guardian | Login",
		PossibleValues: []string{"*"},
		Validator:      ValidateAny,
	}

	Passwords = EnvVariable{
		Name:           "PASSWORDS",
		Required:       true,
		DefaultValue:   "",
		PossibleValues: []string{"algorithm:pass1|pass2|pass3"},
		Validator:      ValidatePasswords,
	}
)

func Initialize() error {
	var envVariables = []*EnvVariable{&Debug, &AuthDomain, &LoginPageTitle, &Passwords}

	for _, variable := range envVariables {
		err := variable.Validate()
		if err != nil {
			return err
		}

		logrus.Info("Config: ", variable.Name, " = ", variable.Value)
	}

	return nil
}
