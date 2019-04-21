package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const sep = string(filepath.Separator)

func main() {
	walkDir := flag.String("d", "."+sep, "The directory you want to walk.")
	outDir := flag.String("o", "."+sep+"output"+sep, "The directory you want to output.")
	ext := flag.String("e", "jpg,jpeg,png,raw", "The extension(s) you want to match.")
	rmSame := flag.Bool("r", false, "If you want to remove same file in output directory.")
	flag.Parse()

	if (*outDir)[len(*outDir)-1] != filepath.Separator {
		*outDir += sep
	}

	extExist := ExtSet(*ext)
	WalkDir(*walkDir, *outDir, extExist)
	if *rmSame {
		RemoveSameFile(*outDir)
	}
}

func ExtSet(ext string) map[string]bool {
	extExist := make(map[string]bool)
	extList := strings.Split(ext, ",")
	for _, ext := range extList {
		extExist["."+ext] = true
	}
	return extExist
}

func WalkDir(walkDir, outDir string, extExist map[string]bool) {
	var pathList []string
	err := filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ext := filepath.Ext(path)
		if !info.IsDir() && extExist[ext] {
			pathList = append(pathList, path)
		}
		return nil
	})
	checkErr(err)
	CopyFile(pathList, outDir)
}

func CopyFile(pathList []string, outDir string) {
	if len(pathList) == 0 {
		return
	}
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err = os.MkdirAll(outDir, 0755)
		checkErr(err)
	}
	for _, path := range pathList {
		fileName := filepath.Base(path)
		bytes, err := ioutil.ReadFile(path)
		checkErr(err)
		err = ioutil.WriteFile(outDir+fileName, bytes, 0644)
		checkErr(err)
	}
}

func RemoveSameFile(workDir string) {
	SHA1List := make(map[string]bool)
	makeSHA1 := func(bytes *[]byte) string {
		h := sha1.New()
		_, err := io.WriteString(h, string(*bytes))
		checkErr(err)
		return fmt.Sprintf("%x", h.Sum(nil))
	}
	err := filepath.Walk(workDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		bytes, err := ioutil.ReadFile(path)
		checkErr(err)
		hash := makeSHA1(&bytes)
		if SHA1List[hash] {
			err := os.Remove(path)
			checkErr(err)
		} else {
			SHA1List[hash] = true
		}
		return nil
	})
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
