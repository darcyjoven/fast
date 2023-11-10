package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".fast" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".fast")
	}
	// initViperDefault()
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err != nil {
			fmt.Println(err)
		}
		// 配置文件解析成功
	} else {
		_, err := fmt.Fprintln(os.Stderr, err.Error())
		if err != nil {
			fmt.Println(err)
		}
	}
}

// initViperDefault Viper 默认值
func initViperDefault() {
	name := filepath.Base(os.Args[0])
	name = strings.Replace(name, filepath.Ext(name), "", 1) // 取运行程序的名称
	viper.SetDefault("logdir", "temp/")                     // 日志目录
	viper.SetDefault("logname", name)                       // 日志名称
	viper.SetDefault("loginterval", "day")                  // 日志名称
}
