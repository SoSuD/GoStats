package apiserver

// Config ...
type Config struct {
	BindAddr    string         `toml:"bind_addr"`
	LogLevel    string         `toml:"log_level"`
	Clouvider   ProviderConfig `toml:"Clouvider"`
	SmsActivate ProviderConfig `toml:"SMSActivate"`
	Asocks      ProviderConfig `toml:"Asocks"`
	DataImpulse ProviderConfig `toml:"DataImpulse"`
	ProxyUrl    string         `toml:"proxy_url"`
	//DatabaseURL string `toml:"database_url"`
	//SessionKey  string `toml:"session_key"`
}

type ProviderConfig struct {
	ApiKey          string `toml:"ApiKey"`
	CriticalBalance int    `toml:"CriticalBalance"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
