package pkg

import "fmt"

type Ant struct {
	ID        int
	Path      []string
	RoomID    int
	Arrived   bool
	QuitStart bool
	pathIdx   int
}

type Distribution struct {
	ants   []int
	Path   []string
	LenAnt int
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
	for i, dist := range distributions {
		distributions[i].LenAnt = len(dist.ants)
	}
	return distributions
}

func MoveAnts(paths [][]string, numAnt int) {

	distributions := Distribute(numAnt, paths)

	ants := []Ant{}
	for i, dist := range distributions {
		for _, antID := range dist.ants {
			ant := Ant{ID: antID, Path: dist.Path[1:], RoomID: 0, Arrived: false, QuitStart: false, pathIdx: i}
			ants = append(ants, ant)
		}
	}

	numAntarrived := 0
	numAntleft := 0
	for numAntarrived < numAnt {
		for i, ant := range ants {
			if ant.QuitStart && !ant.Arrived {
				if ant.RoomID == len(ant.Path[1:]) {
					ants[i].Arrived = true
					numAntarrived++
				} else {
					ants[i].RoomID++
					fmt.Printf("L%d-%v ", ant.ID, distributions[ant.pathIdx].Path[1:][ants[i].RoomID])
				}
			}
		}

		if numAntleft < numAnt {
			for j, dist := range distributions {
				if dist.LenAnt > 0 {
					for i, ant := range ants {
						if !ant.QuitStart && ant.pathIdx == j {
							fmt.Printf("L%d-%s ", ant.ID, dist.Path[1:][ant.RoomID])
							ants[i].QuitStart = true
							distributions[j].LenAnt--
							numAntleft++
							break
						}
					}
				}
			}
		}
		if numAntarrived != numAnt {
			fmt.Println()
		}
	}
}
