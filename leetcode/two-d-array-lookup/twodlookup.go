package twodarraylookup

/*
给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。
给定一个数，判断这个数是否在该二维数组中。
https://github.com/CyC2018/CS-Notes/blob/master/notes/%E5%89%91%E6%8C%87%20Offer%20%E9%A2%98%E8%A7%A3%20-%203~9.md#4-%E4%BA%8C%E7%BB%B4%E6%95%B0%E7%BB%84%E4%B8%AD%E7%9A%84%E6%9F%A5%E6%89%BE
*/

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
