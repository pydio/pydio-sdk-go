package cmd

import (
	"fmt"
	p "path"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/pydio/pydio-sdk-go/config"
	"github.com/pydio/pydio-sdk-go/models"
	"net/url"
)

var SharesCmd = &cobra.Command{
	Use:   "share",
	Short: "Shares-related commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

type listRep struct {
	Shares map[string]*models.Share `json:"data"`
	Cursor map[string]int `json:"cursor"`
}

var ListSharesCmd = &cobra.Command{
	Use:   "ls",
	Short: "List shares",
	Run: func(cmd *cobra.Command, args []string) {

		c := config.DefaultConfig
		httpClient := config.GetHttpClient(c)
		listUri := p.Join(c.Url, c.Path, "api/settings/sharelist-load")
		req, _ := http.NewRequest("POST", fmt.Sprintf("%s://%s/?format=json", c.Protocol, listUri), nil)
		req.SetBasicAuth(c.User, c.Password)
		resp, e := httpClient.Do(req)
		if e != nil {
			log.Fatal(e)
		}
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			log.Fatal("Cannot read body", e)
		}
		var result listRep
		if e := json.Unmarshal(body, &result); e!= nil {
			fmt.Printf("%#v\n", e)
			log.Fatal("Unmarshal Error", e)
		}

		log.Println(fmt.Sprintf("Got %d shares, total %v", len(result.Shares), result.Cursor["total"]))
		for _, s := range result.Shares {
			log.Println(s.SHARETYPE + " - " + s.OWNERID + " - " + s.Metadata.SharedElementHash )
			detailUri := p.Join(c.Url, c.Path, "api", fmt.Sprintf("%v", s.Metadata.SharedElementParentRepository), "load_shared_element_data/", s.Metadata.OriginalPath)
			req1, _ := http.NewRequest("POST", fmt.Sprintf("%s://%s", c.Protocol, detailUri), nil)
			req1.SetBasicAuth(c.User, c.Password)
			req1.PostForm = url.Values{}
			req1.PostForm.Set("owner", s.OWNERID)
			req1.PostForm.Set("merged", "true")
			resp1, e := httpClient.Do(req1)
			if e != nil {
				log.Fatal("Detail request error ", detailUri, " ", e)
			}
			body, _ := ioutil.ReadAll(resp1.Body)
			var element models.ShareElement
			if e := json.Unmarshal(body, &element); e != nil {
				log.Fatal("Parse detail request error ", detailUri, " ", e)
			}
			fmt.Println(element.Description)
		}
	},
}

func init() {
	SharesCmd.AddCommand(ListSharesCmd)
	RootCmd.AddCommand(SharesCmd)
}
