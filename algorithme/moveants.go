package pkg

import "fmt"

type Ant struct {
	id int
}

type Path struct {
	Id       map[int][]string
	Rooms    int
	path     []string
	ants     int
	Assigned []Ant
}

func AssignPath(ant *Ant, path []Path) {

	minPath := &path[0]

	for i := 1; i < len(path); i++ {
		if minPath.Rooms+minPath.ants > path[i].Rooms+path[i].ants {
			minPath = &path[i]
		}
	}

	minPath.Assigned = append(minPath.Assigned, *ant)
	minPath.ants++
}

func MoveAnts(paths map[int][]string, numAnt int) {

	var (
		Mypaths []Path
		stock   = make(map[int][]string)
	)
	for idx, path := range paths {

		stock[idx] = path
		stockPath := Path{
			Id:    stock,
			Rooms: len(path),
		}

		Mypaths = append(Mypaths, stockPath)
	}

	Ants := make([]Ant, numAnt)

	for i := 0; i < numAnt; i++ {
		Ants[i] = Ant{id: i + 1}
	}
	for _, ant := range Ants {
		AssignPath(&ant, Mypaths)
	}

	for idx, path := range Mypaths {
		fmt.Printf("Path%d %v:", idx+1, path.Id[idx+1])
		for i := 0; i < len(path.Assigned); i++ {
			fmt.Printf(" L%d", path.Assigned[i].id)
		}
		fmt.Println()
	}

}
