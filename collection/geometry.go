package collection

import "github.com/lanej/charmed/theme"

// TablePaddingWidth is the horizontal space reserved by NewTable's default
// cell styles for the given number of visible columns. Recalculate it when applying different native cell padding.
func TablePaddingWidth(columns int) int {
	return max(0, columns) * 2 * theme.InlinePadding
}

// DenseTablePaddingWidth is the trailing separation for visible columns in NewDenseTable.
func DenseTablePaddingWidth(columns int) int {
	return max(0, columns) * theme.SpaceSmall
}
