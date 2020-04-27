package main

import (
	"github.com/andlabs/ui"
)

func main() {
	if err := ui.Main(setUp2); err != nil {
		panic(err)
	}
}
func setUp() {
	win := ui.NewWindow("title", 1024, 768, true)
	defer win.Show()
	win.SetMargined(true)
	win.SetBorderless(false)
	win.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		win.Destroy()
		return true
	})
	tab := ui.NewVerticalBox()
	tab.SetPadded(true)
	win.SetChild(tab)

	tr0 := ui.NewHorizontalBox()
	tr0.SetPadded(true)
	tab.Append(tr0, false)

	tr0.Append(ui.NewLabel("host:"), false)
	host := ui.NewEntry()
	tr0.Append(host, false)
	tr0.Append(ui.NewLabel("port:"), false)
	port := ui.NewEntry()
	tr0.Append(port, false)
	tr0.Append(ui.NewLabel("database:"), false)
	database := ui.NewEntry()
	tr0.Append(database, false)

	tr1 := ui.NewHorizontalBox()
	tr1.SetPadded(true)
	tab.Append(tr1, false)

	tr1.Append(ui.NewLabel("user:"), false)
	user := ui.NewEntry()
	tr1.Append(user, false)
	tr1.Append(ui.NewLabel("password:"), false)
	password := ui.NewEntry()
	tr1.Append(password, false)
	tr1.Append(ui.NewLabel("package:"), false)
	pkg := ui.NewEntry()
	tr1.Append(pkg, false)

	tr2 := ui.NewHorizontalBox()
	tr2.SetPadded(true)
	tab.Append(tr2, false)

	jsonAble := ui.NewCheckbox("JsonAble")
	tr2.Append(jsonAble, false)
	nullAble := ui.NewCheckbox("NullAble")
	tr2.Append(nullAble, false)
	gormAble := ui.NewCheckbox("GormAble")
	tr2.Append(gormAble, false)

	tr3 := ui.NewHorizontalBox()
	tr3.SetPadded(true)
	tab.Append(tr3, false)
	reg := ui.NewLabel("regexp:")
	tr3.Append(reg, false)
	tr3.Append(ui.NewEntry(), false)
	ScannerTabs := ui.NewButton("Scan Tables")
	tr3.Append(ScannerTabs, false)
	genCode := ui.NewButton("Generate code")
	tr3.Append(genCode, false)

	tr4 := ui.NewHorizontalBox()
	tr4.SetPadded(true)
	tab.Append(tr4, false)

	tr5 := ui.NewHorizontalBox()
	tr5.SetPadded(true)
	tab.Append(tr5, false)
	multilineEntry := ui.NewMultilineEntry()
	multilineEntry.SetText("sssss")
	tr5.Append(multilineEntry, false)
}

func setUp2() {
	win := ui.NewWindow("", 1024, 768, true)
	defer win.Show()

	table := ui.NewTable(&ui.TableParams{
		Model:                         ui.NewTableModel(H{}),
		RowBackgroundColorModelColumn: 3,
	})
	win.SetChild(table)
	table.AppendTextColumn("00", 0, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("00", 1, ui.TableModelColumnAlwaysEditable, nil)
}

type H struct {
}

func (h H) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{
		ui.TableString(""),
		ui.TableString(""),
		ui.TableString(""),
		ui.TableString(""),
	}
}

func (h H) NumRows(m *ui.TableModel) int {
	return 2
}

func (h H) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	return ui.TableString("")
}

func (h H) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {

}
