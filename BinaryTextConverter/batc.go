package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func txt_to_bin(input_txt string) {
	for i := 0; i < len(input_txt); i++ { //split into letters
		output_txt := input_txt[i]
		output_dec := rune(output_txt)
		var zeros string
		for i := 8; i >= 0; i-- { //how many "0" should be put before each result block
			if 1<<(uint(i+1)) > output_dec && output_dec >= 1<<(uint(i)) {
				for ii := i; ii < 7; ii++ {
					zeros += "0"
				}
			}
		}
		fmt.Printf("%v%b ", zeros, output_dec)
	}
	fmt.Print("\n")
}
func bin_to_txt(input_bin string) {
	input_bin = strings.Replace(input_bin, " ", "", -1) //remove spaces
	var output_int int
	var input_cache string
	for len(input_bin) >= 8 { //split into every ASCII block
		input_cache = input_bin[:8]
		input_bin = input_bin[8:]
		output_int = 0
		for i := 7; i >= 0; i-- { //binary to decimal
			if input_cache[:1] == "1" {
				output_int += (1 << uint(i))
			}
			input_cache = input_cache[1:]
		}
		fmt.Print(string(output_int)) //decimal to letter
	}
	fmt.Print("\n")
}

func main() {
start:
	var choose, input, bot string
	fmt.Println("==========================\n------\n1.binary to text\n2.text to binary\n------\nPlease choose:")
	fmt.Scanln(&choose)
	switch choose {
	case "1":
		bot = "binary"
	case "2":
		bot = "text"
	default:
		fmt.Println("\n#############\n#   ERROR   #\n#############\nPlease restart!")
		goto start
	}
	fmt.Printf("==========================\nPlease enter the %v:\n", bot)
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	input = string(data)
	fmt.Println("==========================\nYour output is:")
	switch choose {
	case "1":
		bin_to_txt(input)
	case "2":
		txt_to_bin(input)
	}
	fmt.Print("==========================\n")
}
