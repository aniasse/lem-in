package main

import (
	"fmt"
	pkg "lem-in/algorithme"
	"os"
)

func main() {

	if len(os.Args) == 2 && os.Args[1] != "" {

		arg := os.Args[1]
		data_file, goodData := pkg.GetDatafile(arg)
		paths, ants := pkg.GetRoomLink(&goodData)
		if len(paths) < 1 {
			fmt.Println("ERROR: Invalid data format")
			return
		}

		for _, data := range data_file {
			fmt.Println(data)
		}
		fmt.Println()
		pkg.MoveAnts(paths, ants)
	}
}
