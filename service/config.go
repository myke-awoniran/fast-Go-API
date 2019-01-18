package service

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config Variable
var Config *viper.Viper

// InitConfig Function
func initConfig() {
	// Set Configuration File Value
	configEnv := strings.ToLower(os.Getenv("CONFIG_ENV"))
	if len(configEnv) == 0 {
		configEnv = "dev"
	}

	// Set Configuration Path Value
	configPath := os.Getenv("CONFIG_PATH")
	if len(configPath) == 0 {
		configPath = "./configs"
	}

	// Set Configuration Type Value
	configType := strings.ToLower(os.Getenv("CONFIG_TYPE"))
	if len(configType) == 0 {
		configType = "yaml"
	}

	// Set Configuration Prefix Value
	configPrefix := strings.ToUpper(os.Getenv("CONFIG_PREFIX"))
	if len(configPrefix) == 0 {
		configPrefix = "CONFIG"
	}

	// Initialize Configuratior
	Config = viper.New()

	// Set Configuratior Configuration
	Config.SetConfigName(configEnv)
	Config.SetConfigType(configType)
	Config.AddConfigPath(configPath)

	// Set Configurator Environment
	Config.SetEnvPrefix(configPrefix)
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set Configurator to Auto Bind Configuration Variables to
	// Environment Variables
	Config.AutomaticEnv()

	// Set Configurator to Load Configuration File
	configLoadFile()

	// Set Configurator to Set Default Value and
	// Parse Configuration Variables
	configLoadValues()
}

// ConfigLoadFile Function to Load Configuration from File
func configLoadFile() {
	// Load Configuration File
	err := Config.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}

// ConfigLoadValues Function to Load Configuration Values
func configLoadValues() {
	// Server IP Value
	Config.SetDefault("SERVER_IP", "0.0.0.0")
	serverCfg.IP = Config.GetString("SERVER_IP")

	// Server Port Value
	Config.SetDefault("SERVER_PORT", "3000")
	serverCfg.Port = Config.GetString("SERVER_PORT")

	// Server Upload Path Value
	Config.SetDefault("SERVER_UPLOAD_PATH", "./uploads")

	// Server Upload Path Value
	Config.SetDefault("SERVER_UPLOAD_LIMIT", 25)

	// CORS Allowed Header Value
	Config.SetDefault("CORS_ALLOWED_HEADER", "X-Requested-With")
	routerCORSCfg.Headers = []string{Config.GetString("CORS_ALLOWED_HEADER")}

	// CORS Allowed Origin Value
	Config.SetDefault("CORS_ALLOWED_ORIGIN", "*")
	routerCORSCfg.Origins = []string{Config.GetString("CORS_ALLOWED_ORIGIN")}

	// CORS Allowed Method Value
	Config.SetDefault("CORS_ALLOWED_METHOD", "'HEAD', 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'OPTIONS'")
	routerCORSCfg.Methods = []string{Config.GetString("CORS_ALLOWED_METHOD")}

	// JWT Signing Key Value
	Config.SetDefault("JWT_SIGNING_KEY", "secrets")
	jwtSigningKey = Config.GetString("JWT_SIGNING_KEY")
}
