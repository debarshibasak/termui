package widgets

import (
	"image"
	"log"
	"strings"

	. "github.com/debarshibasak/termui/v3"
)

type ScrollableTable struct {
	*Block

	Header []string
	Rows   [][]string

	ColWidths []int
	ColGap    int
	PadLeft   int

	ShowCursor  bool
	CursorColor Color

	UniqueCol    int    // the column used to uniquely identify each table row
	SelectedItem string // used to keep the cursor on the correct item if the data changes
	SelectedRow  int
	TopRow       int // used to indicate where in the table we are scrolled at

	ColResizer func()
}

// NewTable returns a new Table instance
func NewTableScrollable() *ScrollableTable {
	return &ScrollableTable{
		Block:       NewBlock(),
		SelectedRow: 0,
		TopRow:      0,
		UniqueCol:   0,
		ColResizer:  func() {},
	}
}

func (self *ScrollableTable) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	self.ColResizer()

	// finds exact column starting position
	colXPos := []int{}
	cur := 1 + self.PadLeft
	for _, w := range self.ColWidths {
		colXPos = append(colXPos, cur)
		cur += w
		cur += self.ColGap
	}

	// prints header
	for i, h := range self.Header {
		width := self.ColWidths[i]
		if width == 0 {
			continue
		}
		// don't render column if it doesn't fit in widget
		if width > (self.Inner.Dx()-colXPos[i])+1 {
			continue
		}
		buf.SetString(
			h,
			NewStyle(Theme.Default.Fg, ColorClear, ModifierBold),
			image.Pt(self.Inner.Min.X+colXPos[i]-1, self.Inner.Min.Y),
		)
	}

	if self.TopRow < 0 {
		log.Printf("table widget TopRow value less than 0. TopRow: %v", self.TopRow)
		return
	}

	// prints each row
	for rowNum := self.TopRow; rowNum < self.TopRow+self.Inner.Dy()-1 && rowNum < len(self.Rows); rowNum++ {
		row := self.Rows[rowNum]
		y := (rowNum + 2) - self.TopRow

		// prints cursor
		style := NewStyle(Theme.Default.Fg)
		if self.ShowCursor {
			if (self.SelectedItem == "" && rowNum == self.SelectedRow) || (self.SelectedItem != "" && self.SelectedItem == row[self.UniqueCol]) {
				style.Fg = self.CursorColor
				style.Modifier = ModifierReverse
				for _, width := range self.ColWidths {
					if width == 0 {
						continue
					}
					buf.SetString(
						strings.Repeat(" ", self.Inner.Dx()),
						style,
						image.Pt(self.Inner.Min.X, self.Inner.Min.Y+y-1),
					)
				}
				self.SelectedItem = row[self.UniqueCol]
				self.SelectedRow = rowNum
			}
		}

		// prints each col of the row
		for i, width := range self.ColWidths {
			if width == 0 {
				continue
			}
			// don't render column if width is greater than distance to end of widget
			if width > (self.Inner.Dx()-colXPos[i])+1 {
				continue
			}
			r := TrimString(row[i], width)
			buf.SetString(
				r,
				style,
				image.Pt(self.Inner.Min.X+colXPos[i]-1, self.Inner.Min.Y+y-1),
			)
		}
	}
}

// Scrolling ///////////////////////////////////////////////////////////////////

// calcPos is used to calculate the cursor position and the current view into the table.
func (self *ScrollableTable) calcPos() {
	self.SelectedItem = ""

	if self.SelectedRow < 0 {
		self.SelectedRow = 0
	}
	if self.SelectedRow < self.TopRow {
		self.TopRow = self.SelectedRow
	}

	if self.SelectedRow > len(self.Rows)-1 {
		self.SelectedRow = len(self.Rows) - 1
	}
	if self.SelectedRow > self.TopRow+(self.Inner.Dy()-2) {
		self.TopRow = self.SelectedRow - (self.Inner.Dy() - 2)
	}
}

func (self *ScrollableTable) ScrollUp() {
	self.SelectedRow--
	self.calcPos()
}

func (self *ScrollableTable) ScrollDown() {
	self.SelectedRow++
	self.calcPos()
}

func (self *ScrollableTable) ScrollTop() {
	self.SelectedRow = 0
	self.calcPos()
}

func (self *ScrollableTable) ScrollBottom() {
	self.SelectedRow = len(self.Rows) - 1
	self.calcPos()
}

func (self *ScrollableTable) ScrollHalfPageUp() {
	self.SelectedRow = self.SelectedRow - (self.Inner.Dy()-2)/2
	self.calcPos()
}

func (self *ScrollableTable) ScrollHalfPageDown() {
	self.SelectedRow = self.SelectedRow + (self.Inner.Dy()-2)/2
	self.calcPos()
}

func (self *ScrollableTable) ScrollPageUp() {
	self.SelectedRow -= (self.Inner.Dy() - 2)
	self.calcPos()
}

func (self *ScrollableTable) ScrollPageDown() {
	self.SelectedRow += (self.Inner.Dy() - 2)
	self.calcPos()
}

func (self *ScrollableTable) HandleClick(x, y int) {
	x = x - self.Min.X
	y = y - self.Min.Y
	if (x > 0 && x <= self.Inner.Dx()) && (y > 0 && y <= self.Inner.Dy()) {
		self.SelectedRow = (self.TopRow + y) - 2
		self.calcPos()
	}
}