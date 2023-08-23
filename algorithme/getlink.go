package pkg

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetRoomLink(array *[]string) ([][]string, int, string) {

	var (
		start         string
		end           string
		index         int
		last          int
		checkstart    int
		checkend      int
		start_count   int
		end_count     int
		arrayRoom     []string
		arrayLinkRoom []string
		graph         = make(Graph)
		paths         [][]string
		numAnts       int
		errAnts       error
	)
	if len(*array) != 0 {
		ants := (*array)[0]
		numAnts, errAnts = strconv.Atoi(ants)
		if errAnts != nil || numAnts < 1 {
			fmt.Println("ERROR: Number of ants is invalid")
			os.Exit(0)
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
				arrayRoom = append(arrayRoom, split[0])
				if split[0][0] == 'L' || split[0][0] == '#' {
					fmt.Printf("ERROR: The %v room is invalid\n", split[0])
					os.Exit(0)
				}
				_, erx := strconv.Atoi(split[1])
				_, ery := strconv.Atoi(split[2])
				if erx != nil || ery != nil {
					fmt.Println("ERROR: Invalid data format")
					os.Exit(0)
				}
			} else {
				fmt.Println("ERROR: Invalid data format")
				os.Exit(0)
			}
		}
		if !strings.HasPrefix(v, "##start") {
			start_count++
		}
		if !strings.HasPrefix(v, "##end") {
			end_count++
		}

	}
	if start_count == len(*array) { // Au cas ou il n'y a pas de ligne ##start
		fmt.Println("ERROR: Invalid data format, no start room found")
		os.Exit(0)
	}
	if end_count == len(*array) {
		fmt.Println("ERROR: Invalid data form, no end room found")
		os.Exit(0)
	}
	if checkstart > checkend && end != "" {
		fmt.Println("ERROR: Invalid data format, no start room found")
		os.Exit(0)
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

	// Remplissage du graph
	for i := index; i < last; i++ {
		split := strings.Split((*array)[i], "-")
		if len(split) != 2 {
			fmt.Println("ERROR: Invalid data format")
			os.Exit(0)
		}
		if split[0] == split[1] {
			fmt.Println("ERROR: Invalid data format")
			os.Exit(0)
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

	// Recuperation des valeurs du room de depart et celui d'arriver
	start_array := strings.Split(start, " ")
	end_array := strings.Split(end, " ")
	start = start_array[0]
	end = end_array[0]

	//Verifier s'il y a un room inconnu (non declaré)
	arrayLinkRoom = append(arrayLinkRoom, start, end)
	for _, room := range arrayLinkRoom {
		if !Verify(arrayRoom, room) {
			fmt.Printf("ERROR: The %s  room is invalid\n", room)
			os.Exit(0)
		}
	}

	paths = FindPaths(graph, start, end) //Recuperation des chemins

	//Recuperation des chemins valides avec comme debut (expl:start) et comme fin (expl:end)
	allValidpaths := [][]string{}
	for _, path := range paths {
		if len(path) != 0 && path[0] == start && path[len(path)-1] == end {
			allValidpaths = append(allValidpaths, path)
		}
	}
	allValidpaths = Sortarray(allValidpaths)

	//Les chemins obtenus avant l'ordonnement des chemins en fonction de leur taille

	triplePaths := [][][]string{}
	validPaths := [][]string{}
	for i := 0; i < len(allValidpaths); i++ {
		tab := ChoicePath(allValidpaths, i)
		for _, path := range tab {
			if !Check(validPaths, path) {
				validPaths = append(validPaths, path)
			}
		}
		triplePaths = append(triplePaths, validPaths)
		validPaths = [][]string{}
	}
	for i := 0; i < len(triplePaths); i++ {
		for j := i + 1; j < len(triplePaths); j++ {
			if len(triplePaths[i]) < len(triplePaths[j]) {
				triplePaths[i], triplePaths[j] = triplePaths[j], triplePaths[i]
			}
		}
	}

	//Choix definitif des chemins
	lastPath := [][]string{}
	if len(triplePaths) != 0 {
		lastPath = triplePaths[0]
	}

	lastPath = Sortarray(lastPath)

	return lastPath, numAnts, end
}
