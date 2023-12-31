package formatting

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func Table(header table.Row, rows []table.Row) string {
	tw := table.NewWriter()
	tw.Style().Format.Header = text.FormatUpper
	tw.Style().Box = table.StyleBoxRounded
	for i, value := range header {
		header[i] = text.Bold.Sprint(value)
	}
	tw.AppendHeader(header)
	tw.AppendRows(rows)
	return tw.Render()
}

func Truncate(value string, maxLength int) string {
	if len(value) <= maxLength {
		return value
	}
	return fmt.Sprintf("%s...", text.Trim(value, maxLength))
}
