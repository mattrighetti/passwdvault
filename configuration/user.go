package configuration

// UserConfiguration struct of user configuration
type UserConfiguration struct {
	Name  string `mapstructure:"name"`
	Email string `mapstructure:"email"`
}
