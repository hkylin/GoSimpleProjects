package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func encode(text string) string {
	var count int = 0
	var bins, out string
	out = " "
	letter := []rune(text)
	for i := 0; i < len(letter); i++ {
		bins = ""
		dec := int(letter[i]) //letter to number
		for ii := 0; ii <= 22; ii++ {
			if dec < 1<<uint(ii) {
				break
			}
			count = ii
		}
		for ii := count; ii >= 0; ii-- { //decimal to binary
			if dec-1<<uint(ii) >= 0 {
				bins += "1"
				dec -= 1 << uint(ii)
			} else {
				bins += "0"
			}
		}
		out += bins + " "
	}
	en := strings.Replace(out, " ", "‌", -1) //replace
	en = strings.Replace(en, "1", "​", -1)
	en = strings.Replace(en, "0", "‍", -1)
	return en
}

func decode(text string) string {
	var en, de string
	var out int
	rtext := []rune(text)
	for i := 0; i < len(rtext); i++ { //remove other content
		if string(rtext[i]) == "‌" {
			en += "‌"
		} else if string(rtext[i]) == "​" {
			en += "​"
		} else if string(rtext[i]) == "‍" {
			en += "‍"
		}
	}
	ren := strings.Replace(en, "‌", " ", -1) //replace
	ren = strings.Replace(ren, "​", "1", -1)
	ren = strings.Replace(ren, "‍", "0", -1)
	enb := strings.Fields(ren)
	for i := 0; i < len(enb); i++ {
		out = 0
		bins := enb[i]
		for ii := len(bins) - 1; ii >= 0; ii-- { //binary to decimal
			if bins[:1] == "1" {
				out += (1 << uint(ii))
			}
			bins = bins[1:]
		}
		de += string(out)
	}
	return de
}

func main() {
	input_en := flag.String("e", "", "what you want to encode")
	input_de := flag.String("d", "", "what you want to decode")
	en_be := flag.String("b", "encoded content ->", "what you want to put before encoded content")
	en_af := flag.String("a", "<- encoded content", "what you want to put before encoded content")
	flag.Parse()

	if *input_en == "" && *input_de == "" {
		fmt.Println("Please use -h for help.")
		os.Exit(128)
	}
	if *input_en != "" {
		fmt.Printf("Encoded:\n%s%v%s", *en_be, encode(*input_en), *en_af)
	}
	if *input_de != "" {
		fmt.Printf("Decoded:\n%v", decode(*input_de))
	}
}
