package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	mark := true
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			mark = false
		}
		if mark {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}
	mark = true
	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			mark = false
		}
		if mark {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] != -1 {
				if obstacleGrid[i-1][j] == -1 {
					obstacleGrid[i-1][j] = 0
				}
				if obstacleGrid[i][j-1] == -1 {
					obstacleGrid[i][j-1] = 0
				}
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
func main() {
	a := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 0, 0}}
	b := uniquePathsWithObstacles(a)
	fmt.Println(b)
}

/*
0 0 0    1 1 1
0 1 0    1 0 1
1 0 0    0 1 2

0 0 0 1 0    1 1 1 0 0
0 0 0 0 0    1 2 3 3 4
*/
