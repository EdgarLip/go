package examplesOnly

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type pair_v3 struct {
	hash string
	path string
}

type fileList_v3 []string
type results_v3 map[string]fileList_v3

func hashFile_v3(path string) pair_v3 {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	hash := md5.New() // fast & good enough

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return pair_v3{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

func processFile_v3(path string, pairs chan<- pair_v3, wg *sync.WaitGroup, limits chan bool) {
	defer wg.Done()

	limits <- true

	defer func() {
		<-limits
	}()

	pairs <- hashFile_v3(path)
}

func collectHashes_v3(pairs <-chan pair_v3, result chan<- results_v3) {
	hashes := make(results_v3)

	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}

	result <- hashes
}

func searchTree_v3(dir string, pairs chan<- pair_v3, wg *sync.WaitGroup, limits chan bool) error {
	defer wg.Done()

	visit := func(p string, fi os.FileInfo, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}

		// ignore dir itself to avoid an infinite loop!
		if fi.Mode().IsDir() && p != dir {
			wg.Add(1)
			go searchTree_v3(p, pairs, wg, limits)
			return filepath.SkipDir
		}

		if fi.Mode().IsRegular() && fi.Size() > 0 {
			wg.Add(1)
			go processFile_v3(p, pairs, wg, limits)
		}

		return nil
	}

	limits <- true

	defer func() {
		<-limits
	}()

	return filepath.Walk(dir, visit)
}

func run_v3(dir string) results_v3 {
	workers := 2 * runtime.GOMAXPROCS(0)
	limits := make(chan bool, workers)
	pairs := make(chan pair_v3)
	result := make(chan results_v3)
	wg := new(sync.WaitGroup)

	// we need another goroutine so we don't block here
	go collectHashes_v3(pairs, result)

	// multi-threaded walk of the directory tree; we need a
	// waitGroup because we don't know how many to wait for
	wg.Add(1)

	err := searchTree_v3(dir, pairs, wg, limits)

	if err != nil {
		log.Fatal(err)
	}

	// we must close the paths channel so the workers stop
	wg.Wait()

	// by closing pairs we signal that all the hashes
	// have been collected; we have to do it here AFTER
	// all the workers are done
	close(pairs)

	return <-result
}

func SearchTreeMainCuncurrent_v3(path string) {
    if strings.TrimSpace(path) == "" {
        log.Fatal("Missing parameter, provide dir path!")
		return
    }
	fmt.Println("SearchTreeMainCuncurrent_v3 - path: ", path)
	if hashes := run_v3(path); hashes != nil {
		for hash, files := range hashes {
			if len(files) > 1 {
				// we will use just 7 chars like git
				fmt.Println(hash[len(hash)-7:], len(files))

				for _, file := range files {
					fmt.Println("  ", file)
				}
			}
		}
	}
}
