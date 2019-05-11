// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"
	"github.com/pcavezzan/sample-apiserver/app"
	"github.com/pcavezzan/sample-apiserver/app/web"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var cfg app.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sample",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: RunApp,
}

func RunApp(cmd *cobra.Command, args []string) {
	web.Start(&cfg)
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
	cfg = app.Config{}

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sample.yaml)")

	rootCmd.PersistentFlags().IntVar(&cfg.Port, "port", 8000, "server port to listen to.")

	rootCmd.PersistentFlags().StringVar(&cfg.DataSource.User, "datasource.user", "postgres", "User to connect to database")
	rootCmd.PersistentFlags().StringVar(&cfg.DataSource.Password, "datasource.password", "postgres", "User password to connect to database")
	rootCmd.PersistentFlags().StringVar(&cfg.DataSource.Database, "datasource.database", "sample", "Database name to connect to database")
	rootCmd.PersistentFlags().StringVar(&cfg.DataSource.Host, "datasource.host", "localhost", "Host to connect to database")
	rootCmd.PersistentFlags().IntVar(&cfg.DataSource.Port, "datasource.port", 5432, "Database port")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

		// Search config in home directory with name ".sample" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".sample")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}


}
