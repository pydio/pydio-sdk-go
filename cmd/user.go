package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pydio/pydio-sdk-go/client/provisioning"
	"github.com/pydio/pydio-sdk-go/config"
)

var (
	createUserLogin string
	createUserPwd   string
	createUserAdmin bool
)


var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "User-related commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func getUsersFromPath(path string) error {
	httpClient := config.GetHttpClient(config.DefaultConfig)
	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)

	format := "json"
	params := &provisioning.GetPeopleParams{
		Context:    ctx,
		HTTPClient: httpClient,
		Format:     &format,
		Path: path,
	}

	result, err :=  apiClient.Provisioning.GetPeople(params, nil)
	if err != nil {
		return err
	}

	if result.Payload == nil || result.Payload.Data == nil || result.Payload.Data.Children == nil {
		return fmt.Errorf("Could not load users")
	}

	for key, node := range result.Payload.Data.Children {
		if node.Type == "collection" {
			err := getUsersFromPath(strings.TrimLeft(key, "/"))
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("%s/%s - %v\n", path, key, node)
		}
	}

	return nil
}

var ListUserCmd = &cobra.Command{
	Use:   "ls",
	Short: "List users",
	Run: func(cmd *cobra.Command, args []string) {

		if err := getUsersFromPath(""); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {

	UserCmd.AddCommand(ListUserCmd)

	RootCmd.AddCommand(UserCmd)
}
