package configuration

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

// Configuration holds default configuration file
// with useful data about the database path and config
type Configuration struct {
	User     UserConfiguration
	Database DatabaseConfiguration
}

const (
	// ConfigFileName name of the configuration file
	ConfigFileName = ".passwdvaultconfig"

	// ConfigFileType configuration file type
	ConfigFileType = "yaml"
)

// DefaultConfig default configuration
var DefaultConfig = Configuration{
	User: UserConfiguration{
		Name:  "someuser",
		Email: "someuser@email.com",
	},
	Database: DatabaseConfiguration{
		Name:      "badger",
		Path:      "/tmp/",
		Encrypted: true,
		MasterKey: MasterKey{
			FromFilePath: "/etc/passwdvault/mk",
			Length:       32,
		},
	},
}

// CheckInitFile checks wether the default configuration file is present or not
// if not it creates it
func CheckInitFile() error {
	configFilePath := path.Join(os.Getenv("HOME"), ConfigFileName+"."+ConfigFileType)
	if !fileExists(configFilePath) {
		err := CreateDefaultFile(configFilePath)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateDefaultFile creates a default configuration file in $HOME/.passwdvaultconfig.yaml
func CreateDefaultFile(completeFilePath string) error {
	viper.AddConfigPath(os.Getenv("HOME"))
	viper.SetDefault("user", DefaultConfig.User)
	viper.SetDefault("database", DefaultConfig.Database)
	viper.WriteConfigAs(completeFilePath)

	if !fileExists(completeFilePath) {
		return fmt.Errorf("configuration file could not have been created")
	}

	return nil
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// ReadMasterKeyFromFile reads masterkey value from file
func ReadMasterKeyFromFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	mkBytes := make([]byte, DefaultConfig.Database.MasterKey.Length)
	byteRead, err := file.Read(mkBytes)
	if err != nil {
		return nil, err
	}

	return mkBytes[:byteRead], nil
}
