package layout

// WidthClass is the layout class of an allocation. Widths snap down to the nearest class so a layout never assumes
// room it was not given. Parse the raw width once with WidthClassOf and pass the class on.
type WidthClass int

const (
	Width40 WidthClass = iota
	Width55
	Width80
	Width100
	Width120
	Width160
)

// WidthClasses lists every class, narrowest first.
func WidthClasses() []WidthClass {
	return []WidthClass{Width40, Width55, Width80, Width100, Width120, Width160}
}

// Columns is the smallest width the class applies to.
func (c WidthClass) Columns() int {
	switch c {
	case Width40:
		return 40
	case Width55:
		return 55
	case Width80:
		return 80
	case Width100:
		return 100
	case Width120:
		return 120
	case Width160:
		return 160
	default:
		return 40
	}
}

// StackWidth is the widest allocation, in columns, that stacks label/value rows and narrow tables.
const StackWidth = 55

// StackedAt reports that an allocation of width columns stacks instead of sitting in columns: StackWidth and
// narrower. WidthClassOf snaps down, so Width55 also covers 56 to 79 columns; use StackedAt, not the class, when the
// rule is "at or below 55".
func StackedAt(width int) bool { return width <= StackWidth }

// Stacked reports that the class is below Width80, which includes every width from 55 to 79. It is a property of the
// class; StackedAt answers the same question for a raw width.
func (c WidthClass) Stacked() bool {
	switch c {
	case Width40, Width55:
		return true
	case Width80, Width100, Width120, Width160:
		return false
	default:
		return true
	}
}

// WidthClassOf is the largest class that fits width; anything under 40 columns uses the narrowest class.
func WidthClassOf(width int) WidthClass {
	class := Width40

	for _, c := range WidthClasses() {
		if width >= c.Columns() {
			class = c
		}
	}

	return class
}
