package files

import (
	"bufio"
	"earmuffs/schemas"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const ProjectName = "earmuffs"

// LoadProfanities - Loads the predefined "bad-words" file into memory.
func LoadProfanities(path string) string {
	log.Println("LoadProfanities::path::", path)
	read, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	var lines []string
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		curr := scanner.Text()

		// Only want words with more than 3 letters or "a$$" (not going overboard)
		aWord := "a" + "s" + "s"
		if len(curr) > 3 || curr == aWord {
			lines = append(lines, curr)
		}
	}

	// Process profanities text file into consumable regex
	badWordsRegex := strings.Join(lines, `\b|(?i)`)
	return badWordsRegex
}

// LoadExcludes - Loads the predefined words to exclude if parsed as profanity.
func LoadExcludes(path string) []string {
	log.Println("LoadExcludes::path::", path)
	read, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get each exclude word separately
	var lines []string
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func GetWordsDirPath() string {

	// Absolute filepath when ran as action is "/"
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if dir == "/" {
		return ""
	}

	// For "go run ..." or "go test ..."
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting pwd::", err.Error())
		os.Exit(1)
	}

	// Handle path for test files.
	path := pwd
	if !strings.HasSuffix(pwd, ProjectName) {
		base := strings.Split(pwd, ProjectName)[0]
		path = base + ProjectName + "/data"
	}

	return path + "/"
}

// GetWordsPayload - Wrapper function to get data from required files.
func GetWordsPayload() schemas.WordsPayload {
	path := GetWordsDirPath()
	payload := schemas.WordsPayload{
		ProfanitiesRegex: LoadProfanities(path + "data/words/bad-words.txt"),
		ExcludeWords:     LoadExcludes(path + "data/words/excludes.txt"),
	}

	return payload
}
