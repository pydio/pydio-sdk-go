package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/pydio-sdk-go/client/provisioning"
	"github.com/pydio/pydio-sdk-go/config"
)

var RoleCmd = &cobra.Command{
	Use:   "role",
	Short: "Role-related commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var ListRoleCmd = &cobra.Command{
	Use:   "ls",
	Short: "List roles",
	Run: func(cmd *cobra.Command, args []string) {

		httpClient := config.GetHttpClient(config.DefaultConfig)
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Fatal(err)
		}

		format := "json"

		params := &provisioning.GetRolesParams{
			Context:    ctx,
			HTTPClient: httpClient,
			Format:     &format,
		}

		result, err := apiClient.Provisioning.GetRoles(params, nil)
		if err != nil {
			fmt.Printf("could not list roles: %s\n", err.Error())
			log.Fatal(err)
		}

		count := 1
		for key, node := range result.Payload.Data.Children {
			fmt.Printf("%d. %s - %v\n", count, key, node)
			count = count + 1
		}
	},
}

func init() {
	RoleCmd.AddCommand(ListRoleCmd)
	RootCmd.AddCommand(RoleCmd)
}
