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
)

type _pair struct {
	hash string
	path string
}

type _fileList []string
type _results map[string]_fileList

func _hashFile(path string) _pair {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	hash := md5.New() // fast & good enough

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return _pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

/*
a process that runs in a goroutine.
takes a file that comes from the "paths" channel and hashes it using the _hashFile function.
then the hash is sent to the "pairs" channel in the format of _pair struct

	when key: is a hash string
	     value: is a path to the file.

then the done channel is sent a true value to indicate that the process is done.
*/
func processFiles(paths <-chan string, pairs chan<- _pair, done chan<- bool) {
	for path := range paths {
		pairs <- _hashFile(path)
	}
	// the indication that the paths channel is closed is:
	// when the function "_searchTree" will finish it's work , and then in the run
	// funtion "close(paths)" will be called, and then the range will stop.
	done <- true
}

func collectHashes(pairs <-chan _pair, result chan<- _results) {
	hashes := make(_results)

	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}
	// what is the indication of the paris range to stop ?
	// in the body of the run funtion, when the function "close(pairs)" will be called, and then the range will stop.
	result <- hashes
}

func _searchTree(dir string, paths chan<- string) error {
	visit := func(p string, fi os.FileInfo, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}

		if fi.Mode().IsRegular() && fi.Size() > 0 {
			// this is the trigger what will make the workers goroutines to start working.
			// once the workers will start getting paths from the paths channel, they will start hashing the files.
			paths <- p
		}

		return nil
	}

	return filepath.Walk(dir, visit)
}

func run(dir string) _results {
	workers := 2 * runtime.GOMAXPROCS(0)
	paths := make(chan string)
	pairs := make(chan _pair)
	done := make(chan bool)
	result := make(chan _results)

	for i := 0; i < workers; i++ {
		//this is the place where i spin up the workers goroutines.
		// each worker will wait for a path from the paths channel, and then will hash it using the _hashFile function.
		// many workers will send their results to the pairs channel.
		// while there will be only one! goroutine that will collect the results from the pairs channel and will put them in the result map.
		go processFiles(paths, pairs, done)
	}

	// we need another goroutine so we don't block here
	go collectHashes(pairs, result)

	if err := _searchTree(dir, paths); err != nil {
		return nil
	}

	// we must close the paths channel so the workers stop
	close(paths)

	// wait for all the workers to be done
	// This loop is needed to ensure that all worker goroutines have completed their work
	// before we proceed to close the 'pairs' channel. Each worker signals completion by sending to 'done'.
	// By waiting for 'workers' number of messages from 'done', we guarantee all workers are finished.
	for i := 0; i < workers; i++ {
		<-done // this line blocks , until the done channel will get a true value , and then the loop will continue.
		// which means that another worker done it's work.
	}

	// by closing pairs we signal that all the hashes
	// have been collected; we have to do it here AFTER
	// all the workers are done
	close(pairs)

	return <-result
}

func SearchTreeMainCuncurrent(path string) {
	if strings.TrimSpace(path) == "" {
		log.Fatal("Missing parameter, provide dir path!")
		return
	}
	fmt.Println("SearchTreeMain - path: ", path)

	if hashes := run(path); hashes != nil {
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
