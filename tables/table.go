package tables

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderAllTableData(data [][]string, url string, autoMerge bool) {
	if len(data) == 0 {
		fmt.Printf("no services was found at %v\r\n", url)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"VIP", "SRV STATE", "REAL", "HEALTH", "PROTO", "ROUTING", "TYPE", "HC TYPE", "HC ADDR", "HC TIMERS"})
	if autoMerge {
		table.SetAutoMergeCellsByColumnIndex([]int{0, 1, 4, 5, 6, 7, 9})
	} else {
		table.SetAutoMergeCellsByColumnIndex([]int{0})
	}
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetRowLine(true)
	for _, d := range data {
		twService := tablewriter.Colors{}
		twServer := tablewriter.Colors{}
		if d[1] == "UP" {
			twService = tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor}
		} else if d[1] == "DOWN" {
			twService = tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor}
		}
		if d[3] == "UP" {
			twServer = tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor}
		} else if d[3] == "DOWN" {
			twServer = tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor}
		}
		table.Rich(d, []tablewriter.Colors{nil, twService, nil, twServer, nil, nil, nil, nil, nil, nil})
	}
	table.Render()
}
