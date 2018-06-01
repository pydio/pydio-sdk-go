package cmd

import (
	"fmt"
	"log"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

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
		rt := httptransport.New(
			config.DefaultConfig.Host,
			config.DefaultConfig.BasePath,
			config.DefaultConfig.Schemes)
		rt.Consumers["text/xml"] = runtime.XMLConsumer()
		rt.DefaultAuthentication = httptransport.BasicAuth(config.DefaultConfig.User, config.DefaultConfig.Password)
		pydioSdkGo := clientapi.New(rt, strfmt.Default)

		fmt.Printf("User credential aa : %s / %s\n", config.DefaultConfig.User, config.DefaultConfig.Password)
        params := provisioning.NewGetRolesParams()

		//rolesOK, err := clientapi.Default.Provisioning.GetRoles(params, basicAuth)
		rolesOK, err := pydioSdkGo.Provisioning.GetRoles(params, httptransport.BasicAuth(config.DefaultConfig.User, config.DefaultConfig.Password))
		fmt.Println(rolesOK, err)

		if err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("Found %d results\n", len(rolesOK.Payload.Data))

	},
}

func init() {
	RolesCmd.AddCommand(ListRolesCmd)
	RootCmd.AddCommand(ListRolesCmd)
}
