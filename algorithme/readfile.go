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
		trim := strings.Trim(scan.Text(), " ")
		if trim != "" {
			array = append(array, trim)
		}
	}
	return array
}
