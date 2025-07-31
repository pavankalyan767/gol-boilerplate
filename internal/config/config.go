package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("github.com/pavan-kalyan/Nova/env")
	viper.SetConfigType(".env")
	fmt.Print("Loading configuration...",viper.ConfigFileUsed())
}
