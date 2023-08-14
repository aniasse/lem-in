package pkg

import (
	"strings"
)

func GetRoomLink(array *[]string) (string, string, int) {

	var (
		start string
		end   string
		index int
	)
	for ind, v := range *array {
		if strings.HasPrefix(v, "##start") && ind < len(*array)-1 {
			start = (*array)[ind+1]
		} else if strings.HasPrefix(v, "##end") && ind < len(*array)-2 {
			end = (*array)[ind+1]
			for j := ind + 2; j < len(*array); j++ {
				if strings.Contains((*array)[j], "-") {
					index = j
					break
				}
			}
		}
	}
	start_array := strings.Split(start, " ")
	end_array := strings.Split(end, " ")

	start = start_array[0]
	end = end_array[0]

	return start, end, index
}
