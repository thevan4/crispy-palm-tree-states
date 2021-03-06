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
		logging.WithFields(logrus.Fields{
			"config file path":          viperConfig.GetString(configFilePathName),
			"rest API url":              viperConfig.GetString(urlName),
			"login":                     viperConfig.GetString(nlbLoginName),
			"auto merge":                viperConfig.GetBool(autoMergeName),
			"Mode for IP + port search": viperConfig.GetString(ipAndPortSearchModeName),
			"Selected columns":          selectedColumns,
			"password":                  "***",
		}).Info("")

		req := api.NewApiRequests(viperConfig.GetString(urlName),
			viperConfig.GetString(nlbLoginName),
			viperConfig.GetString(nlbPasswordName))
		rawServices, err := req.RequestServiceStates()
		if err != nil {
			logging.Fatalf("api request to lb tier 1 error: %v", err)
		}
		if len(rawServices) == 0 {
			fmt.Printf("no services was found at %v\r\n", viperConfig.GetString(urlName))
			os.Exit(0)
		}
		services := api.ModifyServicesToSliceOfStringSlices(rawServices, viperConfig.GetString(ipAndPortSearchModeName), defaultColumnsHeaders)

		if len(viperConfig.GetStringSlice(selectedColumnsName)) == 0 {
			tables.RenderTableData(services, viperConfig.GetString(urlName), viperConfig.GetBool(autoMergeName), defaultColumnsHeaders)
		} else {
			tables.RenderCustomTableData(services, viperConfig.GetString(urlName), viperConfig.GetBool(autoMergeName), selectedColumns, defaultColumnsHeaders)
		}
		os.Exit(0)
	},
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
