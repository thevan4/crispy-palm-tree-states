package run

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thevan4/crispy-palm-tree-states/api"
	"github.com/thevan4/crispy-palm-tree-states/tables"
)

var rootCmd = &cobra.Command{
	Use:   "states",
	Short: "states for lb tier 1 😉",
	Run: func(cmd *cobra.Command, args []string) {

		// validate fields
		logging.WithFields(logrus.Fields{
			"config file path":          viperConfig.GetString(configFilePathName),
			"rest API url":              viperConfig.GetString(urlName),
			"login":                     viperConfig.GetString(nlbLoginName),
			"auto merge":                viperConfig.GetBool(autoMergeName),
			"Mode for IP + port search": viperConfig.GetString(ipAndPortSearchModeName),
			"password":                  "***",
		}).Info("")

		req := api.NewApiRequests(viperConfig.GetString(urlName),
			viperConfig.GetString(nlbLoginName),
			viperConfig.GetString(nlbPasswordName))
		rawServices, err := req.RequestServiceStates()
		if err != nil {
			logging.Fatalf("api request to lb tier 1 error: %v", err)
		}
		services := api.ModifyServicesToSliceOfStringSlices(rawServices, viperConfig.GetString(ipAndPortSearchModeName))
		tables.RenderAllTableData(services, viperConfig.GetString(urlName), viperConfig.GetBool(autoMergeName))
	},
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
