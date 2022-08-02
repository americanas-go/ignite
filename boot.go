package ignite

import (
	"os"

	"github.com/americanas-go/config"
	"github.com/common-nighthawk/go-figure"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Boot() {
	config.Load()

	if config.Bool(bannerEnabled) {
		fig := figure.NewColorFigure("ignite", "standard", "white", false)
		fig.Print()
	}

	if config.Bool(cfgEnabled) {
		var rows []table.Row
		for _, entry := range config.Entries() {
			v := config.Get(entry.Key)
			if entry.Options.Hide {
				v = "****"
			}
			rows = append(rows, table.Row{
				entry.Key, entry.Value, v,
			})
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleColoredBright)
		t.AppendHeader(table.Row{"key", "default value", "value"})
		t.AppendRows(rows)
		t.Render()
	}
}
