package main

import (
	"fmt"

	"github.com/robjporter/go-library/xtable"
)

func main() {
	t, e := xtable.NewTableWriter([]string{"ID", "Name", "Description", "Cost"}, []int{10, 35, 50, 20})
	if e != nil {
		fmt.Println(e)
	}
	t.PrintHeader()
	t.PrintRow([]string{"1", "Test name", "Test description", "cost"}, xtable.AlignLeft)
	t.PrintFooter()
	summaryRow1 := "Id's: 3"
	summaryRow2 := "Count: 1000022"
	t.PrintRowAsOneColumn(summaryRow1, xtable.AlignLeft)
	t.PrintRowAsOneColumn(summaryRow2, xtable.AlignLeft)
	t.PrintFooter()
}
