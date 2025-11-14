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

type pair_v2 struct {
	hash string
	path string
}

type fileList_v2 []string
type results_v2 map[string]fileList_v2

func hashFile_v2(path string) pair_v2 {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	hash := md5.New() // fast & good enough

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return pair_v2{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

func processFile_v2(path string, pairs chan<- pair_v2, wg *sync.WaitGroup, limits chan bool) {
	defer wg.Done()

	limits <- true

	defer func() {
		<-limits
	}()

	pairs <- hashFile_v2(path)
}

func collectHashes_v2(pairs <-chan pair_v2, result chan<- results_v2) {
	hashes := make(results_v2)

	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}

	result <- hashes
}

func searchTree_v2(dir string, pairs chan<- pair_v2, wg *sync.WaitGroup, limits chan bool) error {
	defer wg.Done()

	visit := func(p string, fi os.FileInfo, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}
		// In this block:
		// "dir" is the directory passed to this invocation of searchTree_v2 (the "root" of the tree/subtree we're currently exploring).
		// "p" is the current filesystem path entry examined by filepath.Walk as it descends through the tree.
		// For example, if we're doing searchTree_v2("/some/path", ...),
		//   - On the first call: dir == "/some/path"
		//   - For each recursive goroutine spawned for a subdirectory (say "/some/path/a"), dir == "/some/path/a".
		//
		// Line 72:
		//  if fi.Mode().IsDir() && p != dir {
		// This checks: "Is the current entry a directory (but not the starting directory for this searchTree_v2 invocation)?"
		// Example: While traversing "/some/path", when we encounter "/some/path/a" (a subdirectory), p=="/some/path/a", dir=="/some/path".
		// This block will run for every subdirectory encountered except the root.
		// We then:
		//   - Increment the wait group (wg.Add(1)) to account for the new goroutine we'll launch.
		//   - Start a new goroutine to recursively walk ("searchTree_v2") the subdirectory, with dir as that subdirectory.
		//   - Return filepath.SkipDir which tells filepath.Walk: "Don't descend into this subdirectory right now; we've delegated it to a goroutine."
		if fi.Mode().IsDir() && p != dir {
			wg.Add(1)
			go searchTree_v2(p, pairs, wg, limits)
			return filepath.SkipDir
		}

		// Line 78: If the current entry is a regular file and has a size greater than 0,
		// we increment the wait group counter because we are about to process this file concurrently.
		// Then, we launch "processFile_v2" as a goroutine, which will hash the file and send the result to the "pairs" channel.
		if fi.Mode().IsRegular() && fi.Size() > 0 {
			wg.Add(1)
			go processFile_v2(p, pairs, wg, limits)
		}

		// Why do I need 'return nil' here?
		// Because the 'visit' function used by filepath.Walk must return an error (or nil).
		// Returning nil tells Walk to keep traversing; returning a non-nil error would halt traversal or alter the behavior.
		return nil
	}

	// We need this line to enforce the concurrency limit: sending a value into the 'limits' channel blocks
	// if the channel is already full, thereby limiting the number of concurrent goroutines running this function and processFile_v2.
	// This effectively acts as a semaphore; only 'cap(limits)' goroutines can proceed past this point at once.
	limits <- true

	// The following defer ensures that once this function (searchTree_v2) completes,
	// it frees up a slot in the 'limits' channel, effectively releasing one permit in our
	// concurrency semaphore. This allows another goroutine that's blocked (waiting to
	// send into 'limits') to proceed. Placing this right after acquiring the slot with
	// 'limits <- true' ensures that *every* time we enter searchTree_v2, we also guarantee
	// the permit is eventually released, regardless of how the function exits.
	defer func() {
		<-limits
	}()

	return filepath.Walk(dir, visit)
}

func run_v2(dir string) results_v2 {
	workers := 2 * runtime.GOMAXPROCS(0)
	limits := make(chan bool, workers)
	pairs := make(chan pair_v2)
	result := make(chan results_v2)
	wg := new(sync.WaitGroup)

	// we need another goroutine so we don't block here
	go collectHashes_v2(pairs, result)

	// multi-threaded walk of the directory tree; we need a
	// waitGroup because we don't know how many to wait for
	wg.Add(1)

	err := searchTree_v2(dir, pairs, wg, limits)

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

func SearchTreeMainCuncurrent_v2(path string) {
    if strings.TrimSpace(path) == "" {
        log.Fatal("Missing parameter, provide dir path!")
		return
    }
	fmt.Println("SearchTreeMainCuncurrent_v2 - path: ", path)

	if hashes := run_v2(path); hashes != nil {
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
