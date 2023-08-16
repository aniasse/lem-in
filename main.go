package main

import (
	"fmt"
	pkg "lem-in/algorithme"
	"os"
)

func main() {

	if len(os.Args) == 2 && os.Args[1] != "" {

		arg := os.Args[1]
		data_file := pkg.GetDatafile("./data/" + arg)
		paths := pkg.GetRoomLink(&data_file)

		fmt.Println("Les chemins possibles")
		for _, v := range paths {
			fmt.Println(v)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		//Les chemins obtenus avant l'ordonnement des chemins en fonction de leur taille
		validPaths := [][]string{}
		for _, path := range paths {
			if !pkg.Check(validPaths, path) {
				validPaths = append(validPaths, path)
			}
		}
		//Les chemins obtenus apres ordonnement de l'ensemble des chemins en fonction de leur taille
		sortpaths := pkg.Sortarray(paths)
		sortvalidPaths := [][]string{}
		for _, path := range sortpaths {
			if !pkg.Check(sortvalidPaths, path) {
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

		fmt.Println("Le(s) chemin(s) valides")
		for _, path := range lastPath {
			fmt.Println(path)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")

	}
}
