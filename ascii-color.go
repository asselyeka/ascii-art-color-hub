package ascii

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PrintColor = "\033[38;5;%dm%s\033[39;49m"
)

func IndexOfColoredLetter(s string, colorStr string) []int {
	flagWord := false
	var indexArr []int
	for i, letter := range []byte(s) {
		if letter == colorStr[0] {
			lenColor := 0
			j := i
			for _, l := range []byte(colorStr) {
				if s[j] == l {
					lenColor++
				}
				j++
			}
			if lenColor == len(colorStr) {
				flagWord = true
				for k := i; k < i+len(colorStr); k++ {
					indexArr = append(indexArr, k)
				}
				return indexArr
			}
		}
	}

	if flagWord == false {
		indexStart := 0
		indexEnd := len(s) - 1
		index, err := strconv.Atoi(colorStr)
		if err == nil {
			indexArr = append(indexArr, index)
		} else {
			flagFormat := false

			for i, l := range colorStr {
				if l == ':' {
					flagFormat = true
					if i-1 < 0 && i+1 < len(colorStr) {
						indexEnd1, err2 := strconv.Atoi(colorStr[i+1:])
						indexEnd = indexEnd1
						if err2 != nil {
							fmt.Printf("Parse error: Please provide with last index of letters to be colored in format: \":number\"")
							return indexArr
						}
					} else if i-1 >= 0 && i+1 >= len(colorStr) {
						indexStart1, err1 := strconv.Atoi(colorStr[:i])
						indexStart = indexStart1
						if err1 != nil {
							fmt.Printf("Parse error: Please provide with first index of letters to be colored in format: \"number:\"\n")
							return indexArr
						}
					} else if i-1 >= 0 && i+1 < len(colorStr) {
						indexStart1, err1 := strconv.Atoi(colorStr[:i])
						indexEnd1, err2 := strconv.Atoi(colorStr[i+1:])
						indexStart = indexStart1
						indexEnd = indexEnd1
						if err1 != nil || err2 != nil {
							fmt.Printf("Parse error: Please provide with first and last indexes of letters to be colored in format: \"number:number\"\n")
							return indexArr
						}
					} else {
						fmt.Printf("Parse error: Please provide with first or/and last indexes of letters to be colored in format: \"number:number\"\n")
						return indexArr
					}
					for i := indexStart; i <= indexEnd; i++ {
						indexArr = append(indexArr, i)
					}
					break
				} else if l == ',' {
					flagFormat = true
					indexes := strings.Split(colorStr, ",")
					for _, v := range indexes {
						i, err := strconv.Atoi(v)
						if err == nil {
							indexArr = append(indexArr, i)
						} else {
							fmt.Printf("Parse error: Please provide with indexes of letters to be colored in format: \"number,number,number...\"\n")
							return indexArr
						}
					}
					break
				}
			}
			if flagFormat == false {
				fmt.Printf("Error format: Please provide letters or indexes as in exapmle: 4:6; 1:; :5; 5,6,4,2; 11, etc.\n")
				return indexArr
			}
		}
	}

	return indexArr
}

func Color(s string) int {
	switch s {
	case "blue":
		return 4
	case "green":
		return 2
	case "red":
		return 1
	case "yellow":
		return 3
	case "purple":
		return 5
	case "magenta":
		return 6
	case "orange":
		return 130
	default:
		num, err := strconv.Atoi(s)
		if err != nil {
			return 7
		}
		return num
	}
}
