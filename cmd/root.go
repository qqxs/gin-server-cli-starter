/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"

	"github.com/spf13/cobra"

	config "gin-server-cli/config"
	"gin-server-cli/server"
)

var cfgFile string

// var configChange = make(chan int, 1)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server is a simple restful api server",
	Long: `server is a simple restful api server
    use help get more ifo`,
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value.String()
		mode := cmd.Flag("mode").Value.String()
		server.RunServer(server.ServerConfig{
			Port: port,
			Mode: mode,
		})
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ./config/config.yaml)")

	rootCmd.Flags().StringP("port", "p", "8080", "default server port 8080")
	rootCmd.Flags().StringP("mode", "m", "debug", "default  server running in 'debug' mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {

	} else {
		c := config.Config{
			Name: cfgFile,
		}

		if err := c.InitConfig(); err != nil {
			panic(err)
		}
		log.Printf("载入配置成功")
		// c.WatchConfig(configChange)
	}

	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := homedir.Dir()
	// 	cobra.CheckErr(err)

	// 	// Search config in home directory with name "./config/config.yaml".
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigName("./config/config.yaml")
	// }

	// viper.AutomaticEnv() // read in environment variables that match

	// // If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }
}

// 定义 rootCmd 命令的执行
