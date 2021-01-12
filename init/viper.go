package init

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// inits viper
func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetTypeByDefaultValue(true)

	viper.SetDefault("CONFIG_FILE", "config.yml")

	configfile := viper.GetString("CONFIG_FILE")

	viper.SetConfigName(configfile)
	viper.SetConfigType(filepath.Ext(configfile)[1:])
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		logrus.Debugln(err)
	} else if err != nil {
		logrus.Fatalln(err)
	}
}

// sets the env vars for go-swagger if set with viper
func init() {
	viper.SetDefault("host", "0.0.0.0")
	if host := viper.GetString("host"); host != "" {
		//#nosec
		os.Setenv("HOST", host)
	}

	viper.SetDefault("port", "8080")
	if host := viper.GetString("port"); host != "" {
		//#nosec
		os.Setenv("PORT", host)
	}

	if host := viper.GetString("tls.host"); host != "" {
		//#nosec
		os.Setenv("TLS_HOST", host)
	}

	viper.SetDefault("tls.port", "8443")
	if host := viper.GetString("tls.port"); host != "" {
		//#nosec
		os.Setenv("TLS_PORT", host)
	}

	viper.SetDefault("tls.certificate", "tls/certificate.crt")
	if host := viper.GetString("tls.certificate"); host != "" {
		//#nosec
		os.Setenv("TLS_CERTIFICATE", host)
	}

	viper.SetDefault("tls.key", "tls/key.key")
	if host := viper.GetString("tls.key"); host != "" {
		//#nosec
		os.Setenv("TLS_PRIVATE_KEY", host)
	}

	if host := viper.GetString("tls.ca.certificate"); host != "" {
		//#nosec
		os.Setenv("TLS_CA_CERTIFICATE", host)
	}
}
