package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"strings"

	"github.com/pydio/pydio-sdk-go/config"
	"github.com/pydio/pydio-sdk-go/shares"
)

var (
	sharesListDetails bool
)

var SharesCmd = &cobra.Command{
	Use:   "share",
	Short: "Shares-related commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var ListSharesCmd = &cobra.Command{
	Use:   "ls",
	Short: "List shares",
	Run: func(cmd *cobra.Command, args []string) {

		// TODO : MAKE SURE TO USE NEW CLEAN ADMIN USER
		// Admin user may be the owner of some shared workspaces
		// wich will create issue
		shares.SetConfig(config.DefaultConfig)

		allShares, ignored, err := shares.LoadShares()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(fmt.Sprintf("Found %d shares", len(allShares)))
		if len(ignored) > 0 {
			for _, i := range ignored {
				log.Println("Ignored: " + i.Error())
			}
		}
		if !sharesListDetails {
			return
		}
		for _, s := range allShares {
			fmt.Println("")
			if e := s.LoadElement(); e != nil {
				fmt.Println("[ERROR] Cannot load ShareElement: ", e)
				continue
			}

			links := s.GetElement().GetLinks()
			if len(links) > 0 {
				fmt.Println("Share has links: ")
				for _, link := range links {
					fmt.Println("-- " + link.PublicLink)
				}
			}

			sharedUsers := s.GetSharedUsers()
			if len(sharedUsers) > 0 {
				fmt.Printf("Share has %d shared users: %s\n", len(sharedUsers), strings.Join(s.GetSharedUsersIds(), ","))
			}

			// Rebuild Path
			if fullPath, ws, ownerId, e := shares.RecurseParentPaths(allShares, s); e == nil {
				fmt.Printf("Full Node Path is %s:/%s [top owner:%s]\n", ws.Slug, fullPath, ownerId)
			} else {
				fmt.Println("[ERROR] Cannot rebuild full path: ", e)
			}
		}
	},
}

func init() {
	flags := SharesCmd.PersistentFlags()
	flags.BoolVarP(&sharesListDetails, "details", "d", false, "Load shares workspaces")
	SharesCmd.AddCommand(ListSharesCmd)
	RootCmd.AddCommand(SharesCmd)
}
