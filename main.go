package main

import (
	"fmt"
	"lem-in/pkg"
	"os"
	"strings"
)

type Graph map[string][]string

func DFS(graph Graph, current string, target string, visited map[string]bool, path []string, paths *[][]string) {
	if current == target {
		newPath := make([]string, len(path))
		copy(newPath, path)
		newPath = append(newPath, target)
		*paths = append(*paths, newPath)
		return
	}

	visited[current] = true
	path = append(path, current, "->")

	for _, neighbor := range graph[current] {
		if !visited[neighbor] { //Au cas ou la chambre n'est pas visite
			DFS(graph, neighbor, target, visited, path, paths)
		}
	}

	visited[current] = false
	// path = path[:len(path)-1]
}

func FindPaths(graph Graph, start string, target string) [][]string {
	visited := make(map[string]bool)
	paths := [][]string{}
	path := []string{}
	DFS(graph, start, target, visited, path, &paths)
	return paths
}

func main() {

	if len(os.Args) == 2 && os.Args[1] != "" {

		arg := os.Args[1]

		graph := make(Graph)

		data_file := pkg.GetDatafile(arg)
		start, end, index := pkg.GetRoomLink(&data_file)

		for i := index; i < len(data_file); i++ {

			if data_file[i][0] == '#' {

			} else {
				split := strings.Split(data_file[i], "-")
				if len(split) != 2 {
					fmt.Printf("Ligne %d invalide\n", i+1)
					continue
				}
				graph[split[0]] = append(graph[split[0]], split[1])
				graph[split[1]] = append(graph[split[1]], split[0])
			}
		}

		paths := FindPaths(graph, start, end)
		validpaths := [][]string{}
		fmt.Printf("Chemin à partir du noeud %s au noeud %s:\n", start, end)
		for _, path := range paths {
			if len(path) != 0 && path[0] == start && path[len(path)-1] == end {
				validpaths = append(validpaths, path)
			}
		}
		for _, path := range validpaths {
			fmt.Println(path)
		}
		// fmt.Println("--------------------------------------------------------------------------------------------")
		// newpath := [][]string{}
		// for _, path := range validpaths {
		// 	if ChoicePath(newpath, path) {
		// 		newpath = append(newpath, path)
		// 	}
		// }
		// for _, path := range newpath {
		// 	fmt.Println(path)
		// }
	}
}
