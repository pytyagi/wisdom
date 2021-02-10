package api

var (
	DefaultEnv     = "dev"
	DefaultHost    = "localhost"
	DefaultPort    = 3000
	DefaultAPIPath = "/"
)

// Config struct with necessary info
type Config struct {
	Env     string
	Host    string
	Port    int
	APIPath string
}

// NewConfig return instance of Config
func NewConfig(env, host string, port int, apiPath string) *Config {
	return &Config{
		Env:     env,
		Host:    host,
		Port:    port,
		APIPath: apiPath,
	}
}

// DefaultConfig return instance of Config with defaults
func DefaultConfig() *Config {
	return &Config{
		Env:     DefaultEnv,
		Host:    DefaultHost,
		Port:    DefaultPort,
		APIPath: DefaultAPIPath,
	}
}
