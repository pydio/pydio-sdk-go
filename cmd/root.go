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

	protocol   string
	host       string
	path       string
	user       string
	pwd        string
	skipVerify bool
)

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Rest Client of Pydio Cells",
	Long:  `Pydio Cells client for managing API`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Parse from parameters
		//if host != "" && id != "" && secret != "" && user != "" && pwd != "" {
		if host != "" && user != "" && pwd != "" {
			config.DefaultConfig = &config.SdkConfig{
				Protocol:     protocol,
				Url:          host,
				Path:         path,
				User:         user,
				Password:     pwd,
				SkipVerify:   skipVerify,
			}
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
			fmt.Println("Connecting to " + config.DefaultConfig.Url)
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

	flags.StringVar(&protocol, "protocol", "http", "Http scheme to server (http|https)")
	flags.StringVarP(&host, "url", "u", "", "Host to server (e.g. share.company.com)")
	flags.StringVarP(&path, "path", "t", "", "Path to server (e.g. /pydio)")
	flags.StringVarP(&user, "login", "l", "", "User login")
	flags.StringVarP(&pwd, "password", "p", "", "User password")
	flags.BoolVar(&skipVerify, "skipVerify", false, "Skip SSL certificate verification (not recommended)")

}
