package configuration

import (
	"fmt"
	"log"
	"os"
	"path"

	db "github.com/MattRighetti/passwdvault/database"
	"github.com/MattRighetti/passwdvault/utils"
	badger "github.com/dgraph-io/badger/v2"
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
var (
	configuraton   *Configuration
	ConfigFilePath = path.Join(os.Getenv("HOME"), ConfigFileName+"."+ConfigFileType)

	DefaultConfig = Configuration{
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
)

// ParseConfigurationFile parses the configuration file
func ParseConfigurationFile() error {
	viper.SetConfigType(ConfigFileType)
	viper.SetConfigName(ConfigFileName)
	viper.AddConfigPath(os.Getenv("HOME"))
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %s", err)
	}

	return nil
}

// CreateConfigurationFile creates a configuration file in $HOME/.passwdvaultconfig.yaml with
// values specified in user and database
func CreateConfigurationFile(user *UserConfiguration, database *DatabaseConfiguration) error {
	viper.AddConfigPath(os.Getenv("HOME"))
	viper.SetDefault("user", *user)
	viper.SetDefault("database", *database)

	completeFilePath := path.Join(os.Getenv("HOME"), ConfigFileName+"."+ConfigFileType)
	viper.WriteConfigAs(completeFilePath)

	if !FileExists(completeFilePath) {
		return fmt.Errorf("configuration file could not have been created")
	}

	return nil
}

// CreateDefaultFile creates a default configuration file in $HOME/.passwdvaultconfig.yaml
func CreateDefaultFile(completeFilePath string) error {
	return CreateConfigurationFile(&DefaultConfig.User, &DefaultConfig.Database)
}

// ReadMasterKeyFromFile reads masterkey value from file
func ReadMasterKeyFromFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mkBytes := make([]byte, viper.GetInt("database.masterkey.length"))
	byteRead, err := file.Read(mkBytes)
	if err != nil {
		return nil, err
	}

	return mkBytes[:byteRead], nil
}

// DbInit executes a function that initiates and opens BadgerDB
func DbInit() error {
	options, err := initOptions()
	if err != nil {
		return err
	}

	db.DB, err = badger.Open(options)
	if err != nil {
		return err
	}

	return nil
}

// CloseDb closes BadgerDB
func CloseDb() {
	db.DB.Close()
}

// InitCriticalData utility function that initiates every critical data needed to the program
func InitCriticalData() error {
	if exists := FileExists(ConfigFilePath); exists != true {
		return fmt.Errorf("configuration file is not present")
	}

	if err := ParseConfigurationFile(); err != nil {
		return err
	}

	databaseFilePath := path.Join(os.Getenv("HOME"), viper.GetString("database.path"), viper.GetString("database.name"))
	if exists := FileExists(databaseFilePath); exists != true {
		return fmt.Errorf("database file is not present")
	}

	if err := DbInit(); err != nil {
		return err
	}

	return nil
}

func initOptions() (badger.Options, error) {
	databaseFilePath := path.Join(viper.GetString("database.path"), viper.GetString("database.name"))
	options := badger.DefaultOptions(databaseFilePath).WithLogger(nil)

	if viper.GetBool("database.encrypted") {
		if mk := viper.GetString("database.masterkey"); mk != "" {
			log.Printf("MasterKey is present: %s\n", mk)
			options.EncryptionKey = []byte(mk)
			return options, nil
		}

		handleReadMasterKey(&options)
		return options, nil
	}

	return options, nil
}

func handleReadMasterKey(opt *badger.Options) error {
	mkFilePath := viper.GetString("database.masterkey.fromfilepath")

	var err error
	var masterkey []byte
	if mkFilePath != "" {
		masterkey, err = ReadMasterKeyFromFile(mkFilePath)
	} else {
		masterkey, err = utils.ReadInputStringHideInput("MasterKey: ")
	}

	if err != nil {
		return err
	}

	opt.EncryptionKey = masterkey

	return nil
}

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
