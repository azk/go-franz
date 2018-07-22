package franz


type Provider interface {
	PopulateConfig(*Config) (*Config, error)
}

func Negotiate() (Provider, error) {

	return newEnv(), nil
}
