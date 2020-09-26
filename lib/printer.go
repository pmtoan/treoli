package lib

import (
	"os"

	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
)

func DefaultPrint(data interface{}) {
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "|"
	printer.ColumnSeparator = "|"
	printer.RowSeparator = "-"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor

	printer.Print(data)
}
