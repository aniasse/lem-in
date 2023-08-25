package pkg

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetDatafile(str string) ([]string, []string) {

	var (
		array []string
		contentFile []string
	)
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
		if Contain(trim, '#') && (trim=="##start" || trim == "##end") && trim!=""{
			contentFile = append(contentFile, trim)
		} else if !Contain(trim, '#')&& trim!="" {
			contentFile = append(contentFile, trim)
		}
	}
	return array, contentFile
}
