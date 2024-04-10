package utilsfunctions

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var Utils_helper_var int = 10

func Print_file_name2() {
	fmt.Printf("utilhelper ! \n")
}

func Pring_tag_version_1() {
	fmt.Printf("this is tag version 1  \n")
}

// ****** file functions ******
func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename) // Create the file with the specified filename
	if err != nil {
		return nil, err
	}
	return file, nil
}

func OpenFile(filenamepath string) (*os.File, error) {
	file_content, err := os.OpenFile(filenamepath, os.O_CREATE, 0664) // Open the file for reading (you can change os.O_RDONLY to os.O_RDWR for read-write)
	if err != nil {
		return nil, err // If there's an error, return nil for the file pointer and the error
	}
	return file_content, nil // If successful, return the file pointer and nil error
}

func FileOffsetToStartOfFile(file *os.File) error {
	// Set the file offset to the beginning (offset of 0)
	_, err := file.Seek(0, 0)
	if err != nil {
		return err
	}

	return nil
}

func WriteToFile(file *os.File, content string) error { // writeToFile function writes content to the specified file
	//defer file.Close()
	writer := bufio.NewWriter(file) // Create a new writer for the file

	_, err := writer.WriteString(content) // Write the content to the file
	if err != nil {
		return err
	}

	err = writer.Flush() // Flush the writer to ensure content is written to the file
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(file *os.File) error {
	//defer file.Close() // Ensure the file is closed when we're done with it
	scanner := bufio.NewScanner(file) // Create a new scanner to read the file line by line

	for scanner.Scan() { // Iterate over each line in the file
		line := scanner.Text() // Get the current line
		fmt.Println(line)      // Print the line (you can process the line here)
	}

	if err := scanner.Err(); err != nil { // Check for any scanning errors (like reaching EOF)
		return err
	}

	return nil
}

func DoWork() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func TimeMesureFunc(f func()) string {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	return elapsed.String()
}
