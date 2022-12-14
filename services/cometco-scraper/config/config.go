package config

import (
	"fmt"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
)

type Config struct {
	Log                LogConfig                `env:"SC_COMETCO_SCRAPER_LOG"`
	ProfileCredentials ProfileCredentialsConfig `env:"SC_COMETCO_SCRAPER_PROFILE_CREDENTIALS"`
	ChromeDP           ChromeDPConfig           `env:"SC_COMETCO_SCRAPER_CHROME_DP"`
	DB                 DBConfig                 `env:"SC_COMETCO_SCRAPER_DB"`
	HTTP               HTTPConfig               `env:"SC_COMETCO_SCRAPER_HTTP"`
}

type LogConfig struct {
	Format     string `env:"FORMAT" default:"console"`
	Debug      bool   `env:"DEBUG" default:"false"`
	Stacktrace bool   `env:"STACKTRACE" default:"false"`
	Caller     bool   `env:"CALLER" default:"false"`
}

type ProfileCredentialsConfig struct {
	Email string `env:"EMAIL"`
	Pass  string `env:"PASSWORD"`
}

type ChromeDPConfig struct {
	Debug    bool `env:"DEBUG" default:"false"`
	Headless bool `env:"HEADLESS" default:"true"`
}

type DBConfig struct {
	MongodbURI string `env:"MONGODB_URI"`
}

type HTTPConfig struct {
	Port int `env:"PORT" default:"8080"`
}

func GetConfig() (Config, error) {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		AllowUnknownFields: true,
		AllowUnknownEnvs:   true,
		FailOnFileNotFound: false,
		Files:              []string{".env"},
		FileDecoders:       map[string]aconfig.FileDecoder{".env": aconfigdotenv.New()},
	})
	if err := loader.Load(); err != nil {
		fmt.Println("error:", err)
		return Config{}, err
	}

	return cfg, nil
}
