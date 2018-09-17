// Copyright © 2016 Christian Lipphardt <christian.lipphardt@camunda.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cvm",
	Short: "Setup a Camunda BPM installation",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var CfgFile string
var Debug bool
var DryRun bool

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&CfgFile, "config", "c", "", "config file (default is $HOME/.cvm.yml)")
	RootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "enable debug output")
	RootCmd.PersistentFlags().BoolVar(&DryRun, "dry-run", false, "enable dry-run mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	}
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigName(".cvm") // name of config file (without extension)
	viper.SetConfigType("yml")  // type of the configuration returned, e.g. "json"

	cwd, _ := os.Getwd()
	viper.AddConfigPath(cwd)  // adding local directory as first search path
	viper.AddConfigPath(home) // adding home directory as last search path

	viper.SetEnvPrefix("cvm")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if Debug {
		viper.Debug()
	}
}
