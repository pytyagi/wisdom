package api

var (
	DefaultEnv     = "dev"
	DefaultHost    = "localhost"
	DefaultPort    = 3000
	DefaultAPIPath = "/"
)

type Config struct {
	Env     string
	Host    string
	Port    int
	APIPath string
}

func NewConfig(env, host string, port int, apiPath string) *Config {
	return &Config{
		Env:     env,
		Host:    host,
		Port:    port,
		APIPath: apiPath,
	}
}

func DefaultConfig() *Config {
	return &Config{
		Env:     DefaultEnv,
		Host:    DefaultHost,
		Port:    DefaultPort,
		APIPath: DefaultAPIPath,
	}
}
