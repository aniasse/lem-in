package pkg

type Graph map[string][]string

func DFS(graph Graph, current string, target string, visited map[string]bool, path []string, paths *[][]string) {

	if current == target {
		path = append(path, target)
		newPath := make([]string, len(path))
		copy(newPath, path)
		*paths = append(*paths, newPath)
		return
	}

	visited[current] = true
	path = append(path, current)

	for _, neighbor := range graph[current] {
		if !visited[neighbor] { //Au cas ou la chambre n'est pas visite
			DFS(graph, neighbor, target, visited, path, paths)
		}
	}

	visited[current] = false
	path = path[:len(path)-1]
}

func FindPaths(graph Graph, start string, target string) [][]string {
	visited := make(map[string]bool)
	paths := [][]string{}
	path := []string{}
	DFS(graph, start, target, visited, path, &paths)
	return paths
}
