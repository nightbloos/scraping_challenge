package app

import (
	"fmt"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
)

type Config struct {
	Log                LogConfig                `env:"SCRAPING_CHALLENGE_LOG"`
	ProfileCredentials ProfileCredentialsConfig `env:"SCRAPING_CHALLENGE_PROFILE_CREDENTIALS"`
}

type LogConfig struct {
	Format string `env:"FORMAT" default:"console"`
	Debug  bool   `env:"DEBUG" default:"false"`
}

type ProfileCredentialsConfig struct {
	Email string `env:"EMAIL"`
	Pass  string `env:"PASSWORD" `
}

func GetConfig() (Config, error) {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		AllowUnknownFields: true,
		AllowUnknownEnvs:   true,
		FailOnFileNotFound: false,
		AllFieldRequired:   true,
		Files:              []string{".env"},
		FileDecoders:       map[string]aconfig.FileDecoder{".env": aconfigdotenv.New()},
	})
	if err := loader.Load(); err != nil {
		fmt.Println("error:", err)
		return cfg, err
	}

	return cfg, nil
}
