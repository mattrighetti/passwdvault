package configuration

import (
	"fmt"
	"log"
	"os"
	"path"

	db "github.com/MattRighetti/passwdvault/database"
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
	configuraton *Configuration

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

// CheckInitFile checks wether the default configuration file is present or not
func CheckInitFile() error {
	configFilePath := path.Join(os.Getenv("HOME"), ConfigFileName+"."+ConfigFileType)
	log.Println("checking for file in " + configFilePath)
	if !fileExists(configFilePath) {
		log.Println("file could not be found")
		return fmt.Errorf("no .passwdvaultconfig file was found")
	}
	log.Println("file was found")

	return nil
}

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

	if !fileExists(completeFilePath) {
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

	mkBytes := make([]byte, viper.GetInt("database.masterkey.length"))
	byteRead, err := file.Read(mkBytes)
	if err != nil {
		return nil, err
	}

	return mkBytes[:byteRead], nil
}

// DbInit executes a function that initiates and opens BadgerDB
func DbInit() error {
	var err error
	databaseFilePath := path.Join(viper.GetString("database.path"), viper.GetString("database.name"))
	log.Printf("loading database from %s\n", databaseFilePath)
	if viper.GetBool("database.encrypted") {
		log.Println("database is encrypted")
		mkFilePath := viper.GetString("database.masterkey.fromfilepath")

		var masterKey []byte
		var masterKeyString string

		if mkFilePath != "" {
			log.Printf("reading masterkey from %s\n", mkFilePath)
			masterKey, err = ReadMasterKeyFromFile(mkFilePath)
		} else {
			fmt.Print("MasterKey: ")
			fmt.Scanf("%s", &masterKeyString)
			log.Printf("read %s|\n", masterKeyString)
		}

		if err != nil {
			return err
		}

		log.Printf("Read MasterKey: %s|\n", masterKey)

		db.DB, err = badger.Open(badger.DefaultOptions(databaseFilePath).WithLogger(nil).WithEncryptionKey(masterKey))
	} else {
		log.Println("database unencrypted")
		db.DB, err = badger.Open(badger.DefaultOptions(databaseFilePath).WithLogger(nil))
	}

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
	if err := CheckInitFile(); err != nil {
		return err
	}

	if err := ParseConfigurationFile(); err != nil {
		return err
	}

	if err := DbInit(); err != nil {
		return err
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
