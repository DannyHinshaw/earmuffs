package files

import (
	"bufio"
	"earmuffs/schemas"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var payload = GetWordsPayload()

// doesFileMatch - Handles matching file patterns.
func doesFileMatch(path string, exclude string) bool {
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		excludeRe := regexp.MustCompile(exclude)
		return !excludeRe.Match([]byte(path))
	}

	return false
}

// buildExcludeHandler - Build/preload function to check a word against a list of words to exclude.
func buildExcludeHandler(excludes []string) func(word string) bool {
	return func(word string) bool {
		for _, e := range excludes {
			if strings.ToLower(e) == strings.ToLower(word) {
				return true
			}
		}

		return false
	}
}

// CheckLineProfanity - Parses a line of text for profanity.
func CheckLineProfanity(line string, n int) schemas.LineMatch {
	isExcludeWord := buildExcludeHandler(payload.ExcludeWords)
	re := regexp.MustCompile(payload.ProfanitiesRegex)
	words := strings.Split(line, " ")

	var matches []string
	for _, word := range words {

		allProfanityMatches := re.FindAllString(word, -1)
		for _, match := range allProfanityMatches {

			excluded := isExcludeWord(word)
			if !excluded {
				matches = append(matches, match)
			}
		}
	}

	lineMatch := schemas.LineMatch{
		Profanities: matches,
		Line:        n,
	}

	return lineMatch
}

// scanFile - Opens a file and check it for profanity.
func scanFile(path string) []schemas.LineMatch {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Check each line for profanities
	line := 1
	scanner := bufio.NewScanner(file)
	var matches []schemas.LineMatch
	for scanner.Scan() {

		currLine := scanner.Text()
		lineMatch := CheckLineProfanity(currLine, line)
		if len(lineMatch.Profanities) > 0 {
			matches = append(matches, lineMatch)
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matches
}

// handleScan - Wrapper for scanFile to handle channel responses.
func handleScan(c chan schemas.ResultsMessage, path string) {
	matches := scanFile(path)
	if len(matches) <= 0 {
		empty := schemas.ResultsMessage{
			Matches: nil,
			Path:    "",
		}
		c <- empty
		return
	}

	// Register matches with the file they were found in.
	r := schemas.ResultsMessage{
		Matches: matches,
		Path:    path,
	}

	c <- r
}

// listFiles - List all files in a directory (minus those told to exclude).
func ListFiles(excludeFiles string) ([]string, error) {
	var fileList []string
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		if doesFileMatch(path, excludeFiles) {
			fileList = append(fileList, path)
		}

		return nil
	})

	return fileList, err
}

// parseFiles - Iterate all files and parse for profanities.
func ParseFiles(files []string) []schemas.ResultsMessage {
	c := make(chan schemas.ResultsMessage)
	all := make([]schemas.ResultsMessage, 0)
	stop := len(files)

	// Parse each file as goroutine.
	for _, path := range files {
		go handleScan(c, path)
	}

	// Wait and collect all responses.
	i := 1
	for r := range c {
		all = append(all, r)
		if i == stop {
			close(c)
		}
		i++
	}

	return all
}
