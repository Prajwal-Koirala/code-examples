package main

import (
	"bufio"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Check if file exists
	if fileExists("fileName") {
		fmt.Println("There is a file here.")
	}
	// Get the size of a file
	fileSize, err := fileSize("/")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(fileSize)
	// Get the file extension
	fmt.Println(fileExtension("/random/file/here.json"))
	// Read a file and return the contents
	content, err := readAFile("/random/file/here.json")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(content)
	// Hidden files [.files]
	if hiddenFile(".DS_STORE") {
		fmt.Println("This is a hidden file")
	}
	// Get the sha256 hash of a file
	fmt.Println(generateSHA512("test.file"))
	// Get the path of the current executable
	fmt.Println(getExecutablePath())
	// Read and append the file.
	var exampleContentSlice []string
	exampleContentSlice = readAndAppend("apple.txt", exampleContentSlice)
}

// Check if a file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// Get the size of a file
func fileSize(filepath string) (int64, error) {
	file, err := os.Stat(filepath)
	if err != nil {
		return -1, err
	}
	return file.Size(), nil
}

// Get the file extension
func fileExtension(filename string) string {
	return filepath.Ext(filename)
}

// Read a file and return the contents
func readAFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// Hidden files [.files]
func hiddenFile(filename string) bool {
	return strings.HasPrefix(filename, ".")
}

// Get the sha256 hash of a file
func generateSHA512(filePath string) []byte {
	file, _ := os.Open(filePath)
	defer file.Close()
	hash := sha512.New()
	io.Copy(hash, file)
	return hash.Sum(nil)
}

// Get the path of the current executable
func getExecutablePath() string {
	currentExecutable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return currentExecutable
}

// Read and append the file line by line to a slice.
func readAndAppend(fileLocation string, arrayName []string) []string {
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		arrayName = append(arrayName, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
	}
	return arrayName
}

// get the length of a file
func fileLength(fileName string) int {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return len(content)
}