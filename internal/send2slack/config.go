package send2slack

import (
	"github.com/spf13/viper"
)

const Version = "0.1.2"

type Config struct {
	Token           string
	DefChannel      string
	SendmailChannel string
}

func NewConfig() (*Config, error) {

	viper.SetConfigName("config.yaml")       // name of config file (without extension)
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/send2slack/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.send2slack") // call multiple times to add many search paths
	viper.AddConfigPath(".")                 // optionally look for config in the working directory
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		return nil, err
	}

	cfg := Config{
		Token:           viper.GetString("token"),
		DefChannel:      viper.GetString("default_channel"),
		SendmailChannel: viper.GetString("sendmail_channel"),
	}
	return &cfg, nil
}