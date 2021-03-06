package main

import (
	"os"
	"strings"

	"github.com/arteev/fmttab"
)

func main() {
	tab := fmttab.New("Environments", fmttab.BorderDouble, nil)
	tab.CloseEachColumn = true
	tab.AddColumn("ENV", 25, fmttab.AlignLeft).
		AddColumn("VALUE", 25, fmttab.AlignLeft)
	for _, env := range os.Environ() {
		keyval := strings.Split(env, "=")
		tab.AppendData(map[string]interface{}{
			"ENV":   keyval[0],
			"VALUE": keyval[1],
		})
	}
	tab.Columns.Get(0).Caption = "Environment"
	tab.WriteTo(os.Stdout)
}
