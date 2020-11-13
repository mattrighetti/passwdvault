package configuration

// DatabaseConfiguration struct of the database configuration
type DatabaseConfiguration struct {
	Name      string `mapstructure:"filename"`
	Path      string `mapstructure:"path"`
	Encrypted bool   `mapstructure:"encrypted"`
	MasterKey MasterKey
}

// MasterKey struct of the masterkey needed to unlock the database
type MasterKey struct {
	FromFilePath string `mapstructure:"filepath"`
	Length       int8   `mapstructure:"len"`
}
