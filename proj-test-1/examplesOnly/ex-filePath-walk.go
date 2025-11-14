package examplesOnly

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type pair struct {
    hash, path string
}

type fileList []string
type results map[string]fileList


func hashFile(path string) pair {
    file, err := os.Open(path)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    hash := md5.New() 

    if _, err := io.Copy(hash, file); err != nil {
        log.Fatal(err)
    }

    return pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
}


func searchTree(dir string) (results, error) {
    hashes := make(results)

    err := filepath.Walk(dir, func(p string, fi os.FileInfo, err error) error {
        // ignore the error parm for now
        // Check if fi is nil (happens when there's an error accessing the file/dir)
        if fi == nil || err != nil {
			fmt.Println("SearchTreeMain, error: ", err)
            return err // or return err if you want to propagate the error
        }
        if fi.Mode().IsRegular() && fi.Size() > 0 {
            h := hashFile(p)
            hashes[h.hash] = append(hashes[h.hash], h.path)
        }

        return nil
    })

    return hashes, err
}


// To execute this function, you need to call `searchTreeMain()` from your `main()` function or another function in your Go program.
// For example, in your main.go file, you could add:
//     examplesOnly.searchTreeMain("/c/Users/xxx/go_projects") // for linux use /c/Users/xxx/go_projects
//     examplesOnly.searchTreeMain("C:/Users/xxx/go_projects")  // for windows use C:/Users/xxx/go_projects
// Note: This function expects a directory path as a command-line argument.
func SearchTreeMain(path string) {
    if strings.TrimSpace(path) == "" {
        log.Fatal("Missing parameter, provide dir path!")
		return
    }
	fmt.Println("SearchTreeMain - path: ", path)

    if hashes, err := searchTree(path); err == nil {
        for hash, files := range hashes {
            if len(files) > 1 {
                // we will use just 7 chars like git
                fmt.Printf("for hash %v there are %v files\n",hash[len(hash)-7:], len(files))

                for _, file := range files {
                    fmt.Printf("  %v\n", file)
                }
            }
        }
    }
}