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
	path = append(path, current, "->")

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

func Check(array [][]string, tab1 []string) bool {

	for _, tab := range array {
		for i := 1; i < len(tab)-1; i++ {
			for j := 1; j < len(tab1)-1; j++ {
				if tab[i] == tab1[j] && tab[i] != "->" && tab1[j] != "->" {
					return true
				}
			}
		}
	}
	return false
}

func Sortarray(array [][]string) [][]string {

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if len(array[i]) > len(array[j]) {
				swap := array[i]
				array[i] = array[j]
				array[j] = swap
			}
		}
	}
	return array
}

func Contain(str string, r rune) bool {

	for _, v := range str {
		if r == v {
			return true
		}
	}
	return false
}

func Verify(array []string, str string) bool {

	for _, v := range array {
		if str == v {
			return true
		}
	}
	return false
}
