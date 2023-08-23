package pkg

import "fmt"

type Ant struct {
	ID       int
	Path     []string
	RoomID   int
	Previous string
	Arrived  bool
}

type Distribution struct {
	ants   []int
	Path   []string
	Length int
}

func Distribute(ant int, paths [][]string) []Distribution {
	var distributions []Distribution
	for i := 0; i < len(paths); i++ {
		dist := Distribution{Path: paths[i], Length: len(paths[i])}
		distributions = append(distributions, dist)
	}
	j := 0

	distributions[0].ants = append(distributions[0].ants, 1)
	distributions[0].Length++
	for i := 2; i <= ant; i++ {

		if j+1 < len(distributions) && distributions[j+1].Length <= distributions[j].Length {
			distributions[j+1].ants = append(distributions[j+1].ants, i)
			distributions[j+1].Length++
			j++
		} else {
			j = 0
			distributions[j].ants = append(distributions[j].ants, i)
			distributions[j].Length++
		}
	}

	return distributions
}

func MoveAnts(paths [][]string, numAnt int, end string) {

	distributions := Distribute(numAnt, paths)

	ants := []Ant{}
	for _, dist := range distributions {
		for _, antID := range dist.ants {
			ant := Ant{ID: antID, Path: dist.Path[1:], RoomID: 0, Previous: "", Arrived: false}
			ants = append(ants, ant)
		}
	}

	exit := false
	taken := make(map[string]bool)
	var room string
	for !exit {
		allArrived := true
		for i, ant := range ants {
			if ant.Arrived {
				continue
			}
			if ant.RoomID < len(ant.Path) {
				room = ant.Path[ant.RoomID]
			}
			if taken[room] {
				continue
			}

			fmt.Print("L", ant.ID, "-", room, " ")
			ants[i].RoomID++

			if ant.RoomID >= len(ant.Path) {
				ant.Arrived = true
			} else {
				taken[ant.Previous] = false
				if room != end {
					taken[room] = true
					ants[i].Previous = room
				} else {
					ants[i].Arrived = true
				}
			}
			if !ant.Arrived {
				allArrived = false
			}
		}
		if !allArrived {
			fmt.Println()
		} else {
			exit = true
		}
	}
}
