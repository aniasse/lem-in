package main

import (
	"fmt"
	pkg "lem-in/algorithme"
	"os"
)

func main() {

	if len(os.Args) == 2 && os.Args[1] != "" {

		arg := os.Args[1]
		data_file := pkg.GetDatafile(arg)
		paths, ants, end := pkg.GetRoomLink(&data_file)

		for _, data := range data_file {
			fmt.Println(data)
		}
		fmt.Println()
		pkg.MoveAnts(paths, ants, end)
	}
}
