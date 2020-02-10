package main

import (
	"earmuffs/files"
	"earmuffs/schemas"
	"fmt"
	"log"
	"os"
)

// getEnv - Util to handle retrieving environment vars with defaults.
func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}

func main() {

	// Setup files for parsing.
	excludeFilesEnv := getEnv("INPUT_EXCLUDE_FILES", `\.idea|\.git|node_modules`)
	excludeFiles := excludeFilesEnv + `|bad-words\.txt`
	fileNames, err := files.ListFiles(excludeFiles)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Process and parse results.
	results := files.ParseFiles(fileNames)
	var allResults []schemas.ResultsMessage
	for _, curr := range results {
		if curr.Matches != nil {
			allResults = append(allResults, curr)
		}
	}

	// Iterate all results so we can log them to the user.
	for _, r := range allResults {
		lineMatches := r.Matches

		for _, m := range lineMatches {
			log.Printf(
				`found profanity: %s, in file: %s, at line: %d`,
				m.Profanities,
				r.Path,
				m.Line,
			)
		}
	}

	// Fail the process
	if len(allResults) > 0 {
		panic("error:: found profanities!")
	}

	output := "success!"
	fmt.Println(fmt.Sprintf(`::set-output name=myOutput::%s`, output))
}
