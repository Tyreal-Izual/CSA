package main

import "fmt"

//func countNeighbour(i int, j int, maxi int, maxj int, world [][]byte) int {
//	var count int
//
//	for x := maxi; x <= i+1; i++ {
//		for y := maxj; y <= j+1; j++ {
//			if x == i && y == j {
//				continue
//			}
//
//		}
//	}
//	fmt.Println("count: ", count)
//	return count
//}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	fmt.Println(world)

	newWorld := make([][]byte, p.imageWidth, p.imageHeight)
	for i := range world {
		newWorld[i] = make([]byte, len(world[i]))
		copy(newWorld[i], world[i])
	}
	maxi := len(world) - 1    //一维slice
	maxj := len(world[0]) - 1 //二维slice
	fmt.Println(maxi, maxj)

	for i, _ := range world {
		for j, _ := range world[i] {
			n := countNeighbour(i, j, maxi, maxj, world)

			if (n < 2 || n > 3) && world[i][j] != 0 {
				fmt.Println("killcell:i:", i, "j:", j)
				newWorld[i][j] = 0
			} else if n == 3 && world[i][j] == 0 {
				fmt.Println("newcell:i:", i, "j", j)
				newWorld[i][j] = 0xFF
			} else {
				newWorld[i][j] = 0
			}
		}
	}
	copy(world, newWorld)
	fmt.Println("my:", world)
	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var aliveCells []cell
	for x, vx := range world {
		for y, vy := range vx {
			if vy == 0xFF {
				c := cell{y, x}
				aliveCells = append(aliveCells, c)
			}
		}
	}
	return []cell{}
}
