package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	// the shared DB ORM object
	DB *gorm.DB
	// the error thrown be GORM when using DB ORM object
	DBErr error
	// the server port. Defaults to 1234
	ServerPort int `mapstructure:"server_port"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `mapstructure:"dsn"`
	// the API key needed to authorize to API. required.
	APIKey string `mapstructure:"api_key"`
	// Auth0 JWT Audience. required.
	JWTAudience string `mapstructure:"jwt_audience"`
	// Auth0 JWT Issuer. required.
	JWTIssuer string `mapstructure:"jwt_issuer"`
	// Auth0 JWT JWKS path. required.
	JWTJwks string `mapstructure:"jwt_jwks"`
	// Certificate file for HTTPS
	CertFile string `mapstructure:"cert_file"`
	// Private key file for HTTPS
	KeyFile string `mapstructure:"key_file"`
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("health_survey")
	v.AutomaticEnv()

	Config.DSN = v.Get("DSN").(string)
	Config.APIKey = v.Get("API_KEY").(string)
	Config.JWTAudience = v.Get("JWT_AUDIENCE").(string)
	Config.JWTIssuer = v.Get("JWT_ISSUER").(string)
	Config.JWTJwks = v.Get("JWT_JWKS").(string)
	v.SetDefault("server_port", 1234)

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
