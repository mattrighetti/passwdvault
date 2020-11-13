package configuration

// DatabaseConfiguration struct of the database configuration
type DatabaseConfiguration struct {
	Name      string `mapstructure:"filename"`
	Path      string `mapstructure:"path"`
	Encrypted bool   `mapstructure:"encrypted"`
}
