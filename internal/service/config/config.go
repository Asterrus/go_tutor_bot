package config

type Config struct {
	Domen    *string
	Subdomen *string
}

func NewConfig(
	domen *string,
	subdomen *string,
) *Config {
	return &Config{
		Domen:    domen,
		Subdomen: subdomen,
	}
}
