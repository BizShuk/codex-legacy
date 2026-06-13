package graph

import "github.com/bizshuk/algo/util"

// [Variant]: [DFS] Find cells flow water to both Pacific and Atlantic
// - Left & Top boarder are Pacific Ocean
// - Right & Buttom are Atlantic Ocean
// - MxN height matrix
// When it's raining, water will flow from higher ground to lower ground.
// Find out which ground water will flow into both ocean

// [Solution]: Use DFS and copy the result to the same height ground
func Pacificatlantic(heights [][]int) [][]int {
	rsize := len(heights)
	csize := len(heights[0])

	result := make([][]int, 0)
	calMap := make([][]byte, rsize)
	for i := 0; i < rsize; i++ {
		calMap[i] = make([]byte, csize)
	}

	shiftx := []int{1, 0, -1, 0}
	shifty := []int{0, 1, 0, -1}

	var dfs func(x, y int)
	var copyEqualCell func(x, y int)

	copyEqualCell = func(x, y int) {
		for i := 0; i < 4; i++ {
			nx, ny := x+shiftx[i], y+shifty[i]
			isOcean, _ := checkOcean(nx, ny, rsize, csize)

			if isOcean {
				continue
			}

			if heights[x][y] != heights[nx][ny] {
				continue
			}

			if calMap[nx][ny] == calMap[x][y] {
				continue
			}

			calMap[nx][ny] = calMap[nx][ny] | calMap[x][y]
			copyEqualCell(nx, ny)
		}
	}

	dfs = func(x, y int) {
		if calMap[x][y] > 0 { // visited
			return
		}

		calMap[x][y] = 1

		for i := 0; i < 4; i++ {
			nx, ny := x+shiftx[i], y+shifty[i]
			isOcean, ocean := checkOcean(nx, ny, rsize, csize)

			if isOcean {
				calMap[x][y] = calMap[x][y] | ocean
				continue
			}

			if heights[x][y] < heights[nx][ny] {
				continue
			}

			dfs(nx, ny)
			calMap[x][y] = calMap[x][y] | calMap[nx][ny]
		}

		copyEqualCell(x, y)
	}

	for i := 0; i < rsize; i++ {
		for j := 0; j < csize; j++ {
			dfs(i, j)
		}
	}
	// fmt.Println(calMap)
	for i := 0; i < rsize; i++ {
		for j := 0; j < csize; j++ {
			if calMap[i][j] == 7 {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func checkOcean(x, y, r, c int) (bool, byte) {
	if x < 0 || y < 0 {
		return true, byte(2)
	}

	if x >= r || y >= c {
		return true, byte(4)
	}

	return false, byte(0)
}

// [Solution]: Start from ocean boarders seperately, from Pacific, from Atlantic. Similiar to Flipping 'O'

// [Variant]: [BFS] Flipping 'O' region surrounded by 'X'
// [Hint]: Start from board only
func FlippingOsurroundByX(board [][]byte) {
	rsize, csize := len(board), len(board[0])

	visited := make([][]byte, rsize)
	for rowi := range visited {
		visited[rowi] = make([]byte, csize)
	}

	xshift := []int{1, 0, -1, 0}
	yshift := []int{0, 1, 0, -1}

	var bfs func(x, y int)

	bfs = func(x, y int) {
		if x < 0 || y < 0 || x >= rsize || y >= csize { // out of boarder
			return
		}
		if board[x][y] == 'X' || visited[x][y] == '1' {
			return
		}

		visited[x][y] = '1'

		for i := 0; i < 4; i++ {
			bfs(x+xshift[i], y+yshift[i])
		}
	}

	queue := make([][]int, 0)

	//  row is 0
	for col := 0; col < csize; col++ {
		if board[0][col] == 'O' {
			queue = append(queue, []int{0, col})
		}
	}

	//  col is 0
	for row := 0; row < rsize; row++ {
		if board[row][0] == 'O' {
			queue = append(queue, []int{row, 0})
		}
	}
	//  row is rsize-1
	for col := 0; col < csize; col++ {
		if board[rsize-1][col] == 'O' {
			queue = append(queue, []int{rsize - 1, col})
		}
	}
	//  col is csize-1
	for row := 0; row < rsize; row++ {
		if board[row][csize-1] == 'O' {
			queue = append(queue, []int{row, csize - 1})
		}
	}

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		bfs(x, y)
	}

	for i := 0; i < rsize; i++ {
		for j := 0; j < csize; j++ {
			if visited[i][j] == 0 && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// [Variant]: [BFS] Count number of round all orange rotten.
// [Tips]: if no rotten orange, 0. if some fresh remain, -1.
func RottenRound(grid [][]int) int {

	queue := make([][]int, 0)

	rsize, csize := len(grid), len(grid[0])

	xshift, yshift := util.ShiftMatrix()
	freshCount := 0

	for i := 0; i < rsize; i++ {
		for j := 0; j < csize; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
				grid[i][j] = 0
			}
			if grid[i][j] == 1 {
				freshCount += 1
			}
		}
	}
	if freshCount == 0 {
		return 0
	}
	round := 0
	for len(queue) > 0 {
		nQ := make([][]int, 0)
		round += 1
		// fmt.Println(queue)
		for i := range queue {
			x, y := queue[i][0], queue[i][1]

			for d := 0; d < 4; d++ {
				nx, ny := x+xshift[d], y+yshift[d]
				if nx < 0 || ny < 0 || nx >= rsize || ny >= csize || grid[nx][ny] == 0 {
					continue
				}

				if grid[nx][ny] == 1 {
					nQ = append(nQ, []int{nx, ny})
					grid[nx][ny] = 0
					freshCount -= 1
				}
			}
		}
		queue = nQ
	}

	if freshCount > 0 {
		return -1
	}

	return round - 1
}
