package config

import (
	"log"

	"github.com/spf13/viper"
)

var ViperConfig = viper.New()

func InitConfig() {
	ViperConfig.AddConfigPath(".")
	ViperConfig.SetConfigName("chatwatcher")
	ViperConfig.SetConfigType("yaml")

	ViperConfig.SetDefault("profanityfilter.enabled", true)
	ViperConfig.SetDefault("profanityfilter.permission", "git-fal7.chatwatcher.profanityfilter.bypass")
	ViperConfig.SetDefault("antispam.enabled", true)
	ViperConfig.SetDefault("antispam.permission", "git-fal7.chatwatcher.antispam.bypass")
	ViperConfig.SetDefault("antispam.cooldown", 3)
	ViperConfig.SetDefault("antispam.message", "Please wait %duration%")

	err := ViperConfig.ReadInConfig()
	if err != nil {
		// Create config file
		log.Println("Couldn't find chatwatcher.yml, creating a new config file")
		ViperConfig.WriteConfigAs("./chatwatcher.yml")
	}

}
