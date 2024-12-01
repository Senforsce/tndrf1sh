package config

type Config struct {
	CurrentLang string
	ListenAddr  string
}

func New(lang string, addr string) *Config {
	return &Config{
		CurrentLang: lang,
		ListenAddr:  addr,
	}
}
