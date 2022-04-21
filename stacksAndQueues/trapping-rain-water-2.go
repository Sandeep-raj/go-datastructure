package stacksandqueues

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Trapping Rain Water - 2

Given an m x n integer matrix heightMap representing the height of each unit cell in a 2D elevation map, return the volume of water it can trap after raining.

Sample Input
0 1 4 2 8
4 3 6 5 0
1 2 4 1 4
2 0 7 3 2
3 1 5 9 2

Sample Output
4
*/

func TestTrappingRainWater2() {
	heightMap := make([][]int, 0)
	heightMap = append(heightMap, []int{0, 1, 4, 2, 8})
	heightMap = append(heightMap, []int{4, 3, 6, 5, 0})
	heightMap = append(heightMap, []int{1, 2, 4, 1, 4})
	heightMap = append(heightMap, []int{2, 0, 7, 3, 2})
	heightMap = append(heightMap, []int{3, 1, 5, 9, 2})

	traverse := make([][]bool, len(heightMap))
	for i := 0; i < len(heightMap); i++ {
		traverse[i] = make([]bool, len(heightMap[i]))
	}

	pq := utils.InitPQ()

	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if i == 0 || j == 0 || i == len(heightMap)-1 || j == len(heightMap[0])-1 {
				pq.Add(utils.PQObj{Key: heightMap[i][j], Data: []int{i, j}})
			}
		}
	}

	trappingRainWater(heightMap, traverse, pq)

}

func trappingRainWater(hmap [][]int, traverse [][]bool, pq *utils.PQ) {
	totWater := 0

	direction := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	block, err := pq.Remove()
	for err == nil {
		x := block.Data.([]int)[0]
		y := block.Data.([]int)[1]

		traverse[x][y] = true

		min := 10000
		minx := -1
		miny := -1

		for _, dir := range direction {
			if x+dir[0] < 0 || x+dir[0] > len(hmap)-1 ||
				y+dir[1] < 0 || y+dir[1] > len(hmap[0])-1 || traverse[x+dir[0]][y+dir[1]] {
				continue
			} else {
				if hmap[x+dir[0]][y+dir[1]] < min {
					min = hmap[x+dir[0]][y+dir[1]]
					minx = x + dir[0]
					miny = y + dir[1]
				}
			}
		}

		if minx != -1 && miny != -1 {
			if hmap[x][y] > hmap[minx][miny] {
				totWater = totWater + (hmap[x][y] - hmap[minx][miny])
				hmap[minx][miny] = hmap[x][y]
			}

			pq.Add(utils.PQObj{Key: hmap[minx][miny], Data: []int{minx, miny}})
		}

		block, err = pq.Remove()
	}

	log.Print(totWater)
}
