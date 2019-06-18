package twodarraylookup

var twodarray = make([][]int, 5)

func init() {
	twodarray[0] = []int{1, 3, 5, 6, 8}
	twodarray[1] = []int{4, 9, 20, 30, 32}
	twodarray[2] = []int{8, 20, 40, 80, 100}
	twodarray[3] = []int{100, 110, 120, 200, 300}
	twodarray[4] = []int{102, 120, 121, 201, 301}
}

type table struct {
	twodarray              [][]int
	maxRow, maxCol         int
	currentRow, currentCol int
}

func newTable() *table {
	return &table{
		twodarray:  twodarray,
		maxRow:     4,
		maxCol:     4,
		currentRow: 0,
		currentCol: 4,
	}
}

func (tbl *table) edgeInclude(target int) bool {
	return target >= tbl.twodarray[0][0] && target <= tbl.twodarray[tbl.maxRow][tbl.maxCol]
}

func (tbl *table) moveLeft() (overflow bool) {
	tbl.currentCol--
	if tbl.currentCol < 0 {
		return true
	}
	return false
}

func (tbl *table) moveDown() (overflow bool) {
	tbl.currentRow++
	if tbl.currentRow > tbl.maxRow {
		return true
	}
	return false
}

func (tbl *table) value() int {
	return tbl.twodarray[tbl.currentRow][tbl.currentCol]
}

func (tbl *table) contains(target int) bool {
	if tbl.value() == target {
		return true
	} else if tbl.value() > target {
		if overflow := tbl.moveLeft(); overflow {
			return false
		}
		return tbl.contains(target)
	} else {
		if overflow := tbl.moveDown(); overflow {
			return false
		}
		return tbl.contains(target)
	}
}
