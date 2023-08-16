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
		paths, _, maquette := pkg.GetRoomLink(&data_file)

		fmt.Println("Le(s) chemin(s) valides")
		for _, path := range paths {
			fmt.Println(path)
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		for ind, path := range maquette {
			fmt.Println(ind, path)
		}
	}
}
