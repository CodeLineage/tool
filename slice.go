package tool

// MakeSlice
// 创建二维切片
func MakeIntSlice(row, column int) [][]int {
	data := make([][]int, row)
	for index := range data {
		data[index] = make([]int, column)
	}
	return data
}
