package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/cavaliercoder/grab"
)

var (
	//count download and convert
	count                = 0
	statusList, nameList []string
	total                int
)

func main() {
	urls := flag.String("u", "", "Url(s) that you want to download.")
	urlFile := flag.String("f", "", "File(s) that stores the URL.")
	flag.Parse()
	if *urls == "" && *urlFile == "" {
		log.Println("No URL input.")
		flag.Usage()
		return
	}
	fmt.Println("Please wait ...")
	UrlList := Sp(*urls + "\n" + ReadFile(*urlFile))
	total = len(UrlList)
	statusList = make([]string, total)
	nameList = make([]string, total)
	for id, url := range UrlList {
		go Download(id, url)
	}
	time.Sleep(1e9)
	fmt.Printf("\x1b[2J") //Clear terminal
	for {
		PrintTable(MakeTable())
		if count == total {
			break
		}
		time.Sleep(1e9)
	}
	fmt.Println("\nALL DONE!")
}

func ReadFile(fns string) string {
	var a string
	fnList := Sp(fns)
	for _, fn := range fnList {
		c, err := ioutil.ReadFile(fn)
		check(err)
		a += string(c) + "\n"
	}
	return a
}

func Sp(str string) []string {
	return strings.FieldsFunc(str, func(r rune) bool {
		if r == ' ' || r == '\n' {
			return true
		}
		return false
	})
}

func Download(id int, url string) {
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", url)
	resp := client.Do(req)
	nameList[id] = resp.Filename
	for {
		status := strconv.FormatFloat(resp.Progress()*100, 'f', 4, 64)
		statusList[id] = status + "%"
		if resp.IsComplete() {
			break
		}
		time.Sleep(1e9) //Sleep 1s
	}
	if check(resp.Err()) {
		statusList[id] = "ERROR"
		count++
		return
	}
	count++
}

func MakeTable() string {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "DOWNLOADING STATUS"},
		},
	}

	var data = make([][]interface{}, total)
	//Rebuild data
	for id := 0; id < total; id++ {
		name := nameList[id]
		status := statusList[id]
		data[id] = []interface{}{id + 1, name, status}
	}
	for _, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
			{Align: simpletable.AlignLeft, Text: row[1].(string)},
			{Align: simpletable.AlignCenter, Text: row[2].(string)},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	return table.String()
}

func PrintTable(str string) {
	arr := strings.Split(str, "\n")
	//Made for each line
	for i := 0; i < len(arr); i++ {
		fmt.Printf("\x1b[%d;%dH\n", i+1, 1)
		fmt.Fprintf(os.Stdout, "\r%s", arr[i])
	}
}

func check(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
