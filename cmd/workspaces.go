package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/pydio-sdk-go/client/provisioning"
	"github.com/pydio/pydio-sdk-go/config"
)

var WorkspacesCmd = &cobra.Command{
	Use:   "ws",
	Short: "Workspaces-related commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func getWorkspaces() error {
	httpClient := config.GetHttpClient(config.DefaultConfig)
	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)

	format := "json"
	params := &provisioning.AdminListWorkspacesParams{
		Context:    ctx,
		HTTPClient: httpClient,
		Format:     &format,
	}

	result, err := apiClient.Provisioning.AdminListWorkspaces(params)
	if err != nil {
		return err
	}

	if result.Payload == nil || result.Payload.Data == nil || result.Payload.Data.Children == nil {
		return fmt.Errorf("Could not load workspaces")
	}

	for key, _ := range result.Payload.Data.Children {
		if key == "/" {
			continue
		}
		wsParams := &provisioning.AdminGetWorkspaceParams{
			Context:ctx,
			HTTPClient:httpClient,
			Format:&format,
			WorkspaceID:key,
		}
		if res, e := apiClient.Provisioning.AdminGetWorkspace(wsParams); e == nil {
			ws := res.Payload
			wsParams := ws.Parameters.(map[string]interface{})
			if *ws.AccessType == "fs" {
				fmt.Printf("FS Workspace : %s - %s - %v\n", key, *ws.Display, wsParams["PATH"])
			} else if *ws.AccessType == "s3" {
				fmt.Printf("FS Workspace : %s - %s - %v\n", key, *ws.Display, wsParams["BUCKET"])
			} else {
				fmt.Printf("Unsupported Access Type %s (%s)\n", *ws.AccessType, *ws.Display)
			}
		} else {
			fmt.Printf("could not load workspaces %s: %v\n", key, e.Error())
		}
	}

	return nil
}

var ListWsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List workspaces",
	Run: func(cmd *cobra.Command, args []string) {

		if err := getWorkspaces(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	WorkspacesCmd.AddCommand(ListWsCmd)
	RootCmd.AddCommand(WorkspacesCmd)
}
