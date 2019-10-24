package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	mci "github.com/bingoohuang/haconf"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	chk := pflag.BoolP("check", "", false, "check HAConf")
	ver := pflag.BoolP("version", "v", false, "show version")
	conf := pflag.StringP("config", "c", "./config.toml", "config file path")

	mci.DeclarePflagsByStruct(mci.Settings{})

	pflag.Parse()

	args := pflag.Args()
	if len(args) > 0 {
		fmt.Printf("Unknown args %s\n", strings.Join(args, " "))
		pflag.PrintDefaults()
		os.Exit(1)
	}

	if *ver {
		fmt.Printf("Version: 1.4.0\n")
		return
	}

	viper.SetEnvPrefix("HACONF")
	viper.AutomaticEnv()
	_ = viper.BindPFlags(pflag.CommandLine)

	configFile, _ := homedir.Expand(*conf)
	settings := mustLoadConfig(configFile)

	if *chk {
		settings.CheckHAProxyServers()
	}

	if *chk {
		return
	}

	if _, err := settings.Exec(); err != nil {
		logrus.Errorf("error %v", err)
		os.Exit(1)
	}
}

func findConfigFile(configFile string) (string, error) {
	if mci.FileExists(configFile) == nil {
		return configFile, nil
	}

	if ex, err := os.Executable(); err == nil {
		exPath := filepath.Dir(ex)
		configFile = filepath.Join(exPath, "config.toml")
	}

	if mci.FileExists(configFile) == nil {
		return configFile, nil
	}

	return "", errors.New("unable to find config file")
}

func loadConfig(configFile string) (config mci.Settings, err error) {
	if file, err := findConfigFile(configFile); err != nil {
		return config, err
	} else if _, err = toml.DecodeFile(file, &config); err != nil {
		logrus.Errorf("DecodeFile error %v", err)
	}

	return config, err
}

func mustLoadConfig(configFile string) (config mci.Settings) {
	config, _ = loadConfig(configFile)
	mci.ViperToStruct(&config)

	return config
}
