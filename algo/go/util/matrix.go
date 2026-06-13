package util

func VisitedMatrix(rsize, csize int) [][]bool {
	visited := make([][]bool, rsize)
	for i := range visited {
		visited[i] = make([]bool, csize)
	}
	return visited
}

func VisitedMatrixInt(rsize, csize int) [][]int8 {
	visited := make([][]int8, rsize)
	for i := range visited {
		visited[i] = make([]int8, csize)
	}
	return visited
}

func ShiftMatrix() ([]int, []int) {
	return []int{1, 0, -1, 0}, []int{0, 1, 0, -1}
}
