package pkg

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetDatafile(str string) []string {

	var array []string
	file, err := os.Open(str)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		if strings.Trim(scan.Text(), " ") != "" {
			array = append(array, scan.Text())
		}
	}
	return array
}
