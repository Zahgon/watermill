package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	bigFatGomod := ""

	for _, fileName := range getGomods() {
		dir := filepath.Dir(fileName)
		if dir == "." {
			continue
		}

		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		fileMod := ""

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			txt := scanner.Text()
			if strings.HasPrefix(txt, "go ") {
				continue
			}
			if strings.HasPrefix(txt, "module ") {
				continue
			}

			fileMod += txt + "\n"
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

		if fileMod != "" {
			bigFatGomod += "// " + fileName + "\n"
			bigFatGomod += fileMod + "\n"
		}

		_ = file.Close()

		if err := os.Remove(fileName); err != nil {
			panic(err)
		}
	}

	fmt.Println(bigFatGomod)
}

func getGomods() []string { _ = "STUB: not implemented"; return nil }
