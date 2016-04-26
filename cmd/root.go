// Copyright © 2016 Wolf Bauer <mailsupul@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var cfgFile string
var env string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "uservault",
	Short: "UserVault is a simple server for your users",
	Long: `UserVault is a server designed for storing user for any
	application, api or service.

	The server provides a simple REST api for handling user entities usinging
	JWT and api keys for securing the acccess.

	SSL is not supported but you can easily use any webserver like nginx as
	ssl proxy. You can also use the proxy as loadbalancer for multiple
	UserVault instances. UserVault is designed to work with muliple instances in parallel.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&env, "env", "e", "release", "Set the envirnoment for the command (e.g. release or dev).")
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/uservault/config then $HOME/.uservault/config or ./config)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.Println("Started in Environment:", env)

	var configFileName string = "config"
	if "release" != env {
		configFileName += "_" + env
	}

	log.Println("Config filename:", configFileName)

	viper.BindPFlags(RootCmd.Flags())

	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(configFileName) // name of config file (without extension)
	viper.AddConfigPath("/etc/uservault")  // adding home directory as first search path
	viper.AddConfigPath("$HOME/.uservault")  // adding home directory as first search path
	viper.AddConfigPath(".")  // adding home directory as first search path

	viper.SetEnvPrefix("uservault")
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
