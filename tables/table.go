package tables

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderTable(data [][]string, url string, autoMerge bool) {
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
	table.AppendBulk(data)
	table.Render()
}
