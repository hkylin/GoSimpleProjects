package main

import (
	"flag"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	fpath_p := flag.String("p", "", "The path of your folder which you want to rename all the files under it into UUID file name.")
	flag.Parse()
	if *fpath_p == "" {
		fmt.Println("No path.\n\tUse -h for more information.")
		os.Exit(3)
	}
	fpath := *fpath_p
	files, err := ioutil.ReadDir(fpath)
	check(err)
	for _, file := range files {
		suffix := path.Ext(file.Name())
		uuid := uuid.Must(uuid.NewV4()).String()
		err := os.Rename(fpath+file.Name(), fpath+uuid+suffix)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
