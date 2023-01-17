package ignite

import (
	"fmt"
	"os"
	"sort"

	"github.com/americanas-go/config"
	"github.com/common-nighthawk/go-figure"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Boot() {
	config.Load()

	if config.Bool(bannerEnabled) {
		fig := figure.NewColorFigure(config.String(phrase), config.String(fontName), config.String(color), config.Bool(strict))
		fig.Print()
	}

	if config.Bool(cfgEnabled) {
		var rows []table.Row

		entries := config.Entries()
		sort.Slice(entries[:], func(i, j int) bool {
			return entries[i].Key < entries[j].Key
		})

		for _, entry := range entries {
			v := config.Get(entry.Key)
			if entry.Options.Hide {
				v = "****"
			}

			maxLength := config.Int(maxLength)

			defaultValue := fmt.Sprintf("%v", entry.Value)
			if len(defaultValue) > maxLength {
				defaultValue = defaultValue[:maxLength]
			}

			value := fmt.Sprintf("%v", v)
			if len(value) > maxLength {
				value = value[:maxLength]
			}

			rows = append(rows, table.Row{
				entry.Key, defaultValue, value,
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
