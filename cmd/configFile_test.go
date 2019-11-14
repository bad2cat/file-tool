package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestConfigFile(t *testing.T) {
	configDir := "D:/tempDir/fd.yaml"

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configDir)
	if err := viper.ReadInConfig(); err == nil {
		name := viper.Get("name").(string)
		password := viper.Get("password")
		fmt.Printf("name is %s,password is %s", name, password)
	}
	fmt.Println("test")
}
