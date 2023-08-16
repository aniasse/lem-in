package pkg

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetRoomLink(array *[]string) ([][]string, int, map[int][]string) {

	var (
		start         string
		end           string
		index         int
		last          int
		checkstart    int
		checkend      int
		count         int
		arrayRoom     []string
		arrayLinkRoom []string
		graph         = make(Graph)
		paths         [][]string
		numAnts       int
		errAnts       error
		Maquette      = make(map[int][]string)
	)
	if len(*array) != 0 {
		ants := (*array)[0]
		numAnts, errAnts = strconv.Atoi(ants)
		if errAnts != nil {
			log.Fatal("ERROR: Number of ants is invalid")
		}
	}
	for ind, v := range *array {
		if strings.HasPrefix(v, "##start") && ind < len(*array)-1 {
			checkstart = ind
			start = (*array)[ind+1]
		} else if strings.HasPrefix(v, "##end") && ind < len(*array)-1 {
			checkend = ind
			end = (*array)[ind+1]
		} else if Contain(v, ' ') { //Pour les lignes contenant les coordonnees des rooms
			split := strings.Split(v, " ")
			if len(split) == 3 {
				if strings.ToUpper(string(split[0][0])) == "L" || strings.ToUpper(string(split[0][0])) == "#" {
					log.Fatalf("ERROR: The %v room is invalid", split[0])
				}
				arrayRoom = append(arrayRoom, split[0])
				_, erx := strconv.Atoi(split[1])
				_, ery := strconv.Atoi(split[2])
				if erx != nil || ery != nil {
					log.Fatal("ERROR: Invalid data format")
				}
			} else {
				log.Fatal("ERROR: Invalid data format")
			}
		}
	}
	if count == len(*array)-1 {
		log.Fatal("ERROR: Invalid data format")
	}

	if checkstart > checkend {
		log.Fatal("ERROR: Invalid data format, no start room found")
	}
	// Pour les lignes contenant les liens des rooms
	for ind, v := range *array {
		if Contain(v, '-') {
			index = ind
			break
		}
	}
	last = index
	for i := index; i < len(*array); i++ {
		if Contain((*array)[i], '-') {
			last++
		}
	}

	for i := index; i < last; i++ {
		split := strings.Split((*array)[i], "-")
		if len(split) != 2 {
			log.Fatalf("ERROR: Invalid data format")
			continue
		}
		if split[0] == split[1] {
			log.Fatal("ERROR: Invalid data format")
		}
		if !Verify(arrayLinkRoom, split[0]) {
			arrayLinkRoom = append(arrayLinkRoom, split[0])
		}
		if !Verify(arrayLinkRoom, split[1]) {
			arrayLinkRoom = append(arrayLinkRoom, split[1])
		}
		graph[split[0]] = append(graph[split[0]], split[1])
		graph[split[1]] = append(graph[split[1]], split[0])
	}

	start_array := strings.Split(start, " ")
	end_array := strings.Split(end, " ")
	start = start_array[0]
	end = end_array[0]

	arrayLinkRoom = append(arrayLinkRoom, start, end)

	for _, room := range arrayLinkRoom {
		if !Verify(arrayRoom, room) {
			log.Fatalf("ERROR: The %s  room is invalid", room)
		}
	}

	paths = FindPaths(graph, start, end)

	//Recuperation des chemins valides avec comme debut (expl:start) et comme fin (expl:end)
	allValidpaths := [][]string{}
	for _, path := range paths {
		if len(path) != 0 && path[0] == start && path[len(path)-1] == end {
			allValidpaths = append(allValidpaths, path)
		}
	}

	fmt.Println("Les chemins possibles")
	for _, v := range paths {
		fmt.Println(v)
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
	//Les chemins obtenus avant l'ordonnement des chemins en fonction de leur taille
	validPaths := [][]string{}
	for _, path := range paths {
		if !Check(validPaths, path) {
			validPaths = append(validPaths, path)
		}
	}
	//Les chemins obtenus apres ordonnement de l'ensemble des chemins en fonction de leur taille
	sortpaths := Sortarray(paths)
	sortvalidPaths := [][]string{}
	for _, path := range sortpaths {
		if !Check(sortvalidPaths, path) {
			sortvalidPaths = append(sortvalidPaths, path)
		}
	}
	//Choix definitif des chemins
	lastPath := [][]string{}
	if len(validPaths) >= len(sortvalidPaths) {
		lastPath = validPaths
	} else {
		lastPath = sortvalidPaths
	}

	for ind, path := range lastPath {
		Maquette[ind+1] = path[2:]
	}

	return allValidpaths, numAnts, Maquette
}
