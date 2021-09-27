package config

type Config struct {
	App *AppConfig
}

type AppConfig struct {
	Addr string
	Port int
}

func GetConfig() *Config {
	return &Config{
		App: &AppConfig{
			Addr: "0.0.0.0",
			Port: 8000,
		},
	}
}