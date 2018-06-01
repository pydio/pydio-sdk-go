package cmd

import (
	"fmt"
	"log"

	"github.com/pydio/pydio-sdk-go/client/provisioning"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/spf13/cobra"

	clientapi "github.com/pydio/pydio-sdk-go/client"
	"github.com/pydio/pydio-sdk-go/config"
)

var RolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "Roles-related commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// ListRolesCmd test command on pydio 8 roles
var ListRolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "List roles",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Default config: %s\n", config.DefaultConfig.BasePath)
		transportConfig := clientapi.TransportConfig{
			Host:     config.DefaultConfig.Host,
			BasePath: config.DefaultConfig.BasePath,
			Schemes:  config.DefaultConfig.Schemes,
		}
		pydioSdkGo := clientapi.NewHTTPClientWithConfig(nil, &transportConfig)

		basicAuth := httptransport.BasicAuth(config.DefaultConfig.User, config.DefaultConfig.Password)

		fmt.Printf("User credential : %s / %s\n", config.DefaultConfig.User, config.DefaultConfig.Password)
        params := provisioning.NewGetRolesParams()        
		rolesOK, err := pydioSdkGo.Provisioning.GetRoles(params, basicAuth)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Found %d results\n", len(rolesOK.Payload.Data.Children))

	},
}

func init() {
	RolesCmd.AddCommand(ListRolesCmd)
	RootCmd.AddCommand(ListRolesCmd)
}
