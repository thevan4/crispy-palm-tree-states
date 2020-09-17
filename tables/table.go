package tables

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type reindex struct {
	oldIndex int
	newIndex int
}

func RenderTableData(data [][]string, url string, autoMerge bool, defaultColumnsHeaders []string) {
	if len(data) == 0 {
		fmt.Printf("no services was found at %v\r\n", url)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(defaultColumnsHeaders)
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

func RenderCustomTableData(data [][]string, url string, autoMerge bool, selectedColumns []string, defaultColumnsHeaders []string) {
	if len(data) == 0 {
		fmt.Printf("no services was found at %v\r\n", url)
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(selectedColumns)
	if autoMerge {
		table.SetAutoMergeCellsByColumnIndex([]int{0, 1, 4, 5, 6, 7, 9})
	} else {
		table.SetAutoMergeCellsByColumnIndex([]int{0})
	}
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetRowLine(true)

	//

	var serviceStateNewIndex, serverStateNewIndex int
	reindexes := make([]reindex, len(selectedColumns))
	for i, newColumn := range selectedColumns {
		for j, defaultD := range defaultColumnsHeaders {
			if newColumn == defaultD {
				if newColumn == "SRV STATE" {
					serviceStateNewIndex = i
				} else if newColumn == "HEALTH" {
					serverStateNewIndex = i
				}
				reindexes[i] = reindex{oldIndex: j, newIndex: i}
			}
		}
	}

	newData := make([][]string, len(data))
	for i, d := range data {
		tmpD := make([]string, len(reindexes))
		for _, reindex := range reindexes {
			tmpD[reindex.newIndex] = d[reindex.oldIndex]
		}
		newData[i] = tmpD
	}
	//
	for _, d := range newData {
		twService := tablewriter.Colors{}
		twServer := tablewriter.Colors{}
		if d[serviceStateNewIndex] == "UP" {
			twService = tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor}
		} else if d[serviceStateNewIndex] == "DOWN" {
			twService = tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor}
		}
		if d[serverStateNewIndex] == "UP" {
			twServer = tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor}
		} else if d[serverStateNewIndex] == "DOWN" {
			twServer = tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor}
		}
		tablewriterColors := make([]tablewriter.Colors, len(d))
		tablewriterColors[serviceStateNewIndex] = twService
		tablewriterColors[serverStateNewIndex] = twServer
		table.Rich(d, tablewriterColors)
	}
	table.Render()
}
