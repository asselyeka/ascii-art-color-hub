package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	ascii ".."
)

const (
	PrintColor = "\033[38;5;%dm%s\033[39;49m"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected argument(you can set it's color and colored indexes or letters)")
		os.Exit(1)
	} else {
		str := os.Args[1]
		var indexArr []int
		var argColor *string
		indexOfColoredLetter := ""

		if len(os.Args) >= 3 && os.Args[2][:7] == "--color" {
			argCmd := flag.NewFlagSet(str, flag.ExitOnError)
			argColor = argCmd.String("color", "white", "color in string")
			argCmd.Parse(os.Args[2:])

			if len(os.Args) >= 4 {
				indexOfColoredLetter = argCmd.Arg(0)
			}

			if indexOfColoredLetter == "" {
				for i := range str {
					indexArr = append(indexArr, i)
				}
			} else {
				indexArr = ascii.IndexOfColoredLetter(str, indexOfColoredLetter)
				if indexArr == nil {
					os.Exit(1)
				}
			}

			//fmt.Println(*argColor)
			//fmt.Println(indexOfColoredLetter)
			//fmt.Println(indexArr)
			//return
		}

		// "\n" handling
		splittwo := string(byte(92)) + string(byte(110))
		words := strings.Split(str, splittwo)

		// read from file
		fileName := "standard.txt"

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer file.Close()

		rawBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(rawBytes), "\n")

		//print string to terminal
		for _, word := range words {
			for h := 0; h < 9; h++ {
				for indexLetter, l := range []byte(word) {
					flagPrint := false
					if indexArr != nil {
						for _, indx := range indexArr {
							if indx == indexLetter {
								flagPrint = true
							}
						}
						for i, line := range lines {
							if i == (int(l)-32)*9+h {

								if flagPrint == true {

									fmt.Printf(PrintColor, ascii.Color(*argColor), line)
								} else {
									fmt.Printf(PrintColor, 7, line)
								}
							}
						}
					}
				}
				fmt.Println()
			}
		}
		if ascii.Color(*argColor) == 7 {
			fmt.Printf("Please use this colors:\n")
			fmt.Printf(PrintColor, 1, "red\n")
			fmt.Printf(PrintColor, 2, "green\n")
			fmt.Printf(PrintColor, 3, "yellow\n")
			fmt.Printf(PrintColor, 4, "blue\n")
			fmt.Printf(PrintColor, 5, "purple\n")
			fmt.Printf(PrintColor, 6, "magenta\n")
			fmt.Printf(PrintColor, 7, "white\n")
			fmt.Printf(PrintColor, 130, "orange\n")
		}
	}
}
