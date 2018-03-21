package cleantherules

import (
	"strings"

	"github.com/spf13/viper"
)

var r *Rules

// Rules is the object will conaint all the rules the user want to check
// *** TODO: Handle multiple configuration file
type Rules struct {
	// config file information
	// configFileName      string
	// configFileDirectory string

	// All configuration set by user.
	config *viper.Viper
}

// NewRules return an initialized Rules instances
func NewRules() *Rules {
	r := new(Rules)

	// Set config file name
	viper.SetConfigName("config")
	// Set config file directory
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	r.config = viper.Sub("rules")
	return r
}

// Get retourn a specif key of our config file.
// @param key string Key to get in the config file
// @return map[string]string
func Get(key string) map[string]string { return r.get(key) }
func (r *Rules) get(key string) map[string]string {
	return r.config.GetStringMapString(strings.ToLower(key))
}

// GetGeneralConfig return all general configuration
// @return map[string]string
func (r *Rules) getGeneralConfig() map[string]string {
	return r.get("general")
}
