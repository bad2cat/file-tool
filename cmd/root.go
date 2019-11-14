/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fd",
	Short: "A tool for upload and download file",
	Long:  `A tool which can transport file between os`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig) //这个函数的意思就是在每个命令运行的时候都会执行在这里面设置的函数

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("yaml") // 设置配置文件的类型
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".fd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".yaml")
	}

	//读取环境变量，看是否设置
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	out := os.Stdout
	if err := viper.ReadInConfig(); err == nil { //根据配置文件完成验证，执行每个命令都会进行验证
		name := viper.Get("name").(string)
		password := viper.Get("password").(string)
		isValid := checkConfig(name, password)
		if !isValid {
			fmt.Fprint(out, "username or passwor not correct")
			os.Exit(1)
		}
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Fprint(out, "config file not found")
		os.Exit(1)
	}
}

func checkConfig(name, password string) bool {
	if name == "" || password == "" {
		return false
	}
	if name != "xixi" || password != "haha" {
		return false
	}
	return true
}
