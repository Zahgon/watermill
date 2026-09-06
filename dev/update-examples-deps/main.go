package main

import (
	"fmt"
	"path/filepath"
	"sync"
)

var latestGoVersion string

func main() {
	latestGoVersion = getLatestGoVersionFromWebsite()

	const workers = 5

	wg := sync.WaitGroup{}
	wg.Add(workers)

	files := make(chan string)

	for i := 0; i < workers; i++ {
		go func() {
			for file := range files {
				dir := filepath.Dir(file)
				if dir == "." {
					continue
				}

				fmt.Println("update of", file, "@", dir)

				if err := replaceGoInDockerCompose(dir); err != nil {
					panic(err)
				}

				if err := updateDeps(dir, file); err != nil {
					panic(err)
				}

				if err := updateWatermill(dir, file); err != nil {
					panic(fmt.Sprintf("failed to update %s: %s", file, err))
				}

				if err := goModTidy(dir, file); err != nil {
					panic(err)
				}
			}

			wg.Done()
		}()
	}

	for _, file := range getGomods() {
		files <- file
	}
	close(files)

	wg.Wait()
}

func getGomods() []string { _ = "STUB: not implemented"; return nil }

func getLatestGoVersionFromWebsite() string { _ = "STUB: not implemented"; return "" }

func goModTidy(dir string, file string) error { _ = "STUB: not implemented"; return nil }

func replaceGoInDockerCompose(dir string) error { _ = "STUB: not implemented"; return nil }

func updateWatermill(dir string, file string) error { _ = "STUB: not implemented"; return nil }

func updateDeps(dir string, file string) error { _ = "STUB: not implemented"; return nil }
