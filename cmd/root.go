package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/pydio/pydio-sdk-go/config"
)

var (
	configFile string

	host       string
	basePath   string
	schemes    string
	user       string
	pwd        string
	skipVerify bool
)

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Rest Client of Pydio 8",
	Long:  `Pydio 8 client for managing API`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Parse from parameters
		if host != "" && basePath != "" && schemes != "" && user != "" && pwd != "" {
			config.DefaultConfig = &config.SdkConfig{
				Host:     host,
				BasePath: basePath,
				// TOTO schemes
				Schemes:    []string{schemes},
				SkipVerify: skipVerify,

				// Pydio User Authentication
				User:     user,
				Password: pwd,
			}
			fmt.Println("Connecting to " + config.DefaultConfig.Host)
			fmt.Println("")
			return
		}

		// Parse from config
		if configFile == "" {
			log.Fatal("Please provide a path to the configuration file")
		}
		if data, e := ioutil.ReadFile(configFile); e == nil {
			var c config.SdkConfig
			if e := json.Unmarshal(data, &c); e != nil {
				log.Fatal("Cannot decode config content", e)
			}
			config.DefaultConfig = &c
			fmt.Println("Connecting to " + config.DefaultConfig.Host)
			fmt.Println("")
		} else {
			log.Fatal("Cannot read file ", e)
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVarP(&configFile, "config", "c", "config.json", "Path to the configuration file")

	flags.StringVarP(&host, "host", "a", "", "Server ip address")
	flags.StringVarP(&basePath, "basepath", "b", "", "Base path")
	flags.StringVarP(&schemes, "schemes", "s", "", "schemes https/http")
	flags.StringVarP(&user, "login", "l", "", "User login")
	flags.StringVarP(&pwd, "password", "p", "", "User password")
	
	
	flags.BoolVar(&skipVerify, "skipVerify", false, "Skip SSL certificate verification (not recommended)")
}
