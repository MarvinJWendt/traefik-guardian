package config

import (
	"os"
)

var CONFIG_FILE = ""

const CONFIG_DIR = "/config"

var Cfg Config

func init() {
	CONFIG_FILE = os.Getenv("CONFIG_FILE")
	if CONFIG_FILE == "" {
		CONFIG_FILE = "config.yml"
	}
}

type Config struct {
	AuthDomain string    `yaml:"authDomain"`
	LoginPage  LoginPage `yaml:"loginPage"`
	Groups     []Group   `yaml:"groups"`
	Users      []User    `yaml:"users"`
}

type LoginPage struct {
	Title            string `yaml:"title"`
	DisableParticles bool   `yaml:"disableParticles"`
}

type Group struct {
	Name            string   `yaml:"name"`
	AllowedUriRegex []string `yaml:"allowedUriRegex"`
}

type User struct {
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Groups   []string `yaml:"groups"`
}

var InitConfig = Config{
	AuthDomain: "auth.example.com",
	LoginPage: LoginPage{
		Title:            "Login Page",
		DisableParticles: false,
	},
	Groups: []Group{
		{
			Name:            "admin",
			AllowedUriRegex: []string{".*"},
		},
	},
	Users: []User{
		{
			Username: "admin",
			Password: "password",
			Groups:   []string{"admin"},
		},
	},
}

func DebuggingEnabled() bool {
	return os.Getenv("DEBUG") == "true"
}
