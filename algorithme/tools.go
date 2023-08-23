package pkg

func Check(array [][]string, tab1 []string) bool {

	for _, tab := range array {
		for i := 1; i < len(tab)-1; i++ {
			for j := 1; j < len(tab1)-1; j++ {
				if tab[i] == tab1[j] {
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
				array[i], array[j] = array[j], array[i]
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

func ChoicePath(array [][]string, ind int) [][]string {

	if ind < 0 || ind >= len(array) {
		return array
	}

	tab := make([][]string, len(array))

	copy(tab[0:], array[ind:ind+1])
	copy(tab[1:], array[0:ind])
	copy(tab[ind+1:], array[ind+1:])

	return tab
}
